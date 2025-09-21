## Что проверяет фаззинг:

- Не падает ли функция при случайных входных данных (паника?).
- Соответствует ли поведение ожидаемому: если ошибка — схема/хост пустые; если нет ошибки — они заполнены.
- Найдёт ли он edge-кейсы, которые вы не предусмотрели (например, строку `"://"` или `"https://"`).

---

## Как запустить фаззинг

В терминале, в директории проекта:

```bash
go test -fuzz=FuzzParseURL -fuzztime=30s
```

Go начнёт генерировать случайные строки и передавать их в `FuzzParseURL`, пытаясь найти паники или ошибки в логике.

Если фаззер найдёт входные данные, вызывающие панику или `t.Error`, он сохранит их в `testdata/fuzz/...` и выведет отчёт.

---

##  Зачем это нужно?

- Фаззинг в Go отлично подходит для поиска паник, выходов за границы массивов, деления на ноль и т.п.
- Всегда добавляйте "семена" — хорошие примеры входных данных, чтобы фаззер быстрее нашёл правильные пути.
- После первого запуска Go сохранит найденные "интересные" кейсы и будет использовать их в будущем.

---

## Результата тестирования:
```bash
$go test -fuzz=FuzzParseURL -fuzztime=30s
```

```
fuzz: elapsed: 0s, gathering baseline coverage: 0/7 completed
failure while testing seed corpus entry: FuzzParseURL/dd41bb9a9f2f7568
fuzz: elapsed: 0s, gathering baseline coverage: 4/7 completed
--- FAIL: FuzzParseURL (0.06s)
--- FAIL: FuzzParseURL (0.00s)
urlparser_fuzz_test.go:21: scheme is empty for input: "://0"

FAIL
exit status 1
FAIL    github.com/akozadaev/go_fuzzing_test    0.061s
```
Нужно добавить проверку, что схема не пустая — ведь по смыслу URL "://..." — некорректен.

```go
	if scheme == "" {
		return "", "", errors.New("empty scheme")
	}
```

результат:
```
$ go test -fuzz=FuzzParseURL -fuzztime=30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/7 completed
fuzz: elapsed: 0s, gathering baseline coverage: 7/7 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 212829 (70938/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 6s, execs: 449149 (78775/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 9s, execs: 683800 (78196/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 12s, execs: 919867 (78702/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 15s, execs: 1150210 (76778/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 18s, execs: 1373912 (74581/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 21s, execs: 1607224 (77761/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 24s, execs: 1831657 (74811/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 27s, execs: 2058201 (75510/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 30s, execs: 2277004 (72934/sec), new interesting: 0 (total: 7)
fuzz: elapsed: 30s, execs: 2277004 (0/sec), new interesting: 0 (total: 7)
PASS
ok      github.com/akozadaev/go_fuzzing_test    30.053s

```