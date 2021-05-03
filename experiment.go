package main

import (
	"errors"
	"fmt"
	"time"
)

type Result struct {
	ControlDuration   float64
	CandidateDuration float64
	ResultsAreEqual   bool
	Duration          float64
	controlResult     interface{}
	candidateResult   interface{}
	candidateError    error
}

type Experiment struct {
	Name string
	Result
}

func (e *Experiment) Use(runner func() (result interface{}, err error)) {
	start := time.Now()
	result, err := runner()

	if err != nil {
		panic(err)
	}

	e.Result.controlResult = result

	defer func() {
		diff := float64(time.Since(start) / time.Second)
		fmt.Println("diff use is", diff)
		e.Result.ControlDuration = diff
	}()
}

func (e *Experiment) Try(runner func() (result interface{}, err error)) {
	start := time.Now()
	result, err := runner()

	if err != nil {
		e.Result.candidateError = err
	}

	e.Result.candidateResult = result

	defer func() {
		diff := float64(time.Since(start) / time.Second)
		fmt.Println("diff try is", diff)
		e.Result.CandidateDuration = diff
	}()
}

func (e *Experiment) Run() Result {
	e.Result.Duration = e.Result.ControlDuration + e.Result.CandidateDuration
	return e.Result
}

func main() {

	exp := Experiment{Name: "Loops"}

	control := func() (interface{}, error) {
		return longLoop()
	}
	candidate := func() (interface{}, error) {
		return shortLoop()
	}

	exp.Use(control)

	exp.Try(candidate)

	result := exp.Run()

	fmt.Println("Duration", result.Duration, "control duration", result.ControlDuration, "candidate duration", result.CandidateDuration)
	fmt.Printf("result is %v", result)
	fmt.Println("candidate error", result.candidateError)
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
	return 0, errors.New("short loop error oo")
}
