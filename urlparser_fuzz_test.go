package main

import (
	"testing"
)

func FuzzParseURL(f *testing.F) {
	// Добавляем "семена" — примеры входных данных
	f.Add("https://example.com")
	f.Add("http://localhost:8080")
	f.Add("ftp://files.server.org/path")
	f.Add("")

	// Запускаем фаззинг
	f.Fuzz(func(t *testing.T, input string) {
		scheme, host, err := ParseURL(input)

		// Если нет ошибки — проверяем, что схема и хост не пустые
		if err == nil {
			if scheme == "" {
				t.Errorf("scheme is empty for input: %q", input)
			}
			if host == "" {
				t.Errorf("host is empty for input: %q", input)
			}
		}

		// Дополнительная проверка: если есть ошибка, схема и хост должны быть пустыми
		if err != nil {
			if scheme != "" || host != "" {
				t.Errorf("expected empty scheme/host on error, got scheme=%q host=%q for input: %q", scheme, host, input)
			}
		}
	})
}
