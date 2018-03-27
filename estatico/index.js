
var listaLinks = new Vue({
    delimiters: ['{(', ')}'],
    el: '#listaLinks',
    data: {
        lista: [],
        erro: "Me apague se não quiser ver erro!",
        pagina: 0
    },
    methods: {
        obterMaisLinks: function() {
            listaLinks.pagina++;
            axios.get('/api/links/' + listaLinks.pagina)
            .then(function (response) {
                if(response.data.links == null) {
                    alert('Ops, não há mais links para exibir!');
                    return;
                }
                listaLinks.lista.push(...response.data.links);
            })
            .catch(function (error) {
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
    console.error(error);
});
