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

    // Función para obtener la ruta de la primera imagen del hotel
    const getFirstHotelImage = (hotel_id) => {
        return `/HotelsImages/${hotel_id}/1.jpg`;
    };

    useEffect(() => {
        const fetchHotels = async () => {
            try {
                let response;
                let data = [];
                
                if (startDate === '' && endDate === '' && city === '') {
                    response = await fetch('http://localhost:8000/hotel');

                    if (!response.ok) {
                        throw new Error(`HTTP error! status: ${response.status}`);
                    }
                    data = await response.json();

                } else if (startDate === '' && endDate === '') {
                    response = await fetch(`http://localhost:8000/hotel?city=${city}`);

                    if (!response.ok) {
                        throw new Error(`HTTP error! status: ${response.status}`);
                    }
                    data = await response.json();

                } else {
                    let sqlResponse;
                    sqlResponse = await fetch(`http://localhost:8080/available?city=${city}&startDate=${startDate}&endDate=${endDate}`);
                    if (!sqlResponse.ok) {
                        throw new Error(`HTTP error! status: ${response.status}`);
                    }

                    const sqlData = await sqlResponse.json();
                    
                    for (const hotel of sqlData) {
                        console.log('entro al for');
                        response = await fetch(`http://localhost:8000/hotel/${hotel.id_mongo}`);
                        if (!response.ok) {
                            throw new Error(`HTTP error! status: ${response.status}`);
                        }
                        const hotelData = await response.json();
                        data.push(hotelData);
                    }    
                }

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
                            <img
                                src={getFirstHotelImage(hotel.hotel_id)}
                                alt={`Hotel ${hotel.name}`}
                                style={{width: '200px', height: 'auto'}}
                                onError={(e) => {
                                    e.target.onerror = null;
                                    e.target.src = '/path/to/default/image.jpg';
                                }}
                            />
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
                <div className="noHotels">
                    <h2>No se encontraron hoteles.</h2>
                </div>
            )}
        </div>
    );
};

export default HotelsList;