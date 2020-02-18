package client

import (
	"fmt"
	"log"

	"github.com/arpb2/C-3PO/pkg/ent"
	"github.com/facebookincubator/ent/dialect/sql"
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
