package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/liuzl/gocc"
	"io"
	"log"
	"os"
)

var (
	input  = flag.String("input", "", "file of original text to read")
	output = flag.String("output", "", "file of converted text to write")
	config = flag.String("config", "", "convert config, s2t, t2s, etc")
)

func main() {
	flag.Parse()
	var err error
	var in, out *os.File //io.Reader
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	if *output == "" {
		out = os.Stdout
	} else {
		out, err = os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}

	if *config == "" {
		*config = "s2t"
	}

	conv, err := opencc.New(*config)
	if err != nil {
		log.Fatal(err)
	}

	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			log.Fatal(c)
		}
		str, err := conv.Convert(line)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(out, str)
	}
}
