package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jjg-akers/db-abstraction-patterns/atomic/framework"
	"github.com/jjg-akers/db-abstraction-patterns/atomic/usecase"
)

func main() {
	fmt.Println("starting main")

	// get a db
	db := pretendToMakeADBConn()

	// build your store
	repo := framework.NewAtomicRepo(db)

	// build the thing
	uc := usecase.NewDoer(repo)

	// run it
	uc.DoTheThingThatRequiresMultiple(context.TODO())
}

func pretendToMakeADBConn() *sql.DB {
	return nil
}
