
var listaLinks = new Vue({
    delimiters: ['{(', ')}'],
    el: '#listaLinks',
    data: {
        lista: [],
        erroForm: "Me apague se não quiser ver erro!",
        erroLista: "O mesmo, me apage para sumir!",
        pagina: 0
    },
    methods: {
        obterMaisLinks: function() {
            listaLinks.pagina++;
            axios.get('/api/links/' + listaLinks.pagina)
            .then(function (response) {
                if(response.data.links == null) {
                    let msg = 'Ops, não há mais links para exibir!';
                    listaLinks.erroLista = msg;
                    alert(msg);
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

axios.get('/api/links/' +  + listaLinks.pagina)
.then(function (response) {
    listaLinks.lista = response.data.links;
})
.catch(function (error) {
    listaLinks.erroLista = error;
    console.error(error);
});
