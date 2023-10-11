import React from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import './Components.css';

const BackButton = () => {
    const navigate = useNavigate();
    const buttonReturn = () => {
        navigate(-1);
    };

    return (
        <header id="header">
            <button id="backButton" onClick={buttonReturn}>Atrás</button>
        </header>
    );
};

function ComponenteE() {
    const location = useLocation();

    // Lógica para determinar si el botón debe ocultarse en función de la ruta actual
    const hideButton = location.pathname === '/' || location.pathname === '/home';
    return hideButton ? null : (
            <BackButton />
    );
}

export default ComponenteE;
