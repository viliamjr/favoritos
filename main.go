package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Configurando banco de dados
	var err error
	db, err = sql.Open("sqlite3", "./banco.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco: ", err)
	}
	defer db.Close()

	// Iniciando configuração do servidor web
	r := gin.Default()

	r.Static("/estatico", "estatico/")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "formulario.html", nil)
	})

	r.POST("/adicionarLink", func(c *gin.Context) {

		link := construirLink(c)
		NovoLink(link)

		c.HTML(http.StatusOK, "formulario.html", gin.H{
			"msg": fmt.Sprintf("SUCESSO?! %v", *link),
		})
	})

	/*
		r.GET("/olaAngular", func(c *gin.Context) {

			c.HTML(http.StatusOK, "olaAngular.html", gin.H{
				"msg": "Bem-vindo!!",
			})
		})

		r.GET("/dados", func(c *gin.Context) {

			c.JSON(http.StatusOK,
				struct {
					Mensagem string `json:"msg"`
				}{"Oláaaaaaaaaa, javascript!"},
			)
		})
	*/

	r.Run(":8080")
}

func construirLink(c *gin.Context) *Link {

	var privado bool
	if c.PostForm("inputPrivado") != "" {
		privado = true
	}

	return &Link{
		c.PostForm("inputUrl"),
		c.PostForm("inputTitulo"),
		privado,
		time.Now(),
		c.PostForm("inputTags"),
	}
}
