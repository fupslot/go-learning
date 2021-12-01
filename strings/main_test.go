package main_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer

	_, err := buf.Write([]byte("1"))
	require.NoError(t, err)

	t.Log(fmt.Sprintf("%b", buf.Bytes()))
	require.Equal(t, buf.Bytes(), []byte("1"))
}
