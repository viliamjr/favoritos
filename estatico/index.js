
var modelo = new Vue({
    delimiters: ['{(', ')}'],
    el: '#listaLinks',
    data: {
        lista: [],
        pagina: 0,
        erroLista: null,
        filtroTag: null,
        erroForm: null,
        link: {},
        busca: null
    },
    methods: {
        obterMaisLinks: function() {
            modelo.pagina++;
            let url = '/api/links/' + modelo.pagina;
            if(modelo.filtroTag != null) {
                url += '/' + modelo.filtroTag;
            }
            $.get(url, function( data ) {
                if(data.links == null) {
                    let msg = 'Ops, não há mais links para exibir!';
                    modelo.erroLista = msg;
                    return;
                }
                modelo.lista.push(...data.links);
            }).fail(function() {
                modelo.erroLista = "Opss, algo deu errado! Log registrado no console.";
            });
        },

        salvarNovoLink: function() {
            
            let erro = this.dadosInvalidos();
            if(erro != null) {
                modelo.erroForm = erro;
                return;
            }

            var link = {};
            link.id = modelo.link.inputId;
            link.inputUrl = modelo.link.inputUrl;
            link.inputTitulo = modelo.link.inputTitulo;
            link.inputTags = modelo.link.inputTags;
            link.Privado = modelo.link.Privado;

            $.post("/api/salvar", link, function( data ) {
                if(data.erro != null) {
                    modelo.erroForm = data.msg;
                    return;
                }
                modelo.erroForm = data.msg;
                modelo.link = {};
                obterLinks();
            }).fail(function() {
                modelo.erroForm = "Opss, algo deu errado! Log registrado no console.";
            });
        },

        dadosInvalidos: function() {
			let input = $('[name=inputUrl]');
			let dados = input.val().trim().toLowerCase();
			if( ! (dados.startsWith("http://") || dados.startsWith("https://") || dados.startsWith("ftp://")) ) {
                input.focus();
                return "URL inválida!";
            }
            if(campoVazio('inputTitulo')) {
                return "Título não pode ser vazio!";
            }
            if(campoVazio('inputTags')) {
                return "Deve ser definada ao menos uma Tag!";
            }
            return null;
        },

        removerLink: function(id) {
            if( confirm('Você deseja excluir este link?') ) {
                $.get('/api/remover/' + id, function( data ) {
                    modelo.erroLista = data.msg;
                    obterLinks();
                }).fail(function() {
                    modelo.erroLista = "Opss, algo deu errado! Log registrado no console.";
                });
            }
        },

        editarLink: function(id) {
            modelo.link = {};
            let link = this.obterLink(id);
            if(link == null) {
                modelo.erroLista = "Opss, algo deu errado! Não foi possível obter link para edição.";
                return;
            }
            modelo.link.inputUrl = link.URL;
            modelo.link.inputTitulo = link.Titulo;
            modelo.link.inputTags = link.Tags.toString();
            modelo.link.Privado = link.Privado;
            modelo.link.inputDataCriacao = link.DataCriacao;
            modelo.link.inputId = link.id;
            window.scrollTo(0, 0);
        },

        obterLink: function(id) {
            let encontrado = null;
            modelo.lista.forEach(link => {
                if(link.id == id) {
                    encontrado = link;
                }
            });
            return encontrado;
        },

        filtarTag: function(tag) {
            modelo.erroLista = null;
            if(modelo.filtroTag != tag) {
                modelo.pagina = 0;
                modelo.filtroTag = tag;
            }
            modelo.lista = [];

            $.get('/api/links/' + modelo.pagina + '/' + modelo.filtroTag, function( data ) {
                if(data.links == null) {
                    let msg = 'Ops, não há links para este filtro!';
                    modelo.erroLista = msg;
                    return;
                }
                modelo.lista.push(...data.links);
                modelo.busca = null;
            }).fail(function() {
                modelo.erroLista = "Opss, algo deu errado! Log registrado no console.";
            });
        },

        removerFiltroDeTag: function() {
            modelo.filtroTag = null;
            modelo.busca = null;
            obterLinks();
        }
    }
});

function campoVazio(nome) {
    input = $('[name='+nome+']');
    dados = input.val().trim();
    if(dados == "") {
        input.focus();
        return true;
    }
    return false;
}

function obterLinks() {
    modelo.erroLista = null;
    $.get("/api/links/0", function( data ) {
        modelo.lista = data.links;
    }).fail(function() {
        modelo.erroLista = "Opss, algo deu errado! Log registrado no console.";
    });
}

obterLinks();
