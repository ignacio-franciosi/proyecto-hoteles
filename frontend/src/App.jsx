import { useState } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import React from 'react';
import Header from './components/Header';
import Footer from './components/Footer.jsx';
import Home from './pages/Home.jsx'
import LoginButton from './components/LoginButton.jsx'
import './App.css'

const App = ()=> {

  return (
    <div>
        <Router>
            <Header />
            <LoginButton />
            <Routes>
                <Route path="/" element={<Home />} />
            </Routes>
            <Footer />
        </Router>
    </div>
  );
};

export default App;
