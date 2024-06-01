import './Components.css'
import React, { useEffect, useState } from 'react';

const Footer = () => {
    const [scrolled, setScrolled] = useState(false);

    useEffect(() => {
        const handleScroll = () => {
            const isScrolled = window.scrollY > 5;
            console.log("Scrolling:", isScrolled);
            setScrolled(isScrolled);
        };

        window.addEventListener('scroll', handleScroll);

        return () => {
            window.removeEventListener('scroll', handleScroll);
        };
    }, []);

    return (
        <div className={scrolled ? 'scrolled' : 'block'}>
            <div id="footer">
                {"Autores: Elliott Victoria, Franciosi Ignacio, Havenstein Carolina, Morabito Leonardo "+
                    "Copyright© 2024. Todos los derechos reservados"}
            </div>
        </div>
    );
};

export default Footer;