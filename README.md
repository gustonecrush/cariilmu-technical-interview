# DOCUMENTATION

## HOT TO SET UP PROJECT

- Download this project or clone this repo and save to your local
- Open your computer server to run your server, then create new database on your database
- Configure your .env file, go to database section, and configure the database according to the database you have made before, the user, and the password you are using as below. If you don't make database yet, so make new database in your server. Then configure in .env file.
  ```
      DB_DRIVER=mysql
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