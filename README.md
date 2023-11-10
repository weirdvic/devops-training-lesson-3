# devops-training-lesson-3
Домашнее задание по третьей лекции Тренировок по DevOps

## Пример вывода (пропущено через jq)
```json
{
  "process_info": {
    "fd_count": 13,
    "vm_rss": "unknown",
    "exec_path": "/home/user/dev/devops-training-lesson-3/sysreport/sysreport"
  },
  "cpu_info": {
    "cpu_count": 8,
    "model_name": "11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz"
  },
  "memory_info": {
    "mem_total": "16089484 kB",
    "mem_free": "2645944 kB",
    "mem_available": "5285768 kB",
    "buffers": "315100 kB",
    "cached": "3554820 kB",
    "swap_total": "16777212 kB",
    "swap_cached": "453188 kB",
    "swap_free": "13581520 kB"
  }
}
```
## Локальный запуск
```bash
go run sysreport/*.go
```
или
```bash
cd sysreport
go build -o sysreport
./sysreport
```
## Сборка и запуск в Docker
```bash
docker build -t sysreport .
docker run --rm -it sysreport
```
