package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

// go test -v serialize_properties_test.go
type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize(person Person) string {
	// Получаем значение и тип структуры через рефлексию
	val := reflect.ValueOf(person)
	typ := val.Type()

	var result []string

	// Проходим по всем полям структуры
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Получаем тег properties из поля
		tag := field.Tag.Get("properties")
		if tag == "" {
			continue // Если нет тега, пропускаем поле
		}

		// Разделяем тег на части (имя ключа и опции)
		parts := strings.Split(tag, ",")
		key := parts[0]

		// Проверяем, есть ли опция omitempty
		hasOmitEmpty := false
		for j := 1; j < len(parts); j++ {
			if strings.TrimSpace(parts[j]) == "omitempty" {
				hasOmitEmpty = true
				break
			}
		}

		// Если стоит omitempty и значение пустое - пропускаем поле
		if hasOmitEmpty && isEmptyValue(value) {
			continue
		}

		// Форматируем значение в строку
		strValue := formatValue(value)

		// Добавляем пару "ключ=значение" в результат
		result = append(result, fmt.Sprintf("%s=%s", key, strValue))
	}

	// Объединяем все строки через перенос строки
	return strings.Join(result, "\n")
}

// isEmptyValue проверяет, является ли значение "пустым" (нулевым)
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return false
	}
}

// formatValue преобразует значение любого типа в строку для .properties формата
func formatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	default:
		// Для остальных типов используем стандартное форматирование
		return fmt.Sprintf("%v", v.Interface())
	}
}

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}
