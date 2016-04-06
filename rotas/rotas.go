package rotas

import (
	"favoritos/modelo"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// RegistrarRotas realiza o registro de todas as rotas da aplicação.
func RegistrarRotas(r *gin.Engine, usuario, senha string) {

	// Habilitando esquema de autorização simples
	auth := r.Group("/", gin.BasicAuth(gin.Accounts{usuario: senha}))

	auth.GET("/", Raiz)

	auth.GET("/pagina/:pag", Pagina)

	auth.GET("/formulario", Formulario)

	auth.POST("/salvar", Salvar)

	auth.GET("/remover/:id", Remover)

	auth.GET("/editar/:id", Editar)
}

// Raiz define a rota da página inicial
func Raiz(c *gin.Context) {

	c.HTML(http.StatusOK, "favoritos.html", gin.H{
		"proxPagina": 1,
		"links":      modelo.ObterPagina(0, true),
	})
}

// Pagina define a rota para a paginação dos links
func Pagina(c *gin.Context) {

	pag, _ := strconv.Atoi(c.Param("pag"))
	c.HTML(http.StatusOK, "favoritos.html", gin.H{
		"proxPagina": pag + 1,
		"links":      modelo.ObterPagina(pag, true),
	})
}

// Formulario define a rota para exibir o formulário de link
func Formulario(c *gin.Context) {
	c.HTML(http.StatusOK, "formulario.html", gin.H{
		"novaUrl":    c.Query("url"),
		"novoTitulo": c.Query("titulo"),
	})
}

// Salvar define a rota para salvar o link
func Salvar(c *gin.Context) {

	link := construirLink(c)

	if c.PostForm("inputId") != "" {

		modelo.AtualizarLink(link)
		c.HTML(http.StatusOK, "favoritos.html", gin.H{
			"proxPagina": 1,
			"links":      modelo.ObterPagina(0, true),
			"msg":        "Link atualizado!!",
		})
	} else {

		erro := modelo.NovoLink(link)
		msg := "Link salvo com sucesso!"
		if erro != nil {
			msg = "OPA! Esse link já foi cadastrado O.o"
			log.Printf("Erro ao inserir novo link: %v\n", erro)
		}
		c.HTML(http.StatusOK, "resp-salvar.html", gin.H{
			"error": erro,
			"msg":   msg,
		})
	}
}

// Remover define a rota para a remoção de um link
func Remover(c *gin.Context) {

	modelo.RemoverLink(c.Param("id"))
	c.HTML(http.StatusOK, "favoritos.html", gin.H{
		"msg":   "Link removido!",
		"links": modelo.ObterPagina(0, true),
	})
}

// Editar define a rota para exibir os dados de um link no formulário
func Editar(c *gin.Context) {

	c.HTML(http.StatusOK, "formulario.html", gin.H{
		"link": modelo.ObterLink(c.Param("id")),
	})
}

func construirLink(c *gin.Context) *modelo.Link {

	var privado bool
	if c.PostForm("inputPrivado") != "" {
		privado = true
	}

	id, err := strconv.Atoi(c.PostForm("inputId"))
	if err != nil {
		id = -1
	}

	return &modelo.Link{
		ID:          id,
		URL:         c.PostForm("inputUrl"),
		Titulo:      c.PostForm("inputTitulo"),
		Privado:     privado,
		DataCriacao: modelo.DataFormatada{time.Now()},
		Tags:        modelo.NovasTags(c.PostForm("inputTags")),
	}
}
