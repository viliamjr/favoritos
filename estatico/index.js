
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
            axios.get('/api/links/' + modelo.pagina)
            .then(function (response) {
                if(response.data.links == null) {
                    let msg = 'Ops, não há mais links para exibir!';
                    modelo.erroLista = msg;
                    return;
                }
                modelo.lista.push(...response.data.links);
            })
            .catch(function (error) {
                modelo.erroLista = "Opss, algo deu errado! Log registrado no console.";
                console.error(error);
            });
        },

        salvarNovoLink: function() {
            
            let erro = this.dadosInvalidos();
            if(erro != null) {
                modelo.erroForm = erro;
                return;
            }

            var link = new URLSearchParams();
            link.append('inputUrl', modelo.link.inputUrl);
            link.append('inputTitulo', modelo.link.inputTitulo);
            link.append('inputTags', modelo.link.inputTags);
            link.append('inputPrivado', modelo.link.Privado);
            
            axios.post('/api/salvar', link)
            .then(function (response) {
                if(response.data.erro != null) {
                    modelo.erroForm = response.data.erro;
                    return;
                }
                modelo.erroForm = response.data.msg;
                modelo.link = {};
                obterLinks();
            })
            .catch(function (error) {
                modelo.erroForm = "Opss, algo deu errado! Log registrado no console.";
                console.error(error);
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
    axios.get('/api/links/0')
    .then(function (response) {
        modelo.lista = response.data.links;
    })
    .catch(function (error) {
        modelo.erroLista = "Opss, algo deu errado! Log registrado no console.";
        console.error(error);
    });
}

obterLinks();
