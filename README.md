# Event-driven

![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)


Esse projeto tem como objetivo o estudo de implementação de um sistema de entrega de conteúdo por demanda.

> :construction: Projeto em construção :construction:

## Execução

## 🛠️ Rodar o projeto

Executando e subindo o projeto de "frontend", o script sobre um server Websocket onde manterá conexão com o client.

** Cada client é identificado pelo parametro "session"

```bash
./app 
```
#### Client

Basta acessar a pagina "index.html" dentro da pasta client, preencher o campo "session" e clicar no botão

#### Testar conexão com o client

Requisição local: http://localhost:8080/send

Body da requisição

```bash
{
  from:{session do enviador}
  to: {"session do user destino"},
  message: "Message"
} 
```

