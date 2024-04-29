<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">Rest Go</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()

</div>

---

<p align="center"> A simple GO REST API with Gin and Gorm
    <br> 
</p>

## 📝 Table of Contents

- [Getting Started](#getting_started)
- [Prerequisites](#prerequisites)
- [Installing](#installing)
- [Running the server](#running-the-server)
- [Running the tests](#running-the-tests)
- [Project Structure](#project-structure)



## 🏁 Getting Started <a name = "getting_started"></a>


### Prerequisites
- go 1.22
- air
- postgresql


### Installing


clone the repository

```bash
git clone https://github.com/tst-irfan/rest_go
```

cd into the project directory

```bash
cd rest_go
```

install go dependencies

```bash
go mod download
```

create a .env file in root directory
```bash
touch .env
```

add the following to the .env files

```
# .env
ENV=development
```

duplicate the .env.example file to .env.development and .env.test then fill in the required fields

```bash
cp .env.example .env.development
cp .env.example .env.test
cp .env.example .env.staging
cp .env.example .env.production
```

### Running the server

you can run the server by running the following command

```bash
go run main.go
```

alternatively you can use air to get live reload

```bash
air -c .air.toml
```

End with an example of getting some data out of the system or using it for a little demo.

### Running the tests

```bash
go test ./...
```

<br>

## Project Structure

```
.
├── app
│   ├── controllers
│   ├── middlewares
│   ├── router
│   ├── models
│   ├── services
│   ├── utils
│   ├── types
├── config
│   ├── config.go
├── db
│   ├── query.go
│   ├── setup.go
├── initializers
│   ├── application.go
│   ├── auto.migrate.go
├── tests
│   ├── factories
│   ├── modles
├── main.go
├── .env
├── .env.development
├── .env.test
├── .env.staging
├── .env.production
```
### Directories

- **controllers:** This directory houses the controllers responsible for managing incoming requests and generating responses.
- **models:** Here, you'll find the data models utilized within the application to represent entities and their relationships.
- **services:** Contains the services layer, which encapsulates the business logic of the application, promoting modularity and maintainability.
- **utils:** Within this directory, you'll discover utility functions that offer reusable functionalities across different parts of the application.
- **types:** Stores type definitions that are shared and used throughout the application, ensuring consistency and clarity.
- **middlewares:** This directory holds middleware functions used to intercept and process incoming requests before they reach the main application logic.
- **routers:** Here, you'll find router configurations responsible for directing incoming requests to the appropriate controllers or handlers.

### Files
- **main.go:** Serves as the entry point for the application, orchestrating the initialization and startup processes.
- **initializer/application.go:** This file contains code responsible for initializing the application, such as setting up database connections, initializing logging frameworks, or configuring global settings. It ensures that the application is properly set up and ready to handle incoming requests.  
- **initializer/auto_migrate.go:** This file contains code to automatically migrate database schemas or perform database migrations on application startup. It's responsible for ensuring that the database schema is up-to-date with the application's model definitions, making database management easier and more efficient.
- **db/query.go:** Contains the `QueryHelper` struct and database query methods for interacting with the database. It encapsulates common database operations like querying, creating, updating, and deleting records.
- **db/setup.go:** Contains database setup and initialization logic, such as establishing connections to the database, configuring database settings, and preparing the database for use by the application. It ensures that the database environment is properly configured and ready for use by other components of the application.
- **.env:** Root environment file containing the `ENV` variable, which specifies the environment the application runs in. Valid values are `development`, `test`, `staging`, and `production`. Other environment-specific configurations can be set based on this variable.
- **.env.development:** Development environment file containing configurations specific to the development environment. It may include settings tailored for local development, such as debug options or development database credentials.
- **.env.test:** Test environment file containing configurations specific to the test environment. It includes settings required for running tests, such as test database credentials or test-specific configurations.
- **.env.staging:** Staging environment file containing configurations specific to the staging environment. It includes settings required for staging deployments, such as staging database credentials or staging-specific configurations.
- **.env.production:** Production environment file containing configurations specific to the production environment. It includes settings required for production deployments, such as production database credentials or production-specific configurations.


## Postman Collection
  
  [![Run in Postman](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/30788320/2sA3BuWULJ)

## Authors
👤 [**Irfan Azhar**](http://github.com/tst-irfan)


