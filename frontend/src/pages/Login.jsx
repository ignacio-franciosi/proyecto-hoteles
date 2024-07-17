import React, {useState} from 'react';
import {useNavigate} from "react-router-dom";
import Cookies from 'js-cookie';
import CustomModal from '../components/CustomModal.jsx';
import './../App.css'


const Login = () => {
    const [email, setEmail] = useState(''); // Se inicializan las variables vacías
    const [password, setPassword] = useState('');
    let [emptyRegister] = useState(false);
    const [showAlert1, setShowAlert1] = useState(false);
    const [showAlert2, setShowAlert2] = useState(false);
    const navigate = useNavigate();
    const register = () => { // funbcion que te redirige a  register
        navigate("/register");
    };

    const openAlert1 = () => {
        setShowAlert1(true);
        Cookies.set('user_id', "-1")
        Cookies.set('user_id', "")
        Cookies.set('email', "")
        Cookies.set('token', "")
    };
    const closeAlert1 = () => {
        setShowAlert1(false);
    };
    const openAlert2 = () => {
        setShowAlert2(true);
        Cookies.set('user_id', "-1")
        Cookies.set('user_id', "")
        Cookies.set('email', "")
        Cookies.set('token', "")
    };
    const closeAlert2 = () => {
        setShowAlert2(false);
    };

    function navigateTo(responseData) {
        if (responseData.user_type) {
            navigate("/dashAdmin")
        }
        else{
            navigate("/")
        }
    }

    const handleSubmit = async (e) => {
        e.preventDefault(); // Para que no recargue la página
        if (email === '') {
            document.getElementById('inputEmailLogin').style.borderColor = 'red';
            emptyRegister = true;
        } else {
            document.getElementById('inputEmailLogin').style.borderColor = '';
        }
        if (password === '') {
            document.getElementById('inputPasswordLogin').style.borderColor = 'red';
            emptyRegister = true;
        } else {
            document.getElementById('inputEmailLogin').style.borderColor = '';
        }
        if (!emptyRegister) {

            try {
                // Envía la respuesta al backend (Postman, básicamente)
                const response = await fetch('http://localhost:8080/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({email, password}),
                });
                if (response.ok) {
                    console.log('Usuario válido');
                    const responseData = await response.json();
                    Cookies.set('user_id', responseData.id_user)
                    Cookies.set('email', email)
                    Cookies.set('token', responseData.token)

                    navigateTo(responseData);
                    window.location.reload()
                } else {
                    openAlert1();
                    console.log("hizo el error")
                }
            } catch (error) {
                Cookies.set('user_id', "")
                Cookies.set('email', "")
                Cookies.set('token', "")
                Cookies.set('client_id', "")
                console.log('Error al realizar la solicitud al backend:', error);
            }
        } else {
            openAlert2()
            emptyRegister = true;
        }

    };
    return (
        <div id="body">
            <CustomModal
                showModal={showAlert2}
                closeModal={closeAlert2}
                content="Debes completar todos los campos"
            />
            <CustomModal
                showModal={showAlert1}
                closeModal={closeAlert1}
                content="Usuario no registrado"
            />

            <form id="formLogin" onSubmit={handleSubmit}>
                <div className ={"TitleForm1"}> <h2> ¡Hola de vuelta!</h2></div>
                <input
                    id={'inputEmailLogin'}
                    type="email"
                    placeholder="Correo electrónico"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    id={'inputPasswordLogin'}
                    type="password"
                    placeholder="Contraseña"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />

                <button id="botonLogin" type="submit">Iniciar sesión</button>
                <br/>
                <button id="botonLogin" onClick={register}>Registrarse</button>
            </form>
        </div>

    );
};
export default Login;