
var modeloForm = new Vue({
    delimiters: ['{(', ')}'],
    el: '#modeloForm',
    data: {
        link: {},
        erro: null
    },
    methods: {
        salvarNovoLink: function() {
            
            let erro = this.dadosInvalidos();
            if(erro != null) {
                modeloForm.erro = erro;
                return;
            }

            var link = {};
            link.id = modeloForm.link.inputId;
            link.inputUrl = modeloForm.link.inputUrl;
            link.inputTitulo = modeloForm.link.inputTitulo;
            link.inputTags = modeloForm.link.inputTags;
            link.Privado = modeloForm.link.Privado;

            $.post("/api/salvar", link, function( data ) {
                if(data.erro != null) {
                    modeloForm.erro = data.msg;
                    return;
                }
                modeloForm.erro = data.msg;
                modeloForm.link = {};
                obterLinks();
            }).fail(function() {
                modeloForm.erro = "Opss, algo deu errado! Log registrado no console.";
            });
        },

        dadosInvalidos: function() {
			let input = $('[name=inputUrl]');
			let dados = input.val().trim().toLowerCase();
			if( ! (dados.startsWith("http://") || dados.startsWith("https://") || dados.startsWith("ftp://")) ) {
                input.focus();
                return "URL inválida!";
            }
            if(this.campoVazio('inputTitulo')) {
                return "Título não pode ser vazio!";
            }
            if(this.campoVazio('inputTags')) {
                return "Deve ser definada ao menos uma Tag!";
            }
            return null;
        },

        campoVazio: function(nome) {
            input = $('[name='+nome+']');
            dados = input.val().trim();
            if(dados == "") {
                input.focus();
                return true;
            }
            return false;
        }
    }
});

var modeloLinks = new Vue({
    delimiters: ['{(', ')}'],
    el: '#modeloLinks',
    data: {
        lista: [],
        pagina: 0,
        erro: null,
        filtroTag: null,
        busca: null
    },
    methods: {
        obterMaisLinks: function() {
            modeloLinks.pagina++;
            let url = '/api/links/' + modeloLinks.pagina;
            if(modeloLinks.filtroTag != null) {
                url += '/' + modeloLinks.filtroTag;
            }
            $.get(url, function( data ) {
                if(data.links == null) {
                    let msg = 'Ops, não há mais links para exibir!';
                    modeloLinks.erro = msg;
                    return;
                }
                modeloLinks.lista.push(...data.links);
            }).fail(function() {
                modeloLinks.erro = "Opss, algo deu errado! Log registrado no console.";
            });
        },

        removerLink: function(id) {
            if( confirm('Você deseja excluir este link?') ) {
                $.get('/api/remover/' + id, function( data ) {
                    modeloLinks.erro = data.msg;
                    obterLinks();
                }).fail(function() {
                    modeloLinks.erro = "Opss, algo deu errado! Log registrado no console.";
                });
            }
        },

        editarLink: function(id) {
            modeloForm.link = {};
            let link = this.obterLink(id);
            if(link == null) {
                modeloLinks.erro = "Opss, algo deu errado! Não foi possível obter link para edição.";
                return;
            }
            modeloForm.link.inputUrl = link.URL;
            modeloForm.link.inputTitulo = link.Titulo;
            modeloForm.link.inputTags = link.Tags.toString();
            modeloForm.link.Privado = link.Privado;
            modeloForm.link.inputDataCriacao = link.DataCriacao;
            modeloForm.link.inputId = link.id;
            window.scrollTo(0, 0);
        },

        obterLink: function(id) {
            let encontrado = null;
            modeloLinks.lista.forEach(link => {
                if(link.id == id) {
                    encontrado = link;
                }
            });
            return encontrado;
        },

        filtarTag: function(tag) {
            modeloLinks.erro = null;
            if(modeloLinks.filtroTag != tag) {
                modeloLinks.pagina = 0;
                modeloLinks.filtroTag = tag;
            }
            modeloLinks.lista = [];

            $.get('/api/links/' + modeloLinks.pagina + '/' + modeloLinks.filtroTag, function( data ) {
                if(data.links == null) {
                    let msg = 'Ops, não há links para este filtro!';
                    modeloLinks.erro = msg;
                    return;
                }
                modeloLinks.lista.push(...data.links);
                modeloLinks.busca = null;
            }).fail(function() {
                modeloLinks.erro = "Opss, algo deu errado! Log registrado no console.";
            });
        },

        removerFiltroDeTag: function() {
            modeloLinks.filtroTag = null;
            modeloLinks.busca = null;
            obterLinks();
        }
    }
});

function obterLinks() {
    modeloLinks.erro = null;
    $.get("/api/links/0", function( data ) {
        modeloLinks.lista = data.links;
    }).fail(function() {
        modeloLinks.erro = "Opss, algo deu errado! Log registrado no console.";
    });
}

obterLinks();
