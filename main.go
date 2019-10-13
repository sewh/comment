package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("comment usage: comment op comment_string")
		fmt.Println("")
		fmt.Println("to comment with ' ': comment + ' '")
		fmt.Println("to remove ' ' comment: comment - ' '")
		os.Exit(0)
	}

	op := args[0]
	comment := args[1]

	if op != "+" && op != "-" {
		fmt.Println("op must be '+' to comment or '-' to uncomment")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	var commentRegex *regexp.Regexp
	if op == "-" {
		cR, err := regexp.CompilePOSIX(comment)
		commentRegex = cR
		if err != nil {
			fmt.Println("Bad regex: " + comment)
		}
	} else {
		commentRegex = nil
	}

	for {
		line, err := reader.ReadString('\n')
		dieAfterwards := false
		if err == io.EOF {
			dieAfterwards = true
		}
		if err != nil && err != io.EOF {
			fmt.Println("err: " + err.Error())
			os.Exit(1)
		}

		if op == "+" {
			fmt.Printf("%s%s", comment, line)
		}

		if op == "-" {
			toWrite := commentRegex.ReplaceAll([]byte(line), []byte(""))
			fmt.Printf("%s", string(toWrite))
		}

		if dieAfterwards {
			break
		}

	}
}
