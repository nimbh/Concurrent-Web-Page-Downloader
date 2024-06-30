# Конкурентное скачивание содержимого URL и сохранение в файлы

## Описание
Эта задача заключается в создании программы, которая позволит пользователю вводить список URL-адресов через терминал и одновременно скачивать содержимое каждого адреса в файлы с использованием нескольких потоков.

## Инструкции по установке и запуску
1. Убедитесь, что у вас установлен Go на вашем компьютере.
2. Склонируйте репозиторий с помощью команды `git clone https://github.com/your-repo.git`.
3. Перейдите в каталог проекта с помощью команды `cd your-repo`.
4. Запустите программу с помощью команды `go run main.go -threads <количество_потоков> -urls <список_URL_адресов_через_запятую>`.

Например:
go run main.go -threads 2 -urls https://google.com,https://ya.ru,https://duckduckgo.com

## Используемый стек
- Язык программирования: Go
- Библиотеки: "flag", "fmt", "io/ioutil", "log", "net/http", "strings", "sync"

## Как работает программа
Пользователь указывает количество потоков для одновременной загрузки (N) и вводит список URL-адресов. Программа создает N потоков, каждый загружает содержимое URL и сохраняет в файл. После загрузки каждый поток берет следующий URL из списка, повторяя процесс до окончания всех загрузок.
