package formula

import (
	"strconv"
	"strings"
)

type tseitinConverter struct {
	fv string
	fc int
}

func newTseitinConverter() *tseitinConverter {
	return &tseitinConverter{fv: "Z", fc: 0}
}

func (tc *tseitinConverter) IncFv() string {
	tc.fc++
	return tc.fv + strconv.Itoa(tc.fc)
}

func (tc *tseitinConverter) Conv(e /* expression */ Expression, pop /* parent op */ string) (string, string) {
	var (
		rFormula string
	)
	switch e.(type) {
	case BinOpExpr:
		op := string(rune(e.(BinOpExpr).Operator))
		left, lFv := tc.Conv(e.(BinOpExpr).Left, op)
		right, rFv := tc.Conv(e.(BinOpExpr).Right, op)
		fv := tc.IncFv()
		if op == "&" {
			if lFv == "lit" {
				rFormula = "(~" + fv + "|" + left + ")&(~" + fv + "|" + right + ")"
				return rFormula, fv
			}
			rFormula = left + "&" + right
			return rFormula, fv
		}
		if op == "|" {
			if lFv == "lit" {
				rFormula = "(~" + fv + "|" + left + "|" + right + ")"
				return rFormula, fv
			}
			rFormula = "(~" + fv + "|" + left + "|" + right + ")"
		}
		return rFormula, fv
	case NotOpExpr:
		op := string(rune(e.(BinOpExpr).Operator))
		right, _ := tc.Conv(e.(NotOpExpr).Right, op)
		return right, ""
	case Literal:
		rFormula = e.(Literal).Literal
		if pop == "~" {
			rFormula = "~" + rFormula
		}
		return rFormula, "lit"
	default:
		return "", ""
	}
	return "", ""

}

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
func convImply(e /* expression */ Expression, pop /* parent op */ string) string {
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

// ConvNeg は否定を変数へ寄せ,二重否定を削除する.
func ConvNeg(f /* formula */ string) (string, error) {
	r := strings.NewReader(f)
	// goyaccで構文木を作成する.
	p, err := Parse(r)
	if err != nil {
		return "", err
	}

	fl := convNeg(p, "", 0)
	return fl, nil

}

// convNeg は構文木にある否定を,ドモルガンの法則に従って再帰的に変数へ寄せる. また,二重否定を削除する.
func convNeg(e /* expression */ Expression, pop /* parent op */ string, nc /* negation counter */ int) string {
	var rFormula string
	switch e.(type) {
	case BinOpExpr:
		op := string(rune(e.(BinOpExpr).Operator))
		left := convNeg(e.(BinOpExpr).Left, op, nc)
		right := convNeg(e.(BinOpExpr).Right, op, nc)
		// 二重否定の場合は,&と|の変換を行わない.
		if nc%2 != 0 {
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
		nc++
		right := convNeg(e.(NotOpExpr).Right, op, nc)
		return right
	case Literal:
		rFormula = e.(Literal).Literal
		// 二重否定の場合は否定記号を追加しない.
		if nc%2 != 0 {
			rFormula = "~" + rFormula
		}
		return rFormula
	default:
		return ""
	}
}

// ConvTseitin はxxxする.
func ConvTseitin(f /* formula */ string) (string, error) {
	r := strings.NewReader(f)
	// goyaccで構文木を作成する.
	p, err := Parse(r)
	if err != nil {
		return "", err
	}

	tc := newTseitinConverter()
	fl, _ := tc.Conv(p, "")
	return fl, nil

}
