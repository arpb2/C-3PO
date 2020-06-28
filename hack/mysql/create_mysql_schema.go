package main

import (
	"context"

	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/facebookincubator/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	drv, err := sql.Open("mysql", "mockuser:mockpassword@tcp(3.16.213.100:3380)/testdb")

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
