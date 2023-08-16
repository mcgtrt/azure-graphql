# azure-graphql

Azure SQL Server with GraphQL & Docker

## How to run

1. Run docker-compose. Use -d to run daemon on the backgroud
```
docker-compose up -d
```

2. Store has default values to connect to the database.

If you'd like to add change the configuration, create .env file in the root directory.
```
AZURE_SERVER_URL=
AZURE_SERVER_PORT=
AZURE_USERNAME=
AZURE_PASSWORD=
AZURE_DBNAME=
```

Empty or not provided fields will be replaced by the default values.