package store

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

type storePostgres struct {
	db *sql.DB
}

func NewPostgresMigrator(connectionString string) (Migrator, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &storePostgres{
		db: db,
	}, nil
}

func NewPostgresStore(connectionString string) (Store, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &storePostgres{
		db: db,
	}, nil
}

func (store *storePostgres) MigrateDB() error {
	migrations := allMigrations()
	n, err := migrate.Exec(store.db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Println("db.migrations: applied %d migrations", n)

	return nil
}

func (store *storePostgres) ReverseDB() error {
	migrations := allMigrations()
	n, err := migrate.Exec(store.db, "postgres", migrations, migrate.Down)
	if err != nil {
		return err
	}
	fmt.Println("db.migrations: reversed %d migrations", n)

	return nil
}

func allMigrations() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{`
				CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				name  VARCHAR NOT NULL,
				created_at TIMESTAMPTZ NOT NULL,
				updated_at TIMESTAMPTZ NOT NULL
			)`},
				Down: []string{`DROP TABLE users`},
			},
			{
				Id: "2",
				Up: []string{`
				ALTER TABLE users
				DROP COLUMN name;
				`},
				Down: []string{`
				ALTER TABLE users
				ADD COLUMN name VARCHAR NOT NULL;
				`},
			},
		},
	}
}

var myMigrations = []*migrate.Migration{
	&migrate.Migration{
		Id: "1",
		Up: []string{`
		CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name  VARCHAR NOT NULL,
		email VARCHAR UNIQUE NOT NULL,
		token VARCHAR NOT NULL
		);
		CREATE TABLE people_in_charge (
			user_id INT PRIMARY KEY,
			menu_id INT NOT NULL
		);
		CREATE TABLE menus (
			id SERIAL PRIMARY KEY,
			owner_id INT,
			menu_name VARCHAR,
			created_at TIMESTAMPTZ NOT NULL,
			deadline TIMESTAMPTZ NOT NULL,
			payment_reminder TIMESTAMPTZ NOT NULL,
			status INT
		);
		CREATE TABLE items (
			id SERIAL PRIMARY KEY,
			item_name VARCHAR,
			menu_id INT
		);
		CREATE TABLE orders (
			id SERIAL PRIMARY KEY,
			user_id INT,
			item_id INT,
			created_at TIMESTAMPTZ,
			updated_at TIMESTAMPTZ
		);
		ALTER TABLE people_in_charge ADD FOREIGN KEY (user_id) REFERENCES users(id);
		ALTER TABLE people_in_charge ADD FOREIGN KEY (menu_id) REFERENCES menus(id);
		ALTER TABLE menus ADD FOREIGN KEY (owner_id) REFERENCES users(id);
		ALTER TABLE items ADD FOREIGN KEY (menu_id) REFERENCES menus(id);
		ALTER TABLE orders ADD FOREIGN KEY (item_id) REFERENCES items(id);
		ALTER TABLE orders ADD FOREIGN KEY (user_id) REFERENCES users(id);

	`},
		Down: []string{`
		DROP TABLE people_in_charge;
		DROP TABLE orders;
		DROP TABLE items;
		DROP TABLE menus;
		DROP TABLE users;
		`},
	},
}

func (store *storePostgres) MigrateOneDB() error {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: myMigrations[:1],
	}
	n, err := migrate.Exec(store.db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Println("db.migrations: Migrate %d migrations", n)
	return nil
}

func (store *storePostgres) ReverseOneDB() error {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: myMigrations[:1],
	}
	n, err := migrate.Exec(store.db, "postgres", migrations, migrate.Down)
	if err != nil {
		return err
	}
	fmt.Println("db.migrations: reversed %d migrations", n)
	return nil
}
