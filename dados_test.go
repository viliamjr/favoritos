package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	flag.Parse()

	arquivoBD := "./banco_de_teste.db"
	os.Remove(arquivoBD)

	var err error
	db, err = sql.Open("sqlite3", arquivoBD)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	CriarBanco()

	resultado := m.Run()

	db.Close()

	os.Remove(arquivoBD)

	os.Exit(resultado)
}

func TestNovoLink(t *testing.T) {

	NovoLink(&Link{URL: "www.google.com", Titulo: "Buscador Google", Privado: false, DataCriacao: DataFormatada{time.Now()}, Tags: NovasTags("buscador,site,www,web")})
	NovoLink(&Link{URL: "www.cade.com.br", Titulo: "Buscador Cade", Privado: true, DataCriacao: DataFormatada{time.Now()}, Tags: NovasTags("brasil,buscador,site,www,web")})
}

func TestObterPagina(t *testing.T) {

	encontrados := ObterPagina(0, false)
	if len(encontrados) > 0 {
		for _, link := range encontrados {
			t.Logf("%v\n", link)
		}
	}
}
