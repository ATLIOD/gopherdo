# GopherDo Task Manager

A simple, efficient task management application built with Go and vanilla JavaScript.

## Features

- Create, read, update, and delete tasks
- Task categories (Work, Personal, General)
- Urgency levels
- Due dates
- Task completion tracking
- Sort by:
    - Completion status
    - Due date
    - Task name

## Installation

1. Clone the repository

```bash
git clone https://github.com/yourusername/gopherdo.git
cd gopherdo
```

2. Initialize the database

```bash
# The database will be automatically created in the data directory
# when you first run the application
```

3. Build the application

```bash
make build
```

4. Run the application

```bash
make run
```

The application will be available at `http://localhost:8080`

## Project Structure

```
gopherdo/
├── api/
│   └── static/          # Frontend files
│       └── index.html
├── cmd/
│   └── server/          # Application entry point
│       └── main.go
├── internal/
│   ├── config/          # Configuration
│   ├── db/              # Database operations
│   │   └── migrations/
│   ├── handler/         # HTTP handlers
│   └── model/           # Data models
├── data/                # SQLite database location
├── Makefile
└── README.md
```

## Usage

1. Access the web interface at `http://localhost:8080`
2. Create tasks using the form
3. Use the search bar to find specific tasks
4. Sort tasks using the dropdown menu
5. Mark tasks as complete or delete them as needed

## Development

To run in development mode:

```bash
make run
```

To build for different platforms:

```bash
make build
```

## API Endpoints

- `GET /tasks` - Retrieve all tasks
- `POST /tasks` - Create a new task
- `PATCH /tasks/complete` - Mark a task as complete
- `DELETE /tasks/delete` - Delete a task

## License

This project is licensed under the GPL-3.0

## Acknowledgments

- Built with Go
- SQLite for data storage
- Vanilla JavaScript for frontend
