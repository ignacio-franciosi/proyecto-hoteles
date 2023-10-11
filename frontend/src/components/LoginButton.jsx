import React, {useEffect, useState} from "react";
import './Components.css'
import {useNavigate} from "react-router-dom";

const LoginButton = () => {
    const navigate = useNavigate();
    const userEmail = localStorage.getItem('email');
    const login = () => {
        navigate("/login");
    };
    const [isAuthenticated, setIsAuthenticated] = useState(false); // Supongamos que esta variable controla el estado de autenticación.

    useEffect(() => {
        // Lógica para verificar la autenticación (usando JWT u otra lógica)
        const token = localStorage.getItem("jwt_token"); // Recupera el token JWT almacenado en localStorage
        if (token) {
            // Verifica si el token es válido
            setIsAuthenticated(true);
        }
    }, []);

    return (
        <div>
            {isAuthenticated ? ( // Usamos un operador ternario para mostrar el botón adecuado
                <header id="header">
                    <button id="loginButton">{userEmail || "Bienvenido"}</button>
                </header>
            ) : (
                <header id="header">
                    <button id="loginButton" onClick={login}>Iniciar sesión</button>
                </header>
            )}
        </div>
    );
};

export default LoginButton;
