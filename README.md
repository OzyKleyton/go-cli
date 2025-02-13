# GO-CLI

GO-CLI é uma ferramenta de linha de comando para inicialização rápida de projetos em Go, seguindo uma estrutura organizada e padrão.

Ainda estou desenvolvendo e aos poucos irei melhorando o projeto.

## Requisitos

Antes de utilizar o GO-CLI, verifique se possui os seguintes softwares instalados:

- **Docker**: Para facilitar a execução do ambiente de desenvolvimento.
- **Docker Compose**: Para gerenciar os containers da aplicação.

O Go **não é necessário** para rodar os comandos do `go-cli`, pois ele será compilado e executado dentro do container Docker.

## Instalação

Você pode instalar o GO-CLI de duas formas:

### 1. Instalando via `go install`  

```sh
go install github.com/OzyKleyton/go-cli@latest
```

### 2. Clonando o repositório

```sh
git clone https://github.com/OzyKleyton/go-cli.git
cd go-cli
```

## Uso

Após a instalação, você pode rodar o seguinte comando para iniciar um novo projeto:

```sh
go-cli init nome-do-projeto
```

Isso criará uma estrutura de pastas e arquivos padrão para seu projeto em Go.

## Estrutura Criada

O comando `go-cli init` gerará a seguinte estrutura:

```
nome-do-projeto/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   ├── config.go
│   └── db/
│       └── db.go
├── internal/
│   ├── model/
│   ├── repository/
│   ├── service/
│   ├── api/
│   │   ├── handler/
│   │   ├── router/
│   │   └── api.go
├── .env.example
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
└── makefile
```

### Rodando o projeto

Após os comando de criação vamos rodar os comando `make` para rodar os scrips do docker-compose.

`make up` para começar a buildar a imagem docker.

`make start` para executar o container.
