# Environment variables

Command line flags take precedence over environment variables.

`GLAMOUR_STYLE`: the style to use for rendering Markdown. See
<https://github.com/charmbracelet/glamour#styles> 

`NEO4J_ADDRESS`: specify the address and port to connect to (default: `neo4j://localhost:7687`).
Can also be specified as flag `--address`.

`NEO4J_DATABASE`: specify the database to connect to (default: `neo4j`).
Can also be specified as flag `--database`.

`NEO4J_USERNAME`: specify the username to connect as (default: `neo4j`).
Can also be specified as flag `--username`.

`NEO4J_PASSWORD`: specify the password to use.
Setting this avoids being prompted to authenticate.

`NO_COLOR`: set to any value to avoid printing ANSI escape sequences for color output.
