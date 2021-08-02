# 1 - Simple Database querying

Terdapat sebuah table "USER" yg memiliki 3 kolom: ID, UserName, Parent. Di mana: Kolom ID adalah Primary Key
Kolom UserName adalah Nama User
Kolom Parent adalah ID dari User yang menjadi Creator untuk User tertentu.

## Settings
1. Edit env.example file to `.env` and insert you env data: 
```Shell
cp env.example .env && nano .env
```

2. Import .sql to your mysql table


## Compile
that you can start go applications via 
```Shell
go run main.go
```
but also compile it to an executable with 
```Shell
go build main.go
```

## License
ISC
