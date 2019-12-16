package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ToruMakabe/tseitin-conv/formula"
)

const inputFormatMsg = "Please input a propotitinal formula to convert to CNF.\nNagation: ~, And: &, Or: |, Imply: >\nYou cannot use x(n) as propotional variable due to reserved word for fresh variable of Tseitin convertion.\nSample: A|(B&C&(D|E))\n"

// convertは実質的な主処理である.
func convert() int {

	var (
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
	if err := scanner.Err(); err != nil {
		printError(fmt.Errorf("scanner error"))
		return 1
	}

	// 変換に要した時間を計測するため, 開始時間を取得する.
	st := time.Now()

	// 否定標準形(NNF)への変換を行う.
	nnf, err = formula.ConvNNF(f)
	if err != nil {
		fmt.Println()
		printError(err)
		fmt.Println()
		fmt.Println(inputFormatMsg)
		return 1
	}

	// Tseitin変換を行う.
	cnfm, err = formula.ConvTseitin(nnf)
	if err != nil {
		fmt.Println()
		printError(err)
		fmt.Println()
		fmt.Println(inputFormatMsg)
		return 1
	}

	// Tseitin変換の結果が[][]stringのスライスで得られるため, 内側の選言部分を文字列化する.
	for _, i := range cnfm {
		cnf = append(cnf, "("+strings.Join(i, "|")+")")
	}

	// CNFを表示する.
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
