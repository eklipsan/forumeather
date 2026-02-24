
# Forumeather

**Forumeather** is a Go-based application designed to track youth forums and provide up-to-date weather forecasts at their locations. The project combines a forum database with integration into the Meteoblue weather service.

## Project Essence

The application allows you to store information about various forums (name, location, topics, coordinates) and view the current weather or daily trend for these locations. The project consists of two main parts:

1. **Web Interface**: for visual display of the list of forums and weather.
2. **CLI Tool**: for administrative management of the forum list (adding, deleting, viewing).

## Key Features

* **Meteoblue API Integration**: obtaining a detailed weather forecast based on forum coordinates.
* **Data Management**: a full-fledged CLI interface for working with the SQLite database.
* **Containerization**: availability of Docker configurations for rapid deployment of both the web application and the console utility.
* **Ready-made Database**: the project includes an SQL script with a pre-installed list of the largest youth forums in Russia, such as "Territory of Meanings," "Tavrida.ART," "Mashuk," and others.

## Project Structure

* `cmd/web/` — entry point and logic of the web server.
* `cmd/cli/` — console utility for database management.
* `pkg/meteoblue/` — client for working with the weather forecast API.
* `pkg/storage/` — logic for interacting with the SQLite database.
* `pkg/config/` — configuration and startup flag management.
* `initial_script_db.sql` — SQL script to initialize the DB structure and fill it with initial data.

## Tech Stack

* **Programming Language**: Go 1.23.2.
* **Database**: SQLite 3.
* **API**: Meteoblue.
* **Deployment**: Docker, Docker Compose.

## How to Run

### Prerequisites

* Installed **Go** (version 1.23.2 or higher).
* Installed **Docker** and **Docker Compose** (for running in containers).
* **Meteoblue API Key**: a default key is provided in the configuration, but it may return empty results if disabled.

### Running via Docker Compose (Recommended)

The fastest way to start the project is to use Docker Compose, which will automatically build and configure the web application and the CLI tool:

```bash
docker-compose up --build

```

After startup, the web interface will be available at: `http://localhost:8080`.

### Manual Run (Without Docker)

1. **Start the Web Server**:
```bash
go run ./cmd/web/ --port=:8080 --api-key=YOUR_KEY

```


*The `--port` and `--api-key` flags are optional and have default values*.
2. **Start the CLI Tool**:
```bash
go run ./cmd/cli/

```



## Using the CLI Tool

When running the console utility, you will be prompted to enter an operation code:

* `0` — **Add a new forum**: you will need to enter the name, location, topics, and coordinates (latitude and longitude).
* `1` — **Delete a forum**: you must enter the record ID.
* `2` — **Show all forums**: displays a list of all forums stored in the database.

## Configuration

Configuration parameters are set via command-line flags when starting the application:

* `port`: The port on which the web server will be hosted (default is `:8080`).
* `api-key`: Access key for the Meteoblue API.
* The database path is set in the configuration as `data/forumeather.db`.