package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
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

	CriarBanco()

	// Iniciando configuração do servidor web
	r := gin.Default()

	r.Static("/estatico", "estatico/")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "favoritos.html", gin.H{
			"links": ObterTodos(),
		})
	})

	r.GET("/formulario", func(c *gin.Context) {
		c.HTML(http.StatusOK, "formulario.html", nil)
	})

	r.POST("/salvar", func(c *gin.Context) {

		link := construirLink(c)

		if c.PostForm("inputId") != "" {
			AtualizarLink(link)
		} else {
			NovoLink(link)
		}

		c.HTML(http.StatusOK, "favoritos.html", gin.H{
			"msg":   "Link salvo!",
			"links": ObterTodos(),
		})
	})

	r.GET("/remover/:id", func(c *gin.Context) {

		RemoverLink(c.Param("id"))

		c.HTML(http.StatusOK, "favoritos.html", gin.H{
			"msg":   "Link removido!",
			"links": ObterTodos(),
		})
	})

	r.GET("/editar/:id", func(c *gin.Context) {

		c.HTML(http.StatusOK, "formulario.html", gin.H{
			"link": ObterLink(c.Param("id")),
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

	id, err := strconv.Atoi(c.PostForm("inputId"))
	if err != nil {
		id = -1
	}

	return &Link{
		Id:          id,
		Url:         c.PostForm("inputUrl"),
		Titulo:      c.PostForm("inputTitulo"),
		Privado:     privado,
		DataCriacao: DataFormatada{time.Now()},
		Tags:        NovasTags(c.PostForm("inputTags")),
	}
}
