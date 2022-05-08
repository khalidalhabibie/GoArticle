# GoArticle

GoArticle adalah aplikasi yang berfungsi membuat article dan menampilkan dengan pagination.

### This app has :

- RESTful endpoint for asset's CR operation
- JSON formatted response

&nbsp;

### Tech Stack used to build this app :

- [GoFiber](https://gofiber.io/) web framework<br/>
- [PostgreSQL](https://www.postgresql.org/) as database<br/>
- [gorm](https://gorm.io/index.html) as ORM library <br/>
- [Redis](https://redis.io/) as cache <br/>
- [Docker](https://www.docker.com/)as containerization platform <br/>
- [Docker Compose](https://docs.docker.com/compose/) as container orchestration frameworks.<br/>
- [Prometheus](https://prometheus.io/) monitoring and alerting<br/>
- [Grafana](https://grafana.com/) for to compose observability dashboards with everything from Prometheus<br/>
- [swagger](https://swagger.io/) as Documentation<br/>

&nbsp;

### Requirement :

- [go](https://go.dev/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [go migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) (CLI)
- [swag-go](https://github.com/swaggo/swag)(CLI)(for development process, to regenerate)(optional)
- [mockery](https://github.com/vektra/mockery)(CLI) (for development process, to mock new interface)(optional)
  &nbsp;

## Install and Run 🙌👨‍💻🚀

### my os : Ubuntu 22.04 LTS

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
# install swagger cli
$ go get -u github.com/swaggo/swag/cmd/swag


#install go migrate
# if failed, please check( https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md ) for your device
$ go get -u -d github.com/golang-migrate/migrate/cmd/migrate

#move to directory
$ cd workspace

# Clone
$ git clone https://github.com/khalidalhabibie/GoArticle.git

#move to project
$ cd GoArticle

# rename file .env.example
$ mv .env.example .env


# Build the docker image first
$ make build

# Run the application
$ make run

# check if the containers are running
$ docker ps

# migration database and seed data
$ make migrate.up

```

### Prometheus UI:

http://localhost:9090

### Grafana UI

http://localhost:3000 \
username: admin \
password: admin

### Swagger UI:

http://localhost:8081/swagger/index.html
