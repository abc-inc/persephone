# Contributing

## persephone project layout

At a high level, these areas make up the `github.com/abc-inc/persephone` project:

- [`cmd/`](../cmd) - `main` packages for building the `persephone` executable
- [`comp/`](../comp) - autocompletion rules for Cypher
- [`console/`](../console) - most other CLI code, including different output formats and the REPL
- [`comp/`](../comp) - most other packages, including the implementation for individual gh commands
- [`docs/`](../docs) - documentation for maintainers and contributors
- [`editor/`](../editor) - invokes the parser and provides autocompletion information
- [`go.mod`](../go.mod) - external Go dependencies for this project, automatically fetched by Go at build time
- [`ref`](../ref) - provides references to AST elements visited by the parser

The rest are auxiliary Go packages at the top level of the project for historical reasons.

## Command-line help text

Running `persephone help :template` displays help text for a topic.
In this case, the topic is a specific command, and help text for is embedded in a command's source code.

## Add a new command

1. Create a new file for the new command, e.g. `cmd/persephone/cmd/science/boom.go`
2. The new command should expose a method, e.g. `NewCmdBoom()`, that accepts a `*cmdutil.Factory` type
   and returns a `*cobra.Command`.
   * Any logic specific to this command should be kept within the package and not added to any "global" package.
3. Use the method from the previous step to generate the command and add it to the command tree,
   typically somewhere in the `NewCmdRoot()` method.

## How to write tests

This task might be tricky.
Typically, persephone commands do things like look up information from the database.
Moreover, one does not want to test the CLI framework, rather than the command itself.
To avoid that, you may want to define function (and export them, if necessary),
which do not contain any CLI-related code.

To make your code testable, write small, isolated pieces of functionality that are designed to be composed together.
Prefer table-driven tests for maintaining variations of different test inputs and expectations
when exercising a single piece of functionality.
