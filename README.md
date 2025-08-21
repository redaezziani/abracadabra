# 🪄 Abracadabra

A magical DevOps CLI helper that makes system administration tasks easier and more enjoyable.

## ✨ Features

- **System Information** - Get real-time CPU, memory, and disk usage
- **Log Tailing** - Easily tail systemd service logs
- **Docker Cleanup** - Remove dangling containers and images with one command
- **Beautiful Output** - Emoji-enhanced, human-readable information

## 🚀 Installation

### Prerequisites

- Go 1.19 or higher
- Linux system with systemd (for logs command)
- Docker (optional, for docker-clean command)

### Build from Source

1. Clone or download the repository
2. Navigate to the project directory
3. Initialize Go module and install dependencies:

```bash
go mod init abracadabra
go mod tidy
```

4. Build the application:

```bash
go build -o abracadabra main.go
```

5. (Optional) Install globally:

```bash
sudo cp abracadabra /usr/local/bin/
```

## 📖 Usage

### Available Commands

#### System Information

Display current system resource usage:

```bash
./abracadabra sysinfo
```

**Example output:**

```
🖥️ CPU Usage: 15.32%
💾 RAM: 8.45 GB / 16.00 GB (53%)
```

#### Disk Usage

Show disk space information:

```bash
./abracadabra disk
```

**Example output:**

```
📂 Disk: 245.67 GB free / 500.00 GB total (51% used)
```

#### Service Logs

Tail logs for a systemd service:

```bash
./abracadabra logs nginx
./abracadabra logs docker
./abracadabra logs ssh
```

#### Docker Cleanup

Remove dangling Docker containers and images:

```bash
./abracadabra docker-clean
```

**Output:**

```
🧹 Docker cleaned!
```

### Help

Get help for any command:

```bash
./abracadabra --help
./abracadabra [command] --help
```

## 🛠️ Dependencies

This project uses the following Go packages:

- [`github.com/spf13/cobra`](https://github.com/spf13/cobra) - CLI framework
- [`github.com/shirou/gopsutil/v3`](https://github.com/shirou/gopsutil) - System information library

## 🔧 Development

### Project Structure

```
.
├── main.go          # Main application file
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
└── README.md        # This file
```

### Adding New Commands

To add a new command, follow the Cobra pattern in `main.go`:

```go
var newCmd = &cobra.Command{
    Use:   "command-name",
    Short: "Description of the command",
    Run: func(cmd *cobra.Command, args []string) {
        // Command implementation
    },
}

// Add to root command
rootCmd.AddCommand(newCmd)
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🎯 Roadmap

- [ ] Add more system monitoring features
- [ ] Support for other init systems (OpenRC, runit)
- [ ] Configuration file support
- [ ] Plugin system
- [ ] Network monitoring commands
- [ ] Process management commands

## 🐛 Issues

If you encounter any issues or have suggestions, please [open an issue](https://github.com/redaezziani/abracadabra/issues).

---

Made with ❤️ for DevOps engineers and system administrators.
