# Gophercises
<p align="center" padding="500px">
<img src="https://gophercises.com/img/gophercises_jumping.gif" height="150px">
</p>


Collection of exercises in Golang from the amazing [gophercises.com](https://gophercises.com).
## 1. Quiz Game
* [Problem Statement](https://github.com/gophercises/quiz)
* [My solution](https://github.com/aayush4vedi/Gophercises/tree/master/quiz)
### Packages Used 
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
* [Problem Staement](https://github.com/gophercises/urlshort) 
* [My Solution](https://github.com/aayush4vedi/Gophercises/tree/master/urlshort)
### Packages Used
* `gopkg.in/yaml.v2`

## 3. Choose Your Own Adventure
* [Problem Statement](https://github.com/gophercises/cyoa)
* [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/cyoa)
### Packages Used
* `encoding/json`
* `html/template`


## 4. HTML Parse
* [Problem Statement](https://github.com/gophercises/link) 
* [Solution](https://github.com/aayush4vedi/Gophercises/tree/master/link)
### Packages Used
* `golang.org/x/net/html`
* `strings` => `strings.Join(strings.Fields(ret), " ")`
