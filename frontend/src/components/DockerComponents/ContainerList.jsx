import React, { useEffect, useState } from 'react';
import { getContainers } from './Docker.jsx';

const ContainerList = () => {
    const [containers, setContainers] = useState([]);

    useEffect(() => {
        const fetchContainers = async () => {
            const result = await getContainers();
            setContainers(result);
        };

        fetchContainers();
    }, []);

    return (
        <div>
            <h2>Containers</h2>
            <ul>
                {containers.map(container => (
                    <li key={container.Id}>
                        {container.Names[0]} - {container.State}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ContainerList;
