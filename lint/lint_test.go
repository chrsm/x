package lint

import (
	"testing"

	yaml "gopkg.in/yaml.v3"
)

/*
ideally, we'd be able to write something like dsadsdasd

rule:
	type: import

	# eg, `/pkg/(1: meme)/v(2: 12)`
	regexp: "\/pkg\/(?P<meme>[a-z]+)/v(?P<version>[0-9]+)"

	# a property of the "ast.ImportSpec" struct.
	require_name:
		template: "pkg{{ "${meme}" | upper }}v${version}"

in code, this would be:
&Rule{
	Name: "rule",
	Type: "import",
	Regexp: regexp.MustCompile(`\/pkg\/(?P<meme>[a-z]+)\/v(?P<version>[0-9]+)`),
	Require: importRequirement{
		Name: stringRequirement{
			Template: "pkg{{ "${meme}" | upper }}v${version}",
		},
	},
}
*/
func TestErr(t *testing.T) {
	yml := `
type: import
require:
  x: abcd
  # eg, /pkg/(1: meme)/v(2: 12)\
  regexp: \/pkg\/(?P<meme>[a-z]+)\/v(?P<version>[0-9]+)

  # a property of the "ast.ImportSpec" struct.
  name:
    template: req{{ "${meme}" | upper }}v${version}
`

	r := &Rule{}
	if err := yaml.Unmarshal([]byte(yml), r); err != nil {
		t.Fatalf("error while unmarshalling: %s", err)
	}

	src := `
package main

import (
	reqLULv1 "xyz.org/pkg/lol/v1"
)

`

	defer func() {
		if x := recover(); x != nil {
			t.Logf("recovered from panic: %v", x)
		}
	}()

	errs := Walk(r, src)
	if len(errs) == 0 {
		t.Fatalf("expected >0 errors")
	}
}

func TestBasicOK(t *testing.T) {
	yml := `
type: import
require:
  x: abcd
  # eg, /pkg/(1: meme)/v(2: 12)\
  regexp: \/pkg\/(?P<meme>[a-z]+)\/v(?P<version>[0-9]+)

  # a property of the "ast.ImportSpec" struct.
  name:
    template: req{{ "${meme}" | upper }}v${version}
`

	r := &Rule{}
	if err := yaml.Unmarshal([]byte(yml), r); err != nil {
		t.Fatalf("error while unmarshalling: %s", err)
	}

	src := `
package main

import (
	reqLOLv1 "xyz.org/pkg/lol/v1"

	reqLOLv2 "xyz.org/pkg/lol/v2"
)

`

	defer func() {
		if x := recover(); x != nil {
			t.Logf("recovered from panic: %v", x)
		}
	}()

	errs := Walk(r, src)
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors, got %d", len(errs))
	}
}

func TestBasicMixed(t *testing.T) {
	yml := `
type: import
require:
  x: abcd
  # eg, /pkg/(1: meme)/v(2: 12)\
  regexp: \/pkg\/(?P<meme>[a-z]+)\/v(?P<version>[0-9]+)

  # a property of the "ast.ImportSpec" struct.
  name:
    template: req{{ "${meme}" | upper }}v${version}
`

	r := &Rule{}
	if err := yaml.Unmarshal([]byte(yml), r); err != nil {
		t.Fatalf("error while unmarshalling: %s", err)
	}

	src := `
package main

import (
	reqLOLv1 "xyz.org/pkg/lol/v1"
	reqLULv3 "xyz.org/pkg/notlul/v5"
	reqLOLv2 "xyz.org/pkg/lol/v2"
)

`

	defer func() {
		if x := recover(); x != nil {
			t.Logf("recovered from panic: %v", x)
		}
	}()

	errs := Walk(r, src)
	if len(errs) != 1 {
		t.Fatalf("expected 1 errors, got %d", len(errs))
	}

	err := errs[0]
	if err.Error() != "expected import(\"xyz.org/pkg/notlul/v5\") to be named(reqNOTLULv5), but it was(reqLULv3)" {
		t.Fatalf("error message incorrect, got %s", err)
	}
}
