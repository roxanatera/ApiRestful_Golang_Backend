package routes

import (
	"ApiRestfullgolang/dto"
	"ApiRestfullgolang/utils"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	chi "github.com/go-chi/chi/v5"
)


func RutaGet(w http.ResponseWriter, r *http.Request) {
   respuesta:=map[string]string{
    "Estado":"Todo Ok",
    "mensaje":"Hello desde Get!",
    "Autorization":r.Header.Get("Authorization"),    
    }
   utils.ReturnJSON(w, http.StatusOK, respuesta)
  
}

func RutaPath(w http.ResponseWriter, r *http.Request) {
	id:= chi.URLParam(r,"id")
    respuesta:=map[string]string{"mensaje":"Hello desde Path, | id=" + id,}
    utils.ReturnJSON(w, http.StatusOK, respuesta)
}

func RutaPost(w http.ResponseWriter, r *http.Request) {
    data :=dto.TeamDto{}
    if err:=json.NewDecoder(r.Body).Decode(&data); err!=nil{
        respuesta:=map[string]string{"mensaje":"Ocurrio un error al decodificar el json"}
        utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
        return
    }
   respuesta:=map[string]string{
    "Estado":"Todo Ok",
    "mensaje":"Hello desde Post!",
    "nombre":data.Nombre,
}
   utils.ReturnJSON(w, http.StatusOK, respuesta) 
}
func RutaPut(w http.ResponseWriter, r *http.Request) {
    id:= chi.URLParam(r,"id")
    respuesta:=map[string]string{"mensaje":"Hello desde Put,  | id=" + id,}
    utils.ReturnJSON(w, http.StatusOK, respuesta)
}
func RutaDelete(w http.ResponseWriter, r *http.Request) {
	id:= chi.URLParam(r,"id")
    respuesta:=map[string]string{"mensaje":"Hello desde Delete,  | id=" + id,}
    utils.ReturnJSON(w, http.StatusOK, respuesta)
}
func RutaQuerystring(w http.ResponseWriter, r *http.Request) {
   id:= r.URL.Query().Get("id")
   slug:= r.URL.Query().Get("slug")
   respuesta:=map[string]string{"mensaje":"Hello desde QueryString,  | id=" + id + " | slug=" + slug,}
   utils.ReturnJSON(w, http.StatusOK, respuesta)
   
}
func Rutauploadfile(w http.ResponseWriter, r *http.Request) {
    // Parsea el archivo enviado en el campo "foto"
    file, handler, err := r.FormFile("foto")
    if err != nil {
        respuesta := map[string]string{
        "mensaje": "Error al subir el archivo"}
        utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
        return
    }
    defer file.Close()

    // Genera un nombre único para el archivo
    var extension = strings.Split(handler.Filename, ".")[1]
    time := strings.Split(time.Now().String(), " ")
    foto := string(time[4][6:14]) + "." + extension
    var archivo string = "public/uploads/fotos/" + foto

    // Crea el archivo en el servidor
    f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        respuesta := map[string]string{
        "mensaje": "Error al guardar el archivo"}
        utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
        return
    }
    defer f.Close()

    // Copia el contenido del archivo subido al archivo en el servidor
    _, err = io.Copy(f, file)
    if err != nil {
        respuesta := map[string]string{
        "mensaje": "Error al guardar el archivo"}
        utils.ReturnJSON(w, http.StatusBadRequest, respuesta)
        return
    }

    // Respuesta exitosa
    respuesta := map[string]string{
        "Estado":  "Todo Ok",
        "mensaje": "Archivo subido correctamente",
        "foto":    foto,
    }
    utils.ReturnJSON(w, http.StatusOK, respuesta)
}