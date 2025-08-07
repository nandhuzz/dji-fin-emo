// File: cmd/server/main.go
package main

import (
	"context"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nandhuzz/go-basics/pkg/auth"
	"github.com/nandhuzz/go-basics/pkg/config"
	"github.com/nandhuzz/go-basics/pkg/db"
	"github.com/nandhuzz/go-basics/pkg/logger"

	"github.com/nandhuzz/go-basics/internal/user"
)

func main() {
	logger.Init()
	cfg := config.Load()
	dbPool := db.Connect(cfg.DatabaseURL)
	defer dbPool.Close()
	db.RunMigrations(dbPool, "migrations")

	ctx := context.Background()
	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := dbPool.Ping(ctxTimeout); err != nil {
		log.Fatalf("❌ Unable to ping database: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	// Register additional MIME types
	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".css", "text/css")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Serve static files from dist
	serveStaticFiles(r)

	// Setup API routes
	setupRoutes(r, dbPool)

	port := getEnv("PORT", "8080")
	log.Printf("✅ Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}

func serveStaticFiles(r *chi.Mux) {
	distDir := "./frontend/dist"

	r.HandleFunc("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(distDir, r.URL.Path)
		ext := filepath.Ext(path)

		// Set correct MIME type
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}

		http.ServeFile(w, r, path)
	})

	// SPA fallback
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
	})
}

func setupRoutes(r *chi.Mux, pool *pgxpool.Pool) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Use(withTx(pool))
		r.Mount("/", user.Router(func(ctx context.Context) user.Service {
			tx := txFromContext(ctx)
			return user.NewService(user.NewRepository(tx))
		}))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware)
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
			userID, _ := auth.UserIDFromContext(r.Context())
			w.Write([]byte(`{"message": "Protected /accounts route for user: ` + userID + `"}`))
		})
	})
}

// pgx Tx middleware

type contextKey string

const txKey = contextKey("tx")

func withTx(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tx, err := pool.Begin(r.Context())
			if err != nil {
				http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
				return
			}
			defer tx.Rollback(r.Context())

			ctx := context.WithValue(r.Context(), txKey, tx)
			next.ServeHTTP(w, r.WithContext(ctx))

			if err := tx.Commit(r.Context()); err != nil {
				http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
				return
			}
		})
	}
}

func txFromContext(ctx context.Context) pgx.Tx {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok {
		panic("transaction not found in context")
	}
	return tx
}
