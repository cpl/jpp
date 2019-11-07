package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)


var (
	indent = flag.String("indent", "  ", "string used for indentation")
	prefix = flag.String("prefix", "", "string used as prefix")
	filename = flag.String("file", "", "use a file as input")
)

func main() {
	in := os.Stdin

	flag.Parse()

	if *filename != "" {
		fp, err := os.Open(*filename)
		if err != nil {
			log.Fatalf("failed opening input file, %s\n", err.Error())
		}
		in = fp
	}

	data, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatalf("failed reading input, %s\n", err.Error())
	}

	buffer := bytes.NewBuffer(nil)
	if err = json.Indent(buffer, data, *prefix, *indent); err != nil {
		log.Fatalf("failed indenting data, %s\n", err.Error())
	}

	if _, err := os.Stdout.Write(buffer.Bytes()); err != nil {
		log.Fatalf("failed writing data to output, %s\n", err.Error())
	}

	os.Exit(0)
}
