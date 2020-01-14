package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"bits.chrsm.org/x/lint"
	yaml "gopkg.in/yaml.v3"
)

var (
	srcpath  string
	rulepath string
)

func main() {
	flag.StringVar(&srcpath, "path", "", "path to parse .go files")
	flag.StringVar(&rulepath, "rules", ".xprlint.yml", "path to yaml config")
	flag.Parse()

	if srcpath == "" {
		log.Fatal("provide -path")
	}

	var rules map[string]*lint.Rule
	fp, err := os.Open(rulepath)
	if err != nil {
		log.Fatalf("couldn't open config(%s): %s", rulepath, err)
	}

	if err := yaml.NewDecoder(fp).Decode(&rules); err != nil {
		log.Fatalf("couldn't decode yaml: %s", err)
	}

	buf := new(bytes.Buffer)
	fail := false
	filepath.Walk(srcpath, func(fpath string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return err
		}

		if path.Ext(fi.Name()) != ".go" {
			return err
		}

		buf.Reset()
		fp, err := os.Open(fpath)
		if err != nil {
			return err
		}

		_, err = buf.ReadFrom(fp)
		if err != nil {
			return err
		}

		src := buf.String()
		for i := range rules {
			errs := lint.Walk(rules[i], src)
			for i := range errs {
				fail = true

				fmt.Printf("%s: %s\n", fpath, errs[i])
			}
		}

		return nil
	})

	if fail {
		os.Exit(1)
	}
}
