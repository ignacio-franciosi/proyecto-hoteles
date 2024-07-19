import React, {useEffect, useState} from 'react';
import "../App.css";
import {useNavigate, useParams} from 'react-router-dom';
import Cookies from "js-cookie";

const HotelDetails = () => {
    const {hotel_id} = useParams();
    const [hotel, setHotel] = useState(null);
    const [errorMessage, setErrorMessage] = useState('');
    const [currentImageIndex, setCurrentImageIndex] = useState(0);
    const [images, setImages] = useState([]);
    const navigate = useNavigate();
    const startDate = Cookies.get("startDate");
    const endDate = Cookies.get("endDate");
    const user_id = Number(Cookies.get('user_id'));
    const tokenUser = Cookies.get("token");

    const convertDateFormat = (dateString) => {
        const [year, month, day] = dateString.split('-');
        return `${day}-${month}-${year}`;
    };

    const formattedStartDate = convertDateFormat(startDate);
    const formattedEndDate = convertDateFormat(endDate);

    useEffect(() => {
        const fetchHotel = async () => {
            try {
                const response = await fetch(`http://localhost:8000/hotel/${hotel_id}`);
                const data = await response.json();
                setHotel(data);
                // Aquí asumimos que hay 5 imágenes por hotel, ajusta según sea necesario
                setImages(Array.from({length: 5}, (_, i) => `/HotelsImages/${hotel_id}/${i + 1}.jpg`));
            } catch (error) {
                console.log('Error al obtener el hotel:', error);
            }
        };

        fetchHotel();
    }, [hotel_id]);

    const handleNextImage = () => {
        setCurrentImageIndex((prevIndex) =>
            prevIndex === images.length - 1 ? 0 : prevIndex + 1
        );
    };

    const handlePrevImage = () => {
        setCurrentImageIndex((prevIndex) =>
            prevIndex === 0 ? images.length - 1 : prevIndex - 1
        );
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (user_id === -1 || tokenUser === null) {
            alert("Debes iniciar sesión para poder reservar!");
            navigate("/login")

        }
        if (startDate === '' || endDate === ''){
            alert('completa las fechas');
            navigate("/");
        }
        if (formattedStartDate > formattedEndDate) {
            alert('La "Fecha desde" no puede ser mayor que "Fecha hasta".');
            navigate("/");
        }
        else {
            try {
                const response = await fetch(`http://localhost:8080/booking`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        startDate: formattedStartDate,
                        endDate: formattedEndDate,
                        idMongo: hotel_id,
                        idUser: user_id
                    }),
                });

                if (response.ok) {
                    alert("Su reserva ha sido confirmada");
                    navigate("/");
                } else {
                    alert("No hay habitaciones disponibles");
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        }
    };

    return (
        <div id="backHotelDetails">
            {hotel ? (
                <div id="hotelDetails">
                    <h1 id="h1HotelDetails">Hotel {hotel.name}</h1>
                    <div className="image-carousel">
                        <img
                            src={images[currentImageIndex]}
                            alt={`Hotel ${hotel.name}`}
                            style={{width: '300px', height: 'auto'}}
                            onError={(e) => {
                                e.target.onerror = null;
                                e.target.src = '/path/to/default/image.jpg';
                            }}
                        />
                        <div>
                            <button className={"buttonModal"} onClick={handlePrevImage}>Anterior</button>
                            <button className={"buttonModal"} onClick={handleNextImage}>Siguiente</button>
                        </div>
                    </div>
                    <p id="paragraphDetails">Estrellas: {hotel.stars}</p>
                    <p id="paragraphDetails">Descripción: {hotel.description}</p>
                    <p id="paragraphDetails">Precio por noche: ${hotel.price}</p>
                    <p id="paragraphDetails">Amenities: {hotel.amenities}</p>
                    <form onSubmit={handleSubmit}>
                        <h3 id="confirmacion">Usted está por reservar una habitación del hotel "{hotel.name}" en la
                            ciudad de {hotel.city} desde el día {formattedStartDate} hasta el día {formattedEndDate}</h3>
                        <button id="butonDetails" type="submit">Reservar</button>
                    </form>
                </div>
            ) : (
                <div className={"noHotels"}>
                    <h2>No se encontró el hotel!!!.</h2>
                </div>
            )}
            {errorMessage && <p style={{color: 'red'}}>{errorMessage}</p>}
        </div>
    );
};

export default HotelDetails;
