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
import HotelDetails from "./pages/HotelDetails.jsx"
import ContainerPage from "./pages/ContainerView.jsx"
import './App.css';

const App = ()=> {

  return (
    <div>
        <Router>
            <Header />
            <BackButton />
            <LoginButton />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/home" element={<Home />} />
                <Route path ="/login" element={<Login />} />
                <Route path ="/register" element={<Register />} />
                <Route path="/hotels-list" element={<HotelsList />} />
                <Route path="/hotel-list/:hotel_id" element={<HotelDetails />} />
                <Route path="/dashAdmin" element={<HotelDetails />} />
                <Route path="/dashAdmin/container" element={<ContainerPage />} />
            </Routes>
            <Footer />
        </Router>
    </div>
  );
};

export default App;
