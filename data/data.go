package data

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"

	"entgo.io/ent/dialect"
	"github.com/Tracy-coder/e-mall/configs"
	"github.com/Tracy-coder/e-mall/data/ent"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	redis "github.com/redis/go-redis/v9"
)

var data *Data

const (
	mySqlDsn = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
)

func initData() {
	var err error
	data, err = NewData(configs.Data())
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
	Redis    *redis.Client
}

// NewData .
func NewData(config configs.Config) (data *Data, err error) {
	data = new(Data)
	db, err := NewSqlClient(config)
	if err != nil {
		return nil, err
	}
	data.DBClient = db

	rdb := newRedisDB(config)
	data.Redis = rdb
	return data, nil
}

func newRedisDB(config configs.Config) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	return rdb
}
func NewSqlClient(config configs.Config) (client *ent.Client, err error) {
	drv, err := sql.Open("mysql", fmt.Sprintf(mySqlDsn,
		config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName))
	if err != nil {
		return nil, err
	}
	var drive dialect.Driver = drv
	client = ent.NewClient(ent.Driver(drive))

	if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(true)); err != nil {
		hlog.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}
	return client, nil
}
