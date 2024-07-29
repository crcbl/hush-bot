package data

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func GetDbClient() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("PGHOST") + ":" + os.Getenv("PGPORT"),
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASS"),
		Database: "secrets",
	})
	return db
}

func InitDbForUserClient() error {
	db := GetDbClient()
	defer db.Close()

    fmt.Println("Creating defaults...")
	createRoot(db)
	return createSchema(db)
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*SecretData)(nil),
		(*UserData)(nil),
		(*PermissionData)(nil),
	}

    fmt.Println("Creating tables...")
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
			// `pg:"on_delete:RESTRICT"` on foreign key field. ON UPDATE hook can be added using tag
			// `pg:"on_update:CASCADE"`
			// FKConstraints: false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func createRoot(db *pg.DB) {
	_, _ = db.Exec("CREATE ROLE root LOGIN PASSWORD 'p@$$w0rd';")
	_, _ = db.Exec("CREATE DATABASE root;")
}
