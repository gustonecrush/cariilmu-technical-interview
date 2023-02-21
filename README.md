# DOCUMENTATION

## LIST OF CONTENT

- <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#about">About</a>
- <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#how-to-create-project">How to create project</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#set-up-project">Set up project</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#set-up-laravel-sanctum-package">Set up Laravel Sanctum package</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#create-auth-controller">Create auth controller</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#define-the-route">Define the route</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#trials">Trials</a>
- <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#auth-endpoint">Auth Endpoint</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#register">Register</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#login">Login</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#logout">Logout</a>
   - <a href="https://github.com/gustonecrush/api-token-laravel-sanctum/blob/main/README.md#user">User</a>
   
## ABOUT
Learn how to create token based authentication for rest api using laravel sanctum. This project provides an API endpoint for registering, logging in, logging out, and getting data which is the implementation of one of Laravel Sanctum's uses, namely the API Token. In the HTTP Request process for the registration and login process, the system will generate an api token. This API token will later be used for the process of retrieving user data and also the logout process.

## HOW TO CREATE PROJECT

### SET UP PROJECT

- Download this project or clone this repo and save to your local
- Run `composer install` to install all dependencies needed, if you have not the composer you can go to <a href="https://getcomposer.org
">here</a> to see how to install it
- Download and install Laravel in your computer, go to <a href="https://laravel.com/">here</a> to see how to install it
- Create new laravel project, see the <a href="https://laravel.com/docs/9.x/installation">documentation</a>
- Open your computer server to run your server, then create new database on your database
- Configure your .env file, go to database section, and configure the database according to the database you have made before, the user, and the password you are using as below. If you don't make database yet, so make new database in your server. Then configure in .env file.
  ```
      DB_CONNECTION=mysql
      DB_HOST=your_host
      DB_PORT=your_port
      DB_DATABASE=your_db_name
      DB_USERNAME=your_username
      DB_PASSWORD=your_password
  ```

### TRIALS

- To test it, we first run our project first. Open terminal, then run the following command
  ```
    go mod init
  ```
  ```
    go run main.go
  ```
- Then you can open your postman, and try the endpoint below

## ADMIN AUTH ENDPOINT

### REGISTER

Request :

- Method : POST
- Endpoint : `/api/register`
- Header :
  - Accept: application/json
- Body :

```json
{
    "name": "string",
    "email": "string, email",
    "password": "string",
}
```

Response :

```json
{
  "status": "boolean",
  "status_code": "integer",
  "message": "string",
  "data": {
      "id": "integer, unique",
      "name": "string",
      "email": "string",
      "created_at": "timestamp",
      "updated_at": "timestamp"
  },
  "access_token": "string, unique",
  "token_type": "string",
}
```

### LOGIN

Request :

- Method : POST
- Endpoint : `/api/login`
- Header :
  - Accept: application/json
- Body :

```json
{
    "email": "string, email",
    "password": "string",
}
```

Response :

```json
{
  "status": "boolean",
  "status_code": "integer",
  "message": "string",
  "access_token": "string, unique",
  "token_type": "string",
}
```

### LOGOUT

Request :

- Method : POST
- Endpoint : `/api/logout`
- Header :
  - Accept: application/json
  - Autohorization: Bearer token

Response :

```json
{
  "status": "boolean",
  "status_code": "integer",
  "message": "string",
  "access_token": "string, unique",
  "token_type": "string",
}
```

### USER

Request :

- Method : GET
- Endpoint : `/api/user`
- Header :
  - Accept: application/json
  - Autohorization: Bearer token

Response :

```json
{
  "status": "boolean",
  "status_code": "integer",
  "message": "string",
  "data": {
      "id": "integer, unique",
      "name": "string",
      "email": "string",
      "created_at": "timestamp",
      "updated_at": "timestamp"
  },
  "access_token": "string, unique",
  "token_type": "string",
}
```

### COURSE

### COURSE CATEGORY

### USER COURSES