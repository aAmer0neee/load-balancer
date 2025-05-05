# Load Balancer Test Task

**Описание проекта:**  
HTTP-балансировщик нагрузки на Go с поддержкой Round-Robin алгоритма и Rate-Limiting с алгоритмом Token Bucket.  
Обеспечивает распределение запросов между бэкенд-серверами, защиту от перегрузок и гибкую конфигурацию.

---

## Основной функционал

### Балансировщик нагрузки
- Проксирование HTTP-запросов на бэкенд-серверы через Round-Robin
- Автоматическое исключение недоступных бэкендов
- Конкурентная обработка запросов с использованием горутин
- Конфигурация через YAML-файл

### Rate-Limiting
- Реализация алгоритма Token Bucket
- Автоматическое пополнение токенов
- Потокобезопасные операции с токенами

### Дополнительно
- Health-check бэкенд-серверов
- Логирование запросов и ошибок

---

## Конфигурация

`config.yaml`:
```yaml
server:
  host: localhost
  port: "8888"

services:
  pool:
    - localhost:9001
    - localhost:9002
    - localhost:9003
    - localhost:9004

health:
  timeout_ms: 500
  ticker_ms: 5000

limiter:
  capacity: 1
  ticker_ms: 5000
```

## Основной функционал

- Алгоритм балансировки Round-Robin
- Автоматический Health-Check бэкендов
- Token Bucket Rate-Limiting 
- Гибкая система конфигурации
- Поддержка Graceful Shutdown
- Логирование операций
- Параллелизм и конкурентность

---

## TO DO

- Гранулярное ограничение:
    - Отслеживать состояние каждого клиента (IP/API-ключ)
    - Поддерживать возможность настройки разных лимитов для разных клиентов.
    - Настройки для разных клиентов можно сохранять в базе данных

- Интеграционные тесты 
- Несколько алгоритмов распределения
- Docker-compose для базы состояний клиентов

## Запуск

- Сервис
```bash
go build -o server balancer/cmd/main.go
./server --config=./config.yaml
```
- Тестовые бекэнды
```bash
go build -o backs backends/cmd/main.go
./backs --config=./config.yaml
```


## REST API
- (GET /)

Пример ответа:
```yaml
{
  "from": "localhost:9001",
  "message": "Hello World!"
}
```

## Технологии/пакеты

- Go

- net/http/httputil

- log/slog

- ilyakaznacheev/cleanenv


## Тестирование

Запуск нагрузочного теста:
```bash
ab -n 5000 -c 1000 http://localhost:8888/
```
## Структура проекта

```yaml
.
├── 📄 README.md                # Документация
├── 📄 go.mod                   # Зависимости Go
├── 📄 go.sum                   # Контрольные суммы
├── 📄 config.yaml              # Конфигурация сервиса
├── 📂 backends                 # Примеры бэкенд-серверов
│
└── 📂 balancer                 # Основной сервис балансировщика
    ├── 📂 cmd
    │   └── 📄 main.go          # Запуск приложения
    ├── 📂 domain
    │
    ├── 📂 internal
    │   ├── 📂 balancer         # 🎚 Логика балансировки
    │   │
    │   ├── 📂 health           # 🩺 Health-check
    │   │ 
    │   ├── 📂 limiter          # 🚦 Rate-limiting
    │   │
    │   ├── 📂 proxy            # 🔄 Reverse proxy
    │   │ 
    │   ├── 📂 service          # 🧠 Бизнес-логика
    │   │
    │   └── 📂 transport        # 🌐 HTTP-транспорт
    │ 
    └── 📂 pkg
        ├── 📂 config           # ⚙️ Загрузка конфигурации
        │
        └── 📂 logger           # 📝 Логирование

```