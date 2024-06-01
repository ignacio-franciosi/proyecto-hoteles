import React, { useState } from 'react';
import { createContainer, startContainer, stopContainer, deleteContainer } from './Docker.jsx';

const ContainerActions = () => {
    const [name, setName] = useState('');

    const handleCreate = async () => {
        const config = {
            Image: 'your-image-name', // Reemplaza con el nombre de la imagen que deseas usar
            name,
        };
        const result = await createContainer(config);
        await startContainer(result.Id);
    };

    const handleStop = async (id) => {
        await stopContainer(id);
    };

    const handleDelete = async (id) => {
        await deleteContainer(id);
    };

    return (
        <div>
            <h2>Manage Containers</h2>
            <input
                type="text"
                placeholder="Container Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
            />
            <button onClick={handleCreate}>Create Container</button>
        </div>
    );
};

export default ContainerActions;
