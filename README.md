# Poprev

Poprev is the reward mechanism for Fans all over the world, as such we rely on artist projects and the fans being able to invest in these projects.

## Technology

    - Golang
    - MySQL
    - Redis

Please make sure to install them

## Schema Definition

[db schema](https://dbdiagram.io/d/649b7c3d02bd1c4a5e29e3ce)

# Steps To Run The Project

Clone repository

```bash
git clone https://github.com/tekpriest/poprev.git
```

Donwnload Dependencies

```bash
go mod Donwnload && go mod tidy
```

Setup environment variables

```bash
cp .env.sample .env

```

Sample

```sh
PORT=80
ENV=dev
DATABASE_URL=mysql://root:password@localhost:3306/poprev_dev
REDIS_HOST="127.0.0.1:6379"
REDIS_PORT=6379
REDIS_USER=default
REDIS_PASS=password
```

Please replace the variables with yours

Connect & Create the database name

```sh
CREATE DATABASE `poprev_dev`
```

Migrate the database

```sh
migrate -path="./migrations -database=<DATABASE_URL> up"
```

Run App

```bash
make run/api
```
