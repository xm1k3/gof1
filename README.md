# gof1

gof1 is a comprehensive Formula 1 data management tool built with Go, leveraging a SQLite3 database for storing and retrieving Formula 1 related data, including drivers, constructors, circuits, and race results. It provides a command-line interface (CLI) as well as a RESTful API powered by Gin for easy data manipulation and retrieval.

# Installation

To install gof1, first clone the repository:

```
git clone https://github.com/xm1k3/gof1.git
cd gof1
go build
go install
```

Alternatively, you can install gof1 directly using:

```
go install github.com/xm1k3/gof1@latest
```

# Usage

### Command-Line Interface

The CLI provides commands for managing drivers, constructors, circuits, and race results.

Examples:

Adding a Driver:

```bash
gof1 driver add --name "Lewis Hamilton"
```

Getting a Driver:

```bash
gof1 driver get --id 1
```

Updating a Driver:

```bash
gof1 driver update --id 1 --name "Lewis Hamilton Updated"
```

Deleting a Driver:

```bash
gof1 driver delete --id 1
```

Importing Data from CSV

```bash
gof1 import
```

### RESTful API

`gof1` also exposes a RESTful API for interacting with the Formula 1 data.

Starting the Server

```bash
gof1 api --port 8080
```

Examples api endpoints:

Get Driver by ID:

```bash
curl http://localhost:8080/driver/:id
```

Get Driver by Year:

```bash
curl http://localhost:8080/drivers/:year
```

# Contributing

Contributions to `gof1` are welcome. Please fork the repository and submit a pull request with your proposed changes or improvements.

# License

`gof1` is released under the MIT License. See the LICENSE file for more details.