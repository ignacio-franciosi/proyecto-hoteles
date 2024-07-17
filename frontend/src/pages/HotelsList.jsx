import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';
import { useNavigate } from 'react-router-dom';
import './../App.css';

const HotelsList = () => {
    const [hotels, setHotels] = useState([]);
    const navigate = useNavigate();

    const selectHotels = (hotel_id) => {
        Cookies.set("hotel_id", hotel_id);
        navigate(`/hotels-list/${hotel_id}`);
    };

    const startDate = Cookies.get("startDate") || '';
    const endDate = Cookies.get("endDate") || '';
    const city = Cookies.get("city") || '';

    useEffect(() => {
        const fetchHotels = async () => {
            try {
                let response;
                if (startDate === '' && endDate === '' && city === '') {
                    response = await fetch('http://localhost:8000/hotel');
                } else if (startDate === '' && endDate === '') {
                    response = await fetch(`http://localhost:8000/hotel?city=${city}`);
                } else {
                    response = await fetch(`http://localhost:8080/available?city=${city}&startDate=${startDate}&endDate=${endDate}`);
                }

                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                const data = await response.json();
                console.log('Received data:', data); // Log para verificar los datos recibidos
                setHotels(Array.isArray(data) ? data : []);
            } catch (error) {
                console.log('Error al obtener la lista de hoteles:', error);
                setHotels([]);
            }
        };

        fetchHotels();
    }, [startDate, endDate, city]);

    return (
        <div id="backHotelSearch">
            {hotels.length > 0 ? (
                <div className="hotelContainer">
                    {hotels.map((hotel) => (
                        <div key={hotel.hotel_id} className="hotelCard">
                            <img id="imgSearch" src={hotel.photos} alt={hotel.name}/>
                            <div>
                                <h2 id="h2HotelSearch">{hotel.name}</h2>
                                <h2 id="h2HotelSearch">{hotel.hotel_id}</h2>
                                <p id="paragraphSearch">Estrellas: {hotel.stars}</p>
                                <p id="paragraphSearch">Precio por noche: ${hotel.price}</p>
                                <p id="paragraphSearch">Ciudad: {hotel.city}</p>
                            </div>
                            <button id="butonSearch" type="submit" onClick={() => selectHotels(hotel.hotel_id)}>Ver</button>
                        </div>
                    ))}
                </div>
            ) : (
                <div className="noHotels">
                    <h2>No se encontraron hoteles.</h2>
                </div>
            )}
        </div>
    );
};

export default HotelsList;
