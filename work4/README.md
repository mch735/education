# Многопоточный веб-скрейпер с очередью задач и ограничением параллельных запросов

## Пример использования
`go run . https://medium.com https://google.ru https://ya.ru http://example.com https://habr.com https://google.ru/qwe`

## Параметры запуска
- TIMEOUT=5s - таймаут запроса
- THREAD_COUNT=10 - кол-во потоков
- RETRY_COUNT=3 - число попыток при временно недоступном url