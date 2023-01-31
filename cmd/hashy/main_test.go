package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/rnemeth90/hashy"
)

var (
	binName string = "hashy"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	resultCode := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)

	os.Exit(resultCode)
}

func TestRun(t *testing.T) {
	var b bytes.Buffer

	testCases := []struct {
		name string
		want string
		c    config
		b    *bytes.Buffer
	}{
		{"test_GenerateSHA224", "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809", config{plainText: "test", algorithm: "sha224"}, &b},
		{"test_GenerateSHA256", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", config{plainText: "test", algorithm: "sha256"}, &b},
		{"test_GenerateSHA384", "768412320f7b0aa5812fce428dc4706b3cae50e02a64caa16a782249bfe8efc4b7ef1ccb126255d196047dfedf17a0a9", config{plainText: "test", algorithm: "sha384"}, &b},
		{"test_GenerateSHA512", "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff", config{plainText: "test", algorithm: "sha512"}, &b},
		{"test_GenerateSHA512224", "06001bf08dfb17d2b54925116823be230e98b5c6c278303bc4909a8c", config{plainText: "test", algorithm: "sha512_224"}, &b},
		{"test_GenerateSHA512256", "3d37fe58435e0d87323dee4a2c1b339ef954de63716ee79f5747f94d974f913f", config{plainText: "test", algorithm: "sha512_256"}, &b},
		{"test_GenerateMD5", "098f6bcd4621d373cade4e832627b4f6", config{plainText: "test", algorithm: "md5"}, &b},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if err := run(config{test.c.plainText, test.c.algorithm}, test.b); err != nil {
				t.Fatalf("%v: %v", hashy.ErrRunTimeException, err)
			}

			result := b.String()

			if strings.TrimSuffix(result, "\n\n") != test.want {
				t.Errorf("got %s, wanted %s", result, test.want)
			}
			b.Reset()
		})
	}
}
