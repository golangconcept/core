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

### Summary of the Most Common go mod Commands:
| Command |	Description   |
|---------|---------------|
|`go mod init <module-name>`	| Initialize a new Go module with the specified module name.|
|`go mod tidy` | Clean up go.mod and go.sum by removing unused dependencies.|
|`go mod vendor` |	Creates a vendor/ directory containing dependencies.|
|`go mod verify` | Verifies that dependencies in go.mod match what's on the disk.|
|`go mod edit` | Modify go.mod file (add/remove dependencies, change Go version, etc.)|
|`go mod graph` |	Prints a dependency graph of the module.|
|`go mod download` | Downloads modules into the local cache.|
|`go mod why` |	Explains why a module is included in the current module's dependency graph.|