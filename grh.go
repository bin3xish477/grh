package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
)

const (
	red    = "\u001b[91m"
	green  = "\u001b[32m"
	yellow = "\u001b[33m"
	end    = "\u001b[0m"
	underL = "\u001b[4m"
)

var args struct {
	Url string `arg:"-u,--url" help:"URL to make request to"`
}

func parseHeaders(headers http.Header) map[string]string {
	parsed_headers := make(map[string]string)
	for key, value := range headers {
		parsed_headers[key] = strings.Join(value, " ")
	}
	return parsed_headers
}

func getHTTPHeaders(url string) {
	// Stop redirect to capture headers
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("unable to make GET request to %s", url)
		os.Exit(1)
	}

	headers := parseHeaders(resp.Header)
	printHeaders(headers)

	location := resp.Header.Get("Location")

	if location != "" {
		fmt.Printf("%sREDIRECT%s: %s%s%s\n", green, end, yellow, location, end)
		getHTTPHeaders(location)
	}
}

func printHeaders(headers map[string]string) {
	for k, v := range headers {
		fmt.Printf("  %s%s%s: %s\n", red, k, end, v)
	}
}

func main() {
	arg.MustParse(&args)

	if args.Url != "" {
		url := args.Url
		fmt.Printf("%sGET%s: %s%s%s\n", green, end, yellow, url, end)
		getHTTPHeaders(url)
	} else if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			url := scanner.Text()
			fmt.Printf("%sGET%s: %s%s%s\n", green, end, yellow, url, end)
			getHTTPHeaders(url)
		}
	} else {
		fmt.Println("must use `--url` flag or provide URLs from stdin")
	}
}
