import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './../App.css';

// Parte funcional
const Login = () => {
    const navigate = useNavigate(); // Permite la navegación entre páginas con las rutas
    const [email, setEmail] = useState(''); // Se inicializan las variables vacías
    const [password, setPassword] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault(); // Para que no recargue la página
        if (email === '') {
            document.getElementById('inputEmailLogin').style.borderColor = 'red';
        } else{
            document.getElementById('inputEmailLogin').style.borderColor = '';
        }
        if (password === '') {
            document.getElementById('inputPasswordLogin').style.borderColor = 'red';
        } else{
            document.getElementById('inputEmailLogin').style.borderColor = '';
        }

            try {
                // Envía la respuesta al backend (Postman, básicamente)
                const response = await fetch('http://localhost:8080/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        email: email,
                        password: password
                    }),
                }).then((response) => {
                    if (response.ok) {
                        return response.json();
                    } else {
                        alert('Usuario Inválido');
                        return { "id_user": -1}
                    }
                });
                if (response.id_user) {
                    // Si el usuario existe
                    // El usuario está en la base de datos
                    console.log('Usuario válido');

                    localStorage.setItem('user_id', response.id_user);
                    localStorage.setItem('email', email);
                    navigate(-1);
                }
            } catch (error) {
                console.log('Error al realizar la solicitud al backend:', error);
            }

    };

    // Parte visible
    return (
        <div id="body">
            <h1 id="h1Login">Iniciar sesión</h1>
            <form id="formLogin" onSubmit={handleSubmit}>
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

                <button id="botonLogin" type="submit">
                    Iniciar sesión
                </button>
            </form>
        </div>
    );
};

export default Login;
