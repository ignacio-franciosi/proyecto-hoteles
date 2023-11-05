import React, { useEffect, useState } from 'react';
import Header from '../components/Header.jsx';
import Footer from "../components/Footer.jsx";
import "../App.css"
import { useNavigate, useParams } from 'react-router-dom';

const HotelDetails = () => {
    const { hotel_id } = useParams();
    const [hotel, setHotel] = useState(null);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();
    const startDate1 = localStorage.getItem("startDate");
    const endDate1 = localStorage.getItem("endDate");
    const user_id = Number(localStorage.getItem('user_id'));

    const handleSubmit = async (e) => {
        e.preventDefault();
        const inputValue = e.target[0].value;
        const splitValues = inputValue.split(' - ');


            try {


                const response = await fetch(`http://localhost:8090/booking`, {
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
                    console.log('Se ha registrado su reserva');
                } else {
                    console.log('No hay habitaciones disponibles');
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }

    };

    useEffect(() => {
        const fetchHotel = async () => {
            try {
                const response = await fetch(`http://localhost:8090/hotels/${hotel_id}`);
                const data = await response.json();
                setHotel(data);
            } catch (error) {
                console.log('Error al obtener el hotel:', error);
            }
        };

        fetchHotel();
    }, [hotel_id]);

    return (
        <div style={{ alignItems: 'left', backgroundColor: '#CBE4DE', minHeight: '100vh' }}>
            {hotel ? (
                <div style={{ display: 'flex', alignItems: 'flex-start', justifyContent: 'flex-start', marginTop: '80px', marginBottom: '20px', marginLeft: '60px' }}>
                    <div style={{ alignItems: 'left', maxWidth: '100%' }}>
                        <h1 style={{ textAlign: 'left', color: '#0E8388' }}>Hotel {hotel.name}</h1>
                        <p style={{ textAlign: 'left', color: '#2C3333' }}>Estrellas: {hotel.stars}</p>
                        <p style={{ textAlign: 'left', color: '#2C3333', maxWidth: '80%' }}>Descripción: {hotel.description}</p>
                        <p style={{ textAlign: 'left', color: '#2C3333' }}>Precio por noche: ${hotel.price}</p>
                        <p style={{ textAlign: 'left', color: '#2C3333' }}>Amenities: {hotel.amenities}</p>
                        <form onSubmit={handleSubmit}>
                        <h3 id="confirmacion">Usted está por reservar una habitación del hotel "{hotel.name}" en la ciudad de {hotel.city} desde el día {startDate1} hasta el día {endDate1}</h3>
                            <button id="botonLogin" type="submit" style={{ textAlign: 'right', backgroundColor: '#2E4F4F', marginLeft: '15px' }}>Reservar</button>
                        </form>
                    </div>

                </div>
            ) : (
                <p>No se encontró el hotel</p>
            )}
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </div>
    );
};

export default HotelDetails;
