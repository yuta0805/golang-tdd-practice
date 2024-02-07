package main

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// benchmarkが測れる記法がある。testing.B型を引数に取る関数を作成する。
// この関数は、go test -bench=. で実行される。
func Benchmark(b *testing.B) {
	// b.Nは、ベンチマーク関数が実行される回数を示す。
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
