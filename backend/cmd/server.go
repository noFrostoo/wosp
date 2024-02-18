package main

import (
	"backend/api/v1"
	"backend/router"
	"backend/store"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db := set_up_db()
	us := store.NewUserStore(db)
	ts := store.NewTodoStore(db)
	h, err := v1.NewHandler(us, ts)
	if err != nil {
		log.Fatal(err)
	}

	e := router.New()

	v1 := e.Group("/api/v1")
	h.Register(v1)

	e.Logger.Fatal(e.Start(":8080"))
}

var schema = `
create table if not exists "user"
(
    id            uuid primary key default gen_random_uuid(),
    username      text unique not null,
    password      text        not null
);

create table if not exists "todo"
(
    id            uuid primary key default gen_random_uuid(),
    user_id       uuid not null,
    title         text,
    description   text,
    done          Boolean not null,
    due_at        text not null
);`



func set_up_db() (*sqlx.DB) {
	database_name, ok := os.LookupEnv("DB_DATABASE")
	if !ok {
		log.Fatal("Not database name, cannot start")
	}

	db_username, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		log.Fatal("Not database name, cannot start")
	}

	db_host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("Not database name, cannot start")
	}

	db_password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		log.Fatal("Not database name, cannot start")
	}

	db_string := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", db_host, db_username, database_name, db_password)
	db, err := sqlx.Connect("postgres", db_string)
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	return db
}
