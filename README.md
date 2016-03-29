## Código para exibir o form de add link:

javascript:(function(e,t){
  var url=encodeURIComponent(e.location.href);
  var n=e.document;
  var s=n.createElement("iframe");
  function sair(e) {n.body.removeChild(s);}
  s.id="iframeNovoLink";
  url += "&titulo=" + encodeURIComponent(n.title);
  s.src="http://127.0.0.1:3000/formulario?url="+url; s.style.position="fixed";s.style.top="0";s.style.left="0";s.style.height="40%";s.style.width="50%";
  s.style.zIndex="16777270";s.style.border="none";s.style.visibility="hidden";
  s.onload=function(){this.style.visibility="visible";};
  n.body.appendChild(s);
  var o=e.addEventListener?"addEventListener":"attachEvent";var u=o=="attachEvent"?"onmessage":"message";e[o](u,sair,false);
})(window)
