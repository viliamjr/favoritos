### Código a ser usado como `link` na barra de favoritos do navegador

```javascript

javascript:(function(e,t){
  var url=encodeURIComponent(e.location.href);
  var n=e.document;
  var s=n.createElement("iframe");
  function sair(e) {n.body.removeChild(s);}
  s.id="iframeNovoLink";
  url += "&titulo=" + encodeURIComponent(n.title);
  s.src="https://127.0.0.1:8080/formulario?url="+url; s.style.position="fixed";s.style.top="0";s.style.left="0";s.style.height="40%";s.style.width="50%";
  s.style.zIndex="16777270";s.style.border="none";s.style.visibility="hidden";
  s.onload=function(){this.style.visibility="visible";};
  n.body.appendChild(s);
  var o=e.addEventListener?"addEventListener":"attachEvent";var u=o=="attachEvent"?"onmessage":"message";e[o](u,sair,false);
})(window)

```

Obs.: o acesso ao formulário a partir da barra de favoritos do browser pode não funcionar dependendo das configurações de segurança (CORS?!) do servidor HTTP onde o site alvo está hospedado.

### Rodando servidor em modo de desenvolvimento

Utilizando a ferramenta **gin**, execute:

    gin --logPrefix=monitor -i -p 3000 -a 8083 run -e :8083 -u admin -s 123 -nao-https

O servidor será iniciado sem suporte a HTTPS e com login admin / 123.

**Obs**: devido à dependência com o `go-sqlite3`, para compilar o projeto é necessário ligar o CGO e ter o gcc instalado. Por exemplo, no Windows é necessário:

```powershell
$env:Path="$env:Path;E:\mingw64\bin"
$env:CGO_ENABLED=1
```