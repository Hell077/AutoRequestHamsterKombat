


# AutoRequestHamsterKombat

# AutoRequestHamsterKombat - это простая программа на Go, которая отправляет HTTP POST запросы с заданным интервалом времени.

# AutoRequestHamsterKombat is a simple Go program that sends HTTP POST requests at a specified time interval.

## Описание / Description

Программа отправляет HTTP POST запросы на указанный URL каждые 30 секунд. Пользователь задает максимальное количество кликов и токен авторизации для заголовка `Authorization`  Пример токена `17221144434937L8RsSgFbNtwxSRMLtsKJZffEsXn5vk4FcD9XJ8FZNsa6zpZCrTzxr0YkPHHydEt1097316933` . Программа использует эти значения для формирования тела запроса и заголовка перед отправкой запроса на сервер.

The program sends HTTP POST requests to the specified URL every 30 seconds. The user specifies the maximum number of clicks and an authorization token for the `Authorization` header. Example token is `17221144434937L8RsSgFbNtwxSRMLtsKJZffEsXn5vk4FcD9XJ8FZNsa6zpZCrTzxr0YkPHHydEt1097316933` . The program uses these values ​​to form the request body and header before sending the request to the server.

## Использование / Usage

1. Клонируйте репозиторий / Clone the repository:

    ```bash
    git clone https://github.com/Hell077/AutoRequestHamsterKombat
    cd AutoRequestHamsterKombat
    ```

2. Запустите exe файл /Run exe 

    ```bash
    go build -o AutoRequestHamsterKombat
    ./AutoRequestHamsterKombat
    ```

3. Следуйте инструкциям на экране / Follow the on-screen instructions:

    - Введите максимальное количество кликов / Enter the maximum number of clicks.
    - Введите токен авторизации в формате `Bearer <your-token>` / Enter the authorization token in the format `Bearer <your-token>`.

Программа начнет отправлять HTTP POST запросы каждые 30 секунд.

The program will start sending HTTP POST requests every 30 seconds.

## Запуск / Running

Для работы программы нужно просто запустить скомпилированный файл `requestscheduler.exe`.

To run the program, simply execute the compiled file `requestscheduler.exe`.


