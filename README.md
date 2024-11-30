# GopherDo

GopherDo is a simple, lightweight task management application built with Go and vanilla JavaScript. It provides a clean web interface for managing your daily tasks with features like task categorization, urgency levels, and due dates.

## Features

- Create, complete, and delete tasks
- Task categorization (General, Work, Personal)
- Urgency levels (Urgent, Non-urgent)
- Due date assignment
- Persistent storage using JSON file
- Clean and responsive web interface
- Cross-platform support

## Technologies Used

- Backend: Go 1.23.3
- Frontend: HTML, CSS, JavaScript (Vanilla)
- Data Storage: JSON file-based persistence
- Build System: Make

## Prerequisites

- Go 1.23.3 or higher
- Make (for building)

## Installation

### Option 1: Build from Source

1. Clone the repository:

```bash
git clone https://github.com/ATLIOD/gopherdo.git
cd gopherdo
```

2. Build the application:

```bash
make build
```

This will create a `build` directory containing the executable and static files.

### Option 2: Pre-built Releases

Download the appropriate zip file for your operating system:

- Windows: `gopherdo-windows.zip`

    - Extract and run `gopherdo.exe`
    - Static files are included in the `static` directory

- Linux: `gopherdo-linux.zip`

    - Extract and run `gopherdo`
    - Make sure to set executable permissions: `chmod +x gopherdo`
    - Static files are included in the `static` directory

- macOS: `gopherdo-mac.zip`
    - Extract and run `gopherdo`
    - Make sure to set executable permissions: `chmod +x gopherdo`
    - Static files are included in the `static` directory

Each release package contains:

- The GopherDo executable
- A `static` directory with the web interface files
- Any saved tasks will be stored in `saved_tasks.json` in the same directory as the executable

## Running the Application

1. Navigate to the build directory:

```bash
cd build
```

2. Run the executable:

```bash
./gopherdo
```

3. Open your web browser and visit:

```
http://localhost:8080
```

## Building for Different Platforms

GopherDo supports cross-platform builds for Windows, Linux, and macOS.

- Build for all platforms:

```bash
make build-all
```

- Build for specific platforms:

```bash
make build-windows
make build-linux
make build-mac
```

- Create distribution packages:

```bash
make dist
```

## API Endpoints

- `GET /tasks` - Retrieve all tasks
- `POST /tasks` - Create a new task
- `DELETE /tasks/delete?id={id}` - Delete a task
- `PATCH /tasks/complete?id={id}` - Mark a task as complete

## Task Structure

```json
{
    "id": 1,
    "name": "Task Name",
    "description": "Task Description",
    "complete": false,
    "category": "general",
    "urgency": "non-urgent",
    "dueDate": "2024-01-01T00:00:00Z",
    "creationDate": "2024-01-01T00:00:00Z"
}
```


## License

This project is licensed under the GPL-3.0 license - see the LICENSE file for details.

## Author

ATLIOD
