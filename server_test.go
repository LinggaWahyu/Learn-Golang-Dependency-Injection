package main

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	server := InitializedServer()

	fmt.Println(server)
}
