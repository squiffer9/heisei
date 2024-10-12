# Terminal-based BBS Application (Heisei)

This is a terminal-based bulletin board system (BBS) application implemented in Go. It provides a nostalgic yet modern approach to online discussions.

## Features

- Terminal User Interface (TUI) for a retro feel
- Category and thread management
- Real-time updates using WebSocket
- Multi-language support (English and Japanese)
- Secure and efficient backend API
- PostgreSQL database for data persistence

## Project Structure

```
heisei/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
├── internal/
│   ├── server/
│   │   ├── api/
│   │   ├── config/
│   │   ├── models/
│   │   ├── repositories/
│   │   └── services/
│   ├── client/
│   │   ├── api/
│   │   ├── config/
│   │   └── tui/
│   └── common/
│       └── models/
├── pkg/
│   ├── database/
│   └── utils/
├── migrations/
├── scripts/
├── test/
└── docs/
```

## Setup

### Prerequisites

- Go 1.22 or later
- PostgreSQL 14 or later
- Docker (optional, for containerized setup)

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/squiffer9/heisei.git
   cd heisei
   ```

2. Install dependencies:

   ```
   go mod download
   ```

3. Set up the database:

   - Create a PostgreSQL database
   - Copy `.env.example` to `.env` and update the database connection details

4. Run database migrations:

   ```
   go run scripts/migrate.go up
   ```

5. Build the application:
   ```
   go build -o heisei_server ./cmd/server
   go build -o heisei_client ./cmd/client
   ```

## Configuration

1. Copy `configs/config.sample.yaml` to `configs/config.yaml`.
2. Edit `configs/config.yaml` and set your specific configuration values.
3. For sensitive information like database passwords, use environment variables instead of putting them in the config file:

   ```sh
   export DB_PASSWORD=your_secure_password
   ```

The application will load the configuration from `configs/config.yaml` and override values with environment variables if they are set.

Never commit `configs/config.yaml` to version control, as it may contain sensitive information.

## Usage

1. Start the server:

   ```
   ./heisei_server
   ```

2. In a new terminal, start the client:

   ```
   ./heisei_client
   ```

3. Follow the on-screen instructions to navigate the BBS.

## Development

### Running Tests

Run all tests with:

```
go test ./...
```

### Code Style

We follow the standard Go style guide. Please ensure your code is formatted with `gofmt` before submitting:

```
gofmt -w .
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to all contributors who have helped shape Heisei
- Inspired by classic BBS systems and modern Go practices
