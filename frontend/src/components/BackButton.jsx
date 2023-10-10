import React from 'react';
import { useNavigate } from 'react-router-dom';
import './Components.css'



const BackButton = () => {

    const navigate = useNavigate();
    const buttonReturn = () => {
        navigate(-1);
    };

    return (
        <header id="header">
            <button id="BackButton" onClick={buttonReturn}>Atrás</button>
        </header>

    );
};

export default BackButton;
