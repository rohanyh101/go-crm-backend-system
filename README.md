# CRM Backend System

This project is a backend system for a Customer Relationship Management (CRM) application. It is built using Go with the Gin framework, and it integrates MongoDB for database management. The application is containerized using Docker, ensuring easy deployment and scalability. The backend supports full CRUD operations, JWT-based authentication, and includes comprehensive security measures such as role-based access control.

## System Design

![Screenshot (138)](https://github.com/user-attachments/assets/8dfff095-567c-4a9c-8372-5fc46d29583c)



## Features

- **User and Customer Management:**
  - Full CRUD operations for users and customers.
  - JWT authentication for both users and customers.
  - Role-based access control ensuring proper authorization.

- **Interactions:**
  - Manage interactions between users and customers, including meetings, tasks, and follow-ups.
  - Secure handling of sensitive data.

- **Security:**
  - Role-based access control (RBAC).
  - Password encryption and secure data storage.

- **Email Notifications:**
  - Automatic email notifications for interactions.
  - Configurable SMTP settings for email service.

- **Deployment:**
  - Dockerized application for seamless deployment.
  - Deployed on AWS for production use.

## Technologies Used

- **Backend:** Go (Gin framework)
- **Database:** MongoDB
- **Authentication:** JWT
- **Containerization:** Docker
- **Deployment:** AWS
- **Email Service:** SMTP

## Installation

### Prerequisites

- Docker
- Docker Compose

### Environment Variables

Create a `.env` file in the root directory of the project with the following variables:

```bash
PORT=8080
MONGO_URI=mongodb://mongo:27017/crm_database
SECRET_KEY=your_secret_key
USER_SECRET_KEY=your_user_secret_key
CUSTOMER_SECRET_KEY=your_customer_secret_key

SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_MAIL=your_email@example.com
SMTP_PASSWORD=your_smtp_password
```

### Docker Setup
1. Clone the repository:
```bash
  git clone https://github.com/rohanyh101/matrice.ai-assignment.git <project_name>
  cd <project_name>
```

2. Run the application using Docker Compose:

```bash
docker-compose up --build
```

3. The application will be available at `http://localhost:8080`.

### API Endpoints
**Auth Routes**
 - Register User: POST /auth/register
 - Login User: POST /auth/login
 - Register Customer: POST /auth/register-customer
 - Login Customer: POST /auth/login-customer

### User Routes
 - Get Users: GET /users
 - Get User by ID: GET /users/:id
 - Update User: PATCH /users/:id
 - Delete User: DELETE /users/:id

### Customer Routes
 - Get Customers: GET /customers
 - Get Customer by ID: GET /customers/:id
 - Update Customer: PATCH /customers/:id
 - Delete Customer: DELETE /customers/:id

### Interaction Routes
 - Get Interactions by User ID: GET /interactions/user/:user_id
 - Get Interaction by Interaction ID: GET /interactions/:interaction_id
 - Create Interaction: POST /interactions
 - Update Interaction: PATCH /interactions/:interaction_id
 - Delete Interaction: DELETE /interactions/:interaction_id

### **gin logs**,
```bash
  Connected to MongoDB!
  Connected to MongoDB!
  Connected to MongoDB!
  Connected to MongoDB!
  Connected to MongoDB!
  [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
   - using env:   export GIN_MODE=release
   - using code:  gin.SetMode(gin.ReleaseMode)
  
  [GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (2 handlers)
  [GIN-debug] POST   /users/signup             --> github.com/roh4nyh/matrice_ai/routes.AuthRoutes.UserSignUp.func1 (2 handlers)
  [GIN-debug] POST   /users/login              --> github.com/roh4nyh/matrice_ai/routes.AuthRoutes.UserLogIn.func2 (2 handlers)
  [GIN-debug] POST   /customers/signup         --> github.com/roh4nyh/matrice_ai/routes.AuthRoutes.CustomerSignUp.func3 (2 handlers)
  [GIN-debug] POST   /customers/login          --> github.com/roh4nyh/matrice_ai/routes.AuthRoutes.CustomerLogIn.func4 (2 handlers)
  [GIN-debug] GET    /health                   --> main.main.func1 (2 handlers)
  [GIN-debug] GET    /users                    --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.GetUsers.func2 (3 handlers)
  [GIN-debug] GET    /users/:user_id           --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.GetUser.func3 (3 handlers)
  [GIN-debug] PUT    /users/:user_id           --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.UpdateUser.func4 (3 handlers)
  [GIN-debug] DELETE /users/:user_id           --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.DeleteUser.func5 (3 handlers)
  [GIN-debug] GET    /users/meet/              --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.GetAllInteractions.func6 (3 handlers)
  [GIN-debug] GET    /user/meet/               --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.GetInteractionsByUserID.func7 (3 handlers)
  [GIN-debug] POST   /users/meet/:customer_id  --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.CreateInteractionAndSendEmail.func8 (3 handlers)
  [GIN-debug] DELETE /users/meet/:interaction_id --> github.com/roh4nyh/matrice_ai/routes.UserRoutes.DeleteInteraction.func9 (3 handlers)
  [GIN-debug] GET    /customers                --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.GetCustomers.func2 (4 handlers)
  [GIN-debug] GET    /customers/:customer_id   --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.GetCustomer.func3 (4 handlers)
  [GIN-debug] PUT    /customers/:customer_id   --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.UpdateCustomer.func4 (4 handlers)
  [GIN-debug] DELETE /customers/:customer_id   --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.DeleteCustomer.func5 (4 handlers)
  [GIN-debug] GET    /customers/tickets/       --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.GetAllTickets.func6 (4 handlers)
  [GIN-debug] POST   /customers/ticket/:interaction_id --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.CreateTicket.func7 (4 handlers)
  [GIN-debug] GET    /customers/ticket/:user_id --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.GetTicketsByUserID.func8 (4 handlers)
  [GIN-debug] PUT    /customers/ticket/:ticket_id --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.UpdateTicket.func9 (4 handlers)
  [GIN-debug] DELETE /customers/ticket/:ticket_id --> github.com/roh4nyh/matrice_ai/routes.CustomerRoutes.DeleteTicket.func10 (4 handlers)
  [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
  Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
  [GIN-debug] Listening and serving HTTP on :8080
```

### Example CRUD operations:

### health check
1. **health check/server-status => `GET    /health`**,
    ```bash
      #request
      curl --location --request GET 'localhost:8080/health' \
      --header 'Content-Type: application/json'

      #response
      {"success":"server is up and running..."}
    ```

### **User Routes:**
1. **signup user => `POST   /users/signup`**,
   ```bash
     #request
     curl --location --request POST 'localhost:8080/users/signup' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "name": "rohan", "email": "rohan@gmail.com", "password": "12121212", "role": "ADMIN" }'
   
     #response
    {"InsertedID":"66cc46295c7a983d35a8b228"}
   ```

2. **login user => `POST   /users/login`**,
    ```bash
     #request
     curl --location --request POST 'localhost:8080/users/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "email": "rohan@gmail.com", "password": "12121212" }'
   
     #response
    {"InsertedID":"66cc46295c7a983d35a8b228"}
   ```
    
3. **update user => `PUT /users/:user_id`**,
   ```bash
   #request
   curl --location --request PUT 'localhost:8080/users/66cc87ca6cc87479e44f1443' \
   --header 'Content-Type: application/json' \
   --data-raw '{ "email": "manoj@domain.com" }' \
   --header 'token: <token>'
   #response
   { "MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null }
   
   ```

4. **delete user => `DELETE /users/:user_id`**,
   ```bash
   #request
   curl --location --request DELETE 'localhost:8080/users/66cc8d343557fdb75b7a32b2' \
    --header 'Content-Type: application/json' \
    --header 'token: <token>' | jq

   #response
    {  "DeletedCount": 1 }
   ```

### **customer routes**

1. **signup customer => `POST   /customers/signup`**
   ```bash
   #request
   curl --location --request POST 'localhost:8080/customers/signup' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "name": "rahul", "email": "rahul@domain.com", "password": "12121212", "comapny": "TCS", "phone_no": "2222" }'
   
   #response
    {"InsertedID":"66cc8d343557fdb75b7a32b2"}
   ```

2. **login customer => `POST   /customers/login`**
   ```bash
   #request
   curl --location --request POST 'localhost:8080/customers/login' `
    --header 'Content-Type: application/json' `
    --data-raw '{ "email": "mynewpc2513@gmail.com", "password": "12121212" }' | jq

    #response
    {
      "id": "66cc9d35a7c3ac465fab3599",
      "name": "charan",
      "email": "user@gmail.com",
      "password": "$2a$15$ov28wsJ6vnS9rGU/vV93V.DjVgiIOOAevquF5bR6pREe7zDovcYHG",
      "token": "<token>",
      "created_at": "2024-08-26T15:20:21Z",
      "updated_at": "2024-08-26T15:23:37Z",
      "customer_id": "66cc9d35a7c3ac465fab3599"
    }
   ```

3. **get all customers => `GET    /customers`**
      ```bash
      #request
      curl --location --request GET 'localhost:8080/customers/' \
      --header 'Content-Type: application/json' \
      --header 'token: <token>' | jq

       #response
      [
        {
          "id": "66cc766447ef7e236f14adb3",
          "name": "manish",
          "email": "manish@domain.com",
          "password": "$2a$15$.5RyIOA/CrCSp46b2bPIRur09j7jMAHPGgm32QKfD1dIamJK95Vce",
          "token": "<token>",
          "created_at": "2024-08-26T12:34:44Z",
          "updated_at": "2024-08-26T13:23:37Z",
          "customer_id": "66cc766447ef7e236f14adb3"
        },
        {
          "id": "66cc8ea03557fdb75b7a32b3",
          "name": "rahul",
          "email": "rahul@domain.com",
          "password": "$2a$15$y5.LllEgmns4adXt3zFM2emBTSjXlybN4lGXfiCuxZTDc8EX4kJkm",
          "token": "<token>",
          "created_at": "2024-08-26T14:18:08Z",
          "updated_at": "2024-08-26T14:18:08Z",
          "customer_id": "66cc8ea03557fdb75b7a32b3"
        }
      ]
      ```

  4. **get single customer => `GET    /customers/:customer_id`**
     ```bash
     #request
     curl --location --request GET 'localhost:8080/customers/66cc8d343557fdb75b7a32b2' \
     --header 'Content-Type: application/json' \
     --data-raw '{ "name": "rahulp" }' \
     --header 'token: <token>' | jq


     #response
     {
       "id": "66cc8d343557fdb75b7a32b2",
       "name": "rahul",
       "email": "rahul@domain.com",
       "password": "$2a$15$FWfUlf0q2YWIlJiVQv3kquO0.PWsoEqEqp9wt.3k7WnYK4c5aQ4ym",
       "token": "<token>",
       "created_at": "2024-08-26T14:12:04Z",
       "updated_at": "2024-08-26T14:12:04Z",
       "customer_id": "66cc8d343557fdb75b7a32b2"
     }
     ```

5. **update customer => ` PUT    /customers/:customer_id`**
    ```bash
    #request
    curl --location --request PUT 'localhost:8080/customers/66cc8d343557fdb75b7a32b2' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "name": "rahulp" } \
    --header 'token: <token>' | jq

    #respnse
    { "MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null }
    ```
6. **delete the customer by ID => `DELETE /customers/:customer_id`**
    ```bash
    #request
    curl --location --request DELETE 'localhost:8080/customers/66cc766447ef7e236f14adb3' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "name": "rahulp" }' \
    --header 'token: <token>' | jq

    #response
    {
      "error": "UnAuthenticated to access this resource"
    }
    ```

### **User Services** 
1. **create interaction => `POST   /users/meet/:customer_id`**
   ```bash
   #request
    curl --location --request POST 'localhost:8080/users/meet/66cc9d35a7c3ac465fab3599' \
    --header 'Content-Type: application/json' \
    --header 'token: <token>' \
    --data-raw '{ "title": "demo title", "description": "demo description", "start_time": "2024-08-27T20:03:00Z" }' | jq

   #response
   { "InsertedID": "66ccc4d1e3f9cd0e36da4878 }
   ```

2. **get interaction related to current user => `GET    /user/meet/`**
    ```bash
    #request
    curl --location --request GET 'localhost:8080/user/meet/' \
    --header 'Content-Type: application/json' \
    --header 'token: <token>' | jq
  
    #response
    [
      {
        "id": "66ccac66c7b0b27e26097e60",
        "user_id": "66cc9d6ca7c3ac465fab359a",
        "customer_id": "66cc9d35a7c3ac465fab3599",
        "title": "demo title",
        "description": "demo description",
        "start_time": "2024-08-28T20:03:00Z",
        "created_at": "2024-08-26T16:25:10Z",
        "updated_at": "2024-08-26T16:25:10Z",
        "interaction_id": "66ccac66c7b0b27e26097e60"
      }
    ]
    ```

3. **get all interaction, admin route => `GET    /users/meet/`**
    ```bash
    #request
    curl --location --request GET 'localhost:8080/users/meet/' \
    --header 'Content-Type: application/json' \
    --header 'token: <token>' | jq
  
    #response
   [
    {
      "id": "66ccac66c7b0b27e26097e60",
      "user_id": "66cc9d6ca7c3ac465fab359a",
      "customer_id": "66cc9d35a7c3ac465fab3599",
      "title": "demo title",
      "description": "demo description",
      "start_time": "2024-08-28T20:03:00Z",
      "created_at": "2024-08-26T16:25:10Z",
      "updated_at": "2024-08-26T16:25:10Z",
      "interaction_id": "66ccac66c7b0b27e26097e60"
    },
    {
      "id": "66ccc4d1e3f9cd0e36da4878",
      "user_id": "66cc9d6ca7c3ac465fab359a",
      "customer_id": "66cc9d35a7c3ac465fab3599",
      "title": "demo title",
      "description": "demo description",
      "start_time": "2024-08-27T20:03:00Z",
      "created_at": "2024-08-26T18:09:21Z",
      "updated_at": "2024-08-26T18:09:21Z",
      "interaction_id": "66ccc4d1e3f9cd0e36da4878"
    },
    {
      "id": "66ccc4e6e3f9cd0e36da4879",
      "user_id": "66cc9d6ca7c3ac465fab359a",
      "customer_id": "66cc9d35a7c3ac465fab3599",
      "title": "demo title",
      "description": "demo description",
      "start_time": "2024-08-29T20:03:00Z",
      "created_at": "2024-08-26T18:09:42Z",
      "updated_at": "2024-08-26T18:09:42Z",
      "interaction_id": "66ccc4e6e3f9cd0e36da4879"
    }
    ]
    ```

4. **delete interaction => `DELETE   /users/meet/:interaction_id`**
    ```bash
    #request
    curl --location --request DELETE 'localhost:8080/users/meet/66ccad18fcf4d6bf088a55fd' \
    --header 'Content-Type: application/json' \
    --header 'token: <token>' | jq

    #response
    {
      "message": "Interaction deleted successfully"
    }
    ```

### **Customer Services** 
1. **raise ticket => ` POST   /customers/ticket/:interaction_id`**
   ```bash
   #request
   curl --location --request POST 'localhost:8080/customers/ticket/66ccc4d1e3f9cd0e36da4878' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "description": "demo description", "status": "open" }' \
    --header 'token: <token>' | jq

   #response
    {
      "InsertedID": "66ccc6caeecebcb36684dbe8"
    }
   ```

2. **get all tickets => ` GET    /customers/tickets/`**
   ```bash
   #request
   curl --location --request GET 'localhost:8080/customers/tickets/' `
    --header 'Content-Type: application/json' `
    --header 'token: <token>' | jq

    #response
    [
      {
        "id": "66ccccba62c4b3bbe341d2df",
        "interaction_id": "66ccc98b5f673a3fce7ebecb",
        "customer_id": "66ccc8da5f673a3fce7ebeca",
        "status": "in_progress",
        "description": "new desc",
        "created_at": "2024-08-26T18:43:06Z",
        "updated_at": "2024-08-26T20:24:18.316Z",
        "ticket_id": "66ccccba62c4b3bbe341d2df"
      },
      {
        "id": "66cce6cad8cd633786e93b75",
        "interaction_id": "66ccc4d1e3f9cd0e36da4878",
        "customer_id": "66cc9d35a7c3ac465fab3599",
        "status": "open",
        "description": "demo description",
        "created_at": "2024-08-26T20:34:18Z",
        "updated_at": "2024-08-26T20:34:18Z",
        "ticket_id": "66cce6cad8cd633786e93b75"
      }
    ]
   ```
3. **update ticket => `PUT     /customers/ticket/:ticket_id`**
   ```bash
   #request
   curl --location --request PUT 'localhost:8080/customers/ticket/66cce6cad8cd633786e93b75' `
    --header 'Content-Type: application/json' `
    --data-raw '{ "description": "new demo description", "status": "solved" }' `
    --header 'token: <customer_token>' | jq

   #response
    {
      "message": "ticket updated successfully"
    }
   ```

4. **delete ticket => `DELETE     /customers/ticket/:ticket_id`**
   ```bash
   #request
   curl --location --request DELETE 'localhost:8080/customers/ticket/66cce6cad8cd633786e93b75' `
    --header 'Content-Type: application/json' `
    --header 'token: <customer_token>' | jq

   #response
    {
      "message": "ticket deleted successfully"
    }
   ```

### Email Notifications
The application automatically sends email notifications when interactions are created. The email content can be configured in the code, and the SMTP settings must be provided in the .env file.

### Deployment
The project is deployed on AWS. It can be accessed via the provided AWS endpoint. The Docker image is also available on Docker Hub:

Docker Image: https://hub.docker.com/r/rohanyh/matriceai_crm/tags
Live Application: 

### Challenges
During the development process, I encountered an unexpected power outage, which posed a significant challenge. Despite this, I managed to implement all the requested features and ensured the project was ready for deployment.

### Future Improvements
Enhancing the security measures with more granular permissions.
Implementing a more robust logging and monitoring system.
Adding unit and integration tests for better code coverage.

### Contributing
If you'd like to contribute to this project, please fork the repository and submit a pull request. All contributions are welcome!

### Contact
For any questions or inquiries, please contact or make a pull request
