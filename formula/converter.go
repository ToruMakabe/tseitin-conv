package formula

import (
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

	fl := convImply(p)
	return fl, nil

}

// convImplyは含意を再帰的に変換する.
func convImply(e /* expression */ Expression) string {
	switch e.(type) {
	case BinOpExpr:
		left := convImply(e.(BinOpExpr).Left)
		right := convImply(e.(BinOpExpr).Right)
		if string(rune(e.(BinOpExpr).Operator)) == ">" {
			return "(" + "~" + left + "|" + right + ")"
		}
		return left + string(rune(e.(BinOpExpr).Operator)) + right
	case NotOpExpr:
		return string(rune(e.(NotOpExpr).Operator)) + convImply(e.(NotOpExpr).Right)
	case Literal:
		return e.(Literal).Literal
	default:
		return ""
	}
}

// ConvNeg は否定を変換する.
func ConvNeg(f /* formula */ string) (string, error) {
	r := strings.NewReader(f)
	// goyaccで構文木を作成する.
	p, err := Parse(r)
	if err != nil {
		return "", err
	}

	fl := convNeg(p)
	return fl, nil

}

// convImplyは含意を再帰的に変換する.
func convNeg(e /* expression */ Expression) string {
	switch e.(type) {
	case BinOpExpr:
		left := convNeg(e.(BinOpExpr).Left)
		right := convNeg(e.(BinOpExpr).Right)
		return left + string(rune(e.(BinOpExpr).Operator)) + right
	case NotOpExpr:
		return string(rune(e.(NotOpExpr).Operator)) + convNeg(e.(NotOpExpr).Right)
	case Literal:
		return e.(Literal).Literal
	default:
		return ""
	}
}
