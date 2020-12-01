package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySum(t *testing.T) {
	assert.Equal(t, 11, MySum(6,5))
}
