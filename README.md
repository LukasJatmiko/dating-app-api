# Dating App API

### Application directory structure

    .
    ├── constants               # Application global constants package
    ├── driver                  # Documentation files (alternatively `doc`)
    ├── migrations              # Contains database migrations files
    ├── packages                # Application packages
    ├── tools                   # Tools and utilities
    ├── types                   # Application global types package
    ├── utils                   # Application utilities package
    ├── main.go                 # Main package
    ├── uc_test.go              # Usecase test file
    └── README.md

# How to deploy

Run database migration

    migrate -database postgres://user:password@host:port/dbname?sslmode=disable -path migrations up

For more information about migration tools installation and commands and usage, please check it's [repository](https://github.com/golang-migrate/migrate)

Build docker image

    sudo docker build --tag dating-app-api .

Run container

    sudo docker run -d --name dating-app-api -p 8080:8080 -e DB_URI=postgresql://username:password@host:port/dbname?sslmode=disable dating-app-api

Acceptable environment variable and example of it's value is as follows

 - DB_URI=postgresql://user:password@localhost:5432/datingapp?sslmode=disable
 - DB_DRIVER_TYPE=POSTGRES
 - DB_MAX_OPEN_CONN=3
 - DB_MAX_CONN_LIFETIME=30m
 - DB_MAX_IDLE_CONN=1
 - APP_NAME=dating-app-api
 - APP_PORT=8080
 - AUTH_JWT_PRIVATE_KEY=/path/to/private/key/mykey.key
 - AUTH_JWT_PUBLIC_KEY=/path/to/public/key/mykey.key.pub

Note : if environment variable is not provided, the service will use default value.

To generate private & public key for JWT encryption, please see example below:

    ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key

    openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub


# Run tests

    go test

# Postman Request Collection

Please check example request collection url [here](https://www.postman.com/admestic/workspace/pub/collection/30567526-bd28623d-b045-4c32-95df-5cbfc1f26251?action=share&creator=30567526)