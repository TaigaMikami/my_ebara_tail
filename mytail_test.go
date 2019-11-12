package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFileToTextLines(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	filename := dir + "/txts/hoge.txt"
	textLines := fileToTextLines(filename)

	if len(textLines) != 6 {
		t.Fatalf("Array size is incorrect, but it is %d lines", len(textLines))
	}
}

func TestPickTextLinesFromEnd(t *testing.T) {
	inputTextLines := []string{"x", "y", "z"}
	n := 2
	textLines := pickTextLinesFromEnd(inputTextLines, n)

	for i := range textLines {
		expected := inputTextLines[i+1]
		if textLines[i] != expected {
			t.Fatalf("textLines value is incorrect, expected value is %s, but it is %s", expected, textLines[0])
		}
	}
}

func TestPickTextLinesFromTop(t *testing.T) {
	inputTextLines := []string{"x", "y", "z"}
	n := 2
	textLines := pickTextLinesFromTop(inputTextLines, n)

	for i := range textLines {
		expected := inputTextLines[i]
		if textLines[i] != expected {
			t.Fatalf("textLines value is incorrect, expected value is %s, but it is %s", expected, textLines[0])
		}
	}
}

	func TestPickTextLinesFromRandom(t *testing.T) {
	inputTextLines := []string{"x", "y", "z"}
	n := 2
	textLines := pickTextLinesFromRandom(inputTextLines, n)

	if len(textLines) != n {
		t.Fatalf("size of textLines is incorrect, expected size is %d, but it is %d", n, len(textLines))
	}
}
