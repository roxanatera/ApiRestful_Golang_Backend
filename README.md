## API RESTful con Golang

Este proyecto es una API RESTful construida con Golang utilizando el framework go-chi. La API maneja operaciones CRUD para equipos y jugadores, soporta CORS y permite la subida de archivos.

*** Tecnologías Utilizadas ***
1.	Lenguaje: Go (Golang)
o	Backend eficiente y de alto rendimiento, ideal para APIs RESTful.
2.	Framework de Enrutamiento: Chi
o	Ligero, rápido y compatible con middlewares estándar de Go.
3. Gestor de Base de Datos: MySQL
Base de datos relacional, escalable y robusta, ideal para almacenar información estructurada.
4.	Gestión de Variables de Entorno: godotenv
o	Configuración segura de claves y puertos mediante archivos .env.
5.	Arquitectura: RESTful API
o	Diseño basado en recursos, métodos HTTP semánticos (GET, POST, PUT, DELETE).
____________________________
Características Destacadas:
✅ Endpoints CRUD Completo:
•	Operaciones básicas: Crear (POST), Leer (GET), Actualizar (PUT), Eliminar (DELETE).
•	Ejemplo:
bash
Copy
GET    /api/v1/RutaGet
POST   /api/v1/RutaPost
PUT    /api/v1/RutaPut/{id}
DELETE /api/v1/RutaDelete/{id}
✅ Manejo de Parámetros en Rutas:
•	Uso de {id} para rutas dinámicas (ej: /api/v1/RutaPath/123).
•	Extracción de parámetros con chi.URLParam.
✅ Versionado de API:
•	Prefijo /api/v1/ para evolución controlada de endpoints.
✅ Configuración Portable:
•	Puerto del servidor configurable via variables de entorno (.env).
________________________________________
Arquitectura y Buenas Prácticas:
•	Separación de Preocupaciones:
o	Paquete dedicado (routes) para manejo de endpoints, separado de la lógica de inicio (main.go).
•	Código Limpio:
o	Funciones con responsabilidades únicas (ej: RutaGet, RutaPost).
•	Manejo de Errores:
o	Carga segura de variables de entorno con godotenv y panic controlado.
o	encoding/json: Codificación/decodificación de datos JSON.
o	net/http: Manejo de solicitudes y respuestas HTTP.
•	Arquitectura: API RESTful con respuestas estructuradas.


