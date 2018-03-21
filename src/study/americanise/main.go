package main

import (
	"bufio"
	"errors"
	"extensions/strext"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	dictFilename, inFilename, outFilename, err := filenamesFromCommandline()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var replacer func(string) string
	if replacer, err = generateReplacer(dictFilename); err != nil {
		log.Fatal(err)
	}

	var inFile *os.File
	if inFile, err = os.Open(inFilename); err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	var outFile *os.File
	if outFile, err = os.Create(outFilename); err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err = americanise(inFile, outFile, replacer); err != nil {
		log.Fatal(err)
	}
}

func filenamesFromCommandline() (string, string, string, error) {
	l := len(os.Args)
	if l < 4 {
		return "", "", "", fmt.Errorf("usage: %s dictFile.txt inFile.txt outFile.txt", filepath.Base(os.Args[0]))
	}

	if strext.IsWhitespace(os.Args[1]) || strext.IsWhitespace(os.Args[2]) || strext.IsWhitespace(os.Args[3]) {
		return "", "", "", errors.New("input file names are empty")
	}

	return os.Args[1], os.Args[2], os.Args[3], nil
}

func generateReplacer(dictFilename string) (func(s string) string, error) {
	bytes, err := ioutil.ReadFile(dictFilename)
	if err != nil {
		return nil, err
	}

	dictText := string(bytes)
	dict := make(map[string]string)
	lines := strings.Split(dictText, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			dict[fields[0]] = fields[1]
		}
	}

	return func(s string) string {
		if v, found := dict[s]; found {
			return v
		}
		return s
	}, nil
}

func americanise(inFile io.Reader, outFile io.Writer, replacer func(string) string) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	wordRx := regexp.MustCompile("[A-Za-z]+")
	for {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return err
		}

		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}
