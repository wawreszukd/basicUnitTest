package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	result, _ := add(1, 2)
	assert.Equal(t, 3, result)

	_, err := add(1, 2)
	assert.Error(t, err)
}
