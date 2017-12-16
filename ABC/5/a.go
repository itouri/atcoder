package main

// OK! 3'54

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	/* 初期化開始 */
	sc.Split(bufio.ScanWords)
	x := nextInt()
	y := nextInt()
	/* 初期化終了 */
	fmt.Printf("%d\n", y/x)
}
