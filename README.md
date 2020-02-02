# Gophercises
<p align="center" padding="500px">
<img src="https://gophercises.com/img/gophercises_jumping.gif" height="150px">
</p>


Collection of exercises in Golang from the amazing [gophercises.com](https://gophercises.com).
## 1. Quiz Game
[Problem Statement](https://github.com/gophercises/quiz), [My solution](https://github.com/aayush4vedi/Gophercises/tree/master/quiz)
#### Packages Used 
* flag => for Command Line Flags(returns a pointer) :: [GoByExample](https://gobyexample.com/command-line-flags)
* encoding/csv : nice [medium article](https://medium.com/@barunthapa/working-with-csv-in-go-50a4f540e623)
* strings => `strings.TrimSpace()`
* fmt : Taking i/o input => `fmt.Scanf("%s", &ans)`
* os
  * `os.Open(*path)`
  * `os.Exit(1)`
* time
* math/rand 
  * Shuffle an array:
   ```rand.Seed(time.Now().UnixNano())
    perm := rand.Perm(len(arr))
    for i, v := range perm {
	      newArr[v] = arr[i]
    }```
## 2. URL Shortner
[Problem Staement](https://github.com/gophercises/urlshort) ,[My Solution](https://github.com/aayush4vedi/Gophercises/tree/master/urlshort)
#### Packages Used
* `gopkg.in/yaml.v2`

## 3. Choose Your Own Adventure
[Problem Statement](https://github.com/gophercises/cyoa), [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/cyoa)
#### Packages Used
* `encoding/json`
* `html/template`


## 4. HTML Parse
[Problem Statement](https://github.com/gophercises/link) ,[Solution](https://github.com/aayush4vedi/Gophercises/tree/master/link)
### Highlight
Used DFS
#### Packages Used
* `golang.org/x/net/html`
* `strings` => `strings.Join(strings.Fields(ret), " ")`

## 5. Sitemap
[Problem Statement[(https://github.com/gophercises/sitemap), [Solution](https://github.com/aayush4vedi/Gophercises/blob/master/sitemap/main.go)

### Highlight
Used BFS!!

#### Packages Used
* `strings` => `strings.HasPrefix(targetString, prefixToCheck)`
* `net/http`
  * GETting a page
  * get base url
* `encoding/xml` 
  * similar to json.Look at the [example](https://golang.org/pkg/encoding/xml/)

## 6. Hackerrank Problems
* CamleCase: [Problem](https://www.hackerrank.com/challenges/camelcase/problem) [Solution](https://github.com/gophercises/hr1/tree/solution/camel)
* Caeser-Cipher: [Problem](https://www.hackerrank.com/challenges/caesar-cipher-1/problem), [Solution](https://github.com/gophercises/hr1/tree/solution/caesar)
 ** Learnt about [rune](https://www.geeksforgeeks.org/rune-in-golang/)

## 7. CLI Task Manager
[Problem](https://courses.calhoun.io/lessons/les_goph_36) , [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/task)

### Highlight
* Connected with DB - **BoltDB**
* CLI package- **Cobra**

#### Packages Used

* BoltDB: `github.com/boltdb/bolt`
  * Everything is on [github](https://github.com/boltdb/bolt) page.
  * All buckets have key/value pairs - both in byte slices `[]byte`, so need to encode/decode int->byte
  * `encoding/binary` : [go-page](https://golang.org/pkg/encoding/binary/), [my implementation](https://github.com/aayush4vedi/Gophercises/blob/7aac7de0cab35554f30be51eb3af597e234997fe/task/db/task.go#L71)
* `storm` : provides toolkit for BoltDB. Is like it's ''gorm'' . [doc](https://github.com/asdine/storm)
* `github.com/spf13/cobra`
  ** [Readme](https://github.com/spf13/cobra)
*  `go-homedir`
  ** [readme](https://github.com/mitchellh/go-homedir)
* `strconv` => [Atoi](https://golang.org/pkg/strconv/#example_Atoi)

## 8. Phone Number Normalizer
[Problem](https://courses.calhoun.io/lessons/les_goph_46) , [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/phone)

### Highlight
* DB: postrgesql
* gorm

#### Used packages:
* `database/sql`: [doc](https://golang.org/pkg/database/sql/), all commands from here
* go-sql driver(for postgres):`lib/pq` : [doc](https://godoc.org/github.com/lib/pq)
* `sqlx` : [doc](https://github.com/jmoiron/sqlx), provides extensio to `database/sql` package
* `jinzhu/gorm` : [github](https://github.com/jinzhu/gorm) 