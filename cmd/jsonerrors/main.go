package main

import (
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	t "text/template"

	"github.com/icio/jsonerrors"

	"github.com/davecgh/go-spew/spew"
)

type (
	Package = jsonerrors.Package
	Decoder = jsonerrors.Decoder
)

func main() {
	out := os.Stdout
	render := RenderMarkdownResults

	tests := Tests()
	packages := jsonerrors.Packages()
	results := Suite(tests, packages)

	if err := render(out, results); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func Suite(tests []*Test, pkgs []*Package) []*TestResults {
	tr := make([]*TestResults, 0, len(tests))
	for _, test := range tests {
		pr := make([]*PackageResults, 0, len(pkgs))
		for _, pkg := range pkgs {
			dr := make([]*DecoderResult, 0, len(pkg.Decoders))
			for _, dec := range pkg.Decoders {
				b := []byte(test.Body)
				v := test.Target()
				e := dec.Unmarshal(b, v)
				dr = append(dr, &DecoderResult{test, pkg, dec, v, e})
			}
			pr = append(pr, &PackageResults{test, pkg, dr})
		}
		tr = append(tr, &TestResults{test, pr})
	}
	return tr
}

type TestResults struct {
	*Test
	Results []*PackageResults
}

type PackageResults struct {
	*Test
	*Package
	Results []*DecoderResult
}

type DecoderResult struct {
	*Test
	*Package
	*Decoder

	V   interface{}
	Err error
}

type Test struct {
	Desc  string
	Body  string
	Proto interface{}
}

func (t Test) Target() interface{} {
	return reflect.New(reflect.TypeOf(t.Proto)).Interface()
}

func Tests() []*Test {
	return []*Test{
		{
			Desc: "Unknown object property.",
			Body: `{"cat": "Sammy", "dog": "Loki"}`,
			Proto: struct {
				Cat string `json:"cat"`
			}{},
		},
		{
			Desc: "Using a string instead of an object.",
			Body: `{"cat": "sammy"}`,
			Proto: struct {
				Cat struct{} `json:"cat"`
			}{},
		},
	}
}

const md = `
	# Go JSON Parsers: Errors

	Go Version: {{ .GoVersion }}

	{{ range $tr := .Results }}{{ $test := $tr.Test }}
	## Test: {{ $test.Desc }}
	
	Body:
	{{ $test.Body | code "json" }}
	Target:
	{{ $test.Proto | printf "%T" | code "go" }}

	{{ range $pr := $tr.Results }}{{ $pkg := $pr.Package }}
	* **pkg [{{ $pkg.Name }}](https://godoc.org/{{ $pkg.Name }})** {{ if $pkg.Desc }} ({{ $pkg.Desc }}){{ end }}

	{{ range $dr := $pr.Results }}{{ $dec := $dr.Decoder }}
		* **{{ $dec.Desc }}**

		  Err: ` + "`{{ $dr.Err }}`" + `

			{{ $dr.Err | spew | code "go" | details | indent 2 }}

	{{ end }}
	{{ end }}
	{{ end }}
`

func RenderMarkdownResults(w io.Writer, results []*TestResults) error {
	s := &spew.ConfigState{
		Indent:         "  ",
		DisableMethods: true,
		SpewKeys:       true,
	}
	tpl, err := t.New("jsonerrors-text").Funcs(t.FuncMap{
		"spew": s.Sdump,
		"code": func(lang, c string) string {
			return "```" + lang + "\n" + strings.TrimRight(c, "\n") + "\n```"
		},
		"indent": func(d int, b string) string {
			return strings.Replace(b, "\n", "\n"+strings.Repeat("\t", d), -1)
		},
		"details": func(c string) string {
			return "<details>\n\n" + c + "\n\n</details>"
		},
	}).Parse(unindent(md, "\t", "\n"))
	if err != nil {
		return err
	}
	return tpl.Execute(w, struct {
		Results   []*TestResults
		GoVersion string
	}{results, runtime.Version()})
}

// unindent identifies the leading whitespace at the beginning of a string and
// strips it from all lines in the whole block of text.
func unindent(body, ws, eol string) string {
	body = strings.TrimLeft(body, eol)
	l := len(body)
	unindented := strings.TrimLeft(body, ws)
	indent := body[:l-len(unindented)]
	return strings.Replace(unindented, eol+indent, eol, -1)
}
