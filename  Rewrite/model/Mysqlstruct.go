package model

type Users struct {
	UserID      string `json:"user_id" gorm:"AUTO_INCREMENT"`
	Name        string `json:"user_name" gorm:"user_name"`
	Password    string `json:"user_password" gorm:"user_password"`
	UserPicture string `json:"user_picture" gorm:"user_picture"`
	Identity    string `json:"identity" gorm:"identity"`
}

type Games struct {
	GameID      string `json:"game_id" gorm:"game_id"`
	GameName    string `json:"game_name" gorm:"game_name"`
	GameDate    string `json:"game_date" gorm:"game_date"`
	Place       string `json:"place" gorm:"place"`
	Info        string `json:"Info" gorm:"Info"`
	Appointment string `json:"appointment" gorm:"appointment"`
	TeamA       string `json:"TEAMA" gorm:"TEAMA"`
	TeamB       string `json:"TEAMB" gorm:"TEAMB"`
}

type Players struct {
	PlayerID string `json:"player_ID" gorm:"AUTO_INCREMENT"`
	Name     string `json:"name" gorm:"name"`
	TeamName string `json:"team_name" gorm:"team_name"`
	Image    string `json:"img" gorm:"img"`
	Info     string `json:"Info" gorm:"Info"`
}

type Reservation struct {
	UserName string `json:"user_name" gorm:"user_name"`
	Game     string `json:"game" gorm:"game"`
}

type Teams struct {
	TeamID      string `json:"team_ID" gorm:"AUTO_INCREMENT"`
	TeamName    string `json:"T_name" gorm:"T_name"`
	Logo        string `json:"logo" gorm:"logo"`
	Information string `json:"Info" gorm:"Info"`
}
