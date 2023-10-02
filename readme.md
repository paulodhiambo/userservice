# Go Ecom

![ci-test](https://github.com/paulodhiambo/ecom/actions/workflows/ci.yml/badge.svg)
![badmath](https://img.shields.io/github/languages/top/paulodhiambo/ecom)

## Description

Go UserService is a simple user management project implemented in Go.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Credits](#credits)
- [License](#license)
- [Badges](#badges)
- [Features](#features)
- [How to Contribute](#how-to-contribute)
- [Tests](#tests)

## Installation

To install and run the project locally, follow these steps:

1. Clone the repository: `git clone https://github.com/paulodhiambo/userservice.git`
2. Navigate to the project directory: `cd userservice`
3. Install dependencies: `go mod tidy`
4. Set up the database and configure the connection details in the configuration file.
5. Run the application: `go run main.go`

## Usage

To use the Go userservice API, follow the API documentation and interact with the available endpoints using tools like
cURL,
Postman, or any other HTTP client.

For example, to retrieve a list of users, send a GET request to `/api/users`:

```
GET /api/user
```

To add a user, send a POST request to `/api/user` with the payload below:

```
POST /api/users
{
    "name": "wayne",
    "email": "wayne@mail.id",
    "password": "plain_password",
    "role_id": 1
}
```

For more detailed usage instructions and examples, please refer to the API documentation.

## Credits

- [Odhiambo Paul](https://github.com/paulodhiambo) - Project Developer

## License

This project is licensed under the [MIT License](LICENSE).

## Features

- User management (CRUD operations)
- Role management (CRUD operations)

## How to Contribute

Contributions to Go Ecom are welcome! If you would like to contribute, please follow the guidelines outlined
in [CONTRIBUTING.md](CONTRIBUTING.md).

## Migrate database

```shell
make migrateup
```

## Tests

To run the tests for this project, use the following command:

```shell
make test
```

> Make sure you have the necessary dependencies installed before running the tests.