# dependency management

In Go (Golang), `dependency management` is a critical part of managing libraries and packages that your project relies on. Over time, Go has evolved its approach to dependency management, and several tools and methods have been introduced and refined.

### Go Modules (`go mod`)

This system allows developers to define project dependencies in a go.mod file and manage them with `go get`, `go build`, and other Go commands.

### Key Features:

- `go.mod file`: Defines the module and lists dependencies.
- `go.sum file`: Contains cryptographic hashes of dependencies for security.
- `Versioning`: Supports semantic versioning for managing versions of dependencies.
- `No GOPATH`: With Go modules, you don't need to work in the GOPATH directory anymore.

## Common Commands:

- `go mod init <module-name>`: Initializes a new Go module.
- `go mod tidy`: Cleans up any unnecessary dependencies and adds missing ones.
- `go get <package>`: Adds or updates a dependency.
- `go list -m all`: Lists all the modules required by the current project.
