<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">Rest Go</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()

</div>

---

<p align="center"> A simple REST API in Go with Gin and Gorm
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

cd into the test directory

```bash
cd tests/models

```

run the tests

```bash
go test
```

<br>

## Project Structure

```
.
├── app
│   ├── controllers
│   ├── models
│   ├── services
│   ├── utils
│   ├── types
├── config
│   ├── middlewares
│   ├── router
│   ├── middleware.go
│   ├── router.go
├── tests
main.go
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

- **middleware.go:** This file defines and initializes middleware functions used in request processing, enhancing the application's extensibility and flexibility.
- **router.go:** Defines and initializes router configurations, which play a crucial role in mapping incoming requests to the appropriate controller methods or handlers.
- **main.go:** Serves as the entry point for the application, orchestrating the initialization and startup processes.




