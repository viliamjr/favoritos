
var listaLinks = new Vue({
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
            listaLinks.pagina++;
            obterLinks();
        },
        salvarNovoLink: function() {
            let link = {
                inputId
            };
            axios.post('/api/salvar', link)
            .then(function (response) {
                if(response.data.links == null) {
                    let msg = 'Ops, não há mais links para exibir!';
                    listaLinks.erroLista = msg;
                    return;
                }
                listaLinks.lista.push(...response.data.links);
            })
            .catch(function (error) {
                listaLinks.erroLista = error;
                console.error(error);
            });
        }
    }
});

function obterLinks() {
    axios.get('/api/links/' + listaLinks.pagina)
    .then(function (response) {
        if(response.data.links == null) {
            let msg = 'Ops, não há mais links para exibir!';
            listaLinks.erroLista = msg;
            return;
        }
        listaLinks.lista.push(...response.data.links);
    })
    .catch(function (error) {
        listaLinks.erroLista = error;
        console.error(error);
    });
}

obterLinks();
