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