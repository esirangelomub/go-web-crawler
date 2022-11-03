# API REST IN GO

A simple Rest API made in Go

## Installation

Clone the project [where](git@github.com:esirangelomub/go-api-rest.git).

```bash
git clone git@github.com:esirangelomub/go-api-rest.git
```

Create a new docker container containing MongoDB image

```bash
docker run -d  --name mongodb -p 27017:27017 mongo
```

## Run application

```bash
go run main.go
```
or
```bash
go run main.go -url=<SITE_URL>
```

## Compiling application and run

```bash
go build
```

```bash
./go-web-crawler
```
or
```bash
./go-web-crawler.exe
```

## Access API

See the Postman collection with all CRUD examples.

## Todo
Implement tests.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)