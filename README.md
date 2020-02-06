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
 * Learnt about [rune](https://www.geeksforgeeks.org/rune-in-golang/)

## 7. CLI Task Manager
[Problem](https://courses.calhoun.io/lessons/les_goph_36) , [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/task)

### Highlight
* Connected with DB - *BoltDB*
* CLI package- *Cobra*

#### Packages Used

* BoltDB: `github.com/boltdb/bolt`
  * Everything is on [github](https://github.com/boltdb/bolt) page.
  * All buckets have key/value pairs - both in byte slices `[]byte`, so need to encode/decode int->byte
  * `encoding/binary` : [go-page](https://golang.org/pkg/encoding/binary/), [my implementation](https://github.com/aayush4vedi/Gophercises/blob/7aac7de0cab35554f30be51eb3af597e234997fe/task/db/task.go#L71)
* `storm` : provides toolkit for BoltDB. Is like it's *gorm* . [doc](https://github.com/asdine/storm)
* `github.com/spf13/cobra`
  * [Readme](https://github.com/spf13/cobra)
*  `go-homedir`
  * [readme](https://github.com/mitchellh/go-homedir)
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

## 9. Deck of Cards
[Problem](https://courses.calhoun.io/lessons/les_goph_54) , [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/deck)

### Highlight
* A good sysD problem #interviews
* Tests are written properly

#### Used Packages
* `math/rand`
* `sort` 
  * `sort.Slice(cards, less(cards))` : [implementation](https://github.com/aayush4vedi/Gophercises/blob/82be1cda692638652c555f7f3419a8b8e5dae501/deck/card.go#L75)
* `fmt.stringer` : [Article](https://riptutorial.com/go/example/9983/stringer) , [doc](https://godoc.org/golang.org/x/tools/cmd/stringer)
  * `//go:generate stringer -type=Suit,Rank` , Here these are the names of types we want to generate methods for.
  * `//cmd: $ go generate`
* *Shuffling*:
  * [How to shuffle](https://github.com/aayush4vedi/Gophercises/blob/82be1cda692638652c555f7f3419a8b8e5dae501/deck/card.go#L96)
  * [How to test](https://github.com/aayush4vedi/Gophercises/blob/82be1cda692638652c555f7f3419a8b8e5dae501/deck/card_test.go#L81)

## 10. File Renamer
[Problem](https://courses.calhoun.io/lessons/les_goph_79) , [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/renamer)

### Highlight
* *Files* & *Filepath*
* *Regex*

#### Used Packages:

* `path/filepath`  
  * `filepath.Walk`
  * `filepath.Dir`
  * `filepath.Join`  // as diff OS use diff joining(`\` in windows)
* `regex` : [doc](https://golang.org/pkg/regexp/) , [My implementation](https://github.com/aayush4vedi/Gophercises/blob/aeac8b0a32766c8766328ae89de74f4cff7fb36b/renamer/main.go#L74)
 * `regexp.MustCompile`
* `os` => `os.Rename`
* `strings` 
  * `sort.Strings()`
  * `strings.Split()`
  * `strings.Join()`
  * `strings.Title()`

## 11. Quiet HackerNews
[Problem](https://courses.calhoun.io/lessons/les_goph_85) , [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/quiethn)

### Highlight
* Thorough understanding of **concurrency**(goroutines, channels, mutex)
* **How not to use goroutines**
* **Caching**
* Using Template(`.gohtml`)

#### Used Packages
* `html/template` 
  * `tpl := template.Must(template.ParseFiles("./index.gohtml"))`
* `net/http`
  * `http.HandleFunc("/", handler(numStories, tpl))`
  * Starting server:  `log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))`
  * Loggin errors: `http.Error(w, err.Error(), http.StatusInternalServerError)`
* `time`
  * ticker: `ticker := time.NewTicker(3 * time.Second)`
  * Comparing times: `if time.Now().Sub(sc.expiration) < 0`
* `sync`
 * **mutex** :Load time got reduced from 2 sec to 1.2µs! Caching is always much faster than all the concurrencies.
   * `sc.mutex.Lock()
defer sc.mutex.Unlock()` ~/Always defer a mutex; else app will be frozen & you'll keep on scratching head`

* `sort`
 * Sorting 2 structs based on one key:
  * ``` type result struct {
		idx  int
		item item
		err  error
	}
        sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})```

* You can limit your workers via channels, or with something like the [x/sync/semaphore](https://godoc.org/golang.org/x/sync/semaphore) package

# 12. Recover
- Panic/Recover Middleware for an HTTP server.
[Problem](https://courses.calhoun.io/lessons/les_goph_92) [Solution](https://github.com/aayush4vedi/Gophercises/blob/master/recover/main.go) 

### Highlight
* Middleware
* Defer, panic, recover
* Errors
  * Errors: `w.WriteHeader(http.StatusInternalServerError)`
  * Error-stack : `debug.Stack()`

#### Used Packages
* `runtime` => `stack()`  [doc](https://golang.org/pkg/runtime/#Stack)
* `net/http`
 * ResponseWriter: is just an interface. [doc](https://golang.org/pkg/net/http/#ResponseWriter)
* Midllware Interface:
 * [Hijacker](https://golang.org/pkg/net/http/#Hijacker)
 * [Flusher](https://golang.org/pkg/net/http/#Flusher) =>[implementation](https://github.com/urfave/negroni/blob/master/response_writer.go#L96)

* 3rd party middleware for panic/recover:
* [negroni](https://github.com/urfave/negroni)
* [chi](https://github.com/go-chi/chi)
