Docker compose

If you have issues with pg admin connection to db please use container name as a host (pg_auth) 
Other docker-compose parameters should be valid

To check unused code use 'go build -gcflags -live'
To remove unused dependencies use 'go mod tidy -v'