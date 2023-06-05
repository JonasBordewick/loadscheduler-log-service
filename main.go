package main

import (
	"sync"

	"github.com/JonasBordewick/loadscheduler-log-service/database"
	"github.com/JonasBordewick/loadscheduler-log-service/service"
)

func main() {
	db := database.GetDBInstance()
	var wg sync.WaitGroup
	wg.Add(1)
	go service.StartServer(&wg, 49497, db)
	wg.Wait()
	db.Close()
}