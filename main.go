package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	endereço = flag.String("e", ":8080", "endereço [IP:PORTA] a ser usado")
	usuario  = flag.String("u", "admin", "nome do usuário")
	senha    = flag.String("s", "foobar", "senha do usuário")
	certSSL  = flag.String("c", "cert.pem", "certificado ssl")
	keySSL   = flag.String("k", "key.pem", "chave do certificado ssl")
	naoHTTPS = flag.Bool("nao-https", false, "desabilita o uso de https/ssl")
)

func main() {

	flag.Parse()

	// configurando banco de dados
	var err error
	db, err = sql.Open("sqlite3", "./banco.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco: ", err)
	}
	defer db.Close()

	CriarBanco()

	// iniciando configuração do servidor web
	r := gin.Default()
	r.Static("/estatico", "estatico/")
	r.LoadHTMLGlob("templates/*")
	RegistrarRotas(r)

	if *naoHTTPS {
		log.Fatal(http.ListenAndServe(*endereço, r))
	} else {
		log.Fatal(http.ListenAndServeTLS(*endereço, *certSSL, *keySSL, r))
	}
}
