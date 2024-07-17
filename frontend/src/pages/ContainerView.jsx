import React from 'react';
import ContainerList from '../components/DockerComponents/ContainerList.jsx';
import ContainerActions from '../components/DockerComponents/ContainerActions.jsx';

const ContainerView = () => {
    return (

        <div className="App">
            <h1>Infrastructure Management</h1>
            <ContainerActions />
            <ContainerList />
        </div>
)
}
export default ContainerView