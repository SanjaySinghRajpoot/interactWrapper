Sure, here's how you can structure your README file:

---

# Interactsh Wrapper

This project is a wrapper around [Interactsh](https://github.com/projectdiscovery/interactsh), allowing for easier integration and customization. A working client has been implemented using the cmd/interactsh-client code.

## Description

The Interactsh Wrapper migrates the core logic and provides two additional endpoints for interacting with Interactsh. It includes a Gin server running in a Goroutine to handle requests concurrently, enabling seamless processing without interference with polling requests.

## Installation

### Prerequisites

- Go installed on your system
- Docker installed (optional)

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/your/interactsh-wrapper.git
   ```

2. Navigate to the project directory:

   ```bash
   cd interactsh-wrapper
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

## Usage

### Running with Go

To run the project using Go:

```bash
go run main.go
```

### Running with Docker

Alternatively, you can use Docker for deployment. A multi-stage Dockerfile with hot reloading is provided.

1. Run the Docker container:

   ```bash
   docker-compose up
   ```

The server will start and be accessible at port 8081.


## Endpoints

### GetURL

- **Description**: Endpoint to retrieve a URL.
- **Method**: GET
- **Path**: /api/getURL

### GetInteractions

- **Description**: Endpoint to retrieve interactions for a given URL.
- **Method**: POST
- **Path**: /api/getInteractions
- **Request Payload**: JSON with URL and optional timestamp.
```
{
  "url":"cnbntm7fdljsogenr20gbneogaue656mm",
  "time_stamp":"2024-02-22 17:05:44"   // optional
}
```

## Utilities

### AppendValue

- **Description**: Function to append a value to the map for the given key.
- **Parameters**:
  - `m map[string][]string`: Map to append the value to.
  - `key string`: Key of the map to append the value to.
  - `value string`: Value to append to the map.

### FilterByTimestamp

- **Description**: Function to filter interactions based on the provided timestamp.
- **Parameters**:
  - `interactions []string`: Slice of interactions to filter.
  - `timestamp string`: Timestamp to filter interactions with.

## Screenshots

1. Server Running using Docker Compose
![Screenshot from 2024-02-22 23-07-42](https://github.com/SanjaySinghRajpoot/interactWrapper/assets/67458417/49ce3801-acda-48e3-bd09-752edaa9c074)

2. `api/getUrl` endpoint 
![Screenshot from 2024-02-22 22-37-28](https://github.com/SanjaySinghRajpoot/interactWrapper/assets/67458417/bf62704e-2a2e-40a5-a932-47c451ee5de9)

3. `api/getinteractions` endpoint without Timestamp
![Screenshot from 2024-02-22 22-37-21](https://github.com/SanjaySinghRajpoot/interactWrapper/assets/67458417/5626e811-c98c-4f7e-be75-54bef0e0478c)

4. `api/getinteractions` endpoint with Timestamp
![Screenshot from 2024-02-22 22-37-12](https://github.com/SanjaySinghRajpoot/interactWrapper/assets/67458417/009a7127-40d7-48cc-971e-f43ee3ead15a)


## Testing

Testing is provided for both the endpoints and utility functions. You can run tests using the `go test` command.

## Contributing

Contributions are welcome! If you have any suggestions or improvements, feel free to open an issue or create a pull request.
