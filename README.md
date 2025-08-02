# dumpHTTP - HTTP Request Dumper

[![Go Version](https://img.shields.io/badge/go-1.24%2B-blue.svg)](https://golang.org/doc/install) [![License: BSD-2-Clause](https://img.shields.io/badge/License-BSD--2--Clause-orange.svg)](https://opensource.org/licenses/BSD-2-Clause)

A simple and powerful HTTP request debugging tool that captures, displays, logs and echo all incoming HTTP requests. Perfect for testing webhooks, debugging API calls, or inspecting HTTP traffic during development.

## Features

• 🔍 **Complete Request Capture**: Captures full HTTP requests including headers, body, and metadata
• 📝 **Dual Output**: Displays requests in terminal with colored output and saves to file
• 📊 **Persistent Logging**: Saves all requests to a file with timestamps and unique identifiers
• ⚡ **Lightweight**: Minimal dependencies and fast performance

## Installation

### From Source

```bash
git clone https://github.com/crgimenes/dumpHTTP.git
cd dumpHTTP
go build -o dumpHTTP .
```

### Using Go Install

```bash
go install crg.eti.br/go/dumpHTTP@latest
```

### Requirements

• Go 1.24 or higher

## Quick Start

1. **Start the HTTP dump server:**

   ```bash
   ./dumpHTTP
   ```

2. **Send a test request:**

   ```bash
   curl -X POST http://localhost:8080/test \
        -H "Content-Type: application/json" \
        -d '{"message": "Hello World"}'
   ```

3. **View the captured request** in your terminal and check the `dump.txt` file

## Usage

### Command Line Options

```bash
./dumpHTTP [options]
```

| Option | Default | Description |
|--------|---------|-------------|
| `-listenAddr` | `:8080` | Address and port to listen on |
| `-dumpFile` | `dump.txt` | File to save request dumps |

### Examples

**Listen on a specific port:**

```bash
./dumpHTTP -listenAddr :9000
```

**Save dumps to a custom file:**

```bash
./dumpHTTP -dumpFile requests.log
```

**Custom configuration:**

```bash
./dumpHTTP -listenAddr 0.0.0.0:8080 -dumpFile /var/log/http-dumps.txt
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the BSD 2-Clause License - see the [LICENSE](LICENSE) file for details.

## Author

Created by [Cesar Gimenes](https://github.com/crgimenes)
