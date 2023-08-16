# azure-graphql

Azure SQL Server with GraphQL & Docker

## How to run

1. Run docker-compose. Use -d to run daemon on the backgroud
```
docker-compose up -d
```

2. Store has default values to connect to the database.

If you'd like to add change the configuration, create .env file in the root directory

and modify fields below to adjust your config.
```
AZURE_SERVER_URL=localhost
AZURE_SERVER_PORT=1433
AZURE_USERNAME=username
AZURE_PASSWORD=yourPassword
AZURE_DBNAME=master

HTTP_LISTEN_ADDR=3000
```

Empty or not provided fields will be replaced by the default values.