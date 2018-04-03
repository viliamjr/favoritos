
var modelo = new Vue({
    delimiters: ['{(', ')}'],
    el: '#listaLinks',
    data: {
        lista: [],
        erroForm: null,
        erroLista: null,
        pagina: 0,
        link: {}
    },
    methods: {
        obterMaisLinks: function() {
            modelo.pagina++;
            $.get('/api/links/' + + modelo.pagina, function( data ) {
                if(data.links == null) {
                    let msg = 'Ops, não há mais links para exibir!';
                    modelo.erroLista = msg;
                    return;
                }
                modelo.lista.push(...response.data.links);
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
            link.inputUrl = modelo.link.inputUrl;
            link.inputTitulo = modelo.link.inputTitulo;
            link.inputTags = modelo.link.inputTags;
            link.inputPrivado = modelo.link.Privado;

            $.post("/api/salvar", link, function( data ) {
                if(data.erro != null) {
                    modelo.erroForm = data.erro;
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
			let url = input.val().trim().toLowerCase();
			input.focus();
			if( (url.startsWith("http://") || url.startsWith("https://") || url.startsWith("ftp://")) ) {
                return null;
            }
            return "URL inválida!";
        }
    }
});

function obterLinks() {
    $.get("/api/links/0", function( data ) {
        modelo.lista = data.links;
    }).fail(function() {
        modelo.erroLista = "Opss, algo deu errado! Log registrado no console.";
    });
}

obterLinks();
