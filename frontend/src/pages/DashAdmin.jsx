import React, {useEffect, useState} from 'react';
import {Link, useNavigate} from "react-router-dom";
import "../App.css"
import Cookies from "js-cookie";


const DashAdmin = () => {
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();
    const user_id = Cookies.get("user_id")
    const token = Cookies.get("token")

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await fetch(`http://localhost:8080/user/${user_id}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch admin');
                }
                const userData = await response.json();
                console.log(userData.name, userData.dni, userData.userType)
                if (!userData.userType) {
                    navigate("/")
                }

            } catch (error) {
                console.error('Error fetching user:', error);
                navigate("/");
            }

        };
        fetchUser();
    }, [user_id, token]);

    return (

        <div className="my-service-container">
            <div className="serviceContainer">
                <div className="service-card">
                    <div className="service-header">
                        <Link to="/dashAdmin/container-view" className={"textButton"}>
                            Ver contenedores
                        </Link>
                    </div>
                </div>
            </div>
            <div className="serviceContainer">
                <div className="service-card">
                    <div className="service-header">
                        <Link to="/dashAdmin/edit-hotels" className={"textButton"}>
                            Editar hoteles
                        </Link>
                    </div>
                </div>
            </div>
            <div className="serviceContainer">
                <div className="service-card">
                    <div className="service-header">
                        <Link to="/dashAdmin/users" className={"textButton"}>
                            Ver usuarios
                        </Link>
                    </div>
                </div>
            </div>
        </div>

    )
}

export default DashAdmin;