import React from 'react';
import './Components.css'
import {useNavigate} from "react-router-dom";



const LoginButton = () => {

    const navigate = useNavigate();
    const login = () => {
        navigate("/login");
    };

    return (
        <header id="header">
            <button id="loginButton" onClick={login}>Iniciar sesión</button>
        </header>

    );
};
export default LoginButton;