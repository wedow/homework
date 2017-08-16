package db

import (
	"flag"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source/go-bindata"
	"github.com/wedow/homework/db/migrations"
	"log"
)

func init() {
	// skip running migrations when running tests
	if flag.Lookup("test.v") != nil {
		return
	}

	// wrap assets into Resource
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})

	d, err := bindata.WithInstance(s)
	m, err := migrate.NewWithSourceInstance("go-bindata", d, "postgres://localhost:5432/homework?sslmode=disable")
	if err != nil {
		log.Printf("[ERR] Failed to initialize database migrations: %s", err)
	} else {
		m.Up() // run your migrations and handle the errors above of course
	}
}
