# 💰 Finance API

> 🚧 **Status:** Em desenvolvimento

Uma API REST desenvolvida em **Go (Golang)** para gerenciamento financeiro pessoal.

Este projeto está sendo desenvolvido com o objetivo de aprender a linguagem Go através da construção de uma aplicação real, aplicando boas práticas de Engenharia de Software, arquitetura em camadas, documentação e organização de projetos.

---

# 🎯 Objetivos

Este projeto não tem como objetivo apenas construir uma API.

O principal foco é aprender Go de forma prática, utilizando conceitos encontrados em aplicações reais.

Durante o desenvolvimento serão estudados:

- Organização de projetos Go
- Modelagem de domínio
- Structs
- Methods
- Receivers
- Ponteiros
- Tratamento de erros
- Interfaces
- APIs REST
- JSON
- Banco de Dados
- PostgreSQL
- Docker
- JWT
- Testes
- Concorrência (Goroutines e Channels)

---

# ✨ Funcionalidades Planejadas

- Cadastro de usuários
- Login com autenticação JWT
- Cadastro de receitas
- Cadastro de despesas
- Categorias personalizadas
- Histórico financeiro
- Consulta de transações
- Dashboard financeiro
- Relatórios
- Exportação de dados

---

# 🏛️ Arquitetura

O projeto segue uma arquitetura em camadas.

```text
Cliente
    │
HTTP Request
    │
    ▼
Handler
    │
    ▼
Service
    │
    ▼
Repository
    │
    ▼
Database (PostgreSQL)
```

Toda a documentação visual da arquitetura será mantida em:

```text
docs/
└── architecture/
```

Os diagramas foram desenvolvidos utilizando **Excalidraw**.

---

# 📂 Estrutura do Projeto

```text
finance-api/
│
├── cmd/
├── configs/
├── docs/
│   └── architecture/
│       ├── system.excalidraw
│       ├── domain.excalidraw
│       ├── request-flow.excalidraw
│       └── database.excalidraw
│
├── internal/
│   └── transaction/
│       └── transaction.go
│
├── pkg/
├── scripts/
│
├── .gitignore
├── go.mod
├── main.go
└── README.md
```

---

# 📦 Tecnologias

- Go
- PostgreSQL *(em desenvolvimento)*
- Docker *(em desenvolvimento)*
- JWT *(em desenvolvimento)*
- Git
- GitHub
- Excalidraw

---

# 📚 Conceitos Estudados

## ✅ Go

- [x] Packages
- [x] Modules
- [x] Structs
- [x] Methods
- [x] Receivers
- [x] Ponteiros
- [x] Error Handling

## 🚧 Próximos Estudos

- [ ] Interfaces
- [ ] JSON
- [ ] HTTP
- [ ] Context
- [ ] Middleware
- [ ] PostgreSQL
- [ ] Docker
- [ ] Testes
- [ ] Goroutines
- [ ] Channels

---

# 📌 Modelo de Domínio

A primeira entidade modelada foi:

```text
Transaction
│
├── ID
├── Description
├── Amount
├── Category
├── Merchant
└── Date
```

Novas entidades serão adicionadas durante o desenvolvimento, como:

```text
User

Wallet

Category

Transaction
```

---

# 🚀 Como executar

Clone o repositório:

```bash
git clone https://github.com/NickolasSantosCremasco/finance-api.git
```

Entre na pasta do projeto:

```bash
cd finance-api
```

Execute:

```bash
go run .
```

---

# 🛣️ Roadmap

## ✅ Etapa 1 — Configuração Inicial

- [x] Criar módulo Go
- [x] Configurar GitHub
- [x] Organizar estrutura do projeto
- [x] Criar documentação inicial

---

## ✅ Etapa 2 — Fundamentos da Linguagem

- [x] Structs
- [x] Methods
- [x] Receivers
- [x] Ponteiros
- [x] Tratamento de erros
- [x] Modelagem da entidade Transaction

---

## 🚧 Etapa 3 — API REST

- [ ] HTTP Server
- [ ] Rotas
- [ ] Handlers
- [ ] JSON
- [ ] Validação

---

## 🚧 Etapa 4 — Banco de Dados

- [ ] PostgreSQL
- [ ] SQL
- [ ] Repository Pattern
- [ ] Migrations

---

## 🚧 Etapa 5 — Autenticação

- [ ] Login
- [ ] Cadastro
- [ ] JWT
- [ ] Middleware

---

## 🚧 Etapa 6 — Docker

- [ ] Dockerfile
- [ ] Docker Compose

---

## 🚧 Etapa 7 — Testes

- [ ] Unit Tests
- [ ] Integration Tests
- [ ] Benchmarks

---

## 🚧 Etapa 8 — Concorrência

- [ ] Goroutines
- [ ] Channels
- [ ] Context
- [ ] Worker Pool

---

## 🚧 Etapa 9 — Deploy

- [ ] Deploy da API
- [ ] CI/CD
- [ ] Documentação da API

---

# 📅 Diário de Desenvolvimento

## Aula 01

- Configuração do ambiente Go
- Criação do módulo
- Organização inicial do projeto
- GitHub

## Aula 02

- Structs
- Modelagem da entidade Transaction

## Aula 03

- Methods
- Receivers
- Ponteiros
- Error Handling

> Este diário será atualizado conforme o projeto evolui.

---

# 📖 Aprendizados

Este projeto está sendo utilizado para estudar:

- Engenharia de Software
- Arquitetura em Camadas
- Boas práticas em Go
- Organização de projetos
- Modelagem de domínio
- Documentação técnica
- Git e GitHub
- Desenvolvimento de APIs REST

---

# 🤝 Contribuição

Este projeto está sendo desenvolvido para fins de aprendizado.

Sugestões e melhorias são sempre bem-vindas.

---

# 📖 Licença

Este projeto é de código aberto e está sendo desenvolvido exclusivamente para fins educacionais e de estudo da linguagem Go.