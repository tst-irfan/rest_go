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
