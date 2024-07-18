import {useEffect, useState} from "react";


const ContainerView = () => {
    const [containers, setContainers] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true)


    useEffect(() => {
        const fetchStats = async () => {

            try {
                const response = await fetch("http://localhost:8004/stats");
                if (response.ok) {
                    const data = await response.json();

                    const sortedData = data.sort((a, b) => a.Name.localeCompare(b.Name));
                    setContainers(sortedData);
                } else {
                    const data = await response.json()
                    const errorMessage = data.error || 'Error';
                    throw new Error(errorMessage);
                }
            } catch (error) {
                console.error(error);
                setError(error.message);
            } finally {
                setLoading(false)
            }
        };

        fetchStats();

        const intervalId = setInterval(fetchStats, 1000);

        return () => clearInterval(intervalId);
    }, [])


    if (loading) {
        return (
            <>
                <h1 id="h1Login">Cargando...</h1>

            </>
        )
    }

    if (error) {
        window.location.reload();
        return (
            <>
                <h1 id={"h1Logini"}>Error: {error}</h1>

            </>
        );
    }

    if (!containers) {
        return (
            <>
                <h1 id="h1Login">No hay contenedores en ejecución</h1>
            </>
        )
    }

    return (
        <>
            <h1 id={"h1Home"}>Administración de infraestructura</h1>
            <br/>
            <br/>
            <div id={"backHotelSearch"}>
                <div className="hotelContainer">
                    {containers.map((container) => (

                        <div key={container.ID} className={"hotelCard"}>
                            <div>
                                <p>ID: {container.ID}</p>
                                <p>Nombre: {container.Name}</p>
                                <p>CPU: {container.CPUPerc}</p>
                                <p>Memoria: {container.MemPerc}</p>
                                <p>Uso: {container.MemUsage}</p>
                            </div>
                        </div>
                    ))}


                </div>
            </div>
        </>
    )
}

export default ContainerView;