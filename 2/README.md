# Mission 2 - Please write a microservice to search movies from http://www.omdbapi.com/

## Settings
1. Edit env.example file to `.env` and insert you env data: 

```bash
cp env.example .env && nano .env
```

## Compile
**1.** that you can start go applications via 
```Shell
go run main.go
```
but also compile it to an executable with 
```Shell
go build main.go
```

**2.** Build and run the service with this command
```shell
$ docker-compose up --build -d
```
