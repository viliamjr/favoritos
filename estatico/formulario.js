
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
