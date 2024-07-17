import React, {useEffect, useState} from 'react';
import "../App.css"
import {useNavigate, useParams} from 'react-router-dom';
import Cookies from "js-cookie";

const HotelDetails = () => {
    const {hotel_id} = useParams();
    const [hotel, setHotel] = useState(null);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();
    const startDate1 = Cookies.get("startDate");
    const endDate1 = Cookies.get("endDate");
    const user_id = Number(Cookies.get('user_id'));
    const tokenUser = Cookies.get("token");

    useEffect(() => {
        const fetchHotel = async () => {
            try {
                const response = await fetch(`http://localhost:8000/hotel/${hotel_id}`);
                const data = await response.json();
                setHotel(data);
            } catch (error) {
                console.log('Error al obtener el hotel:', error);
            }
        };

        fetchHotel();
    }, [hotel_id]);



    const handleSubmit = async (e) => {
        e.preventDefault();
        const inputValue = e.target[0].value;
        const splitValues = inputValue.split(' - ');

        if (user_id === -1 || tokenUser === null) {
            alert("Debes iniciar sesión para poder reservar!");
        } else {
            try {

                const response = await fetch(`http://localhost:8080/booking`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        startDate: startDate1,
                        endDate: endDate1,
                        idMongo: hotel_id,
                        idUser: user_id
                    }),
                });

                if (response.ok) {
                    alert("Su reserva ha sido confirmada");
                    navigate("/");

                } else {
                    alert("No hay habitaciones disponibles");
                }
                if (startDate1 === null || endDate1 === null) {
                    alert("no ha completado los días de reserva!")
                    navigate("/")
                }

            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }

        }

    };

    return (
        <div id="backHotelDetails">
            {hotel ? (
                <div id="hotelDetails">
                    <h1 id="h1HotelDetails">Hotel {hotel.name}</h1>
                    <p id="paragraphDetails">Estrellas: {hotel.stars}</p>
                    <p id="paragraphDetails">Descripción: {hotel.description}</p>
                    <p id="paragraphDetails">Precio por noche: ${hotel.price}</p>
                    <p id="paragraphDetails">Amenities: {hotel.amenities}</p>
                    <form onSubmit={handleSubmit}>
                        <h3 id="confirmacion">Usted está por reservar una habitación del hotel "{hotel.name}" en la
                            ciudad de {hotel.city} desde el día {startDate1} hasta el día {endDate1}</h3>
                        <button id="butonDetails" type="submit">Reservar</button>
                    </form>
                </div>
            ) : (
                <div className={"noHotels"}>
                    <h2>No se encontraron hoteles.</h2>
                </div>
            )}
            {errorMessage && <p style={{color: 'red'}}>{errorMessage}</p>}
        </div>
    );
};

export default HotelDetails;