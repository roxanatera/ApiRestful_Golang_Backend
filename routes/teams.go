package routes

import (
	"ApiRestfullgolang/connect"
	"ApiRestfullgolang/dto"
	"ApiRestfullgolang/models"
	"ApiRestfullgolang/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gosimple/slug"

	"github.com/go-chi/chi/v5"
	"github.com/go-rel/rel/where"
)

func TeamsGet(w http.ResponseWriter, r *http.Request) {
	var datos []models.TeamModel
	if err := connect.Connect().FindAll(context.TODO(), &datos); err != nil {
		respuesta := map[string]string{
			"mensaje": "Se produjo un error al intentar obtener los datos",
		}
		utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
		return
	}
	// Respuesta exitosa
	utils.ReturnJSON(w, http.StatusOK, datos)
}

func TeamsGet_parametros(w http.ResponseWriter, r *http.Request) {
    var dato models.TeamModel
    id := chi.URLParam(r, "id")

    // Realiza la consulta a la base de datos
    if err := connect.Connect().Find(context.TODO(), &dato, where.Eq("id", id)); err != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "Se produjo un error al intentar obtener los datos",
        }
        utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
        return
    }

    // Devuelve los datos encontrados
    utils.ReturnJSON(w, http.StatusOK, dato)
}

//Create a new team
func TeamsPost(w http.ResponseWriter, r *http.Request) {
	data := dto.TeamDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error al leer los datos",
		}
		utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
		return
	}
	datos:= models.TeamModel{Nombre: data.Nombre, Slug: slug.Make(data.Nombre)}
	if err := connect.Connect().Insert(context.TODO(), &datos); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error al insertar los datos",
		}
		utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
		return
	}
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Datos insertados correctamente",
	}
	utils.ReturnJSON(w, http.StatusOK, respuesta)
}

//Update a team
func TeamsPut(w http.ResponseWriter, r *http.Request) {
	data := dto.TeamDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error al leer los datos",
		}
		utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
		return
	}
	var datos models.TeamModel
	if errorBuscar:=connect.Connect().Find(context.TODO(), &datos, where.Eq("id", 
	chi.URLParam(r, "id")));errorBuscar != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error al buscar los datos",
		}
		utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
		return
	}
	datos.Nombre = data.Nombre
	datos.Slug = slug.Make(data.Nombre)
	if errorUpdate:=connect.Connect().Update(context.TODO(), &datos);errorUpdate != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error al actualizar los datos",
		}
		utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
		return
	}
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Datos actualizados correctamente",
	}
	utils.ReturnJSON(w, http.StatusOK, respuesta)
}
//Delete team
func TeamsDelete(w http.ResponseWriter, r *http.Request) {
	var datos models.TeamModel
	if errorBuscar:=connect.Connect().Find(context.TODO(), &datos, where.Eq("id", 
	chi.URLParam(r, "id")));errorBuscar != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error al buscar los datos",
		}
		utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
		return
	}
	if err := connect.Connect().Delete(context.TODO(), &datos); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Error al eliminar los datos",
		}
		utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
		return
	}
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Datos eliminados correctamente",
	}
	utils.ReturnJSON(w, http.StatusOK, respuesta)
}