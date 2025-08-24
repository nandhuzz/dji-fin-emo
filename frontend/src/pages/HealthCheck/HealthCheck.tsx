import { useEffect, useState } from "react";

const HealthCheck = () => {
    const [status, setStatus] = useState("Checking...");
    const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;

    useEffect(() => {
        fetch(`${BACKEND_URL}/health`)
            .then((res) => res.text()) // use res.json() if backend returns JSON
            .then((data) => setStatus(data))
            .catch((err) => {
                console.error(err);
                setStatus("Error connecting to backend");
            });
    }, [BACKEND_URL]);

    return <div className="p-4 text-center font-medium">{status}</div>;
};

export {HealthCheck};
