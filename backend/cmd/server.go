package main

import (
	"log"
	"backend/router"
	"backend/store"
	"backend/api/v1"

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
	db, err := sqlx.Connect("postgres", "host=db port=5432 user=wosp dbname=wosp password=wosp sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	return db
}
