package main

import (
	f "fmt"
)

func main() {
	/*
	   要素数に対応した初期スペースを指定
	   スライスにおける容量とは意味が異なり、ランタイムがメモリ領域を
	   確保するための一種の「ヒント」として機能する.
	   要素数が膨大なマップを構築する場合であれば、パフォーマンス向上を期待できる.
	*/
	m := make(map[int]string, 100)
}
