API RESTful con Golang

Este proyecto es una API RESTful construida con Golang utilizando el framework go-chi. La API maneja operaciones CRUD para equipos y jugadores, soporta CORS y permite la subida de archivos.

🚀 Tecnologías Utilizadas

Golang

Go-Chi (Router)

REL (ORM para Go)

MySQL

Docker

Postman (para pruebas de API)

📁 Estructura del Proyecto

ApiRestfullgolang/
│── connect/        # Conexión a la base de datos
│── dto/            # Definición de estructuras para transferencia de datos
│── models/         # Modelos de la base de datos
│── routes/         # Definición de rutas y controladores
│── utils/          # Utilidades para respuesta JSON
│── public/         # Archivos estáticos
│── .env            # Variables de entorno
│── main.go         # Punto de entrada de la aplicación

🔧 Configuración y Ejecución

1️⃣ Clonar el repositorio

git clone https://github.com/tu-repositorio.git
cd ApiRestfullgolang

