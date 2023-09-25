# REST API MVC

* go mod init
```
    go mod init namaproject
```

* Install Echo
```
    go get -u github.com/labstack/echo/v4
```

* Install GORM
```
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/mysql
```

* Install godotenv
```
    go get -u github.com/joho/godotenv
```

* Create `.env` file
```
SERVERPORT=your-port
DBPORT=your-port
DBHOST=your-host
DBUSER=your-user
DBPASS=your-password
DBNAME=your-dbname
```