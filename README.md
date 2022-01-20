# rot13 encoder/decoder

Rotates a byte value by thirteen characters.

## Caveat

The name of this package is something of a misnomer: it does not implement
the traditional [ROT13][rot13] substitution cipher. In that cipher, the
value "13" takes on special meaning because the cipher is supposed to
operate on a 26-letter alphabet (2x13); because of this, one can encode and
decode using the same algorithm (it is its own inverse). In contrast, this
library operates on a full _byte_, so its "alphabet" is 256 characters
instead of 26, and thus its algorithm is not its own inverse. If you are
looking for a library to perform the ROT13 Caesar cipher, this is not for
you.

## Command-line use

There is a simple command line interface to this library. Clone the repo and
run `go build cmd/rot13.go` to build it.

Running `rot13 -h` will give you usage information:

    Usage of ./rot13:
      -encode
            Encode the input string (default is to decode)
      -input string
            Input file; use '-' or leave blank to read from stdin
      -output string
            Output file; use '-' or leave blank to write to stdout

If you do not specify an input file, it will read from `stdin`; with no
output file specified, it will print results to `stdout`. Some examples:

    $ echo 'hello, world!' > hello
    $ rot13 -encode -input hello
    uryy|9-|yq.

    $ rot13 -encode -input hello -output hello_encoded
    $ cat hello_encoded
    uryy|9-|yq.

    $ ./rot13 -input hello_encoded
    hello, world!

    $ echo 'hello, world!' | ./rot13 -encode
    uryy|9-|yq.

    $ echo 'hello, world!' | ./rot13 -encode | ./rot13
    hello, world!


You can also use Go's testing facility to encode and decode. Supply the
input/output file path as env variables and run:

    INFILE=~/.lantern/global.yaml \
      OUTFILE=~/.lantern/expanded_global.yaml \
      go test -v -run TestFunctional

[rot13]: https://en.wikipedia.org/wiki/ROT13
