import React, {useState} from 'react';
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
        if (city === '') {
            document.getElementById("inputCity").style.borderColor = 'red';
        } else {
            document.getElementById("inputCity").style.borderColor = ''; // Restablecer el borde
        }

        if (from === '') {
            document.getElementById("inputFrom").style.borderColor = 'red';
        } else {
            document.getElementById("inputFrom").style.borderColor = ''; // Restablecer el borde
        }

        if (to === '') {
            document.getElementById("inputTo").style.borderColor = 'red';
        } else {
            document.getElementById("inputTo").style.borderColor = ''; // Restablecer el borde
        }

        // Validación de fecha
        const fromDate = new Date(from);
        const toDate = new Date(to);

        if (fromDate > toDate) {
            document.getElementById("inputFrom").style.borderColor = 'red';
            document.getElementById("inputTo").style.borderColor = 'red';
            alert('La fecha "Fecha desde" no puede ser mayor que "Fecha hasta"');
        } else {
            try {   //envía la respuesta al back (postaman basicamente)
                const response = await fetch('http://localhost:8080/search', {
                    method: 'POST', headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({city, from, to}),
                }).then(response => {
                    if (response.ok) {
                        navigate("/search")
                        return response.json();
                    } else {
                        alert("Error, datos inválidos");
                    }
                });
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        }

    };

    return (
        <div id="body">
            <h1 id="h1Home">Encuentra tu hotel ;)</h1>
            <form id="formHome" onSubmit={handleSubmit}>
                <p>Ciudad</p>
                <input id={"inputCity"}
                       type="text"
                       placeholder="Ciudad"
                       value={city}
                       onChange={(e) => setCity(e.target.value)}

                />
                <p>Fecha desde: </p>
                <input id={"inputFrom"}
                       type="date"
                       placeholder="Fecha desde"
                       value={from}
                       onChange={(e) => setFrom(e.target.value)}
                />
                <p>Fecha hasta: </p>
                <input id={"inputTo"}
                       type="date"
                       placeholder="Fecha hasta"
                       value={to}
                       onChange={(e) => setTo(e.target.value)}
                />
                <p></p>
                <button id="ButonSearchHome" type="submit">Buscar</button>
            </form>
        </div>
    );
};

export default Home;