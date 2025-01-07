# Go Application with Hot Reload and Debugging in Docker

This repository sets up a simple Go application with hot reload using [Air](https://github.com/cosmtrek/air) and debugging using [Delve](https://github.com/go-delve/delve). 
The application is containerized with Docker and managed via Docker Compose.

## Features
- **Hot Reload**: Automatically reload the application on code changes.
- **Debugging**: Integrated debugging using Delve.
- **Containerized**: Runs entirely in Docker for consistency and portability.
- **Log Forwarding**: Includes a Fluent Bit container to forward logs to New Relic.

## Prerequisites

Make sure you have the following installed:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

## Setup

### 1. Clone the Repository
```bash
git clone git@github.com:dzemildupljak/go_app_tools.git
cd go_app_tools
```

### 2. Install Dependencies
Ensure all Go dependencies are installed by running:
```bash
go mod tidy
```

### 3. Start the Application
Run the application using Docker Compose:
```bash
docker-compose up --build
```

This will:
- Start the application with Air for hot reload.
- Launch a Delve debugger on port `2345`.
- Start a Fluent Bit container to forward logs to New Relic.

### 4. Debugging with Delve
To debug the application, connect your debugger (e.g., GoLand, VSCode) to `localhost:2345`.

## Configuration

### Air Configuration
The `.air.toml` file contains the configuration for hot reload. You can customize it based on your project requirements.

### Docker Compose
The `docker-compose.yml` file includes services for:
- The Go application with Air and Delve.
- Fluent Bit for log forwarding to New Relic.
- Additional dependencies (e.g., databases) if needed.

### Fluent Bit Configuration
The Fluent Bit container uses a configuration file located in the `fluent-bit-config/` directory. 
Ensure you customize it with your New Relic endpoint and credentials.

## Contributing
Feel free to submit issues or pull requests for improvements.

## License
This project is licensed under the [MIT License](LICENSE).

## References
- [Air Documentation](https://github.com/cosmtrek/air)
- [Delve Documentation](https://github.com/go-delve/delve)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Fluent Bit Documentation](https://docs.fluentbit.io/)

