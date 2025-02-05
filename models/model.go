package models

type TeamModel struct {
	ID int `db:"id" json:"id"`
	Nombre string `db:"nombre" json:"nombre"`
	Slug string `db:"slug" json:"slug"`
	
}

func (b TeamModel) Table() string {
	return "teams"
}


type PlayersModel struct {
	ID          int    `db:"id" json:"id"`
    Nombre      string `db:"nombre" json:"nombre"`
    Slug        string `db:"slug" json:"slug"`
    Descripcion string `db:"descripcion" json:"descripcion"`
    TeamID      int    `db:"teams_id" json:"teams_id"`
	Teams    TeamModel `db:"teams"  json:"teams"` 
}

func (b PlayersModel) Table() string {
	return "players"
}
