package rot13

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	var in, out, rt bytes.Buffer
	for i := byte(0); i <= 255; i++ {
		err := in.WriteByte(i)
		require.NoError(t, err, "Unable to write byte")
		if i == 255 {
			break
		}
	}
	orig := in.Bytes()

	w := NewWriter(&out)
	r := NewReader(&out)

	n, err := io.Copy(w, &in)
	require.NoError(t, err, "Unable to write rot13")
	require.EqualValues(t, 256, n, "Wrong number of bytes written")
	n, err = io.Copy(&rt, r)
	require.NoError(t, err, "Unable to read rot13")
	require.EqualValues(t, 256, n, "Wrong number of bytes read")
	require.Equal(t, len(orig), len(rt.Bytes()), "Size of round-tripped didn't equal original")
	require.Equal(t, orig, rt.Bytes(), "Round-tripped didn't equal original")
}

func TestFunctional(t *testing.T) {
	var err error
	infilePath := os.Getenv("INFILE")
	outfilePath := os.Getenv("OUTFILE")
	require.NotEmpty(t, infilePath)
	require.NotEmpty(t, outfilePath)
	infilePath, err = filepath.Abs(infilePath)
	require.NoError(t, err)
	outfilePath, err = filepath.Abs(outfilePath)
	require.NoError(t, err)

	infile, err := os.Open(infilePath)
	require.NoError(t, err)
	defer infile.Close()
	rot13Reader := NewReader(infile)
	bytes, err := ioutil.ReadAll(rot13Reader)
	require.NoError(t, err)
	require.NotEmpty(t, bytes)

	require.NoError(t, ioutil.WriteFile(outfilePath, bytes, 0644))
	fmt.Printf("Decoded %v to %v successfully\n", infilePath, outfilePath)
}
