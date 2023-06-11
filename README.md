# bookstore-go

Another example practice CRUD app using Go. This practice exposes on the following concept:

- using `gorilla` to manage routing
- arrange the codes in packages
- docker-compose to manage secrets and deploy database locally
- usage of `dotenv` to manage app secrets

## Development

### Running the database locally
```shell
# Start up the database
docker-compose up -d

# Tear down the database
docker-compose down
```

### Running the app
```shell
# Build the app
cd cmd/main && go build -o ../../bookstore

# Run the executable
./bookstore

```