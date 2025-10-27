# Tasks - CLI Todo List Manager

A simple and efficient command-line todo list application written in Go. Manage your tasks directly from the terminal with persistent CSV storage.

## Features

- ✅ Add, list, complete, and delete tasks
- 📝 CSV-based persistent storage with file locking
- ⏱️ Human-readable timestamps (e.g., "2 minutes ago")
- 🎨 Clean tabular output formatting
- 🔒 Concurrent access protection via file locking
- 🛠️ Customizable data file location

## Installation

### Prerequisites

- Go 1.16 or higher

### Install

```bash
go install github.com/Triiltz/ToDo-app/cmd/tasks@latest
```

Or clone and build from source:

```bash
git clone https://github.com/Triiltz/ToDo-app.git
cd ToDo-app
go install ./cmd/tasks
```

Make sure `$HOME/go/bin` is in your PATH:

```bash
export PATH="$HOME/go/bin:$PATH"
```

## Usage

### Add a new task

```bash
tasks add "Buy groceries"
tasks add "Finish the project documentation"
```

### List tasks

List all uncompleted tasks:

```bash
tasks list
```

Output:
```
ID    Task                              Created
1     Buy groceries                     2 minutes ago
2     Finish the project documentation  a few seconds ago
```

List all tasks (including completed):

```bash
tasks list -a
```

Output:
```
ID    Task                              Created          Done
1     Buy groceries                     5 minutes ago    false
2     Finish the project documentation  3 minutes ago    true
3     Review pull requests              a minute ago     false
```

### Complete a task

```bash
tasks complete 1
```

### Delete a task

```bash
tasks delete 2
```

### Custom data file location

By default, tasks are stored in `~/.tasks/tasks.csv`. You can specify a custom location:

```bash
tasks -f /path/to/custom/tasks.csv list
tasks -f /path/to/custom/tasks.csv add "Custom task"
```

## Project Structure

```
ToDo-app/
├── cmd/
│   └── tasks/
│       └── main.go           # Application entry point
├── internal/
│   ├── task/
│   │   ├── task.go          # Task data model
│   │   └── storage.go       # CSV storage with file locking
│   └── cli/
│       ├── root.go          # Root command
│       ├── add.go           # Add command
│       ├── list.go          # List command
│       ├── complete.go      # Complete command
│       └── delete.go        # Delete command                  
└── go.mod
```

## Technical Details

### Storage Format

Tasks are stored in CSV format with the following structure:

```csv
ID,Description,CreatedAt,IsComplete
1,My task,2025-10-26T21:25:10-03:00,false
2,Another task,2025-10-26T21:30:15-03:00,true
```

### File Locking

The application uses `flock` system calls to prevent concurrent read/write conflicts, ensuring data integrity even when multiple processes attempt to access the task file simultaneously.

### Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [github.com/mergestat/timediff](https://github.com/mergestat/timediff) - Human-readable time differences

## Development

### Build locally

```bash
go build -o tasks ./cmd/tasks
./tasks list
```

### Run without building

```bash
go run cmd/tasks/main.go list
go run cmd/tasks/main.go add "Test task"
```

### Run tests

```bash
go test ./...
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is open source and available under the MIT License.

## Author

Built by [Triiltz](https://github.com/Triiltz)