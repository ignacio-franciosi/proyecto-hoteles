import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';
import './../App.css';

const MyBookings = () => {
    const [bookings, setBookings] = useState([]);
    const user_id = Number(Cookies.get('user_id'));

    useEffect(() => {
        const fetchBookings = async () => {
            try {
                const response = await fetch(`http://localhost:8080/bookings/${user_id}`); //revisar
                if (!response.ok) {
                    throw new Error(`Error al obtener la lista de reservas: ${response.status}`);
                }
                const data = await response.json();
                setBookings(data);
            } catch (error) {
                console.error('Error al obtener la lista de reservas:', error);
            }
        };

        fetchBookings();
    }, [user_id]);

    const formatDate = (dateString) => {
        const options = { year: 'numeric', month: 'numeric', day: 'numeric', timeZone: 'UTC' };
        return new Date(dateString).toLocaleDateString(undefined, options);
    };

    return (
        <div className="my-bookings-container">
            {bookings && bookings.length > 0 ? (
                <div className= "bookingContainer">
                    {bookings.map((booking) => (
                        <div key={booking.id} className="booking-card">
                            <h2 className="booking-header">Reserva #{booking.id}</h2>
                            <p className="booking-info">ID del hotel: {booking.id_hotel}</p>
                            <p className="booking-info">ID del usuario: {booking.id_user}</p>
                            <p className="booking-info">Desde: {formatDate(booking.fecha_desde)}</p>
                            <p className="booking-info">Hasta: {formatDate(booking.fecha_hasta)}</p>
                            <p className="booking-info">Precio total: ${booking.precio_total}</p>
                        </div>
                    ))}
                    <div style={{ height: '100px' }}></div>

                </div>
            ) : (
                <div className="no-bookings">
                    <p>No se encontraron reservas.</p>
                </div>
            )}
        </div>
    );
};

export default MyBookings;