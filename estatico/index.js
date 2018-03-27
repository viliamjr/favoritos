
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
            obterLinks();
        },
        salvarNovoLink: function() {
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
                modelo.lista.unshift(modelo.link);
                modelo.link = {}; //XXX funciona para zerar?!
            })
            .catch(function (error) {
                modelo.erroForm = error;
                console.error(error);
            });
        }
    }
});

function obterLinks() {
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
        modelo.erroLista = error;
        console.error(error);
    });
}

obterLinks();
