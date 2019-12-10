package formula

import (
	"fmt"
	"strings"
)

// ConvImply は含意を変換する.
func ConvImply(f /* formula */ string) (string, error) {
	r := strings.NewReader(f)
	// goyaccで構文木を作成する.
	p, err := Parse(r)
	if err != nil {
		return "", err
	}

	fl := convImply(p, "")
	return fl, nil

}

// convImply は構文木にある含意を再帰的に変換する.
func convImply(e /* expression */ Expression, pop /* childs op */ string) string {
	var rFormula string
	switch e.(type) {
	case BinOpExpr:
		op := string(rune(e.(BinOpExpr).Operator))
		left := convImply(e.(BinOpExpr).Left, op)
		right := convImply(e.(BinOpExpr).Right, op)
		if op == ">" {
			rFormula = "(" + "~" + left + "|" + right + ")"
			return rFormula
		}
		rFormula = left + op + right
		if op == pop {
			return rFormula
		}
		return "(" + rFormula + ")"
	case NotOpExpr:
		op := string(rune(e.(NotOpExpr).Operator))
		right := convImply(e.(NotOpExpr).Right, op)
		return op + right
	case Literal:
		return e.(Literal).Literal
	default:
		return ""
	}
}

// ConvNeg は否定をリテラルに寄せ、二重否定を削除する.
func ConvNeg(f /* formula */ string) (string, error) {
	r := strings.NewReader(f)
	// goyaccで構文木を作成する.
	p, err := Parse(r)
	if err != nil {
		return "", err
	}

	fl := convNeg(p, "", false)
	return fl, nil
	//	return strings.Replace(fl, "~~", "", -1), nil

}

// convNeg は構文木にある否定を再帰的にリテラルに寄せる(ドモルガンの法則).
func convNeg(e /* expression */ Expression, pop /* parent op */ string, n /* negation flag */ bool) string {
	fmt.Println(e)
	var rFormula string
	switch e.(type) {
	case BinOpExpr:
		op := string(rune(e.(BinOpExpr).Operator))
		left := convNeg(e.(BinOpExpr).Left, op, n)
		right := convNeg(e.(BinOpExpr).Right, op, n)
		if n {
			if op == "&" {
				rFormula = left + "|" + right
			}
			if op == "|" {
				rFormula = left + "&" + right
			}
		} else {
			rFormula = left + op + right
		}

		if op == pop {
			return rFormula
		}
		return "(" + rFormula + ")"
	case NotOpExpr:
		op := string(rune(e.(NotOpExpr).Operator))
		right := convNeg(e.(NotOpExpr).Right, op, true)
		return right
	case Literal:
		if n {
			rFormula = "~" + e.(Literal).Literal
			if pop == "~" {
				rFormula = "~" + rFormula
			}
			return rFormula
		}
		return e.(Literal).Literal
	default:
		return ""
	}
}
