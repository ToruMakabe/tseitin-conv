package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ToruMakabe/tseitin-conv/formula"
)

const inputFormatMsg = "Please input xxx"

// convertは実質的な主処理である.
func convert() int {

	var (
		r    string
		nnf  string
		cnf  []string
		cnfm [][]string
		err  error
	)

	// 命題論理式の入力を促す.
	fmt.Println(inputFormatMsg)
	fmt.Print("Formula? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	f := scanner.Text()

	// 変換に要した時間を計測するため,開始時間を取得する.
	st := time.Now()

	r, err = formula.ConvImply(f)
	if err != nil {
		printError(err)
		return 1
	}

	nnf, err = formula.ConvNeg(r)
	if err != nil {
		printError(err)
		return 1
	}

	cnfm, err = formula.ConvTseitin(nnf)
	if err != nil {
		printError(err)
		return 1
	}

	for _, i := range cnfm {
		cnf = append(cnf, "("+strings.Join(i, "|")+")")
	}

	// 変換の結果得られたCNFを表示する.
	fmt.Printf("CNF: %v \n", strings.Join(cnf, "&"))

	// 変換に要した時間を表示する.
	et := time.Now()
	fmt.Println("Time: ", et.Sub(st))

	return 0
}

// printErrorはエラーメッセージ出力を統一する.
func printError(err /* error */ error) {
	fmt.Fprintf(os.Stderr, err.Error()+"\n")
}

// mainはエントリーポイントと終了コードを返す役割のみとする.
func main() {
	os.Exit(convert())
}
