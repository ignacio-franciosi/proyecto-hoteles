import React from 'react';
import './Components.css'

import {useNavigate} from "react-router-dom";


const Header = () => {
    const navigate = useNavigate();
        return (
            <header id="header">
                <h1 > hoteles.com</h1>
            </header>

        );
    };
export default Header;