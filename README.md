# Cinematic
Cinematic is an enjoyable pet project where I explored various technologies and tools. In this project, I utilized the Uber FX dependency injection framework to manage dependencies efficiently. Additionally, I experimented with the Swaggo package to generate OpenAPI specifications for APIs and implemented JWT to secure CRUD operations.

Initially, I started with the GORM ORM for database interactions. However, I eventually moved away from this approach in favor of using raw SQL and the PGX package. This decision was driven by my greater familiarity with raw SQL and the enhanced control and understanding PGX provides in building queries.

## How to run and play around

Clone project
```
git clone https://github.com/JulyInSummer/cinematic.git
```
then move inside **cinematic** directory

```
cd cinematic
```

App run inside Docker, so make sure you have Docker & docker-compose installed on your system

> **Warning:** App runs inside Docker, so make sure you have Docker & docker-compose installed on your system.

I'm using migrate for managing migrations in the project. Make sure to install migrate binary. On how to install migrate see the following [link](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).
You'll also need to install [swaggo](https://github.com/swaggo/swag) if you want to play around, add more handlers to generate OpenAPI specification.

And now to run the server simply run:
```
make run-local-docker
make run
```
or
```
docker-compose -f docker-compose.local.yaml up -d --build
go run cmd/app/main.go
```

