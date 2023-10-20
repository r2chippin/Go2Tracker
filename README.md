# r2React

Welcome to the Go2Tracker repository! 

This project aims to create a simple bittorrent tracker application.

Now it can run as a simple bittorrent tracker with full core function。 

## Done

Basic main page response

Basic announce func

Encoded bencode response

## Doing

Fix announce logic

Storage Peer

Performance improvement

## Tech Stack

- **Backend Framework**: [Gin ↗](https://github.com/gin-gonic/gin)
- **Database ORM**: [Gorm ↗](https://github.com/go-gorm/gorm)
- **Database**: SQLite

## Getting Started

### Prerequisites

- Go
- SQLite(Not for now)

### Install Dependencies

Install the project dependencies using the following command:

```shell
go mod download
```

### Start the Application

Start the Go2Tracker application using the following command:

```shell
go run main.go
```

The application will run at `http://localhost:8080`.

## Contributing

Contributions and feedback are welcome! Follow these steps to contribute:

1. Clone the project locally
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

### License

This project is licensed under the [MIT License](LICENSE).