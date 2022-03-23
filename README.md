# Backend App - Books REST API

## Technical Test - PT Atmatech Global Informatika

_Developed By : [Robertus Kristianto Sukoco](https://github.com/krissukoco)_

### **Description**

REST API for GET, POST, PUT, and DELETE books. The routes are protected by requiring Authorization Bearer token. Tech stack: Go (Golang) and PostgreSQL<br><br>

---

## Run the app

---
Using PostgreSQL:
```javascript
// Import database structure and initial values
psql -h localhost -U <username> -d <databasename> -f database.sql

// IMPORTANT: Edit public.env in accordance to your Postgres DB params
// e.g. in public.env file:
DB_USERNAME=postgres
DB_PASSWORD=mysupersecretpassword
DB_HOST=localhost
DB_NAME=mydatabase

// Running go file
go run main.go
```
<br>

___
## Endpoints
___
Host: `http://localhost:3000/`<br><br>
**Authentication (User)**
* Register a user<br>
    `POST /auth/register`<br>

* Login a user<br>
    `POST /auth/login`<br><br>
    Req body for both endpoints:
    ```json
    {
        "username": <username>,
        "password": <password>
    }
    ```
    Response (for both):
    ```json
    {
        "status": 200,
        "message": "Login Successful",
        "data": {
            "is_authenticated": true,
            "access_token": "eyxxxxxxxxxxxx.xxxxxx"
        }
    }
    ```
**Books**<br>
These routes require **Authorization Bearer Token** on Request Header. In the form of:
    `"Authorization: "Bearer eyxxxxxx.xxxxx"` which can be obtained by Registering or Logging In.
* Get all books, with Limit default to 10 books, and Page default to 1<br>
    `GET /books?limit=<limit:int>?page=<page:int>`<br>

* Get a book by ID<br>
    `GET /books/<id:int>`

* Post (create) a book<br>
    `POST /books`

* Update (PUT) a book<br>
    `PUT /books/<id:int>`

* Delete a book<br>
    `DELETE /books/<id:int>`

<br>Request body for POST and PUT requests as follows:<br>
```json
{
    "title": <book title>,
    "description": <book description>,
    "content": <book contents>
}
```
<br>Example Response: `GET /books?limit=5&page=1`<br>
```json
{
    "status": 200,
    "message": "GET ALL BOOKS",
    "data": [
        {
            "id": 1,
            "title": "My Book",
            "description": "My Description",
            "created_at": "2022:01:01 10:00:00 +07:00"
            ....
        },
        {
            .....
        },
        .....
    ]
}
```