### Description

Go http middleware for authenticating incoming http requests with JWT.

#### Run server
```
# get go packages
go mod init
go mod tidy

# run server
export JWT_SECRET_KEY="your-256-bit-secret"
go run main.go

# generate jwt token from https://jwt.io/
# make sure you use the same secret key (JWT_SECRET_KEY) to generate a valid JWT token

# test requests with/without jwt token

JWT_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
curl http://localhost:3000/ --header "Token:$JWT_TOKEN"
```