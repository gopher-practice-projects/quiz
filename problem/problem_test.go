package problem

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	record := []string{"question", "answer"}

	want := Problem{"question", "answer"}
	got := New(record)

	if got != want {
		t.Errorf("expected to create problem %v got %v", want, got)
	}
}

func TestCheckAnswer(t *testing.T) {
	problem := createProblem()

	t.Run("it checks the correct answer", func(t *testing.T) {
		answer := getAnswer(problem, "10\n")
		checkAnswer(t, answer, true)
	})

	t.Run("it checks incorrect answer", func(t *testing.T) {
		answer := getAnswer(problem, "2\n")
		checkAnswer(t, answer, false)
	})
}

func createProblem() Problem {
	record := []string{"7+3", "10"}
	return New(record)
}

func getAnswer(problem Problem, input string) bool {
	r := bytes.NewBufferString(input)
	return problem.CheckAnswer(r)
}

func checkAnswer(t *testing.T, got, want bool) {
	if want != got {
		t.Errorf("Expected to return %v got %v", want, got)
	}
}
