# Pomodoro Timer

A simple golang pomodoro cli :) 

## Features

- **Full-screen terminal UI** with centered layout
- **Visual progress bar** (green → red as time runs out)
- **Automatic work/break session switching**
- **Customizable work and break durations**
- **Simple keyboard controls**
- **Responsive design** that adapts to terminal size

## Installation

### Quick Install
```bash
# Clone the repository
git clone https://github.com/lexO-dat/pomodoro-cli
cd pomodoro

# Install
make install
```

### Manual Build
```bash
# Download dependencies
make deps

# Build the application
make build

# The binary will be in ./bin/pomodoro
```

## Usage

### Command Line Options
```bash
# Default timing (25min work, 5min break)
pomodoro

# Custom work time only
pomodoro -w 30          # 30min work, 5min break (default)

# Custom work and break time
pomodoro -w 45 -b 15    # 45min work, 15min break

# Show help
pomodoro -h
```

### Controls
- `SPACE` - Start/Pause timer
- `R` - Reset to work session
- `H` or `?` - Toggle help screen
- `Q/ESC` - Quit application

## Project Structure

```
pomodoro/
├── cmd/pomodoro/          # Application entry point
│   └── main.go
├── internal/              # Private application code
│   ├── timer/            # Timer logic
│   │   └── timer.go
│   ├── ui/               # User interface
│   │   └── ui.go
│   └── sound/            # Sound notifications
│       └── sound.go
├── bin/                  # Built binaries (created by build)
├── go.mod               # Go module definition
├── Makefile            # Build automation
└── README.md          # This file
```

## Development

### Available Make Commands
```bash
make build      # Build the application
make clean      # Clean build artifacts
make test       # Run tests
make deps       # Download dependencies
make install    # Install (requires sudo)
make uninstall  # Uninstall from system (requires sudo)
make run        # Build and run locally
make run-custom # Build and run with custom timing
make all        # Clean, deps, test, and build
make help       # Show help message
```

### Running Locally
```bash
# Run with default settings
make run

# Run with custom settings (30min work, 10min break)
make run-custom

# Or build and run manually
make build
./bin/pomodoro -w 25 -b 5
```

## How It Works

1. **Timer starts** with a full green progress bar
2. **Bar gradually turns red** as time runs out
3. **Sound notification** plays when session completes
4. **Automatically switches** between work and break sessions
5. **Press R** to reset back to work session anytime

## Requirements

- Go 1.21 or later
- Terminal with color support
- For sound: System audio capabilities
