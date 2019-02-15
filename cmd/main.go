package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/liuzl/gocc"
	"github.com/liuzl/goutil"
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

	conv, err := gocc.New(*config)
	if err != nil {
		log.Fatal(err)
	}

	err = goutil.ForEachLine(br, func(line string) error {
		str, e := conv.Convert(line)
		if e != nil {
			return e
		}
		fmt.Fprint(out, str+"\n")
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
