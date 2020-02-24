package client

import (
	"fmt"
	"log"

	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/facebookincubator/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateMysqlClient(dsn string) (client *ent.Client, drv *sql.Driver) {
	drv, err := sql.Open("mysql", fmt.Sprintf("%s?parseTime=True", dsn))
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client = ent.NewClient(ent.Driver(drv))
	//_ = client.Schema.Create(context.Background())
	return
}
