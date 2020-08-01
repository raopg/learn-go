## Learning Go With TDD

This repo contains all of my exercises for the TDD-based Go learning course. Source: https://quii.gitbook.io/learn-go-with-tests/

## Things I've learned so far (That are unique to Go besides syntax)

1. Go is very opinionated about folder structure - it is usually contained within $GOPATH.
2. Newer versions of Go will now support modules. In order to use one of these we need to run `go mod init <project-name>` within the folder where your project resides.
3. Tests are always written in files called `<module-to-be-tested>_test.go`
4. By default, all functions within a module are exported out to other files (sounds like a nightmare for functions with same names)
5. Good Go practice involves writing a comment for each function that starts with the function name (linter rules)
6. If-else-if statements: the else and else if statements need to be in the same line as the closing bracket of the if block. Ex:

```
    CORRECT
    if <condition>{
        //stuff
    } else if <condition-2> {
        //stuff
    } else {

    }
    INCORRECT
    if <condition>{
        //stuff
    } 
    else if <condition-2> {
        //stuff
    } 
    else {

    }
```
7. Named return values are those variables you can specify in the function signature. They will be initialized with a default value. We can return at any point in the function without specifying the return variable, and the specified named return value will be returned to the point of the function call. Ex:

```
    func callingFn(...){
        //do stuff
        val := calledFn(...)
        //do stuff
    }
    func calledFn(...)(namedReturn string) {
        //do stuff
        return
    }
```
8. In Go, there are no while, do-while loops. All iterative processes involve a for loop.
9. We can write benchmark tests for functions using the in-built `testing` module
10. Arrays - arrays are of fixed sizes and type. Slices are resizable arrays - there exists a 'slice' data type.
11. Go provides a test coverage tool. Ideally, we want this to be as close to 100% as possible. In order to check, run `go test -cover`
12. len => provides the length of an array or slice.
13. make() can be used to create slices/arrays with specific data types (e.g slice of slices)
14. append() method can be used to work with slices (just like Python). The syntax is `append(slice, newItem)`
15. Slices can be sliced! nums[i:j] will return a slice of nums from index i to j - 1
16. You cannot overload methods in Go!
17. Solution to overloading -> METHODS! Methods are functions bound to a specific instance.
18. Method syntax:

```
    func (receiver ReceiverType) functionName(args) rtype {
        //do stuff
    }
```
19. Interfaces -> this can represent high level structures that combine common functionality between multiple structs.
Much like a parent class, an interface will capture common functionality. So if the type you pass in wherever an interface type is required, if all the fields that are present in the interface is present in the passed type, Go compiles the code.
20. Table driven tests -> when we want to test our function on a set of different inputs. If the function to be tested belongs to an interface, we can test it against multiple data types/structs that belong to the interface.
21. Pointers -> same concept in C/C++, with ONE exception -> We dont dereference pointers. So if we try to access p.data, we can. We don't need p->data or (*p).data.
22. Maps -> very similar to dicts in Python. They are, however, passed by reference. So there is no need to create pointers for them. However, maps can be nil, and might cause panic. In order to make sure that map is initialzed to empty map, be sure to initialize like - 

```
    DO
    var myMap map[string]string{}
    or
    var myMap = make(map[string]string)
    DONT
    var myMap map[string]string
```