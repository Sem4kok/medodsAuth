# medodsAuth test backend task 👨‍💻

## Run Locally

Clone the project

```bash
  git clone https://github.com/Sem4kok/medodsAuth
```

Go to the project directory

```bash
  cd medodsAuth
```

You can run the application in docker from here 

```bash
  make start
```
## API Reference

## register

```http
  POST api/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `first_name` | `string` | **Required**. User name|
| `last_name` | `string` | **Required**. User surname|
| `email` | `string` | **Required**. User email|
| `password` | `string` | **Required**. User password|

#### request:
```JSON
{
  "first_name": "Semyon",
  "last_name": "Kremnev",
  "email": "example@mail.ru",
  "password": "qwerty123"
}
```

#### response:
```JSON
{
  "message": "message"
  "guid": : "user_guid"
}
```

## get
```http
  GET api/token/get?guid=user_guid
```

#### response:
```JSON
{
  "message": "message"
  "guid": : "user_guid"
}

```

## refresh

```http
  POST api/token/refresh
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `access_token`      | `string` | **Required**. user jwt-token|
| `refresh_token`      | `string` | **Required**. user refresh token|

#### request:
```JSON
{
    "access_token": "token"
    "refresh_token": "token"
}
```

#### response:
```JSON
{
    "access_token": "token"
    "refresh_token": "token"
}
```



## Contact me

- Telegram: @peso69
- e-mail:   w1usis@mail.ru
