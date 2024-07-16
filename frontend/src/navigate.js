import { lazy } from "react";
import Home from "pages/Home.jsx";
import HotelsList from "frontend/src/pages/HotelsList.jsx";
import HotelDetails from "pages/HotelDetails.jsx";
import hotel_id from "frontend/src/pages/HotelsList.jsx";
import Login from "pages/Login.jsx";
import Register from "pages/Register.jsx";

const Home = lazy(() => import("pages/Home.jsx"));
//estas son las rutas de las paginas de nuestro programa a
export const navigation = [
    {
        id: 0,
        path: "/",
        Element: Home,
    },
    {
        id: 1,
        path: '/home',
        Element: Home,
    },
    {
        id: 2,
        path: '/login',
        Element: Login,
    },
    {
        id: 5,
        path: '/register',
        Element: Register,
    },
    {
        id: 3,
        path: '/hotels-list',
        Element: HotelsList,
    },
    {
        id: 4,
        path: `/hotels-list/'${hotel_id}`,
        Element:HotelDetails,
    },

];