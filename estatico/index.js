
var listaLinks = new Vue({
    delimiters: ['{(', ')}'],
    el: '#listaLinks',
    data: {
        lista: [{id: 0, URL: 'Vegetables', Titulo: 'Teste foo', DataCriacao: '2018-03-23', Privado: false, Tags: ['aaa','bbb'] }],
        erro: "Me apague se não quiser ver erro!"
    }
});

axios.get('/api/links')
.then(function (response) {
    listaLinks.lista = response.data.links;
})
.catch(function (error) {
    console.error(error);
});
