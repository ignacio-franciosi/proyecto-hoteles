import React, {useState} from 'react';
import {useNavigate} from 'react-router-dom';
import './../App.css';
import CustomModal from '../components/CustomModal.jsx';

const Register = () => {
    const navigate = useNavigate();
    const [name, setName] = useState('');
    const [lastName, setLastName] = useState('');
    const [dni, setDni] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    let [emptyRegister] = useState(false);

    const [showAlert1, setShowAlert1] = useState(false);
    const [showAlert2, setShowAlert2] = useState(false);
    const [showAlert3, setShowAlert3] = useState(false);
    const openAlert1 = () => {
        setShowAlert1(true);
    };
    const closeAlert1 = () => {
        setShowAlert1(false);
    };

    const openAlert2 = () => {
        setShowAlert1(true);
    };
    const closeAlert2 = () => {
        setShowAlert1(false);
    };

    const openAlert3 = () => {
        setShowAlert3(true);
    };
    const closeAlert3 = () => {
        setShowAlert3(false);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (name === '') {
            document.getElementById('inputNameRegister').style.borderColor = 'red';
            emptyRegister = true;
        } else {
            document.getElementById('inputNameRegister').style.borderColor = '';
        }
        if (lastName === '') {
            document.getElementById('inputLastNameRegister').style.borderColor = 'red';
            emptyRegister = true;
        } else {
            document.getElementById('inputLastNameRegister').style.borderColor = '';
        }
        if (dni === '') {
            document.getElementById('inputDniRegister').style.borderColor = 'red';
            emptyRegister = true;
        } else {
            document.getElementById('inputDniRegister').style.borderColor = '';

        }
        if (email === '') {
            document.getElementById('inputEmailRegister').style.borderColor = 'red';
            emptyRegister = true;
        } else {
            document.getElementById('inputEmailRegister').style.borderColor = '';
        }
        if (password === '') {
            document.getElementById('inputPasswordRegister').style.borderColor = 'red';
            emptyRegister = true;
        } else {
            document.getElementById('inputPasswordRegister').style.borderColor = '';
        }
        if (!emptyRegister) {
            try {
                const response = await fetch('http://localhost:8080/register', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        name: name,
                        lastName: lastName,
                        email: email,
                        password: password,
                        dni: parseInt(dni),

                    }),
                });
                if (response.ok) {
                    openAlert2()
                    navigate('/login');
                } else {
                    openAlert3()
                    console.log('el usuario ya existe');
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }
        } else {
            openAlert1()
            emptyRegister = true;
        }
    };

    return (
        <div id="body">
            <CustomModal
                showModal={showAlert1}
                closeModal={closeAlert1}
                content="Debes completar todos los campos"
            />
            <CustomModal
                showModal={showAlert2}
                closeModal={closeAlert2}
                content="Gracias por registrarte! ahora puedes iniciar sesión :)"
            />
            <CustomModal
                showModal={showAlert3}
                closeModal={closeAlert3}
                content="El correo electronico ya está registrado"
            />
            <h1 id="h1Register">Registrarse</h1>
            <form id="formRegister" onSubmit={handleSubmit}>
                <input
                    id="inputNameRegister"
                    type="text"
                    placeholder="Nombre"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
                <input
                    id="inputLastNameRegister"
                    type="text"
                    placeholder="Apellido"
                    value={lastName}
                    onChange={(e) => setLastName(e.target.value)}
                />
                <input
                    id="inputDniRegister"
                    type="text"
                    placeholder="DNI"
                    value={dni}
                    onChange={(e) => setDni(e.target.value)}
                />
                <input
                    id="inputEmailRegister"
                    type="email"
                    placeholder="Correo electrónico"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    id="inputPasswordRegister"
                    type="password"
                    placeholder="Contraseña"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <br/>
                <button id="botonLogin" type="submit">
                    Registrarse
                </button>
            </form>
        </div>
    );
};

export default Register;
