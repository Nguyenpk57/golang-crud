"# golang-crud" 

**1. database**

```sql

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(40) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime
) 

```

**2. go.mod**

```bash

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.7.0
)

```

**3. Structure**

--auth
--controller
--middlewares
--models
--responses

**4. API


	4.1. server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")
  
  ```json
  
  request: 
  
  {
    "email": "npk2@gmail.com",
    "password": "123456"
  }
  
  response: 
  
  {
    "message": "Success",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzgzNTA1MTAsInVzZXJfaWQiOjl9.f3nIncrf-6jTkL21CTZ2uCTvGSogXqv87rVXuYy-bgc"
}

```

	4.2. server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
  
  
  ```json
  
 response:
 
  {
    "message": "Success",
    "data": [
        {
            "id": 9,
            "name": "npkUpdate2",
            "email": "npk2@gmail.com",
            "password": "",
            "created_at": "2023-03-09T02:39:17Z",
            "updated_at": "2023-03-09T02:42:06Z"
        },
        {
            "id": 10,
            "name": "npk",
            "email": "npk@gmail.com",
            "password": "",
            "created_at": "2023-03-09T07:23:54Z",
            "updated_at": "2023-03-09T07:23:54Z"
        }
    ]
}

```
  
	4.3. server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
  
  ```json
  response:
  
  {
    "message": "Success",
    "data": {
        "id": 10,
        "name": "npk",
        "email": "npk@gmail.com",
        "password": "",
        "created_at": "2023-03-09T07:23:54Z",
        "updated_at": "2023-03-09T07:23:54Z"
    }
}
```
 
	4.4. server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")

```json
  request:
     
  
  {
    "name": "npk",
    "email": "npk@gmail.com",
    "password": "123456"
}

  response:
  
  {
    "message": "Success",
    "data": {
        "id": 10,
        "name": "npk",
        "email": "npk@gmail.com",
        "password": "$2a$10$df.6dzH71Gg0j250YNk/1.ppa1TCQ0l23J7LAXYgfYnKByK8FQeP2",
        "created_at": "2023-03-09T14:23:54.2423551+07:00",
        "updated_at": "2023-03-09T14:23:54.2423551+07:00"
    }
  }
```
 
	4.5. server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")
  
  ```json
  request:
  
  {
    "name": "npkUpdate2",
    "email": "npk2@gmail.com"
  }

  set Authorization: Bearer Token
  Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzgzMzU4OTIsInVzZXJfaWQiOjl9.eE26o1FV80XvVM898qpO_Ns3ZA79455n93obgCpvH6U
  
  response: 
  
  {
    "id": 9,
    "name": "npkUpdate2",
    "email": "npk2@gmail.com",
    "password": "",
    "created_at": "2023-03-09T14:29:53.5490068+07:00",
    "updated_at": "2023-03-09T14:29:53.5490068+07:00"
}
```
  
	4.6. server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.DeleteUser)).Methods("DELETE")
  
  ```json
response:
    
  {
    "message": "Success",
    "data": 2
  }

 ```
