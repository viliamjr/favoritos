package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ListaTags []string

func (l ListaTags) String() string {
	var s string
	for i := 0; i < len(l); i++ {
		s = s + l[i]
		if i < (len(l) - 1) {
			s = s + ","
		}
	}
	return s
}

func NovasTags(s string) *ListaTags {
	tags := make(ListaTags, 0)
	for _, i := range strings.Split(s, ",") {
		tags = append(tags, i)
	}
	return &tags
}

type DataFormatada struct {
	time.Time
}

func (d DataFormatada) String() string {
	return fmt.Sprintf("%02d/%02d/%d", d.Day(), d.Month(), d.Year())
}

type Link struct {
	Id          int
	Url         string
	Titulo      string
	Privado     bool
	DataCriacao DataFormatada
	Tags        *ListaTags
}

var db *sql.DB

func CriarBanco() {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS link (
		url text not null unique,
		titulo text not null,
		tags text not null,
		data_criacao timestamp not null,
		privado bool not null);`)

	if err != nil {
		log.Fatal("Erro na criação do banco: ", err)
	}
}

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

	_, err = stmt.Exec(link.Url, link.Titulo, link.Tags.String(), link.DataCriacao.Time.Unix(), link.Privado)
	if err != nil {
		log.Fatal("Erro ao executar um insert no banco: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Erro ao commitar transação de insert: ", err)
	}
}

func AtualizarLink(link *Link) {

	_, err := db.Exec(`update link set url=?, titulo=?, tags=?, privado=? where rowid = ?;`,
		link.Url, link.Titulo, link.Tags.String(), link.Privado, link.Id)

	if err != nil {
		log.Fatal("Erro no update de link: ", err)
	}
}

func RemoverLink(id string) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Erro ao criar transação: ", err)
	}

	stmt, err := tx.Prepare("delete from link where rowid = ?")
	if err != nil {
		log.Fatal("Erro ao preparar a query de delete: ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal("Erro ao executar delete no banco: ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Erro ao commitar transação de insert: ", err)
	}
}

func ObterTodos() []*Link {

	encontrados := make([]*Link, 0)

	rows, err := db.Query("select rowid,url,titulo,tags,data_criacao,privado from link order by data_criacao desc;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		link := &Link{}
		var tags string
		rows.Scan(&link.Id, &link.Url, &link.Titulo, &tags, &(link.DataCriacao.Time), &link.Privado)
		link.Tags = NovasTags(tags)
		encontrados = append(encontrados, link)
	}

	return encontrados
}

func ObterLink(id string) *Link {

	link := &Link{}

	rows, err := db.Query("select rowid,url,titulo,tags,data_criacao,privado from link where rowid = ?;", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		var tags string
		rows.Scan(&link.Id, &link.Url, &link.Titulo, &tags, &(link.DataCriacao.Time), &link.Privado)
		link.Tags = NovasTags(tags)
	}

	return link
}

func ProcurarLinkPorTag(tag string) []*Link {

	encontrados := make([]*Link, 0)

	//TODO popula slice com itens encontrados

	return encontrados
}
