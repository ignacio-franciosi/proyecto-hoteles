import React, { useEffect, useState } from 'react';
import "./../App.css";

const EditHotels = () => {
    const [hotels, setHotels] = useState([]);

    const deleteHotel = async (hotel_id) => {
        try{
            const response = await fetch(`http://localhost:8080/hotel/${hotel_id}`, {
                method: 'DELETE',
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            // Actualizar la lista de hoteles después de eliminar uno
            setHotels(hotels.filter(hotel => hotel.hotel_id !== hotel_id));
        } catch (error) {
            console.log('Error al eliminar el hotel:(uba)', error);
        }
        try {
            const response = await fetch(`http://localhost:8090/hotel/${hotel_id}`, {
                method: 'DELETE',
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            // Actualizar la lista de hoteles después de eliminar uno
            setHotels(hotels.filter(hotel => hotel.hotel_id !== hotel_id));
        } catch (error) {
            console.log('Error al eliminar el hotel:(hotels)', error);
        }
    };

    useEffect(() => {
        const fetchHotels = async () => {
            try {
                const response = await fetch('http://localhost:8000/hotel');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                const data = await response.json();
                setHotels(Array.isArray(data) ? data : []);
            } catch (error) {
                console.log('Error al obtener la lista de hoteles:', error);
                setHotels([]);
            }
        };

        fetchHotels();
    }, []);

    return (
        <div id="backHotelSearch">
            {hotels.length > 0 ? (
                <div className="hotelContainer">
                    {hotels.map((hotel) => (
                        <div key={hotel.hotel_id} className="hotelCard">
                            <img id="imgSearch" src={hotel.photos} alt={hotel.name} />
                            <div>
                                <h2 id="h2HotelSearch">{hotel.name}</h2>
                                <h2 id="h2HotelSearch">{hotel.hotel_id}</h2>
                                <p id="paragraphSearch">Ciudad: {hotel.city}</p>
                            </div>
                            <button
                                id="butonSearch"
                                type="submit"
                                onClick={() => deleteHotel(hotel.hotel_id)}
                            >
                                Eliminar
                            </button>
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

export default EditHotels;
