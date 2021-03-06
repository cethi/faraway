1. Клиент посылает сообщение вида `hello$<client_id>` на сервер и получает задание-строку.
2. Используя полученное задание клиент вычисляет значение `hashcash` и отдает его на сервер в виде `result$<client_id>$<hash>`. Сервер проверяет полученное значение и, если оно корректно, возвращает клиенту цитату.

Точки роста или что не сделано, чтобы уложиться в пару часов:

- Клиенты могут получать задания, но не возвращаться с ответом. Информация о полученных заданиях будет храниться на сервере, занимая место в памяти. Что с этим делать: добавить в `MemoryStore` горутину, которая будет запускаться раз в `N` минут и удалять задания, которые были выданы клиентам, но с решением клиенты так и не вернулись. Для того, чтобы это проще было делать в `MemoryStore` надо добавить `LinkedList` в котором хранить указатели на элементы `map`.
- Сервер не переживает перезапуск (теряет информацию о выданных заданиях). Чтобы этого не происходило, можно в дополнение к `MemoryStore` добавить, например, `RedisStore`. Заодно решится предыдущая проблема, так как для ключей можно будет устанавливать TTL.
- Никак не рассмотрен вопрос с масштабированием и обеспечением отказоустойчивости. Если у нас есть общий `Store`, то можно поднять несколько серверов и балансировать запросы клиенты между ними. Если общего `Store` нет - тоже можно, но надо следить, чтобы один и тот же клиент всегда попадал на один и тот же сервер. Для этого можно использовать хеш функцию от `client_id`, а чтобы иметь возможность менять количество серверов использовать `consistent hashing`.
- Тесты далеко не полные. Хорошо бы добавить те, которые тестируют работу сервера при конкурентных обращениях клиента.
- Порт, на котором работает сервер, длину строки в задании и прочие константы можно вынести в переменные окружения, чтобы конфигурировать работу сервера через `docker-compose`.


# Prerequisites 

- `make`
- `docker`
- `docker-compose`. 

# Build

```shell
# make build
```

# Example

```shell
# make run
docker-compose up --force-recreate
Recreating tcp_server ... done
Recreating tcp_client ... done
Attaching to tcp_server, tcp_client
tcp_client | 2022/06/05 14:59:43 Task was received from server: pTNeoCcitc
tcp_server | 2022/06/05 14:59:42 Starting server...
tcp_server | 2022/06/05 14:59:43 [clientId=Alice] Created task: pTNeoCcitc
tcp_client | 2022/06/05 14:59:44 Hash was calculated: 1:20:220605145943:pTNeoCcitc::8Ep1v3I3r/8=:OTU1OTQ4
tcp_server | 2022/06/05 14:59:44 [clientId=Alice] Validate solution: 1:20:220605145943:pTNeoCcitc::8Ep1v3I3r/8=:OTU1OTQ4
tcp_client | 2022/06/05 14:59:44 You have got the quote: Here is your quote: Believe and act as if it were impossible to fail.
tcp_server | 2022/06/05 14:59:44 [clientId=Alice] Validation result: true
tcp_server | 2022/06/05 14:59:44 connection was closed
tcp_client exited with code 0
```