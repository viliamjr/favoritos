package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/viliamjr/favoritos/modelo"
	"github.com/viliamjr/favoritos/rotas"
)

var (
	endereço     = flag.String("e", ":8080", "endereço [IP:PORTA] a ser usado")
	usuario      = flag.String("u", "admin", "nome do usuário")
	senha        = flag.String("s", "foobar", "senha do usuário")
	caminhoBanco = flag.String("b", "./banco.db", "caminho do arquivo onde de banco de dados (sqlite)")
	certSSL      = flag.String("c", "cert.pem", "certificado ssl")
	keySSL       = flag.String("k", "key.pem", "chave do certificado ssl")
	naoHTTPS     = flag.Bool("nao-https", false, "desabilita o uso de https/ssl")
)

func main() {

	flag.Parse()

	db := modelo.CarregarBanco(caminhoBanco)
	defer db.Close()

	// iniciando configuração do servidor web
	r := gin.Default()
	r.Static("/estatico", "estatico/")
	r.LoadHTMLGlob("templates/*")
	rotas.RegistrarRotas(r, *usuario, *senha)

	if *naoHTTPS {
		log.Fatal(http.ListenAndServe(*endereço, r))
	} else {
		log.Fatal(http.ListenAndServeTLS(*endereço, *certSSL, *keySSL, r))
	}
}
