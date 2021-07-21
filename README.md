# Websockets

Alguns exemplos práticos de como implementarmos websockets utilizando Go ❤️.
Enviar mensagens e obter uma respostas de forma instantânea sem atualizar a página com famoso F5 legal não é. No passado não tão distante fazer isto era um custo para o backend e para o frontend porque utilizavamos somente o http puro para isto usando AJAX.

O WebSocket é uma tecnologia que surgiu em 2011 [RFC 6455](https://datatracker.ietf.org/doc/html/rfc6455) e permite comunicação bidirecional por canais full-duplex sobre um único soquete Transmission Control Protocol (TCP). Ele é projetado para ser executado em browsers e servidores web que suportem o HTML5. A API WebSocket está sendo padronizada pelo W3C e o protocolo WebSocket está sendo padronizado pelo IETF.

Websocket foi desenvolvido para ser implementado em browsers web e servidores web, mas pode ser usado por qualquer cliente ou aplicação servidor. O protocolo Websocket é um protocolo independente baseado em TCP. Sua única relação com o HTTP é que seu handshake é interpretado por servidores HTTP como uma requisição de Upgrade.

O WebSockets, que permitem abrir uma sessão interativa entre o navegador do usuário e um servidor. Os WebSockets permitem que um navegador envie mensagens a um servidor e receba respostas orientadas a eventos sem ter que consultar o servidor para obter uma resposta, lindão não é ? 😍

Por enquanto, WebSockets são a solução número um para construir aplicativos em tempo real: jogos online, mensageiros instantâneos, aplicativos de rastreamento, gráficos real time e assim por diante.


## Websocket Standard Library

Temos varias formas de construir um Websocket utilizando Go.
 - Gorilla
 - Stdlib (standard libray)
 - Gobwas
 - Gowebsockets
 - go-socket.io
 - Fiber (framework Web - utiliza fasthttp)
  
Em nosso repo temos StdLib e Fiber para apresentar o comportamento do Websocket.

Temos clients Go e HTML usando HTML5 para simularmos conneções e a comunicação entre eles.


