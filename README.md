# RODAR O PROJETO: Clean Architecture

- Subir o banco de dados:
    docker-compose up -d --build

- Executar migrations:
    make migration

- Subir o projeto GO:
    cd cmd/ordersystem && go run main.go wire_gen.go


# DESAFIO: Clean Architecture

Agora é a hora de botar a mão na massa. Pra este desafio, você precisará criar o usecase de listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)

- Service ListOrders com GRPC

- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.
