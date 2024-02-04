package main

import (
	"log"
	"backend/router"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	set_up_db()

	e := router.New()
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



func set_up_db() {
	db, err := sqlx.Connect("postgres", "host=db port=5432 user=wosp dbname=wosp password=wosp sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

}
