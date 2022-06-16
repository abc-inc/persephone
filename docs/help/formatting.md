# Formatting

## CSV/JSON/YAML

Some Persephone commands support exporting the data as CSV/JSON/YAML as an alternative to their
usual line-based plain text output. This is suitable for passing structured data to scripts.
The output format can be changed with the `:format` command, followed by the desired format.
The following formats are supported:

- `auto`: Use `table` for interactive terminals and `json` otherwise.
- `csv`: Comma-separated values
- `json`: JSON string
- `jsonc`: Colorized JSON
- `raw`: JSON as returned from Neo4j without post-processing
- `rawc`: JSON as returned from Neo4j without post-processing, but colorized
- `table`: ASCII table
- `text`: Name and value pairs, separated by a single whitespace
- `tsv`: Tab-separated name and value pairs (useful for grep, sed, or awk)
- `yaml`: YAML, a machine-readable alternative to JSON
- `yamlc`: Colorized YAML

## Custom Templates

With the `:template` command, you can specify templates to render query results.
For the syntax of Go templates, see: <https://golang.org/pkg/text/template/>
