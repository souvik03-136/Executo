# Executo an Automatic App Launcher

An efficient and customizable application launcher built with Go. This project aims to provide users with an easy-to-use tool for managing and launching applications, along with advanced features like shortcuts, batch launching, and more.

---

## Features

### Core Features
- **App Scanning and Listing**: Automatically detect and list installed applications.
- **Customizable Shortcuts**: Allow users to assign hotkeys or aliases for apps.
- **Search Functionality**: Quickly find apps using a search bar.
- **Recent Apps**: Display recently opened applications.
- **Categorization**: Organize apps into categories for easy navigation.
- **Multi-Platform Support**: Supports Windows, macOS, and Linux.

### User Interaction
- **Custom Commands**: Define specific launch actions (e.g., open with parameters).
- **Favorite Apps**: Mark frequently used apps as favorites.
- **Drag-and-Drop**: Add apps to the launcher via drag-and-drop.

### Automation
- **Startup Apps**: Configure apps to launch at system startup.
- **Batch Launching**: Launch multiple apps simultaneously.
- **Integration with Schedulers**: Schedule apps to launch at specific times.

### Performance and Usability
- **Lightweight Design**: Minimal resource usage.
- **Theme Support**: Light, dark, and customizable themes.
- **Notifications**: Inform users when an app fails to launch.

### Advanced Features
- **Command-Line Interface**: Enable CLI-based app launching for advanced users.
- **Portability**: Run as a standalone executable.
- **System Tray Integration**: Quick access via the system tray.

### Customization
- **Personalized Layout**: Customize the launcher’s interface layout.
- **Shortcut Management**: Edit, delete, or reorder app shortcuts.

### Security
- **Password Protection**: Lock specific apps behind a password or PIN.
- **Permission Management**: Ensure user consent for sensitive actions.

### Maintenance
- **Update Notifications**: Notify users about new versions.
- **Error Logging**: Log errors for troubleshooting.

---

## Directory Structure

```
Executo/
│
├── cmd/                     # Entry points
│   └── main.go              # Main application logic
│
├── internal/                # Core internal logic (not exposed outside)
│   ├── appscanner/          # App scanning and listing logic
│   │   ├── scanner.go       # Detect installed apps
│   │   └── types.go         # App metadata structs
│   ├── launcher/            # Logic to launch applications
│   │   ├── launch.go        # Launch app by name or path
│   │   └── shortcuts.go     # Handle app shortcuts
│   ├── config/              # Configuration handling
│   │   ├── config.go        # Load and save user settings
│   │   └── defaults.go      # Default configurations
│   └── logger/              # Logging functionality
│       └── logger.go        # Centralized logging utility
│
├── ui/                      # GUI-related files (if implementing a UI)
│   ├── assets/              # Icons, fonts, and other assets
│   └── ui.go                # GUI logic and event handling
│
├── pkg/                     # Reusable packages (if needed)
│   ├── utils/               # Utility functions
│   │   └── fileutils.go     # File handling utilities
│   └── cli/                 # CLI logic
│       └── cli.go           # CLI parser for app launching
│
├── tests/                   # Unit and integration tests
│   ├── appscanner_test.go   # Tests for app scanning logic
│   ├── launcher_test.go     # Tests for app launching logic
│   └── config_test.go       # Tests for configuration handling
│
├── build/                   # Build and deployment files
│   ├── Dockerfile           # Docker container setup
│   ├── docker-compose.yml   # Docker Compose configuration (if needed)
│   └── Taskfile.yml         # Taskfile for automation
│
├── .github/                 # GitHub Actions for CI/CD
│   └── workflows/
│       └── go.yml           # GitHub Actions workflow for Go
│
├── .gitignore               # Git ignore file
├── README.md                # Project documentation
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies file
└── LICENSE                  # License file
```

---

## Getting Started

### Prerequisites
- Install [Go](https://go.dev/) on your system.
- Ensure Git is installed for version control.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/souvik03-136/Executo.git
   cd app-launcher
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run cmd/main.go
   ```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
