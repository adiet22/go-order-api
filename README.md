# go-order-api
Simple app product order

## 🔗 Description

This Backend Application is used for simple order product, in this application there are two models / ERD Schema likes User / Costumer & Products.
Also have several features like JWT, Authentification & Authorization.
There are 3 main modules :
1. Customer Management (Get with paginate, Get Detail, Insert, Update, Delete,
Search)
2. Order Management (Get with paginate, Get Detail, Insert, Update, Delete,
Search)
3. Authentikasi Management (Get Login Data, Insert Login Data)

Notes :
1. I'am using UUID for user_id, don't forget to create extenxion in SQL console after create the database with this query below :
```bash
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

2. In this application there are two types of users (Roles). admins and costumer. 
   Admin can do *Costumer Management* but Role Costumer can't, Registration page can only register Costumer roles, Admins can only be registered through seeding data with this command below :

```bash
go run . seed
```
<h1 align="center">
 ERD (Entity Relation Database)
</h1>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1665874871/erd_c15gne.png" alt="erd.jpg" /></p>

## POST /login
_Purpose to login as Admin's Role_
```
{
  "email": "admin@gmail.com",
  "password": "admin12345678"
}
```

## Several command you must know in this app :
```bash
1. go run . serve //to run the app / server
2. go run . migrate -u //for database migration
# or
go run . migrate -d //for rollback
3. go run . seed // to seeding data Role admin if u want Email : "admin@gmail.com" Pass : admin12345678
```

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone git@github.com:adiet95/go-order-api.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```
3. Database Migration and Rollback

```bash
go run main.go migrate --up //for database migration
# or
go run main.go migrate --down //for rollback
```
4. Add Env

```sh
  DB_USER = Your DB User
  DB_HOST = Your DB Host
  DB_NAME = Your DB Name
  DB_PASS = Your DB Password
  JWT_KEYS = Your JWT Key
  PORT = Your Port
```
5. Run the app

```bash
go run . serve
```

## 🔗 RESTful endpoints
### POST /register

> Create new user

_Request Header_
```
not needed
```

_Request Body_
```
{
  "user_name": <your username> (STRING),
  "email": <your email> (STRING),
  "password": <your password> (STRING),
  "address": <your address> (STRING),
  "phone": <your phone> (STRING)
}
```

#### Success Response: ####
_Response (201 - Created)_
```
{
  "id": <id>,
  "user_name": <your username>,
  "email": <your email>,
  "password": <your encrypted password>,
  "address": <your address>,
  "phone": <your phone>,
  "updatedAt": <date>,
  "createdAt": <date>
}
```

#### Error Response: ####
_Response (400 - Bad Request)_
```
[
  "message": <detail message>
]
```

_Response (409 - conflict)_
```
{
  "message": "Email Already registered!"
}
```

_Response (500 - Internal Server Error)_
```
{
  "message": "Internal Server Error"
}
```

### POST /login

> Process Login

_Request Header_
```
not needed
```

_Request Body_
```
{
  "email": <your email> (STRING),
  "password": <your password> (STRING)
}
```

#### Success Response: ####
_Response (200 - Ok)_
```
{
  "token": <your access token>
}
```

#### Error Response: ####

_Response (400 - Bad Request)_
```
[
  "message": <detail message>
]
```

_Response (404 - Not Found)_
```
{
  "message": "user not registered!"
}
```

_Response (500 - Internal Server Error)_
```
{
  "message": "Internal Server Error"
}
```

## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [gorilla/mux](https://github.com/gorilla/mux): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS


## 🚀 About Me

- Linkedin : [Achmad Shiddiq](https://www.linkedin.com/in/achmad-shiddiq-alimudin/)
