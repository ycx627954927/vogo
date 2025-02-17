// Copyright 2019 The vogo Authors. All rights reserved.

package vos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProcessUser(t *testing.T) {
	u, err := GetProcessUser(os.Getpid())
	assert.Nil(t, err)
	t.Logf("user: %s", u)
}
