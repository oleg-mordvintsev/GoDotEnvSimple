# Простая читалка файлов .env

## Начало работы

- Загрузите необходимый репозиторий `go get -u github.com/oleg-mordvintsev/GoDotEnvSimple`.
- Подключите в проекте

```go
import (
env "github.com/oleg-mordvintsev/GoDotEnvSimple"
)
```

- Инициализация автоматически при первом обращении.

## Описание методов

- Одно значение в .env

```go
env.Get("NAME_ENV_CONST")
```

- Одно булево значение в .env

```go
env.GetBool("IS_TEST")`
```

- Проверка существования значения в .env

```go
env.Exists("NAME_ENV_CONST")
```

- Все значения в .env

```go
env.All()
```
