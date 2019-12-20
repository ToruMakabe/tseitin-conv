package formula

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestConvNNF(t *testing.T) {

	var (
		f   string
		r   string
		err error
	)

	// Convert A to A
	f = "A"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "A" {
		t.Errorf("(Convert A to A: Failed. The result is %v", r)
	}

	// Convert A&A to (A&A)
	f = "A&A"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(A&A)" {
		t.Errorf("(Convert A to A: Failed. The result is %v", r)
	}

	// Convert A&B&C to (A&B&C)
	f = "A&B&C"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(A&B&C)" {
		t.Errorf("(Convert A&B&C to A&B&C: Failed. The result is %v", r)
	}

	// Convert (A&B&C|D&E&F) to ((A&B&C)|(D&E&F))
	f = "(A&B&C)|(D&E&F)"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "((A&B&C)|(D&E&F))" {
		t.Errorf("(Convert (A&B&C|D&E&F) to ((A&B&C)|(D&E&F)): Failed. The result is %v", r)
	}

	// Convert A>B to (~A|B)
	f = "A>B"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|B)" {
		t.Errorf("(Convert A>B to (~A|B)): Failed. The result is %v", r)
	}

	// Convert A>(A>B) to (~A|(~A|B))
	f = "A>(A>B)"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|~A|B)" {
		t.Errorf("(Convert A>(A>B) to (~A|~A|B): Failed. The result is %v", r)
	}

	// Convert ~(A&B&C) to (~A|~B|~C)
	f = "~(A&B&C)"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|~B|~C)" {
		t.Errorf("(Convert ~(A&B&C) to (~A|~B|~C): Failed. The result is %v", r)
	}

	// Convert ~(~A&B&C) to (A|~B|~C)
	f = "~(~A&B&C)"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(A|~B|~C)" {
		t.Errorf("(Convert ~(~A&B&C) to (A|~B|~C): Failed. The result is %v", r)
	}

	// Convert ~(~A&~(~B&C)) to (A|~B|~C)
	f = "~(~A&~(~B&C))"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(A|(~B&C))" {
		t.Errorf("(Convert ~(~A&~(~B&C)) to (A|(~B&C)): Failed. The result is %v", r)
	}

	// Convert ~(A&(B|C)) to (~A|(~B&~C))
	f = "~(A&(B|C))"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|(~B&~C))" {
		t.Errorf("(Convert ~(A&(B|C)) to (~A|(~B&~C)): Failed. The result is %v", r)
	}

	// Convert ~(A&B&C)|~(A&B&C) to ((~A|~B|~C)|(~A|~B|~C))
	f = "~(A&B&C)|~(A&B&C)"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "((~A|~B|~C)|(~A|~B|~C))" {
		t.Errorf("(Convert ~(A&B&C)|~(A&B&C) to ((~A|~B|~C)|(~A|~B|~C)): Failed. The result is %v", r)
	}

	// Convert ~~A to A
	f = "~~A"

	r, err = ConvNNF(f)
	if err != nil {
		printError(err)
	}

	if r != "A" {
		t.Errorf("(Convert ~~A to A: Failed. The result is %v", r)
	}

}

func TestConvTseitin(t *testing.T) {

	var (
		f   string
		r   []string
		rs  [][]string
		err error
	)

	// Convert (A&B&C)|(D&E&F) to (~x1|A)&(~x1|B)&(~x2|x1)&(~x2|C)&(~x3|D)&(~x3|E)&(~x4|x3)&(~x4|F)&(~x5|x2|x4)&(x5)
	f = "(A&B&C)|(D&E&F)"

	rs, err = ConvTseitin(f)
	if err != nil {
		printError(err)
	}

	r = nil
	for _, i := range rs {
		r = append(r, "("+strings.Join(i, "|")+")")
	}

	if strings.Join(r, "&") != "(~x1|A)&(~x1|B)&(~x2|x1)&(~x2|C)&(~x3|D)&(~x3|E)&(~x4|x3)&(~x4|F)&(~x5|x2|x4)&(x5)" {
		t.Errorf("(Convert (A&B&C)|(D&E&F) to (~x1|A)&(~x1|B)&(~x2|x1)&(~x2|C)&(~x3|D)&(~x3|E)&(~x4|x3)&(~x4|F)&(~x5|x2|x4)&(x5): Failed. The result is %v", strings.Join(r, "&"))
	}

	// Convert A|(B&C&(D|E)) to (~x1|B)&(~x1|C)&(~x2|D|E)&(~x3|x1)&(~x3|x2)&(~x4|A|x3)&(x4)
	f = "A|(B&C&(D|E))"

	rs, err = ConvTseitin(f)
	if err != nil {
		printError(err)
	}

	r = nil
	for _, i := range rs {
		r = append(r, "("+strings.Join(i, "|")+")")
	}

	if strings.Join(r, "&") != "(~x1|B)&(~x1|C)&(~x2|D|E)&(~x3|x1)&(~x3|x2)&(~x4|A|x3)&(x4)" {
		t.Errorf("(Convert A|(B&C&(D|E)) to (~x1|B)&(~x1|C)&(~x2|D|E)&(~x3|x1)&(~x3|x2)&(~x4|A|x3)&(x4): Failed. The result is %v", strings.Join(r, "&"))
	}

	// Convert ((~A&~B)|~C)|(~D) to (~x1|~A)&(~x1|~B)&(~x2|x1|~C)&(~x3|x2|~D)&(x3)
	f = "((~A&~B)|~C)|(~D)"

	rs, err = ConvTseitin(f)
	if err != nil {
		printError(err)
	}

	r = nil
	for _, i := range rs {
		r = append(r, "("+strings.Join(i, "|")+")")
	}

	if strings.Join(r, "&") != "(~x1|~A)&(~x1|~B)&(~x2|x1|~C)&(~x3|x2|~D)&(x3)" {
		t.Errorf("(Convert ((~A&~B)|~C)|(~D) to (~x1|~A)&(~x1|~B)&(~x2|x1|~C)&(~x3|x2|~D)&(x3): Failed. The result is %v", strings.Join(r, "&"))
	}

}

// printErrorはエラーメッセージ出力を統一する.
func printError(err /* error */ error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}
