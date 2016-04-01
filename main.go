package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	endereço := flag.String("e", ":8080", "endereço [IP:PORTA] a ser usado")
	usuario := flag.String("u", "admin", "nome do usuário")
	senha := flag.String("s", "foobar", "senha do usuário")
	flag.Parse()

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

	// Habilitando esquema de autorização simples
	auth := r.Group("/", gin.BasicAuth(gin.Accounts{*usuario: *senha}))

	auth.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "favoritos.html", gin.H{
			"proxPagina": 1,
			"links":      ObterPagina(0, true),
		})
	})

	auth.GET("/pagina/:pag", func(c *gin.Context) {

		pag, _ := strconv.Atoi(c.Param("pag"))
		c.HTML(http.StatusOK, "favoritos.html", gin.H{
			"proxPagina": pag + 1,
			"links":      ObterPagina(pag, true),
		})
	})

	auth.GET("/formulario", func(c *gin.Context) {
		c.HTML(http.StatusOK, "formulario.html", gin.H{
			"novaUrl":    c.Query("url"),
			"novoTitulo": c.Query("titulo"),
		})
	})

	auth.POST("/salvar", func(c *gin.Context) {

		link := construirLink(c)
		var erro error

		if c.PostForm("inputId") != "" {
			AtualizarLink(link)
		} else {
			erro = NovoLink(link)
		}

		msg := "Link salvo com sucesso!"
		if erro != nil {
			msg = "OPA! Esse link já foi cadastrado O.o"
			log.Printf("Erro ao inserir novo link: %v\n", erro)
		}

		c.HTML(http.StatusOK, "resp-salvar.html", gin.H{
			"error": erro,
			"msg":   msg,
		})
	})

	auth.GET("/remover/:id", func(c *gin.Context) {

		RemoverLink(c.Param("id"))
		c.HTML(http.StatusOK, "favoritos.html", gin.H{
			"msg":   "Link removido!",
			"links": ObterPagina(0, true),
		})
	})

	auth.GET("/editar/:id", func(c *gin.Context) {

		c.HTML(http.StatusOK, "formulario.html", gin.H{
			"link": ObterLink(c.Param("id")),
		})
	})

	r.Run(*endereço)
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
