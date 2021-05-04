<h1 align="center">Scientist</h1>

A naive Go implementation of the Scientist Ruby gem. Use it to experiment with the performance of new function/methods implementation.


## What you get
- The ability to create experiments and compare new implementations to already existing implementations
- Run series of experiments.
- Randomize the execution of experiments

## How to use
```bash
go get github.com/ChukwuEmekaAjah/scientist
```

```go

    package main

    import (
        "errors"
        "fmt"

        "github.com/ChukwuemekaAjah/scientist"
    )

    func main() {

        exp := scientist.New("Loops", true)

        control := func() (interface{}, error) {
            return longLoop()
        }
        candidate := func() (interface{}, error) {
            return shortLoop()
        }

        exp.Use(control)

        exp.Try(candidate)

        val, _ := exp.Run()

        fmt.Printf("value is %v", val)

        fmt.Println("Duration", exp.Result.Duration, "control duration", exp.Result.ControlDuration, "candidate duration", exp.Result.CandidateDuration)
        fmt.Printf("exp.Result is %v", exp.Result)
        fmt.Println("candidate error", "exp.Result.candidateError")
    }

    func longLoop() (int, error) {

        for i := 0; i < 100000; i++ {
            fmt.Println(i)
        }
        return 2, nil
    }

    func shortLoop() (int, error) {
        for i := 0; i < 100; i++ {
            _ = i * 2
        }
        return 2, errors.New("short loop error oo")
    }
```

## API 
Methods on the <b>Experiment</b> struct:
- exp.Use: Runs the original implementation of the function or method to be experimented on
- exp.Try: Runs the new implementation of the function or method
- exp.Run: Executes `.Try` and `.Use` and then returns the values from `.Use`
- scientist.New: Creates a new experiment struct and returns a pointer to the created experiment.

Properties of the <b>Experiment</b> struct:
- exp.Result: It's a struct that contains fields on the duration of the experiment, duration of the control experiment and that of the candidates.


## Understanding API methods and properties
WIP


## Contributing
In case you have any ideas, features you would like to be included or any bug fixes, you can send a PR.

- Clone the repo

```bash
git clone https://github.com/ChukwuEmekaAjah/scientist.git
```

## Todo
- Write more documentation on the explanation of experiment methods parameters.