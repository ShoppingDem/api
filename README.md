# Shopping API

This project is a Shopping API built with Go, using the Gin framework and Okta for authentication.

## Setup and Installation

### Prerequisites

- Go (version 1.16 or later)
- Git

### Clone the Repository

```bash
git clone https://github.com/ShoppingDem/api.git
cd apigo get -u github.com/gin-gonic/gin

### Install Dependencies

go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/okta/okta-jwt-verifier-golang
go get -u github.com/joho/godotenv

### Generate Swagger Documentation

swag init


### Set Up Environment Variables

OKTA_ISSUER=https://{yourOktaDomain}/oauth2/default
OKTA_CLIENT_ID={yourClientId}

### Build and Run


go build -o api
./api


###API Documentation
####Security

This API uses Okta for authentication. All protected endpoints require a valid Okta token to be included in the request header.

GET /api/protected HTTP/1.1
Host: localhost:8080
Authorization: Bearer your_okta_token_here


Replace your_okta_token_here with a valid Okta token.


### Swagger Documentation
http://localhost:8080/swagger/index.html

### Additional Information


This README.md file now includes all the necessary information for setting up the project, understanding the API documentation, and using the authentication system. The security definitions from the selected code are incorporated into the "API Documentation" section.
