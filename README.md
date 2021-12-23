# Client Service
### by Dmitry Taranov
<br></br>
#### Client Service GitHub Repo: [Client Service](https://github.com/TaranovDmitry/ClientService)
#### Domain Service GitHub Repo: [Domain Service](https://github.com/TaranovDmitry/DomainService)

### Prerequisites:
- docker
- running Postgres
- running Domain Service 

### To start please use the following commands:
1. Create image: ``docker build -t client-service``
2. Run docker container: ``docker run -d -p 8080:8080 --name client client-service``

