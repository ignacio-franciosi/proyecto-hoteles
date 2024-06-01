import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';
import { useNavigate } from 'react-router-dom';
import './../App.css';

const HotelsList = () => {
    const [hotels, setHotels] = useState([]);
    const navigate = useNavigate();
    const selectHotels = (hotel_id) => {
        Cookies.set("hotel_id", hotel_id);
        navigate(`/details/${hotel_id}`);
    };


    useEffect( () => {
        const fetchHotels = async () => {
            try {
                const response = await fetch('http://localhost:8090/hotels');
                const data = await response.json();
                setHotels(data);
            } catch (error) {
                console.log('Error al obtener la lista de hoteles:', error);
            }
        };

        fetchHotels();
    }, []);

    return (
        <div id="backHotelSearch">
            {hotels.length > 0 ? (
                <div className="hotelContainer">
                    {hotels.map((hotel) => (
                        <div id="hotelSearch">
                            <img id="imgSearch" src={hotel.photos}/>
                            <div>
                                <h2 id="h2HotelSearch">{hotel.name}</h2>
                                <p id="paragraphSearch">Estrellas: {hotel.stars}</p>
                                <p id="paragraphSearch">Precio por noche: ${hotel.price}</p>
                                <p id="paragraphSearch">Ciudad: {hotel.city}</p>
                            </div>
                            <button id="butonSearch" type="submit" onClick={() => selectHotels(hotel.hotel_id)}>Ver</button>

                        </div>
                    ))}
                </div>
            ) : (
                <div className={"noHotels"}>
                    <h2>No se encontraron hoteles.</h2>
                </div>
            )}
        </div>
    );
};

export default HotelsList;
