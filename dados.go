package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Link struct {
	Url         string
	Titulo      string
	Privado     bool
	DataCriacao int64
	Tags        string
}

var db *sql.DB;

func NovoLink(link *Link) {

	stmt, err := tx.Prepare("insert into foo(nome) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if err != nil {
		log.Fatal("Erro ao criar novo link: ", err)
	}
}

func ProcurarLinkPorTag(tag string) []*Link {

	encontrados := make([]*Link, 0)

	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("LinksAlgorix"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			link := Link{}
			link.popular(v)
			if strings.Contains(link.Tags, tag) {
				encontrados = append(encontrados, &link)
			}
		}
		return nil
	})

	return encontrados
}

func ObterTodos() []*Link {

	encontrados := make([]*Link, 0)

	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("LinksAlgorix"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			link := Link{}
			link.popular(v)
			encontrados = append(encontrados, &link)
		}
		return nil
	})

	return encontrados
}
