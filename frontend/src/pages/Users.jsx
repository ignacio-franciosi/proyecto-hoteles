import React, { useEffect, useState } from 'react';
import "./../App.css";

const Users = () => {
    const [users, setUsers] = useState([]);
    const [error, setError] = useState(null);
    let type = 0;
    useEffect(() => {
        const fetchUsers = async () => {
            let id = 1;
            let usersList = [];
            let keepFetching = true;

            while (keepFetching) {
                try {
                    const response = await fetch(`http://localhost:8080/user/${id}`);
                    if (!response.ok) {
                        if (response.status === 404) {
                            keepFetching = false;
                            break;
                        } else {
                            throw new Error(`HTTP error! status: ${response.status}`);
                        }
                    }

                    const data = await response.json();
                    type = data.userType;
                    usersList.push(data);
                    id++;
                } catch (error) {
                    setError('Error al obtener la lista de usuarios: ' + error.message);
                    keepFetching = false;
                }
            }

            setUsers(usersList);
        };

        fetchUsers();
    }, []);

    return (
        <div id="backHotelSearch">
            {error && <div className="error">{error}</div>}
            {users.length > 0 ? (
                <div className="hotelContainer">
                    {users.map((user) => (
                        <div key={user.id_user} className="hotelCard">
                            <div>
                                <h2 id="h2HotelSearch">Id: {user.id_user}</h2>
                                <h2 id="h2HotelSearch">Nombre: {user.name}</h2>
                                <h2 id="h2HotelSearch">Apellido: {user.lastName}</h2>
                                <h2 id="h2HotelSearch">DNI: {user.dni}</h2>
                                <h2 id="h2HotelSearch">Tipo: {user.userType}</h2>
                            </div>
                        </div>
                    ))}
                </div>
            ) : (
                <div className="noHotels">
                    <h2>No se encontraron usuarios!</h2>
                </div>
            )}
        </div>
    );
};

export default Users;
