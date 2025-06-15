# Singo

Singo: Simple Single Golang Web Service

go-crud has been officially renamed to Singo!

Developing Web Services with Singo: Using the simplest architecture to implement a sufficient framework for serving massive users

https://github.com/Gourouting/singo

## Changelog

1. API testing support added
2. Now supports go1.20, please install this version of golang to use this project

## Video Tutorial

[Let's Build a G-Site! Golang Full Stack Programming Live](https://space.bilibili.com/10/channel/detail?cid=78794)

## Example Projects Built with Singo

Bilibili-inspired G-Site: https://github.com/Gourouting/giligili

Singo framework example with Token login for mobile: https://github.com/bydmm/singo-token-exmaple

## Purpose

This project incorporates several popular Golang components and can be used as a foundation to quickly build Restful Web APIs.

## Features

This project has integrated many essential components for API development:

1. [Gin](https://github.com/gin-gonic/gin): Lightweight Web framework, claims to have the fastest routing in golang
2. [GORM](https://gorm.io/index.html): ORM tool. This project requires MySQL
3. [Gin-Session](https://github.com/gin-contrib/sessions): Session management tool provided by Gin framework
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis client
5. [godotenv](https://github.com/joho/godotenv): Environment variable tool for development, making it easy to use environment variables
6. [Gin-Cors](https://github.com/gin-contrib/cors): CORS middleware provided by Gin framework
7. [httpexpect](https://github.com/gavv/httpexpect): API testing tool
8. Basic internationalization (i18n) functionality implemented
9. This project uses cookie-based sessions to maintain login state, which can be modified to token authentication if needed

This project has pre-implemented some commonly used code for reference and reuse:

1. User model created
2. Implemented ```/api/v1/user/register``` user registration endpoint
3. Implemented ```/api/v1/user/login``` user login endpoint
4. Implemented ```/api/v1/user/me``` user profile endpoint (requires session after login)
5. Implemented ```/api/v1/user/logout``` user logout endpoint (requires session after login)

This project has pre-created a series of folders to organize the following modules:

1. api folder serves as the controller in the MVC framework, responsible for coordinating components to complete tasks
2. model folder stores database models and database operation related code
3. service handles complex business logic, modeling business code can effectively improve code quality (e.g., user registration, recharge, order placement)
4. serializer stores common JSON models, converting database models from model into JSON objects needed by the API
5. cache handles Redis cache related code
6. auth folder for permission control
7. util contains common utility tools
8. conf stores static configuration files, with locales containing translation related configuration files

## Godotenv

The project depends on the following environment variables at startup, but you can also create a .env file in the project root directory to set environment variables for easier use (recommended for development environment)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # MySQL connection address
REDIS_ADDR="127.0.0.1:6379" # Redis port and address
REDIS_PW="" # Redis connection password
REDIS_DB="" # Redis database from 0 to 10
SESSION_SECRET="setOnProducation" # Session secret key, must be set and not leaked
GIN_MODE="debug"
```

## Go Mod

This project uses [Go Mod](https://github.com/golang/go/wiki/Modules) for dependency management.

```shell
go mod init singo
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // automatic installation
```

## Running

```shell
go run main.go
```

The project runs on port 3000 (can be modified, refer to gin documentation)

## API Testing
[New] This project includes built-in API testing

#### Usage
0. Make sure you are in the project root directory
1. Create a test-specific environment variable file in the test directory

```
cp test/.env.example test/.env
```

2. Modify the environment variables in ```test/.env``` file to ensure proper connection to mysql/redis
3. Execute tests in the project root directory with ```-v``` to check if tests are running correctly

```
go test -v ./test
```

4. After confirming tests run correctly, remove the -v parameter to check if tests pass
```
go test ./test
ok      singo/test      (cached)
```