Docker compose

If you have issues with pg admin connection to db please use container name as a host (pg_auth) 
Other docker-compose parameters should be valid

To check unused code use 'go build -gcflags -live'
To remove unused dependencies use 'go mod tidy -v'


#GO SDK issue
if you faced with "Failed to download SDK Unpacked SDK is corrupted" issue downloading go sdk through intellij idea
you need to go to sdk folder (for example C:\Users\{user}\sdk\{go sdk version}\src\runtime\internal\sys) and
add to zversion.go 
const StackGuardMultiplierDefault = 1 
const theVersion = `{go version}` 

(theVersion for me was go1.20.2)
 
Then you need to invoke File | Settings | Go | GOROOT, select Local and specify the path to the Go SDK.
 
#'go' is not recognized as an internal or external command
You need to go to system properties / env varible and add to Path : C:\Users\{user}\sdk\{go package}\bin

Also, you can read some useful information about syncing dependencies etc here 
'https://www.jetbrains.com/help/go/create-a-project-with-go-modules-integration.html#notify-about-replacements-of-local-paths-in-go-mod-file'

#Postman
You can use `AUTH.postman_collection.json` to trigger auth endpoints

#Swagger
to update swagger docs use `swag init -g src/main.go`
swagger url : `http://localhost:9993/swagger/index.html`