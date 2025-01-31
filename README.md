# CLI-GO

CLI-GO é uma ferramenta de linha de comando para inicialização rápida de projetos em Go, seguindo uma estrutura organizada e padrão.

## Requisitos

Antes de utilizar o CLI-GO, verifique se possui os seguintes softwares instalados:

- **Docker**: Para facilitar a execução do ambiente de desenvolvimento.
- **Docker Compose**: Para gerenciar os containers da aplicação.

O Go **não é necessário** para rodar os comandos do `cli-go`, pois ele será compilado e executado dentro do container Docker.

## Instalação

Você pode instalar o CLI-GO de duas formas:

### 1. Instalando via `go get`  

```sh
go get github.com/OzyKleyton/project-go-cli@latest
```

### 2. Clonando o repositório

```sh
git clone https://github.com/OzyKleyton/project-go-cli.git
cd cli-go
```

## Uso

Após a instalação, você pode rodar o seguinte comando para iniciar um novo projeto:

```sh
cli-go init nome-do-projeto
```

Isso criará uma estrutura de pastas e arquivos padrão para seu projeto em Go.

## Estrutura Criada

O comando `cli-go init` gerará a seguinte estrutura:

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

##



