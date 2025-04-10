// Module
// A Go module is a collection of Go packages stored in a single directory tree.
// A module is defined by a go.mod file that defines the module's path and its dependencies.
// An application can have multiple modules.

package project

// Creating a Go module
// To create a Go module, you need to create a directory for your module and run the go mod init command.
// Usually, the module file is located in the root of the module/project directory.
// The module name will be the base for all packages imports.
// Note: Usually, the module name follows the format below:
// github.com/<username>/<repository-name>.
var _ = "go mod init <module-name>"

// The go.mod File
// The go.mod file defines the module's path and its dependencies.
// The file has the statements below which is used to configure the module.
// - module: The module's path.
// - go: The Go version the module is written in.
// - require: The dependencies of the module.
// - replace: Replaces a module with another module.
// - exclude: Excludes a module from the module's dependencies.
// A simplified structure of the go.mod file is as follows:
var _ = `
module <module-name>

go <version>

require (
	<module-name> <version>
	...
)
replace (
	<module-name> => <module-name> <version>
	...
)
exclude (
	<module-name> <version>
	...
)
`

// Adding Dependencies
// Dependencies are registered in the go.mod file.
// To add a dependency, you can use the go get command.
// The go get command will add the dependency to the go.mod file and download the dependency.
// If version is not specified, the latest version will be used.
var _ = "go get <module-name>@<version>"

// Updating Dependencies
// To update a dependency, we can use the go get command with the -u flag.
// The -u flag will update the dependency to the latest version.
var _ = "go get -u <module-name>"

// Removing Dependencies
// To remove a dependency, we can use the go mod tidy command.
// The go mod tidy command will remove any dependencies that are not used in the module.
var _ = "go mod tidy"

// Installing Dependencies
// Go install all the dependencies in the go.mod file automatically when the application is built or run.
var _ = "go run ."

// Workspaces
// A workspace is a collection of modules, and it is defined by a go.work file.
// The difference between a module and a workspace is that a module is a single unit of code,
// while a workspace is a collection of modules.
// The go.work file has the statements below which is used to configure the workspace.
// To create a workspace, we need to create a directory for our workspace and run the "go work init" command.
var _ = `
go: <version>

use (
	<module-name>
	...
)
replace (
	<module-name> => <module-name>
	...
)
exclude (
	<module-name>
	...
)
`
