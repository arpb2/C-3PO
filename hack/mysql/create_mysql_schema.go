package main

import (
	"context"
	"fmt"
	"os"

	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/facebookincubator/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn, exists := os.LookupEnv("MYSQL_DSN")
	if !exists {
		panic("No MYSQL_DSN provided")
	}

	ctx := context.Background()
	drv, err := sql.Open("mysql", fmt.Sprintf("%s/testdb", dsn))

	if err != nil {
		panic(err)
	}
	defer drv.Close()

	client := ent.NewClient(ent.Driver(drv))

	err = client.Schema.Create(ctx)
	if err != nil {
		panic(err)
	}
}
