import axios from 'axios';

const API_URL = 'http://localhost:2375'; // URL del domonio de Docker

export const getContainers = async () => {
    const response = await axios.get(`${API_URL}/containers/json`);
    return response.data;
};

export const createContainer = async (config) => {
    const response = await axios.post(`${API_URL}/containers/create`, config);
    return response.data;
};

export const startContainer = async (id) => {
    await axios.post(`${API_URL}/containers/${id}/start`);
};

export const stopContainer = async (id) => {
    await axios.post(`${API_URL}/containers/${id}/stop`);
};

export const deleteContainer = async (id) => {
    await axios.delete(`${API_URL}/containers/${id}`);
};
