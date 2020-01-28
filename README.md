# Gophercises
<p align="center" padding="500px">
<img src="https://gophercises.com/img/gophercises_jumping.gif" height="150px">
</p>


Collection of exercises in Golang from the amazing [gophercises.com](https://gophercises.com).
## 1. Quiz Game
[Problem Statement]() . [My solution](https://github.com/aayush4vedi/Gophercises/tree/master/quiz)
### Packages Used 
* flag => for Command Line Flags(returns a pointer) :: [[GoByExample|https://gobyexample.com/command-line-flags]]
* encoding/csv : nice [[medium article|https://medium.com/@barunthapa/working-with-csv-in-go-50a4f540e623]]
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
