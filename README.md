# go-base

go-base is my own personal base code to start a new back end server project, feel free to use it if you find this base code helpful 😄

## Installation

```bash
go mod tidy
```

## Usage

```bash

# migration guide

go run cmd/main.go gobase make:migration <file_name> #to make migration file (use create_table_name_table or update_table_name_table to automatically generate sql)
go run cmd/main.go gobase migrate #run all the up file of the migration you made
go run cmd/main.go gobase migrate:reset #to reset migration or to run all the down fil of the migration file you made
go run cmd/main.go gobase migrate:refresh #to refresh database (run rollback, then migrate database again)
go run cmd/main.go gobase migrate:rollback <step> #to rollback migration or to run the down file of migration file according to the number of step inserted
go run cmd/main.go gobase migrate:status #to check the migration status


#Seeder guide (still in process)
go run cmd/main.go gobase db:seed #to run all the seeder file (make sure you create and set up the seeder file in the database folder first)

```

## Library used

- Gorm: https://gorm.io/
- Echo: https://echo.labstack.com/
- Viper: https://github.com/spf13/viper
- Cobra: https://github.com/spf13/cobra
- migrate: https://github.com/golang-migrate/migrate/v4

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
