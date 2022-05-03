# http  var
HTTP_PORT=4000

# AuthEndpoint  var
BASE_URL=http://localhost:8000/api/authenticate/public/preview1
REVOKE_URL=http://localhost:8000/api/auth/revoke
REDIRECT_URL=http://localhost:8010/callback
AUTH_URL=http://localhost:8080/auth/login
TOKEN_URL=http://localhost:8000/api/auth/token
CLIENT_ID=VSRPUJobSY
CLIENT_SECRET=xQgwBhPJke

# Database  var
DB_SSL_MODE=disable
DB_HOST=localhost
DB_PORT=5432
DB_NAME=inventory-toll
DB_PASSWORD=""
DB_USER=postgres


run:
	go mod tidy
	time go run main.go

graph-build:
	go get -u github.com/99designs/gqlgen/internal/imports@v0.13.0
	go get -u github.com/99designs/gqlgen/internal/code@v0.13.0
	go get -u github.com/99designs/gqlgen/cmd@v0.13.0
	go get -u github.com/vektah/gqlparser/v2@v2.1.0
	rm -f internal/gql/generated/exec.go \
    		internal/gql/models/generated.go \
    		internal/gql/resolvers/generated/generated.go
	time go run -v github.com/99designs/gqlgen $1
	printf "\nDone.\n\n"