import { lazy } from "react";
import Congrats from "pages/Congrats.jsx";
import Home from "pages/Home.jsx";
import Result from  "pages/Result.jsx";
import HotelDetails from "pages/HotelDetails.jsx";
import id from "pages/Home.jsx";
import Login from "pages/Login.jsx";

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
        id: 3,
        path: "/result",
        Element: Result,
    },
    {
        id: 4,
        path: `/result/'${id}`, //muy dudoso de que sea así
        Element:HotelDetails,
    },
    {
        id: 5,
        path: "/congrats",
        Element: Congrats,
    },
];