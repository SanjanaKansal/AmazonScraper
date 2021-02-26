package models

type ConfigJson struct {
	HOST					string		`json:"HOST"`
	PORT					int			`json:"PORT"`
	MONGODBNAME				string		`json:"MONGODBNAME"`
	MONGOCOLLECTIONNAME		string		`json:"MONGOCOLLECTIONNAME"`
	MONGOURI				string		`json:"MONGOURI"`
}

