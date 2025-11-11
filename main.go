package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

func main() {
	assert.Equal(nil, 1, 1, "test")
	fmt.Println("Hello, go mod tidy!")
}
