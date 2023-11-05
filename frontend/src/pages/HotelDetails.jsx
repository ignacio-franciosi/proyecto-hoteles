import React, { useEffect, useState } from 'react';

import Header from '../components/Header.jsx';
import Footer from "../components/Footer.jsx";
import { useNavigate, useParams } from 'react-router-dom';

const HotelDetails = () => {
    const { hotel_id } = useParams();
    const [hotel, setHotel] = useState(null);
    const [startDate, setStartDate] = useState(null);
    const [endDate, setEndDate] = useState(null);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const back = () => {
        navigate(-1);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const inputValue = e.target[0].value;
        const splitValues = inputValue.split(' -');

        if (splitValues.length === 2) {
            try {
                const startDateString = splitValues[0].trim().replace(/\//g, '-');
                const endDateString = splitValues[1].trim().replace(/\//g, '-');
                const tempStartDate = startDateString + 'T00:00:00-03:00';
                const tempEndDate = endDateString + 'T00:00:00-03:00';

                const response = await fetch(`http://localhost:8090/hotels/${hotel_id}`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        fecha_desde: tempStartDate,
                        fecha_hasta: tempEndDate,
                        id_hotel: hotel_id, // Usar hotel_id en lugar de hotel
                        id_user: Number(localStorage.getItem('user_id')),
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
        } else {
            setErrorMessage('Formato de fechas incorrecto. Debe ser "Fecha Inicio - Fecha Fin".');
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
                    <div style={{ alignItems: 'left', maxWidth: '50%' }}>
                        <h1 style={{ textAlign: 'left', color: '#0E8388' }}>Hotel {hotel.name}</h1>
                        <p style={{ textAlign: 'left', color: '#2C3333' }}>Estrellas: {hotel.stars}</p>
                        <p style={{ textAlign: 'left', color: '#2C3333', maxWidth: '80%' }}>Descripción: {hotel.description}</p>
                        <p style={{ textAlign: 'left', color: '#2C3333' }}>Precio por noche: ${hotel.price}</p>
                        <p style={{ textAlign: 'left', color: '#2C3333' }}>Amenities: {hotel.amenities}</p>
                    </div>
                    <form onSubmit={handleSubmit}>
                        <input type="text" placeholder="Fecha Inicio - Fecha Fin" />
                        <button type="submit" style={{ textAlign: 'right', backgroundColor: '#2E4F4F', marginLeft: '15px' }}>
                            Reservar
                        </button>
                    </form>
                </div>
            ) : (
                <p>No se encontró el hotel</p>
            )}
            <button id="botonAtras" onClick={back}>
                Atrás
            </button>
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </div>
    );
};

export default HotelDetails;
