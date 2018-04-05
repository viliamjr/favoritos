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
	auth := gin.BasicAuth(gin.Accounts{usuario: senha})

	raiz := r.Group("/", auth)
	{
		raiz.GET("/", Raiz)
		raiz.GET("/formulario", Formulario)
		raiz.GET("/editar/:id", Editar)
		raiz.GET("/tag/:tag", Pagina)
	}

	api := r.Group("/api", auth)
	{
		api.GET("/links/:pag", Links)
		api.GET("/remover/:id", Remover)
		api.POST("/salvar", Salvar)
	}
}

// Links define a rota da página inicial
func Links(c *gin.Context) {

	pag, _ := strconv.Atoi(c.Param("pag"))

	c.JSON(http.StatusOK, gin.H{
		"links": modelo.ObterPagina(pag, true),
	})
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

	tag := c.Param("tag")
	log.Println("TAAaaaaaaaag: " + tag)

	pag, _ := strconv.Atoi(c.Param("pag"))
	c.HTML(http.StatusOK, "favoritos.html", gin.H{
		"proxPagina": pag + 1,
		"links":      modelo.ObterPaginaPorTag(pag, true, tag),
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

	if c.PostForm("id") != "" {

		modelo.AtualizarLink(link)
		c.JSON(http.StatusOK, gin.H{
			"erro": nil,
			"msg":   "Link atualizado!!",
		})
	} else {

		erro := modelo.NovoLink(link)
		msg := "Link salvo com sucesso!"
		if erro != nil {
			msg = "OPA! Esse link já foi cadastrado O.o"
			log.Printf("Erro ao inserir novo link: %v\n", erro)
		}
		c.JSON(http.StatusOK, gin.H{
			"erro": erro,
			"msg":   msg,
		})
	}
}

// Remover define a rota para a remoção de um link
func Remover(c *gin.Context) {

	modelo.RemoverLink(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"erro": nil,
		"msg":   "Link removido!",
	})
}

// Editar define a rota para exibir os dados de um link no formulário
func Editar(c *gin.Context) {

	c.HTML(http.StatusOK, "formulario.html", gin.H{
		"link": modelo.ObterLink(c.Param("id")),
	})
}

func construirLink(c *gin.Context) *modelo.Link {

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		id = -1
	}

	return &modelo.Link{
		ID:          id,
		URL:         c.PostForm("inputUrl"),
		Titulo:      c.PostForm("inputTitulo"),
		Privado:     (c.PostForm("Privado") == "true"),
		DataCriacao: modelo.DataFormatada{time.Now()},
		Tags:        modelo.NovasTags(c.PostForm("inputTags")),
	}
}
