package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rnemeth90/hashy"
)

type config struct {
	// stdin, flag, os.arg[1]
	plainText string
	algorithm string
}

var (
	text   string
	cypher string
)

func init() {
	flag.StringVar(&text, "t", "", "the plain text to cypher")
	flag.StringVar(&cypher, "c", "sha256", "the cypher to use;\nsha224, sha256, sha384, sha512, sha512_224, sha512_256, md5")
	flag.Usage = usage
}

func usage() {
	fmt.Printf("%s\n\n", os.Args[0])

	fmt.Println("Usage:")
	fmt.Printf("  hashy -t hashme -c sha256\n")
	fmt.Printf("  hashy -t hashme -c sha512_224\n\n")

	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	cliArgs := flag.Args()
	stdin, _ := os.Stdin.Stat()
	var t string

	// determine where we are getting our plaintext from
	if text != "" {
		t = text
	} else if (stdin.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		t = scanner.Text()
	} else if len(cliArgs) > 0 {
		t = cliArgs[0]
	}

	c := config{
		plainText: t,
		algorithm: cypher,
	}

	if err := run(c, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// using io.writer here makes this function easily testable.
// in our test, we can use a bytes buffer to capture output, instead
// of sending to the console
func run(c config, w io.Writer) error {

	var operation hashy.OpFunc

	algorithm := c.algorithm
	text := c.plainText

	switch strings.ToLower(algorithm) {
	case "sha224":
		operation = hashy.GenerateSHA224
	case "sha256":
		operation = hashy.GenerateSHA256
	case "sha384":
		operation = hashy.GenerateSHA384
	case "sha512":
		operation = hashy.GenerateSHA512
	case "sha512_224":
		operation = hashy.GenerateSHA512224
	case "sha512_256":
		operation = hashy.GenerateSHA512256
	case "md5":
		operation = hashy.GenerateMD5
	default:
		return fmt.Errorf("%w: %s", hashy.ErrInvalidOperation, algorithm)
	}

	_, err := fmt.Fprintln(w, operation(text))
	return err
}
