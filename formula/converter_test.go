package formula

import (
	"fmt"
	"os"
	"testing"
)

func TestConv(t *testing.T) {

	var f string
	var r string
	var err error

	// Convert A>B to (~A|B)
	f = "A>B"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|B)" {
		t.Errorf("(Convert A>B to (~A|B)): Failed. The result is %v", r)
	}

	// Convert A>(A>B) to (~A|(~A|B))
	f = "A>(A>B)"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|(~A|B))" {
		t.Errorf("(Convert A>(A>B) to (~A|(~A|B)): Failed. The result is %v", r)
	}

	// Convert A>(A&B&C) to (~A|(A&B&C))
	f = "A>(A&B&C)"

	r, err = ConvImply(f)
	if err != nil {
		printError(err)
	}

	if r != "(~A|(A&B&C)" {
		t.Errorf("(Convert A>(A&B&C) to (~A|(A&B&C)): Failed. The result is %v", r)
	}

	/*
		// Convert ~A>(A>B) to (~~A|(~A|B))
		f = "~A>(A>B)"

		r, err = Conv(f)
		if err != nil {
			printError(err)
		}

		if r != "(~~A|(~A|B))" {
			t.Errorf("(Convert ~A>(A>B) to (~~A|(~A|B))): Failed. The result is %v", r)
		}

		// Convert ~(A>(A>B)) to (~(~A|(~A|B)))
		f = "~(A>(A>B))"

		r, err = Conv(f)
		if err != nil {
			printError(err)
		}

		if r != "~(~A|(~A|B))" {
			t.Errorf("(Convert ~A>(A>B) to (~(~A|(~A|B)))): Failed. The result is %v", r)
		}
	*/
}

// printErrorはエラーメッセージ出力を統一する.
func printError(err /* error */ error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}
