package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"encoding/json"
)

type Translate struct {
	Head Head  `json:"head"`
	Def  []Def `json:"def"`
}

type Head struct{}

type Def struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
	Ts   string `json:"ts"`
	Tr   []Tr   `json:"tr"`
}

type Tr struct {
	Attr
	Syn  []Syn  `json:"syn"`
	Mean []Mean `json:"mean"`
	Ex   []Ex   `json:"ex"`
}

type Syn struct {
	Attr
}

type Mean struct {
	Attr
}

type Ex struct {
	Attr
	Tr
}

type Attr struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
	Gen  string `json:"gen"`
}

type Word struct {
	Id          int64  `db:"id"`
	CreatedDate string `db:"created_date"`
	Value       string `db:"value"`
	LangFrom    string `db:"lang_from"`
	LangTo      string `db:"lang_to"`
	Translate   json.RawMessage
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=password dbname=translator sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	var word []Word
	err = db.Select(&word, "select * from word")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", word[0])
}








