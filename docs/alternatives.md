# Alternatives to persephone

## Cypher Shell

Cypher Shell is the official command-line tool that comes with the Neo4j distribution.
It can also be downloaded from Neo4j Download Center and installed separately.

Cypher Shell CLI is used to run queries and perform administrative tasks against a Neo4j instance.
By default, the shell is interactive, but you can also use it for scripting, by passing cypher
directly on the command line or by piping a file with cypher statements (requires PowerShell on Windows).
It communicates via the Bolt protocol.

## Why didn't you just extend Cypher Shell?

Go has a rich ecosystem of command line libraries such as [cobra][], which is not only used by Google, Microsoft,
Amazon and other big tech companies, but by thousands of command line applications.
Thus, it provides a semi-standard user-experience, which is important to foster adoption.

Besides, Cypher Shell is released under the GNU General Public License (GPL),
and it does not seem to play an important role in the product line of the company.
Hence, any derivative work must be distributed under the same or equivalent license terms.

## What's next for persephone?

persephone aims to achieve feature parity with Cypher Shell, neo4j-client
and Neo4j Browser (where applicable) throughout 2022.
This includes better support for query profiling and logging, built-in guides,
syntax highlighting and an overall better command line experience.
Additional authentication methods, an extension manager and auto-updates will roundup the product.

On the long-run, persephone aims to replace the official Cypher Shell as the Neo4j CLI of choice by the Neo4j community.

In case you miss some features, don't hesitate to file a [feature request][issues].

## What does it mean that Cypher Shell is official and persephone is unofficial?

Cypher Shell is built and maintained by a small team at Neo4j., Inc and a few community members.
When there's something wrong with it, people can reach out to Neo4j support,
or create an issue in the issue tracker, where an employee at Neo4j will respond. 

persephone is a project, whose maintainer works in a large company, which uses Neo4j to power their CI/CD pipeline.
He chooses to maintain persephone in his spare time, along with other [libraries and projects][abc-inc],
as many others do with open source projects.

## Should I use Cypher Shell or persephone?

People should use whatever set of tools makes them happiest and most productive working with Neo4j. 

If you are set on using a tool for repetitive execution of statements, Cypher Shell is likely a better choice.

If you want a tool that's more opinionated and intended to support your Neo4j development from the command line,
you might give persephone a try.
It's intended to be responsive to people's concerns and needs and improve it based on how people are using it over time.

persephone is not intended to be an exact replacement for Cypher Shell and likely never will be.
However, the hope is that the vast majority of the Neo4j community, who use the CLI,
will find more and more value in using persephone, which leads to even more improvements.

[abc-inc]: https://github.com/abc-inc
[cobra]: https://github.com/spf13/cobra
[issues]: https://github.com/abc-inc/persephone/issues
