package main_test

import (
	"bytes"
	"fmt"
	"io"
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

func TestByteReader(t *testing.T) {
	var buf bytes.Buffer

	_, err := buf.Write([]byte("Hello World"))
	require.NoError(t, err)

	r := bytes.NewReader(buf.Bytes())
	t.Logf("len: %d", r.Len())

	var out bytes.Buffer
	var b byte
	for {
		b, err = r.ReadByte()
		if err == io.EOF {
			break
		}
		if b == 0x6c {
			err := out.WriteByte(b)
			require.NoError(t, err)
		}
	}
	require.Equal(t, out.Bytes(), []byte{'l', 'l', 'l'})
}
