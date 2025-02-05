package dto

type TeamDto struct {
	Nombre      string `json:"nombre"`  
}

type PlayersDto struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	TeamID int `json:"teams_id"`

}


