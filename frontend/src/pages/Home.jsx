import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './../App.css';
import Cookies from 'js-cookie';

const Home = () => {
    const navigate = useNavigate();
    const [city, setCity] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();

        const startDate1 = new Date(startDate);
        const endDate1 = new Date(endDate);

        if (startDate1 > endDate1) {
            document.getElementById("inputFrom").style.borderColor = 'red';
            document.getElementById("inputTo").style.borderColor = 'red';
            alert('La fecha de ingreso no puede ser posterior a la fecha de salida');
        } else {
            Cookies.set('city', city);
            Cookies.set('startDate', startDate);
            Cookies.set('endDate', endDate);
            navigate("/hotels-list");
        }
    };

    return (
        <div id="bodyHome" className="row">
            <form className={"formHome"} onSubmit={handleSubmit}>
                <div className={"TitleForm"}>
                    <h2>Elija cual será su destino soñado</h2>
                </div>
                <div className={"TitleForm1"}>
                    <h2>Seleccione Ciudad,<br />Fecha desde y Fecha Hasta</h2>
                </div>
                <input
                    id={"inputCity"}
                    type="text"
                    placeholder="Ciudad"
                    value={city}
                    onChange={(e) => setCity(e.target.value)}
                />
                <input
                    id={"inputFrom"}
                    type="date"
                    placeholder="Fecha desde"
                    value={startDate}
                    onChange={(e) => setStartDate(e.target.value)}
                />
                <input
                    id={"inputTo"}
                    type="date"
                    placeholder="Fecha hasta"
                    value={endDate}
                    onChange={(e) => setEndDate(e.target.value)}
                />
                <button id="ButonSearchHome" type="submit">Buscar</button>
            </form>
        </div>
    );
};

export default Home;
