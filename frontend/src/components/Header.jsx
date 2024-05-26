import React from 'react';
import './Components.css'
import {Link} from "react-router-dom";


const Header = () => {
        return (
            <header id="header">
                <Link to= '/'>
                <h1> hoteles.com</h1>
                </Link>
            </header>

        );
    };
export default Header;