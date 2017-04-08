package main

import (
	"os"
	"testing"
)

func isExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func TestLissajous(t *testing.T) {
	fileName := "test.png"
	file, err := os.Create(fileName)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer file.Close()
	lissajous(file)
	if !isExist(fileName) {
		t.Errorf("%v", err)
	}
}
