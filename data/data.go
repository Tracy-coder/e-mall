package data

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"

	"entgo.io/ent/dialect"
	"github.com/Tracy-coder/e-mall/data/ent"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
)

var data *Data

const (
	mySqlDsn = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
)

func initData() {
	var err error
	data, err = NewData()
	if err != nil {
		hlog.Fatal(err)
	}
}

// Default Get a default database and cache instance
func Default() *Data {
	return data
}

// Data .
type Data struct {
	DBClient *ent.Client
}

// NewData .
func NewData() (data *Data, err error) {
	data = new(Data)
	drv, err := sql.Open("mysql", fmt.Sprintf(mySqlDsn,
		"mall", "password", "127.0.0.1", 3306, "e-mall"))
	if err != nil {
		return nil, err
	}
	var drive dialect.Driver = drv
	client := ent.NewClient(ent.Driver(drive))

	if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
		hlog.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}

	data.DBClient = client
	return data, nil
}
