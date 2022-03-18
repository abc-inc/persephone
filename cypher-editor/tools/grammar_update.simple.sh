#!/bin/bash

java \
    -Xmx500M \
    -cp "tools/antlr-4.7-complete.jar" \
    org.antlr.v4.Tool \
    -Dlanguage=Go \
    -o go-cypher-editor-support/src/_generated.simple \
    cypher-editor-support/src/_generated.simple/Cypher.g4

rm -f go-cypher-editor-support/src/_generated.simple/Cypher.tokens
rm -f go-cypher-editor-support/src/_generated.simple/CypherLexer.tokens
