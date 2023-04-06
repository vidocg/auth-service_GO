## Table of contents
* [Docker compos](#docker-compose)
* [GO SDK issue](#go-sdk-issue)
* [Postman](#postman)
* [Swagger](#swagger)
* [Libraries that used](#libraries-that-used)
* [Environment variables that used](#environment-variables)

# Docker compose

If you have issues with pg admin connection to db please use container name as a host (pg_auth) 
Other docker-compose parameters should be valid

To check unused code use 'go build -gcflags -live'
To remove unused dependencies use 'go mod tidy -v'

# GO SDK issue
## Failed to download SDK Unpacked SDK is corrupted
if you faced with "Failed to download SDK Unpacked SDK is corrupted" issue downloading go sdk through intellij idea
you need to go to sdk folder (for example C:\Users\{user}\sdk\{go sdk version}\src\runtime\internal\sys) and
add to zversion.go 
const StackGuardMultiplierDefault = 1 
const theVersion = `{go version}` 

(theVersion for me was go1.20.2)
 
Then you need to invoke File | Settings | Go | GOROOT, select Local and specify the path to the Go SDK.
 
## 'go' is not recognized as an internal or external command
You need to go to system properties / env varible and add to Path : C:\Users\{user}\sdk\{go package}\bin

Also, you can read some useful information about syncing dependencies etc here 
'https://www.jetbrains.com/help/go/create-a-project-with-go-modules-integration.html#notify-about-replacements-of-local-paths-in-go-mod-file'

# Postman
You can use `AUTH.postman_collection.json` to trigger auth endpoints

# Swagger
to update swagger docs use `swag init -g src/main.go`
(if it doesn't work please use `go install github.com/swaggo/swag/cmd/swag@latest` before)
swagger url : `http://localhost:9993/swagger/index.html`

# Libraries that used
## Migration
As database version control tool is used `golang-migrate`. It is executed it db_configurer. Change sets path: `migrations`

## ORM
Gorm is used as ORM library. Link: `https://gorm.io/`

## HTTP Web
Gin is used as http web framework. Link: `https://github.com/gin-gonic/gin`

## Ioc
For inversion of control golobby/container is used. Link: `https://github.com/golobby/container`

## Validation
For validation go-playground/validator is used. You can use custom wrapper: `custom_validator.CustomValidator`. Link: `https://github.com/go-playground/validator`

## Mapper
For dto->entity\entity->dto mapping devfeel/mapper is used. Link: `https://github.com/devfeel/mapper` 

# Environment variables
For the development is used [app.env](app.env) from the root which is parsed on bootstrap by spf13/viper. Link: `https://github.com/spf13/viper`