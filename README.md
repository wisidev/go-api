# REST API com Golang e GIN

## Introdução

Este projeto é uma API REST robusta desenvolvida em Golang utilizando o framework GIN. A API permite a criação e consulta de produtos armazenados em um banco de dados PostgreSQL. Tanto a aplicação quanto o banco de dados são executados dentro de containers Docker, garantindo maior portabilidade e escalabilidade.

## Funcionalidades

- Criar produtos
- Buscar todos os produtos salvos
- Buscar um produto por ID específico

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para desenvolver a API.
- **PostgreSQL**: Banco de dados relacional onde os produtos serão armazenados.
- **Docker**: Ferramenta para conteinerização da aplicação e banco de dados.
- **GIN**: Framework web para criação da API REST.

## Arquitetura

O projeto segue uma arquitetura limpa, separando claramente as responsabilidades:

1. **Controllers**: Responsáveis por definir as rotas e lidar com as requisições HTTP.
2. **Use Cases**: Contêm a lógica de negócio da aplicação.
3. **Repository**: Responsável pelas operações de acesso ao banco de dados.
4. **Model**: Define a estrutura dos dados utilizados pela aplicação.

## Configuração e Execução

### 1. Clonar o repositório

```sh
$ git clone https://github.com/wisidev/go-api.git
$ cd go-api
```

### 2. Construir e rodar os containers

```sh
$ docker compose up -d
```

Isso iniciará tanto a API quanto o banco de dados PostgreSQL dentro de containers Docker.

### 3. Testar a API

Usando o Postman, você pode testar os seguintes endpoints:

- **Criar um produto**:

```sh
POST /product
Content-Type: application/json
{
  "name": "Produto A",
  "price": 100
}
```

- **Buscar todos os produtos**:

```sh
GET /products
```

- **Buscar um produto por ID**:

```sh
GET /products/{id}
```

## Dependências

A aplicação utiliza dependências gerenciadas via `go.mod`, incluindo o framework GIN e a biblioteca para PostgreSQL:

```sh
require (
    github.com/gin-gonic/gin
    github.com/lib/pq
)
```

## Conclusão

Este projeto fornece uma base sólida para desenvolver APIs REST com Golang, utilizando boas práticas de arquitetura e conteinerização. Expande-se facilmente para novos casos de uso e integrações.

