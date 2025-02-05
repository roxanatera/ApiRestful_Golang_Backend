package routes

import (
	"ApiRestfullgolang/connect"
	"ApiRestfullgolang/dto"
	"ApiRestfullgolang/models"
	"ApiRestfullgolang/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/gosimple/slug"
)

//Rutas para los jugadores
func Playersget(w http.ResponseWriter, r *http.Request) {
	var datos []models.PlayersModel
	if err := connect.Connect().FindAll(context.TODO(), &datos, rel.Select("*", "teams.*").JoinAssoc("teams"), sort.Desc("id")); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
		return
	}
	utils.ReturnJSON(w, http.StatusOK, datos)
}

func Playersget_prametros(w http.ResponseWriter, r *http.Request) {
    var datos models.PlayersModel
    id, _ := strconv.Atoi(chi.URLParam(r, "id"))
    
    // Pasar un puntero a 'datos'
    if errBuscar := connect.Connect().Find(context.TODO(), &datos, rel.Select("*", "teams.*").JoinAssoc("teams"), where.Eq("id", id)); errBuscar != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "Ocurrió un error inesperado",
        }
        utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
        return
    }

    // Retornar los datos en la respuesta
    utils.ReturnJSON(w, http.StatusOK, datos)
}



func Playerspost(w http.ResponseWriter, r *http.Request) {
    data := dto.PlayersDto{} 

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "Error al leer los datos",
        }
        utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
        return
    }
	log.Printf("Datos recibidos: Nombre=%s, Descripcion=%s, TeamID=%d", data.Nombre, data.Descripcion, data.TeamID)
    datos := models.PlayersModel{
        Nombre:      data.Nombre,
        Slug:        slug.Make(data.Nombre),
        Descripcion: data.Descripcion,
        TeamID:      data.TeamID, 
    }

    if errorInsert := connect.Connect().Insert(context.TODO(), &datos); errorInsert != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "Ocurrió un error inesperado",
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

//Ruta players put
func Playersput(w http.ResponseWriter, r *http.Request) {
    // Decodificar el JSON recibido
    data := dto.PlayersDto{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "Error al leer los datos",
        }
        utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
        return
    }

    // Buscar el registro por ID
    var datos models.PlayersModel
    id, _ := strconv.Atoi(chi.URLParam(r, "id"))
    if errBuscar := connect.Connect().Find(context.TODO(), &datos, where.Eq("id", id)); errBuscar != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "El registro no fue encontrado",
        }
        utils.ReturnJSON(w, http.StatusNotFound, respuesta)
        return
    }

    // Actualizar los datos
    datos.Nombre = data.Nombre
    datos.Slug = slug.Make(data.Nombre)
    datos.Descripcion = data.Descripcion
    datos.TeamID = data.TeamID

    // Guardar los cambios en la base de datos
    if errUpdate := connect.Connect().Update(context.TODO(), &datos); errUpdate != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "Error al actualizar el registro",
        }
        utils.ReturnJSON(w, http.StatusInternalServerError, respuesta)
        return
    }

    // Respuesta de éxito
    respuesta := map[string]string{
        "estado":  "ok",
        "mensaje": "Registro actualizado correctamente",
    }
    utils.ReturnJSON(w, http.StatusOK, respuesta)
}

func Playersdelete(w http.ResponseWriter, r *http.Request) {
    var datos models.PlayersModel
    id, _ := strconv.Atoi(chi.URLParam(r, "id"))
    if errBuscar := connect.Connect().Find(context.TODO(), &datos, where.Eq("id", id)); errBuscar != nil {
        respuesta := map[string]string{
            "estado":  "error",
            "mensaje": "El registro no fue encontrado",
        }
        utils.ReturnJSON(w, http.StatusNotFound, respuesta)
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