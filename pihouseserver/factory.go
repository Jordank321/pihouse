package main

import (
	"github.com/Jordank321/pihouse/pihouseserver/control"
	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/jinzhu/gorm"
)

func ProvideDB() (*gorm.DB, error) {
	return gorm.Open("mssql",
		GetSqlConnectionString())
}

func ProvideTemperaureRepository() db.TemperatureRepository {
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	return &db.SQLTemperatureRepository{Connection: dbret}
}

func ProvideNodeRepository() db.NodeRepository {
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	return &db.SQLNodeRepository{Connection: dbret}
}

func ProvideHumidityRepository() db.HumidityRepository {
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	return &db.SQLHumidityRepository{Connection: dbret}
}

func ProvideAIRepository() db.AIRepository {
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	return &db.SQLAIRespository{Connection: dbret}
}

var clientController control.ClientController

func ProvideClientController() control.ClientController {
	if clientController == nil {
		clientController = &control.WebSocketClientController{}
	}
	return clientController
}
