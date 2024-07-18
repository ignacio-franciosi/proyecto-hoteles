import React, {useEffect, useState} from "react";
import Cookies from "js-cookie";

const ContainerStats = () => {
    const [containers, setContainers] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true)
    const user_id = Cookies.get("user_id")
    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await fetch(`http://localhost:8080/user/${user_id}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch admin');
                }
                const userData = await response.json();
                console.log(userData.name, userData.dni, userData.userType)
                if (!userData.userType) {
                    navigate("/")
                }

            } catch (error) {
                console.error('Error fetching user:', error);
                navigate("/");
            }

        };
        fetchUser();
    }, [user_id, token]);
    useEffect(() => {
        const fetchStats = async () => {

            try {
                const response = await fetch("http://localhost:8004/info");
                if (response.ok) {
                    const data = await response.json();

                    const sortedData = data.sort((a, b) => a.Name.localeCompare(b.Name));
                    setContainers(sortedData);
                } else {
                    const data = await response.json()
                    const errorMessage = data.error || 'Error';
                }
            } catch (error) {
                console.error(error);
                setError(error.message);
            } finally {
                setLoading(false)
            }
        };

        fetchStats();

        const intervalId = setInterval(fetchStats, 10000);

        return () => clearInterval(intervalId);
    }, [])

    if (loading) {
        return (
            <>
                <h2>Contenedores</h2>
                <div className="fullscreen">Cargando...</div>
            </>
        )
    }

    if (error) {
        return (
            <>
                <div className="fullscreen">Error: {error}</div>
            </>
        );
    }

    if (!containers) {
        return (
            <>
                <div className="fullscreen">No hay contenedores corriendo</div>
            </>
        )
    }

    return (
        <>
            <h2>Contenedores</h2>
            <table className="container-table fullscreen">
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Nombre</th>
                    <th>CPU (%)</th>
                    <th>Memoria (%)</th>
                    <th>Memoria</th>
                </tr>
                </thead>
                <tbody>
                {containers.map((container) => (
                    <tr key={container.ID}>
                        <td>{container.ID}</td>
                        <td>{container.Name}</td>
                        <td>{container.CPUPerc}</td>
                        <td>{container.MemPerc}</td>
                        <td>{container.MemUsage}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </>
    )
}

export default ContainerStats;