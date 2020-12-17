package main

import (
	"reflect"
	"testing"
)

type flaggerMock struct {
	stringValCalls  int
	intVarCalls     int
	varNames        []string
	varUsages       []string
	varStringValues []string
	varIntValues    []int
}

func (f *flaggerMock) StringVar(p *string, name, value, usage string) {
	f.stringValCalls++
	f.varNames = append(f.varNames, name)
	f.varStringValues = append(f.varStringValues, value)
	f.varUsages = append(f.varUsages, usage)
}

func (f *flaggerMock) IntVar(p *int, name string, value int, usage string) {
	f.intVarCalls++
	f.varNames = append(f.varNames, name)
	f.varIntValues = append(f.varIntValues, value)
	f.varUsages = append(f.varUsages, usage)
}

func TestConfigFlag(t *testing.T) {
	flagger := &flaggerMock{}

	ConfigFlags(flagger)
	assertStringCalls(t, flagger)
	assertIntCalls(t, flagger)
	assertsFlags(t, flagger)
}

func assertStringCalls(t *testing.T, flagger *flaggerMock) {
	t.Helper()
	if flagger.stringValCalls != 1 {
		t.Errorf("it should call StringVar %d times, called %d", 1, flagger.stringValCalls)
	}
}

func assertIntCalls(t *testing.T, flagger *flaggerMock) {
	t.Helper()
	if flagger.intVarCalls != 1 {
		t.Errorf("it should call IntVar %d times, called %d", 1, flagger.intVarCalls)
	}
}

func assertsFlags(t *testing.T, flagger *flaggerMock) {
	t.Helper()

	expectedNames := []string{FileFlag, TimerFlag}
	expectedUsages := []string{FileFlagUsage, TimerFlagUsage}
	expectedStringValues := []string{FileFlagValue}
	expectedIntValues := []int{TimerFlagValue}

	if !reflect.DeepEqual(expectedNames, flagger.varNames) {
		t.Errorf("it should setup flag names to be %v, got %v", expectedNames, flagger.varNames)
	}
	if !reflect.DeepEqual(expectedUsages, flagger.varUsages) {
		t.Errorf("it should setup flag usages to be %v, got %v", expectedUsages, flagger.varUsages)
	}
	if !reflect.DeepEqual(expectedStringValues, flagger.varStringValues) {
		t.Errorf("it should setup string value to be %v, got %v", expectedStringValues, flagger.varStringValues)
	}
	if !reflect.DeepEqual(expectedIntValues, flagger.varIntValues) {
		t.Errorf("it should set int value to be %v, got, %v", expectedIntValues, flagger.varIntValues)
	}
}
