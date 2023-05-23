package main

import (
	initializers "heptaber/notification/app/initializers"
	"heptaber/notification/domain/service"
	"sync"
	// db "heptaber/notification/infrastructure/database"
)

func init() {
	initializers.LoadEnvVariables()
	// db.ConnectToDB()
	// db.SyncDatabase()
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
