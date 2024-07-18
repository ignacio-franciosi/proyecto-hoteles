import React, { useEffect, useState } from 'react';
import "./../App.css";
import {useNavigate} from "react-router-dom";

const EditHotels = () => {
    const [hotels, setHotels] = useState([]);
    const [selectedHotel, setSelectedHotel] = useState(null);
    const [showForm, setShowForm] = useState(false);

    const updateHotel = async (hotel) => {
        const endpoints = [
            `http://localhost:8080/hotel/${hotel.hotel_id}`,
            `http://localhost:8090/hotel/${hotel.hotel_id}`
        ];

        for (let endpoint of endpoints) {
            try {
                const response = await fetch(endpoint, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(hotel)
                });

                if (!response.ok) {
                    throw new Error(`[${endpoint.split(':')[1]}]HTTP error! status: ${response.status}`);
                }

                const updatedHotels = hotels.map(h => h.hotel_id === hotel.hotel_id ? hotel : h);
                setHotels(updatedHotels);
                setShowForm(false);
            } catch (error) {
                console.log(`[${endpoint.split(':')[1]}]Error al actualizar el hotel:`, error);
            }
        }
    };

    const deleteHotel = async (hotel_id) => {
        const endpoints = [
            `http://localhost:8080/hotel/${hotel_id}`,
            `http://localhost:8090/hotel/${hotel_id}`
        ];

        for (let endpoint of endpoints) {
            try {
                const response = await fetch(endpoint, {
                    method: 'DELETE',
                });

                if (!response.ok) {
                    throw new Error(`[${endpoint.split(':')[1]}]HTTP error! status: ${response.status}`);
                }

                setHotels(hotels.filter(hotel => hotel.hotel_id !== hotel_id));
            } catch (error) {
                console.log(`[${endpoint.split(':')[1]}]Error al eliminar el hotel:`, error);
            }
        }
    };

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

    useEffect(() => {
        fetchHotels();
    }, []);

    const handleEditClick = (hotel) => {
        setSelectedHotel(hotel);
        setShowForm(true);
    };

    const handleFormChange = (e) => {
        const { name, value } = e.target;
        setSelectedHotel({ ...selectedHotel, [name]: name === 'rooms' || name === 'stars' || name === 'price' ? Number(value) : value });
    };

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
                                onClick={() => handleEditClick(hotel)}
                            >
                                Editar
                            </button>
                            <br/><br/>
                            <button
                                id="butonSearch"
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
            {showForm && selectedHotel && (
                <div className="popup">
                    <div className="popup-content">
                        <span className="close" onClick={() => setShowForm(false)}>&times;</span>
                        <form onSubmit={(e) => {
                            e.preventDefault();
                            updateHotel(selectedHotel);
                        }}>
                            <label>
                                Nombre:
                                <input
                                    type="text"
                                    name="name"
                                    value={selectedHotel.name}
                                    onChange={handleFormChange}
                                />
                            </label>
                            <br/>
                            <label>
                                Habitaciones:
                                <input
                                    type="number"
                                    name="rooms"
                                    value={selectedHotel.rooms}
                                    onChange={handleFormChange}
                                />
                            </label>
                            <br/>
                            <label>
                                Descripción:
                                <input
                                    type="text"
                                    name="description"
                                    value={selectedHotel.description}
                                    onChange={handleFormChange}
                                />
                            </label>
                            <br/>
                            <label>
                                Ciudad:
                                <input
                                    type="text"
                                    name="city"
                                    value={selectedHotel.city}
                                    onChange={handleFormChange}
                                />
                            </label>
                            <br/>
                            <label>
                                Valoración:
                                <input
                                    type="number"
                                    name="stars"
                                    value={selectedHotel.stars}
                                    onChange={handleFormChange}
                                />
                            </label>
                            <br/>
                            <label>
                                Precio:
                                <input
                                    type="number"
                                    name="price"
                                    value={selectedHotel.price}
                                    onChange={handleFormChange}
                                />
                            </label>
                            <br/>
                            <label>
                                Servicios:
                                <input
                                    type="text"
                                    name="amenities"
                                    value={selectedHotel.amenities}
                                    onChange={handleFormChange}
                                />
                            </label>
                            <br/>
                            <button type="submit">Actualizar</button>
                        </form>
                    </div>
                </div>
            )}
        </div>
    );
};

export default EditHotels;
