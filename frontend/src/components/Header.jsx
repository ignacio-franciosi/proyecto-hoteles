import './Components.css'
import {Link, useLocation, useNavigate} from "react-router-dom";
import React, {useEffect, useState} from "react";
import Cookies from 'js-cookie';
import CustomModal2 from "./CustomModal2.jsx";

const Nav = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const userName = Cookies.get("email");
    const user_id = Cookies.get("user_id");
    const token = Cookies.get("token");

    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [showAlert, setShowAlert] = useState(false);
    const [menuOpen, setMenuOpen] = useState(false);

    const openAlert = () => {
        setShowAlert(true);
    };

    const closeAlert = () => {
        Cookies.remove('user_id');
        Cookies.remove('type');
        Cookies.remove('email');
        Cookies.remove('token');
        Cookies.remove('city');
        Cookies.remove('startDate');
        Cookies.remove('endDate');
        navigate("/");
        window.location.reload();
    };

    const CancelUnlogin = () => {
        setShowAlert(false);
    };

    const login = () => {
        navigate("/login");
    };

    useEffect(() => {
        const token = Cookies.get('token');
        if (token) {
            setIsAuthenticated(true);
        }
    }, []);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await fetch(`http://localhost:8080/user/${user_id}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch user');
                }
                const userData = await response.json();
                const type_user = userData.userType;
                if (type_user) {
                    setIsAdmin(true);
                }
            } catch (error) {
                console.error('Error fetching user:', error);
            }
        };

        if (isAuthenticated) {
            fetchUser();
        }
    }, [isAuthenticated, user_id]);

    const shouldHideBackButton = location.pathname === '/' || location.pathname === '/home';

    const toggleMenu = () => {
        setMenuOpen(!menuOpen);
    };

    return (
        <div className={'externDiv'}>
            <div className="logoDiv">
                <Link to="/">
                    <h1>Hoteles.com</h1>
                </Link>
                <button className="menuButton" onClick={toggleMenu}>☰</button>
            </div>

            <div className={`linksDiv ${menuOpen ? 'open' : ''}`}>
                <div className="listDiv">
                    <ul className="listStyles">
                        {!shouldHideBackButton && (
                            <button id="backButton" onClick={() => navigate(-1)}>Atrás</button>
                        )}
                        <Link to="/hotels-list">
                            <li>Ver hoteles</li>
                        </Link>

                        {isAuthenticated ? (
                            <>
                                <button id="loginButton" onClick={openAlert}>
                                    {userName ? `Bienvenido, ${userName}` : "Bienvenido"}
                                </button>
                                {isAdmin && (
                                    <Link to="/dashAdmin/">
                                        <li>Panel de control</li>
                                    </Link>
                                )}
                            </>
                        ) : (
                            <button id="loginButton" onClick={login}>Iniciar sesión</button>
                        )}
                        <CustomModal2
                            showModal2={showAlert}
                            closeModal2={closeAlert}
                            closeModal22={CancelUnlogin}
                            content2="¿Cerrar sesión?"
                        />
                    </ul>
                </div>
            </div>
        </div>
    );
};

export default Nav;
