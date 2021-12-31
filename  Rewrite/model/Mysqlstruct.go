package model

type Users struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
	Identity int    `json:"identity"`
}

type Games struct {
	GameID      int    `json:"game_id"`
	GameName    string `json:"game_name"`
	GameDate    string `json:"game_date"`
	Place       string `json:"place"`
	Info        string `json:"Info"`
	Appointment string `json:"appointment"`
	TeamA       string `json:"TEAMA"`
	TeamB       string `json:"TEAMB"`
}

type Players struct {
	PlayerID int    `json:"player_ID"`
	Name     string `json:"name"`
	TeamName string `json:"team_name"`
	Image    string `json:"img"`
	Info     string `json:"Info"`
}

type Reservation struct{

}

type Teams struct{

}

type Token