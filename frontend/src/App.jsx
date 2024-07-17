import { useState } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import React from 'react';
import Header from './components/Header';
import Footer from './components/Footer.jsx';
import Home from './pages/Home.jsx';
import Login from './pages/Login.jsx';
import Register from './pages/Register.jsx';
import HotelsList from "./pages/HotelsList.jsx";
import LoginButton from './components/LoginButton.jsx';
import BackButton from "./components/BackButton.jsx";
import Buttons from "./components/Buttons.jsx";
import HotelDetails from "./pages/HotelDetails.jsx"
import ContainerView from "./pages/ContainerView.jsx"
import DashAdmin from "./pages/DashAdmin.jsx"
import EditHotels from "./pages/EditHotels.jsx"
import Users from "./pages/Users.jsx"
import './App.css';

const App = ()=> {

  return (
    <div>
        <Router>
            <Header />
            <BackButton />
            <LoginButton />
            <Buttons />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/home" element={<Home />} />
                <Route path ="/login" element={<Login />} />
                <Route path ="/register" element={<Register />} />
                <Route path="/hotels-list" element={<HotelsList />} />
                <Route path="/hotels-list/:hotel_id" element={<HotelDetails />} />
                <Route path="/dashAdmin" element={<DashAdmin />} />
                <Route path="/dashAdmin/container-view" element={<ContainerView />} />
                <Route path="/dashAdmin/edit-hotels" element={<EditHotels />} />
                <Route path="/dashAdmin/users" element={<Users />} />
            </Routes>
            <Footer />
        </Router>
    </div>
  );
};

export default App;
