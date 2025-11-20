package app

import (
	"Service-for-assigning-reviewers-for-Pull-Requests/config"
	"Service-for-assigning-reviewers-for-Pull-Requests/pkg/database/postgres"
	"fmt"
)

/*
func InitDB(db *postgres.DatabaseSource, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	if _, err = db.Pool.Exec(context.Background(), string(file)); err != nil {
		return fmt.Errorf("error executing sql: %v", err)
	}
	givenOrder, err := utils.GetGivenOrder()
	if err != nil {
		return fmt.Errorf("error getting givenOrder: %v", err)
	}

	if err = postgres.AddOrdersToDB(db, givenOrder); err != nil {
		return fmt.Errorf("error adding orders to database: %v", err)
	}
	return nil
}*/

func Run(cfg *config.Config) error {
	db, err := initPostgres(cfg)
	if err != nil {
		return err
	}
	defer func(db *postgres.DatabaseSource) {
		db.Close()
	}(db)
	
	pgRepository := initDBRepository(db)
	service := CreateNewOrderService(pgRepository)
	orderController := controller.CreateNewOrderController(orderService)
	if err = startServer(cfg, orderController); err != nil {
		return fmt.Errorf("start server error: %v", err)
	}
	return nil
}
