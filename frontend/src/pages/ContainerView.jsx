import React, {useEffect} from 'react';
import ContainerList from '../components/DockerComponents/ContainerList.jsx';
import ContainerActions from '../components/DockerComponents/ContainerActions.jsx';
import {useNavigate} from "react-router-dom";
import Cookies from "js-cookie";


const ContainerView = () => {
    const token = Cookies.get("token")
    const user_id = Cookies.get("user_id")
    const navigate = useNavigate();

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
        <div id="backHotelSearch">
            <div className={"hotelContainer"}>
                <h1 id="h1Container">Infrastructure Management</h1>
                <br/>
                <div className="hotelCard">

                    <ContainerActions/>

                </div>
                <div className={"hotelCard"}>

                    <ContainerList/>

                </div>
            </div>
        </div>
    )
}
export default ContainerView