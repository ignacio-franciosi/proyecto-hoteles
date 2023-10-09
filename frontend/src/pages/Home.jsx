import React, { useEffect, useState } from 'react';
import {useNavigate} from "react-router-dom";
import './../App.css'

const Home = () => {
    const navigate = useNavigate(); //permite la navegación entre paginasd con las rutas
     //se inicializan las variables vacias
    const [city, setCity] = useState('');
    const [from, setFrom] = useState('');
    const [to, setTo] = useState('');
    const handleSubmit = async (e) => { //recibe los datos del formulario a

        e.preventDefault(); // para que no recarga la página

        try {   //envía la respuesta al back (postaman basicamente)
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST', headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({city, from, to}),
            }).then(response => {
                if (response.ok) {
                    return response.json();
                    navigate("/result")
                } else {
                    alert("Error, datos inválidos");
                }
            });
        } catch (error) {
            console.log('Error al realizar la solicitud al backend:', error);
        }
};

    return(
        <div id="body">
            <h1 id="h1Home">Encuentra tu hotel ;)</h1>
            <img src="../images/home/suiza.jpg" alt="imagen de suiza" className="imagen"/>
            <form id="formHome" onSubmit={handleSubmit} >
                <input id={"inputCity"}
                       type="text"
                       placeholder="Ciudad"
                       value={city}
                       onChange={(e) => setCity(e.target.value)}

                />
                <input id={"inputFrom"}
                       type="date"
                       placeholder="Fecha desde"
                       value={from}
                       onChange={(e) => setFrom(e.target.value)}
                />
                <input id={"inputUntil"}
                       type="date"
                       placeholder="Fecha hasta"
                       value={to}
                       onChange={(e) => setTo(e.target.value)}
                />

                <button id="SearchHome" type="submit">Buscar</button>
            </form>
        </div>
    );
};

export default Home;