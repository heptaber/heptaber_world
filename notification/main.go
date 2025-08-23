package main

import (
	initializers "heptaber/notification/app/initializers"
	"heptaber/notification/domain/service"
	"heptaber/notification/infrastructure/database"
	"sync"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()
	database.RunMigrations()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		service.ConsumeVerificationEmailNotifications()
	}()

	wg.Wait()
}
