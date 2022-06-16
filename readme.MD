## Project based on:
- https://github.com/markpenaranda/echo-gorm
- https://tutorialedge.net/golang/golang-orm-tutorial/ (not used yet but might be usefull to implement DB interactions and endpoints)

## GCC compiller is required to run this project. To install, follow this guide:
- https://code.visualstudio.com/docs/cpp/config-mingw

## To run this project execute following commands while in project directory

- `cd path/to/project/directory/ebinzes_go`
- `go mod init (repository name)`, in this case `go mod init github.com/RobertOchmanek/ebiznes_go`
- `go get github.com/labstack/echo/v4`
- `go get github.com/mattn/go-sqlite3`
- `go get -u github.com/jinzhu/gorm`
- `go build main.go`
- `./main`

## CRUD operations:
- `curl -v -X GET http://localhost:8080/users` to get all users
- `curl -v -X POST http://localhost:8080/users -H 'Content-Type: application/json' -d '{"Username":"username","Email":"example@mail.com"}` to add new user
