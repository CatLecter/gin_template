# Gin Template

For starting:

```shell
cp ./src/build/example.env ./src/build/.env
docker compose up -d --build
```

Update swagger documentation:

```shell
cd src && swag init -g ./cmd/app/main.go -o ./docs
```

[Swagger](http://localhost:8000/swagger/index.html)

[Instructions](https://github.com/swaggo/swag/blob/master/README.md) for describing the documentation
