import { useState } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import React from 'react';
import Header from './components/Header';
import Footer from './components/Footer.jsx';
import Home from './pages/Home.jsx';
import Login from './pages/Login.jsx';
import Search from "./pages/Search.jsx";
import LoginButton from './components/LoginButton.jsx';
import BackButton from "./components/BackButton.jsx";
import HotelDetails from "./pages/HotelDetails.jsx"
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
                <Route path="/search" element={<Search />} />
                <Route path="/hotelDetails/:hotel_id" element={<HotelDetails />} />
            </Routes>
            <Footer />
        </Router>
    </div>
  );
};

export default App;
