package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ToruMakabe/tseitin-conv/formula"
)

const inputFormatMsg = "Please input xxx"

// convertは実質的な主処理である.
func convert() int {

	// 命題論理式の入力を促す.
	fmt.Println(inputFormatMsg)
	fmt.Print("Formula? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// シーケントを "|-" を区切り記号として,前提と結論に分解する.
	f := scanner.Text()

	// 変換に要した時間を計測するため,開始時間を取得する.
	st := time.Now()

	r, err := formula.Conv(f)
	if err != nil {
		printError(err)
		return 1
	}

	// 変換の結果得られたCNFを表示する.
	fmt.Printf("CNF: %v \n", r)

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
