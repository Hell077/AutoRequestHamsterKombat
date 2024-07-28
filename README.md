


# AutoRequestHamsterKombat

# AutoRequestHamsterKombat - это простая программа на Go, которая отправляет HTTP POST запросы с заданным интервалом времени.

# AutoRequestHamsterKombat is a simple Go program that sends HTTP POST requests at a specified time interval.

## Описание / Description

Программа отправляет HTTP POST запросы на указанный URL каждые 30 секунд. Пользователь задает максимальное количество кликов и токен авторизации для заголовка `Authorization`. Программа использует эти значения для формирования тела запроса и заголовка перед отправкой запроса на сервер.

The program sends HTTP POST requests to a specified URL every 30 seconds. The user sets the maximum number of clicks and the authorization token for the `Authorization` header. The program uses these values to form the request body and header before sending the request to the server.

## Использование / Usage

1. Клонируйте репозиторий / Clone the repository:

    ```bash
    git clone https://github.com/ваш-аккаунт/RequestScheduler.git
    cd RequestScheduler
    ```

2. Запустите exe файл /Run exe 

    ```bash
    go build -o requestscheduler
    ./requestscheduler
    ```

3. Следуйте инструкциям на экране / Follow the on-screen instructions:

    - Введите максимальное количество кликов / Enter the maximum number of clicks.
    - Введите токен авторизации в формате `Bearer <your-token>` / Enter the authorization token in the format `Bearer <your-token>`.

Программа начнет отправлять HTTP POST запросы каждые 30 секунд.

The program will start sending HTTP POST requests every 30 seconds.

## Запуск / Running

Для работы программы нужно просто запустить скомпилированный файл `requestscheduler.exe`.

To run the program, simply execute the compiled file `requestscheduler.exe`.


