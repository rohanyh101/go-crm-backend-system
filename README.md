# CRM Backend System

This project is a backend system for a Customer Relationship Management (CRM) application. It is built using Go with the Gin framework, and it integrates MongoDB for database management. The application is containerized using Docker, ensuring easy deployment and scalability. The backend supports full CRUD operations, JWT-based authentication, and includes comprehensive security measures such as role-based access control.

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

### Example CRUD operations:

### health check
1. **health check/server-status**,
    ```bash
      #request
      curl --location --request GET 'localhost:8080/health' \
      --header 'Content-Type: application/json'

      #response
      {"success":"server is up and running..."}
    ```

### **user routes:**
1. **signup user**,
   ```bash
     #request
     curl --location --request POST 'localhost:8080/users/signup' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "name": "rohan", "email": "rohan@gmail.com", "password": "12121212", "role": "ADMIN" }'
   
     #response
    {"InsertedID":"66cc46295c7a983d35a8b228"}
   ```

2. **login user**,
    ```bash
     #request
     curl --location --request POST 'localhost:8080/users/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "email": "rohan@gmail.com", "password": "12121212" }'
   
     #response
    {"InsertedID":"66cc46295c7a983d35a8b228"}
   ```
    
3. **update user**,
   ```bash
   #request
   curl --location --request PUT 'localhost:8080/users/66cc87ca6cc87479e44f1443' \
   --header 'Content-Type: application/json' \
   --data-raw '{ "email": "manoj@domain.com" }' \
   --header 'token: <token>'
   #response
   { "MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null }
   
   ```

4. **delete user**,
   ```bash
   #request
   curl --location --request DELETE 'localhost:8080/users/66cc8d343557fdb75b7a32b2' \
    --header 'Content-Type: application/json' \
    --header 'token: <token>' | jq

   #response
    {  "DeletedCount": 1 }
   ```

### **customer routes**

1. **signup customer**
   ```bash
   #request
   curl --location --request POST 'localhost:8080/customers/signup' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "name": "rahul", "email": "rahul@domain.com", "password": "12121212", "comapny": "TCS", "phone_no": "2222" }'
   
   #response
    {"InsertedID":"66cc8d343557fdb75b7a32b2"}
   ```

2. **login customer**
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

3. **getAll customers**
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

  4. **get single customer**
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

5. **update customer**
    ```bash
    #request
    curl --location --request PUT 'localhost:8080/customers/66cc8d343557fdb75b7a32b2' \
    --header 'Content-Type: application/json' \
    --data-raw '{ "name": "rahulp" } \
    --header 'token: <token>' | jq

    #respnse
    { "MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null }
    ```
10. **delete the customer by ID**
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

### 
11.  **


### Email Notifications
The application automatically sends email notifications when interactions are created. The email content can be configured in the code, and the SMTP settings must be provided in the .env file.

### Deployment
The project is deployed on AWS. It can be accessed via the provided AWS endpoint. The Docker image is also available on Docker Hub:

Docker Image: 
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
