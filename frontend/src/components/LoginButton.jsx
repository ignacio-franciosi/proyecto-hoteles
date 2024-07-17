import React, {useEffect, useState} from "react";
import './Components.css'
import {useNavigate} from "react-router-dom";
import Cookies from "js-cookie";

const LoginButton = () => {
    const navigate = useNavigate();
    const userEmail = Cookies.get("email");
    const login = () => {
        navigate("/login");
    };
    const [isAuthenticated, setIsAuthenticated] = useState(false); // Supongamos que esta variable controla el estado de autenticación.

    function logoutButton() {
        const message = "Deseas cerrar sesión?";
        const result = window.confirm(message);

        if (result) {
            Cookies.set('user_id', "-1");
            Cookies.set('token', "");
            Cookies.set('email', "");
            Cookies.set('city', "");
            Cookies.set('startDate', "");
            Cookies.set('endDate', "");
            alert("Sesión cerrada! vuelva pronto ;)");
            navigate("/")
            document.location.reload()

        }
    }

    useEffect(() => {
        // Lógica para verificar la autenticación (usando JWT u otra lógica)
        const token = Cookies.get("token"); // Recupera el token JWT almacenado en localStorage
        if (token) {
            // Verifica si el token es válido
            setIsAuthenticated(true);
        }
    }, []);

    return (
        <div>
            {isAuthenticated ? ( // Usamos un operador ternario para mostrar el botón adecuado
                <header id="header">
                    <button id="loginButton" onClick={logoutButton}>{userEmail ? `Bienvenido, ${userEmail}` : "Bienvenido"}</button>

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
