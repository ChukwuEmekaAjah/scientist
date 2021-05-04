package scientist

import (
	"math/rand"
	"time"
)

type Result struct {
	ControlDuration   float64
	CandidateDuration float64
	ResultsAreEqual   bool
	Duration          float64
	ControlResult     interface{}
	CandidateResult   interface{}
	CandidateError    error
	ControlError      error
}

type Experiment struct {
	Name string
	Result
	functions map[string]func() (interface{}, error)
	random    bool
}

func (e *Experiment) Use(runner func() (interface{}, error)) {

	e.functions["control"] = func() (interface{}, error) {

		start := time.Now()
		result, err := runner()

		if err != nil {
			e.Result.ControlResult = nil
			return nil, err
		} else {
			e.Result.ControlResult = result
		}

		defer func() {
			diff := float64(time.Since(start) / time.Second)
			e.Result.ControlDuration = diff
		}()

		return result, nil
	}

}

func (e *Experiment) Try(runner func() (interface{}, error)) {

	e.functions["candidate"] = func() (interface{}, error) {
		start := time.Now()
		result, err := runner()

		if err != nil {
			e.Result.CandidateError = err
		}

		e.Result.CandidateResult = result

		defer func() {
			diff := float64(time.Since(start) / time.Second)
			e.Result.CandidateDuration = diff
		}()

		return result, nil
	}

}

func (e *Experiment) Run() (interface{}, error) {

	defer func() {
		e.Result.Duration = e.Result.ControlDuration + e.Result.CandidateDuration
		e.Result.ResultsAreEqual = e.Result.ControlResult == e.Result.CandidateResult
	}()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	if e.random == true {
		value := r1.Intn(100)

		if value < 50 { // run candidate or control first
			e.functions["candidate"]()
		} else {
			e.functions["control"]()
		}

		if value < 50 { // run control or candidate next
			e.functions["control"]()
		} else {
			e.functions["candidate"]()
		}
	} else {
		e.functions["candidate"]()
		return e.functions["control"]()
	}

	return e.Result.ControlResult, e.Result.ControlError
}

func New(name string, random bool) *Experiment {

	exp := Experiment{Name: name, random: random, functions: make(map[string]func() (interface{}, error))}

	return &exp
}
