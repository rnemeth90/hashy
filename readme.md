# hashy [![build-release-binary](https://github.com/rnemeth90/hashy/actions/workflows/build.yaml/badge.svg)](https://github.com/rnemeth90/hashy/actions/workflows/build.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/rnemeth90/hashy/)](https://goreportcard.com/report/github.com/rnemeth90/hashy/)
## Description
hashy is a tool for generating cryptographic hashes of strings.

## Getting Started
Hashy is simple. Pass it a string from stdin, as an argument to the program itself, or using the `-t` flag. You can specify the hashing algorithm using the `-c` flag.

### Dependencies
* to build yourself, you must have Go v1.14+ installed

### Installing

Download the latest release [here](https://github.com/rnemeth90/hashy/releases)

### Executing program
```
ryan:hashy/  |main ✓|$ hashy -h
hashy

Usage:
  hashy -t hashme -c sha256
  hashy -t hashme -c sha512_224

Options:
  -c string
        the cypher to use;
        sha224, sha256, sha384, sha512, sha512_224, sha512_256, md5 (default "sha256")
  -t string
        the plain text to cypher
```

### Supported Algorithms
sha224, sha256, sha384, sha512, sha512_224, sha512_256, md5. More to come.

## Help
If you need help, submit an issue

## To Do
- [x] Finish readme
- [x] reading from stdin doesn't work, fix it
- [x] add hashing algorithms and functions (md5)
- [x] add test for run() func
- [x] add docopt style help
- [x] input should be validated and converted to tolower()
- [x] Ensure builds are working

## Version History
* 1.0.4
    * More tests
* 1.0.3
    * Tests
* 1.0.2
    * Fix issue with stdin
* 1.0.1
    * Initial Release

## License
This project is licensed under the MIT License - see the LICENSE.md file for details
