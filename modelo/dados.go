package modelo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// ListaTags representa os tags de marcação dos link
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

// NovasTags cria uma nova lista de tags com base numa string que usa
// vírgulas para separar as tags.
func NovasTags(s string) *ListaTags {
	tags := make(ListaTags, 0)
	for _, i := range strings.Split(s, ",") {
		tags = append(tags, i)
	}
	return &tags
}

// DataFormatada representa a data a ser armazenada no BD.
type DataFormatada struct {
	time.Time
}

func (d DataFormatada) String() string {
	return fmt.Sprintf("%02d/%02d/%d", d.Day(), d.Month(), d.Year())
}

// Link representa o link a ser salvo.
type Link struct {
	ID          int `json:"id"`
	URL         string
	Titulo      string
	Privado     bool
	DataCriacao DataFormatada
	Tags        *ListaTags
}

var db *sql.DB

// CarregarBanco carrega a variável de pacote 'db' com o banco de dados.
func CarregarBanco(caminhoBanco *string) *sql.DB {

	// configurando banco de dados
	var err error
	db, err = sql.Open("sqlite3", *caminhoBanco)
	if err != nil {
		log.Fatal("Erro ao abrir o banco: ", err)
	}

	CriarBanco()

	return db
}

// CriarBanco cria o esquema do BD, caso o banco ainda não exista.
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

// NovoLink cria um novo link no BD.
func NovoLink(link *Link) error {

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Erro ao criar transação: %v", err)
	}

	stmt, err := tx.Prepare("insert into link values(?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Erro ao preparar a query de insert: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(link.URL, link.Titulo, link.Tags.String(), link.DataCriacao.Time.Unix(), link.Privado)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Erro ao executar um insert no banco: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Erro ao commitar transação de insert: %v", err)
	}

	return nil
}

// AtualizarLink atualiza o link no BD.
func AtualizarLink(link *Link) {

	_, err := db.Exec(`update link set url=?, titulo=?, tags=?, privado=? where rowid = ?;`,
		link.URL, link.Titulo, link.Tags.String(), link.Privado, link.ID)

	if err != nil {
		log.Fatal("Erro no update de link: ", err)
	}
}

// RemoverLink o link do BD.
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

// ObterPagina lista os links conforme algoritmo de paginação.
func ObterPagina(pag int, listarPrivados bool) []*Link {
	return ObterPaginaPorTermos(pag, listarPrivados, "")
}

// ObterPaginaPorTermos lista os links filtrando por termo (tags ou títulos) e conforme algoritmo de paginação.
func ObterPaginaPorTermos(pag int, listarPrivados bool, termos string) []*Link {

	var encontrados []*Link
	offset := 20

	cmdSQL := "select rowid,url,titulo,tags,data_criacao,privado from link where privado = 0 and "
	if listarPrivados {
		cmdSQL = "select rowid,url,titulo,tags,data_criacao,privado from link where "
	}

	cmdSQL += prepararTermos(termos)

	cmdSQL += " order by data_criacao desc limit ?,?;"

	rows, err := db.Query(cmdSQL, (pag * offset), offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		link := &Link{}
		var tags string
		rows.Scan(&link.ID, &link.URL, &link.Titulo, &tags, &(link.DataCriacao.Time), &link.Privado)
		link.Tags = NovasTags(tags)
		encontrados = append(encontrados, link)
	}

	return encontrados
}

// ObterLink retorna link a partir de seu ID.
func ObterLink(id string) *Link {

	link := &Link{}

	rows, err := db.Query("select rowid,url,titulo,tags,data_criacao,privado from link where rowid = ?;", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		var tags string
		rows.Scan(&link.ID, &link.URL, &link.Titulo, &tags, &(link.DataCriacao.Time), &link.Privado)
		link.Tags = NovasTags(tags)
	}

	return link
}

// ProcurarLinkPorTag busca links com base em uma tag
func ProcurarLinkPorTag(tag string) []*Link {

	var encontrados []*Link

	//TODO popula slice com itens encontrados

	return encontrados
}

// prepararTermos retorna trecho SQL para filtar tags e titulos.
func prepararTermos(termos string) string {

	lista := strings.Split(termos, ",")
	tam := len(lista)

	var sql string
	for i := 0; i < tam; i++ {
		sql += "("

		// filtro por tags
		sql += " (tags like '" + strings.TrimSpace(lista[i]) + "%' "
		sql += " or tags like '%," + strings.TrimSpace(lista[i]) + "%')"
		// filtro por titulo
		sql += " or titulo like '%" + strings.TrimSpace(lista[i]) + "%'"

		sql += ")"
		if i < (tam - 1) {
			sql += " and "
		}
	}

	return sql
}
