package main

import (
	"reflect"
	"testing"

	"./github"
)

func TestNewClient(t *testing.T) {
	var c github.Client
	if reflect.TypeOf(github.NewClient("token")) != reflect.TypeOf(&c) {
		t.Errorf("NewClient(token string) is not returned correct type. %v, %v",
			reflect.TypeOf(github.NewClient("token")),
			reflect.TypeOf(c))
	}
}
