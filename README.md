# API Tester

Утилита для проверки и регрессионного тестирования HTTP API.

Проект ориентирован на:

- контрактное тестирование JSON API;
- проверку структуры и содержимого ответов;
- регрессионное тестирование после изменений SQL/бизнес-логики;
- дальнейшее сравнение baseline/current response;
- проверку аналитических и агрегирующих API.

---

# Основные возможности

## Выполнение HTTP-запросов

Поддерживается:

- HTTP requests;
- headers;
- JSON body;
- ENV-переменные через `${VAR}`.

---

## YAML как DSL

Тест-кейсы описываются в YAML.

Пример:

```yaml
request:
  method: POST
  url: ${BASE_URL}/v2/aggregator/get-table

  headers:
    Content-Type: application/json

  body:
    type: TABLE_COLLECTION
    step: WEEKLY

checks:
  - type: status_code
    expected: 200

  - type: json_path_exists
    path: $.metrics

  - type: json_path_count
    path: $.metrics[*]
    expected: 2
```

---

# Реализованные проверки

## status_code

Проверка HTTP status code.

```yaml
- type: status_code
  expected: 200
```

---

## json_path_exists

Проверка существования JSONPath.

```yaml
- type: json_path_exists
  path: $.metrics[*].tableField
```

---

## json_path_count

Проверка количества найденных элементов.

```yaml
- type: json_path_count
  path: $.metrics[*]
  expected: 2
```

---

## json_path_type

Проверка типа всех найденных значений.

Поддерживаемые типы:

- string
- number
- bool
- object
- array
- null

```yaml
- type: json_path_type
  path: $.metrics[*].tableField.values[*].fact
  expected: number
```

---

## json_path_not_empty

Проверка, что значения не пустые.

Для:

- string → `!= ""`
- array → `len > 0`
- object → `len > 0`
- null → fail

```yaml
- type: json_path_not_empty
  path: $.metrics
```

---

## json_path_equals

Проверка равенства значений.

```yaml
- type: json_path_equals
  path: $.status
  expected: ok
```

Если path содержит wildcard (`[*]`), проверка применяется ко всем найденным элементам.

---

# Архитектура

Проект разделён на независимые слои.

## config

Загрузка YAML и ENV resolution.

---

## executor

Выполнение HTTP requests.

Executor:

- не знает о checks;
- не знает о JSON validation;
- только выполняет запрос.

---

## jsonutil

Парсинг JSON и построение структуры документа.

Содержит:

- JSON parser;
- nodes;
- path index;
- extraction helpers.

---

## checks

Слой validation/checks.

Checks:

- анализируют response;
- используют parsed JSON;
- возвращают validation results.

---

## report

Формирование console report.

(в разработке)

---

# Структура проекта

```text
cmd/
    tester/

internal/
    app/
    checks/
    config/
    executor/
    jsonutil/
    report/

examples/
```

---

# Запуск

## ENV

Создать `.env`:

```env
BASE_URL=http://localhost:8080
ORG_ID=...
FIELD_ID=...
METRIC_ID=...
```

---

## Запуск

```bash
go run ./cmd/tester examples/aggregator.yaml
```

---

# Roadmap

## MVP

-  YAML loader
-  ENV resolver
-  HTTP executor
-  JSON parser
-  JSONPath extraction
-  checks engine
-  generic JSON checks
-  console report


# Цели проекта

Основная цель — создать инструмент для:

- проверки аналитических API;
- поиска regressions;
- сравнения текущего и эталонного response;
- проверки бизнес-консистентности данных.