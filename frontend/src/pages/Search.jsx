import React, { useEffect, useState } from 'react';
import './../App.css'
import {useNavigate} from "react-router-dom";

const Search = () => {
    const [hoteles, setHoteles] = useState([]);
    const navigate = useNavigate();
    const selectHotels = (hotel_id) => {
        navigate(`/hotelDetails/${hotel_id}`);
    };


    useEffect(key => {
        // Realizar la solicitud al backend para obtener la lista de hoteles a
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

        <div style={{ alignItems: 'left', backgroundColor: '#CBE4DE', minHeight: '100vh' }}>
            <h1 style={{ textAlign: 'center', color:'#0E8388'}}>Hoteles:</h1>
            {hoteles.length > 0 ? (
                <div>
                    {hoteles.map((hotel) => (
                        <div key={hotel.hotel_id} style={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-start', marginBottom: '20px' }}>
                            <img src={hotel.photos} style={{ width: '150px', height: '150px', marginRight: '10px', marginLeft:'30px' }} />
                            <div>
                                <h2 style={{color: '#0E8388' }}>{hotel.name}</h2>
                                <p style={{ color: '#2C3333' }}>Estrellas: {hotel.stars}</p>
                                <p style={{ color: '#2C3333', marginRight: 'auto' }}>Precio por noche: ${hotel.price}</p>
                                <p style={{ color: '#2C3333', marginRight: 'auto' }}>Ciudad: {hotel.city}</p>
                            </div>
                            <button style={{ marginRight: 'auto', backgroundColor:'#2E4F4F' }} type="submit" onClick={() => selectHotels(hotel.hotel_id)}>Ver</button>

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