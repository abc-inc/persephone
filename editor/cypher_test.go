// Copyright 2022 The Persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package editor_test

import "strings"

var cypherDefault = strings.Split(`CREATE ()§
CREATE( )§
CREATE (n:Person { name : 'Andres', title : 'Developer' })§
MATCH (node1:Label1)

WHERE node1.propertyA = $value

RETURN node2.propertyA, node2.propertyB§
MATCH (tom:Person)-[:ACTED_IN]->(tomHanksMovies) RETURN tom§
MATCH (tom:Person)  -[:ACTED_IN]->(tomHanksMovies) RETURN tom§
MATCH (tom:Person  )-[:ACTED_IN]->(tomHanksMovies) RETURN tom§
match(n)return n        §
MATCH (n)
WHERE (n.name = 'Peter')
RETURN 1§
MATCH (n)
WHERE (  n.name = 'Peter'  )
RETURN 1§
MATCH (n)
WHERE (  n  .  name   =   'Peter'  )
RETURN 1§
MATCH ()-[r:ACTED_IN*  ]-()
RETURN r§
MATCH ()-[r:ACTED_IN*  2  ]-()
RETURN r§
MATCH ()-[r:ACTED_IN*  2  ..  ]-()
RETURN r§
MATCH ()-[r:ACTED_IN*  2  ..  10  ]-()
RETURN r§
MATCH ()-[r:ACTED_IN*  ..  10  ]-()
RETURN r§
MATCH ()-[r:ACTED_IN*  ..  ]-()
RETURN r§
RETURN -1§
RETURN ----+-+-1§
RETURN -   -    + 1 - +2§
RETURN 1-+2§
RETURN -1.0§
RETURN ----+-+-1.0§
RETURN -   -    + 1 - +2.0§
RETURN 0xd34db33f§
RETURN 0777§
RETURN 1e10§
RETURN 1.5e10§
RETURN -1.5e10§
RETURN .5§
RETURN .0§
CREATE (a)§
CREATE (a1)§
MATCH (n1),(n2) RETURN n1, n2§
MATCH (n1), (n2) RETURN n1, n2§
MATCH (n1),   (n2) RETURN n1, n2§
MATCH (n1) , (n2) RETURN n1, n2§
MATCH (n1) ,   (n2) RETURN n1, n2§
MATCH (n1)   ,   (n2) RETURN n1, n2§
MATCH (n) WITH n.prop AS p RETURN toInteger(DISTINCT p)§
MATCH (n) WITH n.prop AS p RETURN toInteger(DISTINCT  p)§
MATCH (n) WITH n.prop AS p RETURN toInteger( DISTINCT  p)§
MATCH (n) WITH n.prop AS p RETURN toInteger(   DISTINCT   p  )§
RETURN toInteger()§
RETURN toInteger( )§
RETURN toInteger(   )§
RETURN toInteger ()§
RETURN toInteger ( )§
RETURN toInteger   ()§
RETURN toInteger   (  )§
MATCH (n) RETURN cOuNt(*)§
MATCH (n) RETURN cOuNt( * )§
MATCH (n) RETURN cOuNt( *)§
MATCH (n) RETURN cOuNt(* )§
MATCH (n) RETURN cOuNt (*)§
MATCH (n) RETURN cOuNt ( * )§
MATCH (n) RETURN cOuNt ( *)§
MATCH (n) RETURN cOuNt (* )§
MATCH (n) RETURN cOuNt  (*)§
MATCH (a:A)
RETURN a AS a
UNION
MATCH (b:B)
RETURN b AS a§
MATCH (a:A)
RETURN a AS a
UNION
MATCH (b:B)
RETURN b AS a
UNION
MATCH (c:C)
RETURN c AS a§
MATCH (a:A)
RETURN a AS a
UNION ALL
MATCH (b:B)
RETURN b AS a§
MATCH (a:A)
RETURN a AS a
UNION ALL
MATCH (b:B)
RETURN b AS a
UNION ALL
MATCH (c:C)
RETURN c AS a§
MATCH (a:A) RETURN a AS a UnIoN     MATCH (b:B) RETURN b AS a§
MATCH (a:A) RETURN a AS a uNiOn   AlL   MATCH (b:B) RETURN b AS a§
RETURN toInteger(  a  ,  b  ,  c  )§
MATCH (s:Store) WHERE s.storeFormat = 'Open' RETURN DISTINCT s.storeFormat ORDER BY s.storeFormat DESC LIMIT 5§
MATCH (s:Store) WHERE s.storeFormat = 'Open' RETURN DISTINCT s.storeFormat ORDER BY s.storeFormat DESCENDING LIMIT 5§
MATCH (s:Store) WHERE s.storeFormat = 'Open' RETURN DISTINCT s.storeFormat ORDER BY s.storeFormat ASC LIMIT 5§
MATCH (s:Store) WHERE s.storeFormat = 'Open' RETURN DISTINCT s.storeFormat ORDER BY s.storeFormat ASCENDING LIMIT 5§
RETURN '' AS n ORDER BY n    DESC§
RETURN '' AS n ORDER BY   n    ASC§
RETURN '' AS n ORDER BY n   DESCENDING§
RETURN '' AS n ORDER BY  n    ASCENDING§
MATCH (other:Person) WHERE NOT other.age > 25 RETURN other.name§
MATCH (other:Person) WHERE  NOT other.age > 25 RETURN other.name§
MATCH (other:Person) WHERE    NOT   other.age > 25 RETURN other.name§
MATCH (other:Person) WHERE (NOT other.age > 25) RETURN other.name§
MATCH (other:Person) WHERE (NOT   other.age > 25) RETURN other.name§
MATCH (other:Person) WHERE (   NOT   other.age > 25) RETURN other.name§
MATCH (other:Person) WHERE ( NOT other.age > 25) RETURN other.name§
MATCH (other:Person) WHERE ( NOT   other.age  > 25   ) RETURN other.name§
MATCH (other:Person) WHERE   NOT  (  other.age > 25) RETURN other.name§
RETURN count(*)§
WITH {a: 1} AS map RETURN count(map.a)§
RETURN 1 AS e§
RETURN 1.0e10§
WITH 1 AS one WHERE one > 0 RETURN one§
WITH 1 AS one   WHERE one > 0 RETURN one§
WITH 1 AS one    WHERE    one > 0 RETURN one§
WITH 1 AS one
WHERE one > 0
RETURN one§
MATCH (a), (b), (c)
DELETE a, b, c§
MATCH (a), (b), (c)
DELETE   a, b, c§
MATCH (a), (b), (c)
DELETE   a ,  b,c §
MATCH (a), (b), (c)
DETACH DELETE a,b,c§
MATCH (a), (b), (c)
DETACH   DELETE a    ,b   ,c§
MATCH (a), (b), (c)
DETACH   DELETE    a   ,   b   ,  c§
RETURN [x IN [1, '', n] | toString(x) ] AS list§
RETURN [x IN [1, '', n]    | toString(x) ] AS list§
RETURN [x IN [1, '', n]    |    toString(x) ] AS list§
RETURN [x IN [1, '', n]    |    toString(x)    ] AS list§
RETURN [  x   IN   [  1  ,   ''  ,   n  ]    |    toString(x)   +   '-string'  ] AS list§
RETURN []§
RETURN [ ]§
RETURN [   ]§
RETURN [  1  ]§
RETURN [  1, 2  ]§
RETURN [  1   ,  2 ]§
RETURN [  1   ,  2]§
RETURN [1   ,  2]§
RETURN [1,  2]§
RETURN [1,2]§
RETURN 0§
RETURN 0000§
RETURN 012345670§
RETURN 000012345670§
RETURN 1234567890§
RETURN 0x123§
RETURN 0x1234567890abcdef§
RETURN 1e10§
RETURN 1e-10§
RETURN 1.0e-10§
RETURN .1e-10§
RETURN .1e10§
RETURN .1E10§
RETURN -.01E10§
RETURN -590.01E5§
RETURN 1.0§
RETURN 1.00001§
RETURN .00001§
RETURN .0020123§
RETURN 10101.0020123§
RETURN 10101.4§
RETURN 99882.04000§
CREATE ()-[v:R{prop:42}]->()§
CREATE ()-[   v:R{prop:42}]->()§
CREATE ()-[   v   :R{prop:42}]->()§
CREATE ()-[   v   :R{prop:42}]->()§
CREATE ()-[   v   :R   {prop:42}]->()§
CREATE ()-[   v   :R   {   prop:42}]->()§
CREATE ()-[   v   :R   {   prop   :42}]->()§
CREATE ()-[   v   :R   {   prop   :   42}]->()§
CREATE ()-[   v   :R   {   prop   :   42   }]->()§
CREATE ()-[   v   :R   {   prop   :   42   }   ]->()§
MATCH (n) SET n.prop=10§
MATCH (n) SET   n.prop=10§
MATCH (n) SET   n.prop   =10§
MATCH (n) SET   n.prop   =   10§
MATCH (n) SET n:Label§
MATCH (n) SET n   :Label§
MATCH (n) SET n   :Label§
MATCH (n) SET n=$param§
MATCH (n) SET n   =$param§
MATCH (n) SET n   =   $param§
MATCH (n) SET n+=$param§
MATCH (n) SET n   +=$param§
MATCH (n) SET n   +=   $param§
RETURN [()-->()|1]§
RETURN   [()-->()|1]§
RETURN   [   ()-->()|1]§
RETURN   [   ()-->()   |1]§
RETURN   [   ()-->()   |   1]§
RETURN   [   ()-->()   |   1   ]§
RETURN [(a)-->()WHERE a:Label|55 + a.prop]§
RETURN [   (a)-->()WHERE a:Label|55 + a.prop]§
RETURN [   (a)-->()   WHERE a:Label|55 + a.prop]§
RETURN [   (a)-->()   WHERE   a:Label|55 + a.prop]§
RETURN [   (a)-->()   WHERE   a   :Label|55 + a.prop]§
RETURN [   (a)-->()   WHERE   a   :Label|55 + a.prop]§
RETURN [   (a)-->()   WHERE   a   :Label   |55 + a.prop]§
RETURN [   (a)-->()   WHERE   a   :Label   |   55 + a.prop]§
RETURN [   (a)-->()   WHERE   a   :Label   |   55 + a.prop   ]§
RETURN [p=(a)-->()WHERE a:Label|55 + a.prop]§
RETURN [   p=(a)-->()WHERE a:Label|55 + a.prop]§
RETURN [   p   =(a)-->()WHERE a:Label|55 + a.prop]§
RETURN [   p   =   (a)-->()WHERE a:Label|55 + a.prop]§
RETURN 1 // comment§
RETURN 1 // * §
RETURN 1 // */ §
WITH 1 AS one // */
RETURN 2 AS two // another§
RETURN 1 /* blockcomment */§
RETURN 1 /* blockcomment
over several
lines */§
RETURN 1 /* blockcomment
with //
and \\
chars */§
MATCH (a)
WHERE all(a IN []WHERE a.prop )
RETURN 1§
MATCH (a)
WHERE any   (a IN []WHERE a.prop )
RETURN 1§
MATCH (a)
WHERE none   (   a IN []WHERE a.prop )
RETURN 1§
MATCH (a)
WHERE single   (   a   IN []WHERE a.prop )
RETURN 1§
MATCH (a)
WHERE all   (   a   IN   []WHERE a.prop )
RETURN 1§
MATCH (a)
WHERE all   (   a   IN   []   WHERE a.prop )
RETURN 1§
MATCH (a)
WHERE all   (   a   IN   []   WHERE   a.prop )
RETURN 1§
MATCH (a)
WHERE all   (   a   IN   []   WHERE   a.prop   )
RETURN 1§
CALL proc(  a  ,  b  ,  c  ) RETURN *§
CALL proc(a,b  ,  c  ) RETURN *§
CALL proc(  a  ) RETURN *§
CALL proc() RETURN *§
CALL proc() YIELD a AS v RETURN *§
CALL proc(   ) RETURN *§
MATCH (a) CALL proc(   )§
MATCH (a) CALL proc( a, b   ,   c  )§
MATCH (a) CALL proc( a, b   ,   c  ) YIELD a, b , c§
CALL proc(   ) RETURN *
`, "§\n")
