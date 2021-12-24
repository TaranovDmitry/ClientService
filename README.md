# Client Service
### by Dmitry Taranov
<br></br>
#### Client Service GitHub Repo: [Client Service](https://github.com/TaranovDmitry/ClientService)
#### Domain Service GitHub Repo: [Domain Service](https://github.com/TaranovDmitry/DomainMicroservice)

### Prerequisites:
- domain service docker image
- migration tool

### How to start Client Service with docker-compose
1. Start docker-compose: ``docker-compose up``
2. Execute DB Scheme Migration: ``migrate -path /Users/dmitrytaranov/GolandProjects/DomainMicroservice/schema -database 'postgresql://postgres:12345@localhost:5432/postgres?sslmode=disable' up``

### API Request examples
1. Get all ports: <br></br>``curl -v --location --request GET 'http://localhost:8080/client/v1/ports'``
1. Upload ports: <br></br>``curl --location --request POST 'localhost:8080/client/v1/ports' --header 'Content-Type: multipart/form-data' --form 'file=@"/Users/dmitrytaranov/GolandProjects/ClientService/ports.json"'``