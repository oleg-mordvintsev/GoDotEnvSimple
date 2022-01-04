# Простая читалка файлов .env

## Начало работы

- Загрузите необходимый репозиторий `go get -u github.com/oleg-mordvintsev/go-dot-env-simple/v1`.
- Подключите в проекте

```go
import (
    env "github.com/oleg-mordvintsev/go-dot-env-simple"
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

## Текущий функционал

- Пропуск пустых строк
- Очистка от незначащих символов в начале и в конце в названиях и значениях
- Если строка в `.env` не содержит знака `=`, то будет выведена ошибка
- Если строка содержит множество знаков `=`, то любое последующее вхождение символа будет единым целым со значением
- Ключ/название переменной не может быть пустым
- Значения в кавычках, одинарных и двойных, очищаются от кавычек в начале и конце значения
