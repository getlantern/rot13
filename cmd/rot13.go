package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/getlantern/rot13"
)

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func main() {
	var err error

	inFile := flag.String("input", "", "Input file; use '-' or leave blank to read from stdin")
	outFile := flag.String(
		"output",
		"",
		"Output file; use '-' or leave blank to write to stdout",
	)
	encode := flag.Bool("encode", false, "Encode the input string (default is to decode)")
	flag.Parse()

	var in io.Reader
	if *inFile == "-" || *inFile == "" {
		in = os.Stdin
	} else {
		file, err := os.Open(*inFile)
		if err != nil {
			die("cannot open input file: %v\n", err)
		}
		defer file.Close()
		in = file
	}

	if !*encode {
		in = rot13.NewReader(in)
	}

	var out io.Writer

	if *outFile == "-" || *outFile == "" {
		out = os.Stdout
	} else {
		file, err := os.Create(*outFile)
		if err != nil {
			die("cannot create output file: %v\n", err)
		}
		defer file.Close()
		out = file
	}

	if *encode {
		out = rot13.NewWriter(out)
	}

	_, err = io.Copy(out, in)
	if err != nil {
		die("error writing to output file: %v\n", err)
	}
}
