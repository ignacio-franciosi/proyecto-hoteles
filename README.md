# 🏨 Hotel Booking Platform

Plataforma para buscar y reservar hoteles, desarrollada con una arquitectura basada en microservicios y un frontend hecho en React. La idea fue construir algo realista, desacoplado y escalable, integrando distintas tecnologías como bases de datos, mensajería, búsqueda y caché.

---

## ⚙️ ¿Qué hace?

- Permite buscar hoteles por ciudad y fechas desde una pantalla inicial.
- Muestra los resultados disponibles con detalles claros y actualizados.
- Permite ver más información de cada hotel seleccionado.
- Y confirma si la reserva fue exitosa o rechazada, desde el frontend.

---

## 🧩 Tecnologías y Arquitectura

El sistema está compuesto por varios microservicios:

- Uno que maneja la información de los hoteles, conectándose a MongoDB.
- Otro que gestiona las búsquedas usando Solr como motor de búsqueda.
- Uno encargado de las reservas, que usa MySQL, Memcached y valida disponibilidad externa con el servicio Amadeus.
- Y todo esto conectado por eventos a través de RabbitMQ para mantener los datos sincronizados.

Cada servicio está dockerizado y todo el sistema se levanta fácilmente con Docker Compose. Se implementó manejo de errores claros, estructura MVC en el backend y comunicación entre servicios usando eventos.

![image](https://github.com/user-attachments/assets/0d960a29-c65c-4f84-a98c-da89b7a101bf)




