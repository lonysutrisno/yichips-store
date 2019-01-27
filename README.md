## Package Require

As a library

1) Echo framework
```shell
go get gopkg.in/echo.v3
```

2) ORM
```shell
go get -u github.com/jinzhu/gorm
```

3) Validator helper
```shell
go get gopkg.in/go-playground/validator.v9
```

4) Testing helper
```shell
go get github.com/stretchr/testify
```

5) Env helper
```shell
go get github.com/joho/godotenv
```

6) map[string]interface{} to struct converter
```shell
go get github.com/mitchellh/mapstructure
```

## Preparation before running

1. Create file .env for main env on root folder
2. Create file .env.test for test env on root folder
3. Copy file .env.test content to (all) /services/*/test/.env