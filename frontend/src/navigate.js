import { lazy } from "react";
import Home from "pages/Home.jsx";
import HotelsList from "frontend/src/pages/HotelsList.jsx";
import HotelDetails from "pages/HotelDetails.jsx";
import hotel_id from "frontend/src/pages/HotelsList.jsx";
import Login from "pages/Login.jsx";
import Register from "pages/Register.jsx";
import DashAdmin from "./pages/DashAdmin.jsx";
import ContainerView from "./pages/ContainerView.jsx";
import EditHotels from "./pages/EditHotels.jsx"
import UploadHotels from "./pages/UploadHotels.jsx"
import Users from "./pages/Users.jsx"

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
    {
        id: 5,
        path: `/dashAdmin`,
        Element:DashAdmin,
    },
    {
        id: 6,
        path: `/dashAdmin/container-view`,
        Element:ContainerView,
    },
    {
        id: 7,
        path: `/dashAdmin/edit-hotels`,
        Element:EditHotels,
    },
    {
        id: 8,
        path: `/dashAdmin/users`,
        Element:Users,
    },
    {
        id: 9,
        path: `/dashAdmin/upload-hotels`,
        Element:UploadHotels,
    },
];