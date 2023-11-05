import React, {useState} from 'react';
import {useNavigate} from "react-router-dom";
import './../App.css'

const Home = () => {
    const navigate = useNavigate(); //permite la navegación entre paginasd con las rutas
    //se inicializan las variables vacias
    const [city, setCity] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');
    const handleSubmit = async (e) => { //recibe los datos del formulario a
        console.log(startDate);
        e.preventDefault(); // para que no recarga la página
        if (city === '') {
            document.getElementById("inputCity").style.borderColor = 'red';
        } else {
            document.getElementById("inputCity").style.borderColor = ''; // Restablecer el borde
        }

        if (startDate === '') {
            document.getElementById("inputFrom").style.borderColor = 'red';
        } else {
            document.getElementById("inputFrom").style.borderColor = ''; // Restablecer el borde
        }

        if (endDate === '') {
            document.getElementById("inputTo").style.borderColor = 'red';
        } else {
            document.getElementById("inputTo").style.borderColor = ''; // Restablecer el borde
        }

        // Validación de fecha
        const startDate1 = new Date(startDate);
        const endDate1 = new Date(endDate);

        if (startDate1 > endDate1) {
            document.getElementById("inputFrom").style.borderColor = 'red';
            document.getElementById("inputTo").style.borderColor = 'red';
            alert('La fecha "Fecha desde" no puede ser mayor que "Fecha hasta"');
        } else {
            try {   //envía la respuesta al back (postaman basicamente)
                const response = await fetch('http://localhost:8080/hotels/', {
                    method: 'POST', headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        city: city,
                        startDate: startDate,
                        endDate: endDate
                    }),

                }).then(response => {
                    if (response.ok) {
                        navigate("/search")
                        return response.json();
                    } else {
                        alert("Error, datos inválidos");
                    }
                });
                localStorage.setItem('city', city);
                localStorage.setItem('startDate', startDate);
                localStorage.setItem('endDate', endDate);
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
                       value={startDate}
                       onChange={(e) => setStartDate(e.target.value)}
                />
                <p>Fecha hasta: </p>
                <input id={"inputTo"}
                       type="date"
                       placeholder="Fecha hasta"
                       value={endDate}
                       onChange={(e) => setEndDate(e.target.value)}
                />
                <p></p>
                <button id="ButonSearchHome" type="submit">Buscar</button>
            </form>
        </div>
    );
};

export default Home;