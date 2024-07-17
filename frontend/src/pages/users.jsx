import React, { useEffect, useState } from 'react';
import "./../App.css";

const Users = () => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await fetch(`http://localhost:8080/user/${id}`);
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                const data = await response.json();
                setUsers(Array.isArray(data) ? data : []);
            } catch (error) {
                console.log('Error al obtener la lista de usuarios:', error);
                setUsers([]);
            }
        };

        fetchUsers();
    }, []);

    return (
        <div id="backHotelSearch">
            {users.length > 0 ? (
                <div className="hotelContainer">
                    {users.map((user) => (
                        <div key={user.user_id} className="hotelCard">
                            <div>
                                <h2 id="h2HotelSearch">{user.name}</h2>
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

export default Users;
