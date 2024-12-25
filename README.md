# Heptaber world!

Open-source version of a blog application with multiple roles and ability for user to create articles.</br>
Built with Golang to be lightweight.</br>
Clean Architecture was used as the core design.</br>
The app is designed in a microservices way (but not real microservices, yet) for easier features implementations and support.</br>
Additionally, the app uses RabbitMQ for service communication.</br>
Docker-compose is implemented and ready to go (run) using the defaults.</br>

## How to run for the first time?

1. Open the project folder and run `docker-compose up` </br>
2. Migrate data to set up databases:

```
cd auth
go run infrastructure/database/migrate/migrate.go
cd ../notification
go run infrastructure/database/migrate/migrate.go
cd ../blog
go run infrastructure/database/migrate/migrate.go
```

3. Enjoy ;)

## Backend

Services:

- auth service
- blog service
- notification service

**Common stack used:** golang 1.23, gin-gonic, gorm.
