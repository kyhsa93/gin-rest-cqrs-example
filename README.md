# Gin CQRS Rest api example

This is sample Go project.

Go: [https://golang.org/](https://golang.org/)

Gin: [https://gin-gonic.com/](https://gin-gonic.com/)

Gorm: [http://gorm.io/](http://gorm.io/)

Docker: [https://www.docker.com/](https://www.docker.com/)

## Getting started

This is REST api made by Gin, redis, mysql with Go.

### Prerequisites

Please install Go and docker.

I recommand to use docker for your environment.

 * Install Go: [https://golang.org/dl/](https://golang.org/dl/)

 * Install Docker for MAC: [https://docs.docker.com/docker-for-mac/install/](https://docs.docker.com/docker-for-mac/install/)

 * Install Docker for Windows: [https://docs.docker.com/docker-for-windows/install/](https://docs.docker.com/docker-for-windows/install/)

 * Install compose: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

## Create development environment

Firstable, create directory `$GOPATH/src/github.com/kyhsa93/` and clone this repository into your under that.

```bash
  mkdir $GOPATH/src/github.com/kyhsa93
  cd $GOPATH/src/github.com/kyhsa93
  git clone https://github.com/kyhsa93/gin-rest-cqrs-example.git
```

And then, build this program for set all dependencies.


```bash
  go build
```

`go.mod` and `go.sum` is already exists. so you don't need run `go mod init`.

Next up, generate mysql and redis.

If you already have mysql and redis in your environment, you can use that.

But if you don't have one or both of them, try followed process.

Install docker for your OS from link in top of this document.

If your docker is successfully installed, you can use docker cli.

```bash
  docker run --name gin-rest-cqrs-example -d -p 3306:3306 -e MYSQL_DATABASE=gin-rest-cqrs-example -e MYSQL_ROOT_PASSWORD=test -v ~/database/gin-rest-cqrs-example:/var/lib/mysql mysql:5.7
  docker run --name redis -d -p 6379:6379 redis:alpine

  OR

  docker-compose -f docker-compose.dev.yml up -d # create mysql, redis container for development environment
  docker-compose -f docker-compose.dev.yml down  # remove created containers
```

> Note. If you use docker container or any other docker resources included docker compose, recommended remove that after you use.

Now you can connect mysql in localhost:3306, that user 'root' and password is 'test'.

Finally, your development environment is created.

And now you can start api with followed command.

```bash
  go run main.go
```

If you want apply your code change into running process, save all changes and rerun `go run main.go`.

## Start with docker

If you can use docker cli, you can build docker image.

```bash
  docker build -t gin-rest-cqrs-example .
  docker images # list up docker images
```

And then you can create and run docker container using builded image.

```bash
  docker run -d -p 5000:5000 gin-rest-cqrs-example
  docker ps # list up runnint docker containers
```

And now you can connect api through http://localhost:5000.

## Start with docker compose

Docker compose in this project is include api, redis and mysql.

Run followed command in this project directory root.

```bash
  docker-compose up -d # pull images, create and run containers in background process
```

If all containers are created, you can access api on http://localhost, and database also you can connect through by http://localhost:3306.

Default database user is 'root' and password is 'test'.

After use docker-compose, you have to stop and remove all resources created by docker-compose in this project.

Run followed command in project root.

```bash
  docker-compose down # stop and remove containers in defined docker-compose.yml
```

> Note. docker-compose in this project, does not build from this source code. If you want to build and use image from this code, you have to modify docker-compose.yml

## Configurations

All configurations are in [./config](https://github.com/kyhsa93/gin-rest-cqrs-example/tree/master/config)

Most default configuration can use with your environment values.

And also you can modify configurations.

## Documentation

Document about this project us made swagger.

Start this api and connect api host in your browser.

 * swagger config: [./config/swagger.go](https://github.com/kyhsa93/gin-rest-cqrs-example/blob/master/config/swagger.go)

> Note. Swagger in this project is use swag ([https://github.com/swaggo/swag](https://github.com/swaggo/swag)). Please check before you use it.

## Scripts

```bash
  git clone https://github.com/kyhsa93/gin-rest-cqrs-example.git # clone this project
  
  go build # build this project
  
  docker run --name gin-rest-cqrs-example -d -p 3306:3306 -e MYSQL_DATABASE=gin-rest-cqrs-example -e MYSQL_ROOT_PASSWORD=test -v ~/database/
  gin-rest-cqrs-example:/var/lib/mysql mysql:5.7 # create mysql container
  
  docker run --name redis -d -p 6379:6379 redis:alpine # create redis container
  
  docker-compose -f docker-compose.dev.yml up -d # create mysql, redis container for development environment
  
  docker-compose -f docker-compose.dev.yml down  # remove created containers

  go run main.go # start 

  docker build -t gin-rest-cqrs-example . # build docker image
  
  docker images # list up docker images

  docker run -d -p 5000:5000 gin-rest-cqrs-example
  
  docker ps # list up runnint docker containers

  docker-compose up -d # build images, create and run containers in background process

  docker-compose down # stop and remove containers in defined docker-compose.yml

  swag init # build swagger
```

## Links
Github: [https://github.com/kyhsa93/gin-rest-cqrs-example](https://github.com/kyhsa93/gin-rest-cqrs-example)

Dockerhub: [https://hub.docker.com/repository/docker/kyhsa93/gin-rest-cqrs-example](https://hub.docker.com/repository/docker/kyhsa93/gin-rest-cqrs-example)
