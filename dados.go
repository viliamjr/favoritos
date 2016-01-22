package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Link struct {
	Url         string
	Titulo      string
	Privado     bool
	DataCriacao time.Time
	Tags        string
}

var db *sql.DB

func NovoLink(link *Link) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Erro ao criar transação: ", err)
	}

	stmt, err := tx.Prepare("insert into link values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Erro ao preparar a query de insert: ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(link.Url, link.Titulo, link.Tags, link.DataCriacao, link.Privado)
	if err != nil {
		log.Fatal("Erro ao executar um insert no banco: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Erro ao commitar transação de insert: ", err)
	}
}

func ObterTodos() []*Link {

	encontrados := make([]*Link, 0)

	rows, err := db.Query("select * from link")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		link := &Link{}
		rows.Scan(&link.Url, &link.Titulo, &link.Tags, &link.DataCriacao, &link.Privado)
		encontrados = append(encontrados, link)
	}

	return encontrados
}

func ProcurarLinkPorTag(tag string) []*Link {

	encontrados := make([]*Link, 0)

	//TODO popula slice com itens encontrados

	return encontrados
}
