import Header from "@layout/Header";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(false);

    const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;
  const navigate = useNavigate();

    console.log("Backend URL:", BACKEND_URL);

    const handleEmailChange = (e) => {
        setEmail(e.target.value);
    };

    const handlePasswordChange = (e) => {
        setPassword(e.target.value);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError(null);
        setLoading(true);

        try {
            const response = await fetch(`${BACKEND_URL}/users/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ email, password })
            });

            if (!response.ok) {
                throw new Error("Invalid email or password");
            }

            const data = await response.json();
            console.log("Login successful:", data);

            if (data.token) {
                localStorage.setItem("token", data.token);
                navigate("/")
            }

        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            <Header />
            <form onSubmit={handleSubmit} className="flex flex-col gap-2 max-w-xs mx-auto mt-8">
                <input
                    type="text"
                    placeholder="Email Id"
                    value={email}
                    onChange={handleEmailChange}
                    className="border border-gray-300 rounded p-2"
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={handlePasswordChange}
                    className="border border-gray-300 rounded p-2"
                />
                <button
                    type="submit"
                    className="bg-blue-500 text-white rounded p-2 hover:bg-blue-600"
                >
                    Login
                </button>
            </form>
        </>
    );
};

export default Login;
