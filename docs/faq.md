# Frequently Asked Questions

## Is Persephone a complete command line interface (CLI)?

Persephone is designed to provide the best experience when working with Neo4j from the command line.
This includes, but is not limited to,

- autocompletion for
  - keywords
  - labels and types
  - properties
  - variables
  - functions and procedures
  - and more
- different output formats like CSV, JSON, YAML or formatted tables
- commands for querying and managing databases
- persistent history of executed statements

## Does Persephone support multi-line statements?

Yes, Cypher statements can span multiple lines and must end with a semicolon (`;`).
If there is no semicolon, then the enter key will start a new line.

## How about support for scripting?

Persephone provides a non-interactive  mode that supports reading input from a file.
You can use the `:source` command or directly pipe statements to Persephone.
It automatically detects whether an interactive terminal is used and changes its behavior providing meaningful defaults.
For example, in non-interactive mode, it provides plain JSON output instead of colorful output and tables.
This can be customized if desired.

## Does Persephone support transactions?

Simply use the standard commands `:begin`, `:commit` and `:rollback`.
By default, every query is executed in an implicit transaction.
