package main

import (
	"testing"
)

func TestStringSplit_No_Split(t *testing.T) {
	testOne := "Test One" // Should be byte len 8
	testOneNoNewLine := []byte(testOne)

	if len(testOneNoNewLine) != 8 {
		t.Error("Incorrect Testing Input")
	}

	outputOne, outputTwo := splitStringByNewline(testOneNoNewLine)

	if outputOne != testOne && outputTwo != "" {
		t.Errorf("Split failed: outputOne: %s\n outputTwo %s\n", outputOne, outputTwo)
	}
}

func TestStringSplit_Split(t *testing.T) {
	testTwo := "Test\nOne" // Should be byte len 8
	testTwoNewLine := []byte(testTwo)

	if len(testTwoNewLine) != 8 {
		t.Error("Incorrect Testing Input")
	}

	outputOne, outputTwo := splitStringByNewline(testTwoNewLine)

	if outputOne != "Test" && outputTwo != "One" {
		t.Errorf("Split failed, outputOne: %s and outputTwo %s\n", outputOne, outputTwo)
	}
}
