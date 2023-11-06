import React, { useEffect, useState } from 'react';
import './../App.css'
import {useNavigate} from "react-router-dom";

const Search = () => {
    const [hoteles, setHoteles] = useState([]);
    const navigate = useNavigate();
    const selectHotels = (hotel_id) => {
        navigate(`/hotelDetails/${hotel_id}`);
    };
    const city = localStorage.getItem("city");

    useEffect(key => {
        // Realizar la solicitud al backend para obtener la lista de hoteles
        const fetchHoteles = async () => {
            try {
                const response = await fetch('http://localhost:8090/hotels');
                const data = await response.json();
                setHoteles(data);
            } catch (error) {
                console.log('Error al obtener la lista de hoteles:', error);
            }
        };

        fetchHoteles();
    }, []);

    return (

        <div id="backHotelSearch">
            {hoteles.length > 0 ? (
                <div>
                    {hoteles.map((hotel) => (
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
                <div id="noHotels">
                    <p >No se encontraron hoteles.</p>
                </div>
            )}
        </div>

    );
};

export default Search;