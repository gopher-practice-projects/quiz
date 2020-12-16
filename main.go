package main

import "flag"

const (
	// FileFlag is used to set a file for the question
	FileFlag = "file"
	// FileFlagValue is the value used when no FileFlag is provided
	FileFlagValue = "problems.csv"
	// FileFlagUsage is the help string for the FileFlag
	FileFlagUsage = "Question file"

	// TimerFlag is used for setting a timer for the quiz
	TimerFlag = "timer"
	// TimerFlagValue is the valus used when no TImerFlag is provided
	TimerFlagValue = 30
	// TimerFlagUsage is the help string for the TimerFlag
	TimerFlagUsage = "Amount of seconds the quiz will allow"
)

// Flagger configures the flags used
type Flagger interface {
	StringVar(p *string, name, value, usage string)
	IntVar(p *int, name string, value int, usage string)
}

type quizFlagger struct{}

func (q *quizFlagger) StringVar(p *string, name, value, usage string) {
	flag.StringVar(p, name, value, usage)
}

func (q *quizFlagger) IntVar(p *int, name string, value int, usage string) {
	flag.IntVar(p, name, value, usage)
}

// TimerSeconds is the amount of time allowed for the quiz
var TimerSeconds int
var file string

// ConfigFlags sets all the flags used by the application
func ConfigFlags(f Flagger) {
	f.StringVar(&file, FileFlag, FileFlagValue, FileFlagUsage)
	f.IntVar(&TimerSeconds, TimerFlag, TimerFlagValue, TimerFlagUsage)
}

func init() {
	flagger := &quizFlagger{}

	ConfigFlags(flagger)

	flag.Parse()
}
