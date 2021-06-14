# Member list web app

## Build and run

1. Clone repo `git clone https://github.com/dittohead/yalantis_golang_school`
1. Navigate to cloned repo`cd member_list_web_app`
1. Install depencies `go mod download`
1. Run `go build` to create the binary (`member_list_app`)
1. Set env variable `export PORT=":8080"`, if variable not set will be used default port 4747
1. Run binary : `./member_list_app`
1. Navigate in browser to http://localhost:[port]/assets/
* If set 8080: http://localhost:8080/assets/
* If not set env: http://localhost:4747/assets/
8. Have fun

## Run tests
To run tests, run `go test ./...` in terminal

## Heroku
https://member-list-web-app.herokuapp.com/assets/
