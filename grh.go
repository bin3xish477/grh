package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

const (
	red    = "\u001b[91m"
	green  = "\u001b[32m"
	blue   = "\u001b[94m"
	yellow = "\u001b[33m"
	end    = "\u001b[0m"
	underL = "\u001b[4m"
)

var args struct {
	Url string `arg:"-u,--url" help:"URL to make request to"`
}

func getHTTPHeaders(url string) map[string]string {
	return make(map[string]string)
}

func printHeaders() {
}

func main() {
	arg.MustParse(&args)

	var headers map[string]string
	if args.Url != "" {
		headers = getHTTPHeaders(args.Url)
	} else {
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeCharDevice == 0 && stat.Size() == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				url := scanner.Text()
				headers = getHTTPHeaders(url)
				fmt.Println(headers)
			}
		}
	}
}
