package env

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var env map[string]string

func Get(key string) string {
	return env[key]
}

func GetBool(key string) bool {
	if env[key] == "1" || strings.ToLower(env[key]) == "true" {
		return true
	}
	return false
}

func Exists(key string) bool {
	_, isExists := env[key]
	return isExists
}

func All() map[string]string {
	return env
}

func init() {
	// Читаем файл (весь сразу, т.к. он вряд ли когда-то будет больше 10кб)
	fileContent, err := ioutil.ReadFile(".env")

	// На случай ошибок чтения, сразу панику, т.к. дальше действовать нет смысла
	if err != nil {
		panic(err)
	}

	// Разделяем файл на строки
	fileContentStrings := strings.Split(string(fileContent), "\n")

	// Инициализируем глобальную переменную
	env = make(map[string]string)

	// Разделяем строки на переменную и её значение
	for key, value := range fileContentStrings {
		// Удаляем пробельные символы в начале и в конце строки (где предполагаемые и ключ и значение)
		value = strings.TrimSpace(value)

		// Пропуск пустых строк
		if "" == value {
			continue
		}

		// Если данные не пусты, то добавляем в общую мапу
		if strings.Contains(value, "=") {
			// Разбираем строку на ключ и значение
			keyValue := strings.Split(value, "=")

			// Ключ не должен быть с пробельными символами
			keyValue[0] = strings.TrimSpace(keyValue[0])

			// Получаем ключ и сцепляем значение, т.к. оно могло содержать знак разбиения(=)
			getKeyValueWithParts(keyValue)

			// Очистка от кавычек значения с сохранением пробельных символов внутри кавычек
			env[keyValue[0]] = strings.Trim(strings.TrimSpace(env[keyValue[0]]), "'\"")

			// Удаляем не нужную переменную
			keyValue = nil
		} else {
			// Цветной вывод алертов https://golangbyexample.com/print-output-text-color-console/
			fmt.Println("\033[33m", "ОШИБКА ПАРСИНГА .env, не удалось получить данные строки", key+1, ":", value, "\033[0m")
		}
	}

	// Удаляем не нужное
	fileContentStrings = nil
}

// Проверяем есть ли ещё части, т.к. если есть, то со второго и последующие элементы нужно обратно сцепить
func getKeyValueWithParts(keyValue []string) {
	if len(keyValue) > 2 {
		env[keyValue[0]] = strings.Join(keyValue[1:], "=")
	} else {
		env[keyValue[0]] = keyValue[1]
	}
}
