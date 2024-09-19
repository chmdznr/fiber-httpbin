# Fiber HTTPBin API

A web service similar to httpbin.org built with Fiber.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/chmdznr/fiber-httpbin.git
    cd fiber-httpbin
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `localhost:3000`.

## API Endpoints

### /anything

- **GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD**: Echoes the request data.
    - **Response**: `200 OK`
    - **Schema**: `models.AnythingResponse`

### /delay

- **GET**: Delays the response by a specified number of seconds.
    - **Query Parameter**: `seconds` (integer, required)
    - **Response**: `200 OK`
    - **Schema**: `object`

### /get

- **GET**: Echoes the GET request data.
    - **Response**: `200 OK`
    - **Schema**: `object`

### /headers

- **GET**: Echoes the request headers.
    - **Response**: `200 OK`
    - **Schema**: `object`

### /ip

- **GET**: Returns the IP address of the requester.
    - **Response**: `200 OK`
    - **Schema**: `object`

### /post

- **POST**: Echoes the POST request data.
    - **Body Parameter**: `data` (object, required)
    - **Response**: `200 OK`
    - **Schema**: `object`

### /status

- **GET**: Returns a response with the given HTTP status code.
    - **Query Parameter**: `code` (integer, required)
    - **Response**: `200 OK`
    - **Schema**: `object`

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).