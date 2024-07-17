import React, {useEffect, useState} from "react";
import Cookies from 'js-cookie';
import './Components.css';
import {useNavigate} from "react-router-dom";

const Buttons = () => {
    const navigate = useNavigate();
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const user_id = Cookies.get('user_id');

    const ShowHotels = () => {
        navigate(`/hotels-list`);
    };

    const DashboardAdmin = () => {
        navigate(`/dashAdmin`);
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

    return (
        <div>
            <header id="header">
                {isAuthenticated ? (
                    <>
                        <button id="hotelButton" onClick={ShowHotels}>Ver hoteles</button>
                        {isAdmin && (
                            <button id="adminButton" onClick={DashboardAdmin}>Dashboard</button>
                        )}
                    </>
                ) : (
                    <p></p>
                )}
            </header>
        </div>
    );
};

export default Buttons;
