# digital_wallet
Desafio Carteira Virtual | [Devgym](https://www.devgym.com.br) - Uma API de transferência de dinheiro como a do PicPay

## Tecnologias utilizadas:
```
uuid v1.3.0
testify v1.8.1
pq v1.10.7
godotenv v1.4.0
gin v1.8.2
```

## Como rodar o projeto?
Devemos criar um arquivo 2 arquivos, .env e .db.env com as seguintes configurações:

.env deve conter o `BASE_URL` para a conexão com o banco de dados
```
BASE_URL="postgres://username:password@host:port/dbname?sslmode=disable"
```

.db.env
```
POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_DB=dbname
```
Devemos rodar o comando `docker compose up -d` para subir o banco de dados

## Comandos utilizados:
```
go test ./test/...
go run ./...
```

## API

Buscar conta pelo id:
```
curl --location --request GET 'http://localhost:3000/accounts/a188a603-1725-47fb-9f6a-5b1ec0aa91cc'
```
Exemplo de retorno após buscar conta:
```
{
    "account_id": "a188a603-1725-47fb-9f6a-5b1ec0aa91cc",
    "balance": 5000,
    "transactions": [
        {
            "id": "7280a38e-64b9-4e70-a082-d6d32e40ff5d",
            "operation": "Credit",
            "amount": 5000,
            "date": "0001-01-01T00:00:00Z"
        }
    ]
}
```

Transferir:
```
curl --location --request POST 'http://localhost:3000/accounts/transfer' \
--data-raw '{
    "from": "a188a603-1725-47fb-9f6a-5b1ec0aa91cc",
    "to": "145695f3-7438-4911-a69d-51074b21c8ec",
    "amount": 300
}'
```

Exemplo de retorno após transferência:
```
{
    "account_id": "a188a603-1725-47fb-9f6a-5b1ec0aa91cc",
    "balance": 4700,
    "transactions": [
        {
            "id": "7280a38e-64b9-4e70-a082-d6d32e40ff5d",
            "operation": "Credit",
            "amount": 5000,
            "date": "0001-01-01T00:00:00Z"
        },
        {
            "id": "35dda4c4-9e7c-417c-850f-7779c1087eaf",
            "operation": "Debit",
            "amount": 300,
            "date": "0001-01-01T00:00:00Z"
        }
    ]
}

{
    "account_id": "145695f3-7438-4911-a69d-51074b21c8ec",
    "balance": 300,
    "transactions": [
        {
            "id": "d767dbed-aab1-4a08-ab9b-9dc41f7711e6",
            "operation": "Credit",
            "amount": 300,
            "date": "0001-01-01T00:00:00Z"
        }
    ]
}
```
