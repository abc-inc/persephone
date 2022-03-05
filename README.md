

https://neo4j-client.net/

https://github.com/cleishm/libneo4j-client

http://manpages.ubuntu.com/manpages/bionic/man1/neo4j-client.1.html



# Examples

## Map Nodes with Custom Template

```text
Query:  MATCH (n) WHERE n.name=$name RETURN n
Params: {name: "ABC"}

Template:
{{range $a := .}}
{{index (index $a "n") "Id"}}: {{index (index (index $a "n") "Props") "name"}}
{{end}}

Output: 
0: ABC
```

## Map Key Value Pairs with Custom Template

```text
Query:  MATCH (n) WHERE n.name=$name RETURN id(n) AS id, n.name AS name
Params: {name: "ABC"}

Template:
{{range $a := .}}
>{{$a}}<
{{index $a "id"}}: {{index $a "name"}}
{{end}}`

Output: 
0: ABC
```
