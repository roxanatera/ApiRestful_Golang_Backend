package main

import (
	"ApiRestfullgolang/routes"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//CORS
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	

	//archivos estaticos
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fs))

	//prefijo
	prefijo := "/api/v1/"

	//rutas gerenales
	router.Get(prefijo+"RutaGet", routes.RutaGet)
	router.Get(prefijo+"RutaPath/{id}", routes.RutaPath)
	router.Post(prefijo+"RutaPost", routes.RutaPost)
	router.Put(prefijo+"RutaPut/{id}", routes.RutaPut)
	router.Delete(prefijo+"RutaDelete/{id}", routes.RutaDelete)
	router.Get(prefijo+"query-string", routes.RutaQuerystring)
	router.Post(prefijo+"Rutaupload", routes.Rutauploadfile)
	//Rutas de teams
	router.Get(prefijo+"Rutateams", routes.TeamsGet)
	router.Get(prefijo+"Rutateams/{id}", routes.TeamsGet_parametros)
	router.Post(prefijo+"Rutateams", routes.TeamsPost)
	router.Put(prefijo+"Rutateams/{id}", routes.TeamsPut)
	router.Delete(prefijo+"Rutateams/{id}", routes.TeamsDelete)

	//Rutas de players
	router.Get(prefijo+"Rutaplayers", routes.Playersget)
	router.Get(prefijo+"Rutaplayers/{id}", routes.Playersget_prametros)
	router.Post(prefijo+"Rutaplayers", routes.Playerspost)
	router.Put(prefijo+"Rutaplayers/{id}", routes.Playersput)
	router.Delete(prefijo+"Rutaplayers/{id}", routes.Playersdelete)

	
	



	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		
	}

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
