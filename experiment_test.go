package scientist

import (
	"testing"
	"time"
)

const expectedResult int = 2

func longLoop() (int, error) {
	for i := 0; i < 1000; i++ {
		_ = i * 2
	}
	return expectedResult, nil
}

func shortLoop() (int, error) {
	for i := 0; i < 100; i++ {
		_ = i * 2
	}
	return expectedResult, nil
}

func longSummation(start, end int) (int, error) {
	time.Sleep(time.Second * 2)
	var result int
	for i := start; i < end; i++ {
		result += i
	}
	return result, nil
}

func shortSummation(start, end int) (int, error) {
	var result int
	for i := start; i < end; i++ {
		result += i
	}
	return result, nil
}

func TestNewExperiment(t *testing.T) {
	experimentName := "Loops"
	exp := New(experimentName, true)

	if exp.Name != experimentName {
		t.Log("Incorrect experiment name set")
		t.Log("Expected", experimentName, "\n Got", exp.Name)
		t.Fail()
	}
}

func TestRandomExperimentRun(t *testing.T) {
	experimentName := "Loops"
	exp := New(experimentName, true)

	exp.Use(func() (interface{}, error) {
		return longLoop()
	})
	exp.Try(func() (interface{}, error) {
		return shortLoop()
	})

	result, err := exp.Run()

	if err != nil {
		t.Log("Experiment should not return error")
		t.Log("Expected", nil, "\n Got", err)
		t.Fail()
	}

	if result != expectedResult {
		t.Log("Experiment did not return expected result")
		t.Log("Expected", expectedResult, "\n Got", result)
	}
}

func TestNormalExperimentRun(t *testing.T) {
	experimentName := "Summation"
	exp := New(experimentName, false)
	start := 0
	end := 1000
	expectedResult := float32(end * (end - 1) / 2)

	exp.Use(func() (interface{}, error) {
		return longSummation(start, end)
	})
	exp.Try(func() (interface{}, error) {
		return shortSummation(start, end)
	})

	result, err := exp.Run()

	if err != nil {
		t.Log("Experiment returned an error")
		t.Log("Expected", nil, "\n Got", err)
		t.Fail()
	}

	if float32(result.(int)) != expectedResult {
		t.Log("Experiment did not return expected result")
		t.Log("Expected", expectedResult, "\n Got", result)
	}
}
