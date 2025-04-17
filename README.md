# GO-CLI

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/OzyKleyton/go-cli" alt="Go Version">
  <img src="https://img.shields.io/github/v/release/OzyKleyton/go-cli" alt="Release">
  <img src="https://img.shields.io/badge/license-MIT-blue" alt="License">
</p>

GO-CLI é uma ferramenta de linha de comando para inicialização rápida de projetos em Go com estrutura organizada.

## 📋 Pré-requisitos

- [Go 1.20+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## ⚡ Instalação Rápida

```bash
# Instalar o CLI
go install github.com/OzyKleyton/go-cli@latest

# Configurar PATH (Linux/macOS)
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.$(basename $SHELL)rc && source ~/.$(basename $SHELL)rc

# Verificar instalação
go-cli version
```

🛠️ Instalação Manual

```bash
git clone https://github.com/OzyKleyton/go-cli.git
cd go-cli
go build -o go-cli .
sudo mv go-cli /usr/local/bin  # Linux/macOS
```

🚀 Como Usar
Criar novo projeto

```bash
go-cli init meu-projeto
```

Estrutura gerada

```bash
meu-projeto/
├── cmd/
│ └── server/
│ └── main.go
├── config/
│ ├── config.go
│ └── db/
│ └── db.go
├── internal/
│ ├── model/
│ ├── repository/
│ ├── service/
│ ├── api/
│ │ ├── handler/
│ │ ├── router/
│ │ └── api.go
├── .env.example
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
└── Makefile
```

🔧 Solução de Problemas
Comando não encontrado

```bash
# Executar com caminho completo
$(go env GOPATH)/bin/go-cli --help

# Verificar instalação
ls $(go env GOPATH)/bin | grep go-cli
```

🤝 Contribuição
Faça um fork do projeto

Crie sua branch (git checkout -b feature/nova-feature)

Commit suas mudanças (git commit -m 'Adiciona nova feature')

Push para a branch (git push origin feature/nova-feature)

Abra um Pull Request

📄 Licença
Este projeto está licenciado sob a licença MIT - veja o arquivo LICENSE para detalhes.
