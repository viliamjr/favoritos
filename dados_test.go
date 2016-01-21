package main

import (
	"database/sql"
	"testing"
	"time"
)

func TestDados(t *testing.T) {

	db, err := sql.Open("sqlite3", "./banco_de_teste.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	NovoLink(&Link{"www.google.com", "Buscador Google", false, time.Now().Unix(), "buscador,site,www,web"})
	NovoLink(&Link{"www.cade.com.br", "Buscador Google", false, time.Now().Unix(), "buscador,site,www,web"})

	encontrados := ObterTodos() //ProcurarLinkPorTag("www")
	if len(encontrados) > 0 {
		for _, link := range encontrados {
			t.Logf("%v\n", link)
		}
	}

	db.Close()
}
