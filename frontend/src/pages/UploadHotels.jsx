import React, {useState} from 'react';
import Cookies from 'js-cookie';
import './../App.css';
import {useNavigate} from "react-router-dom";
import CustomModal from "../components/CustomModal.jsx";
import CustomModal2 from "../components/CustomModal2.jsx";

const UploadHotel = () => {
    const navigate = useNavigate(); // Permite la navegación entre páginas con las rutas
    const [hotelName, setHotelName] = useState('');
    const [description, setDescription] = useState('');
    const [rooms, setRooms] = useState('');
    const [stars, setStars] = useState('');
    const [price, setPrice] = useState('');
    const [amenities, setAmenities] = useState('');
    let emptyInput = false; // Utilizamos useState para emptyInput

    const [showAlert1, setShowAlert1] = useState(false);
    const [showAlertImage, setShowAlertImage] = useState(false);

    const openAlert1 = () => {
        setShowAlert1(true);
    };
    const closeAlert1 = () => {
        setShowAlert1(false);
    };
    const openAlertImage = () => {
        setShowAlertImage(true);
    };
    const noUploadImage = () => {
        navigate("/imageLinks");
    };
    const uploadImage = () => {
        navigate('/home')
    };


    const handleSubmit = async (e) => {
        e.preventDefault(); // Para que no recargue la página
        emptyInput = false; // Restablecemos el estado de emptyInput a false al inicio de cada validación
        if (hotelName === '') {
            document.getElementById('inputHotelName').style.borderColor = 'red';
            emptyInput = true; // Cambiamos el estado a true si hay un campo vacío
        } else {
            document.getElementById('inputHotelName').style.borderColor = '';
        }
        if (description === '') {
            document.getElementById('inputDescription').style.borderColor = 'red';
            emptyInput = true;
        } else {
            document.getElementById('inputDescription').style.borderColor = '';
        }
        if (rooms === '' || rooms <= 0) {
            document.getElementById('inputRooms').style.borderColor = 'red';
            emptyInput = true;
        } else {
            document.getElementById('inputRooms').style.borderColor = '';
        }
        if (stars === '' || stars <= 0 || stars > 5) {
            document.getElementById('inputStars').style.borderColor = 'red';
            emptyInput = true;
        } else {
            document.getElementById('inputStars').style.borderColor = '';
        }
        if (price === '' || price <= 0) {
            document.getElementById('inputPrice').style.borderColor = 'red';
            emptyInput = true;
        } else {
            document.getElementById('inputPrice').style.borderColor = '';
        }
        if (amenities === '') {
            document.getElementById('inputAmenities').style.borderColor = 'red';
            emptyInput = true;
            console.log(emptyInput)
        } else {
            document.getElementById('inputAmenities').style.borderColor = '';
        }


        if (!emptyInput) {
            try {
                // Envía la respuesta al backend (Postman, básicamente)
                const response = await fetch('http://localhost:8080/hotel', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        nombre: hotelName,
                        descripcion: description,
                        cant_habitaciones: parseInt(rooms),
                        valoracion: parseInt(stars),
                        precio: parseFloat(price),
                        amenities: amenities,
                    }),
                });

                if (response.ok) {
                    const data = await response.json();
                    Cookies.set('hotel_id', data.id);
                    openAlertImage()
                } else {
                    alert('Ocurrió un error al crear el hotel');
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        } else {
            openAlert1();

        }
    };

    return (
        <div id="body">
            <CustomModal
                showModal={showAlert1}
                closeModal={closeAlert1}
                content="Completa todos los campos"
            />
            <CustomModal2
                showModal2={showAlertImage}
                closeModal2={noUploadImage}
                closeModal22={uploadImage}
                content2="Hotel cargado correctamente, ¿desea cargar imágenes?"
            />
            <h1 id="h1Login">Estás creando un nuevo hotel ;)</h1>
            <form id="formLogin" onSubmit={handleSubmit}>
                <input
                    id={'inputHotelName'}
                    type='text'
                    placeholder='Nombre del Hotel'
                    value={hotelName}
                    onChange={(e) => setHotelName(e.target.value)}
                />
                <input
                    id={'inputDescription'}
                    type='text'
                    placeholder='Descripción del hotel'
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                />
                <input
                    id={'inputRooms'}
                    type='number'
                    placeholder='Cantidad de habitaciones'
                    value={rooms}
                    onChange={(e) => setRooms(e.target.value)}
                />
                <input
                    id={'inputStars'}
                    type='number'
                    placeholder='¿De cuantas estrellas es el hotel?'
                    value={stars}
                    onChange={(e) => setStars(e.target.value)}
                />
                <input
                    id={'inputPrice'}
                    type='number'
                    placeholder='Precio por noche'
                    value={price}
                    onChange={(e) => setPrice(e.target.value)}
                />
                <input
                    id={'inputAmenities'}
                    type='text'
                    placeholder='Amenities separadas por coma'
                    value={amenities}
                    onChange={(e) => setAmenities(e.target.value)}
                />
                <button id="botonLogin" type="submit">Crear!</button>
            </form>
        </div>
    );
};

export default UploadHotel;