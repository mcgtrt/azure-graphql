# azure-graphql

Azure SQL Server with GraphQL & Docker

## How to run

1. Run docker-compose. Use -d to run daemon on the backgroud
```
docker-compose up -d
```

2. Environment variables (adjust to your local config).

If not provided, will use default (as per docker-compose setup).
```
AZURE_SERVER_URL=localhost
AZURE_SERVER_PORT=1433
AZURE_USERNAME=sa
AZURE_PASSWORD=superStrong(!)Password
AZURE_DBNAME=master

HTTP_LISTEN_ADDR=3000
JWT_SECRET=superstrongpassword
```

3. Run command to your server to create schema, tables & initial employee you can use for auth testing.

If you don't have sqlcmd, you can install with brew (on Mac):
```
brew install sqlcmd
```

Then run (change port, username and password to adjust your settings, below default)
```
sqlcmd -S localhost:1433 -U sa -P 'superStrong(!)Password' -d master -i ./CreateSchema.sql
```

4. Run server
```
make run
```

## Endpoints

- "/employee" - CRUD operations for Employees, secured with JWT authentication

- "/login" - authenticate with email and password. If you executed step 3. properly, you can now make a POST request
```
{
    "email": "test@go.com",
    "password": "superstrongpassword"
}
```
to get the token from initial employee account.

- "/" - GraphQL playground. You can find handy queries for testing in "/playground.graphql" file.

# Enjoy!
