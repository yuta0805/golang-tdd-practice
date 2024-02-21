package arrays

import (
	"reflect"
	"testing"
)

// 全てのテストケースを考え実行することは困難なので、テストケースの取捨選択が重要
// go test -coverd でカバレッジを確認することができる.
func TestSum(t *testing.T) {
	t.Run("sum of 5 elements", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		sum := Sum(arr)
		expected := 15
		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
	
	t.Run("collection of any size", func(t *testing.T) {
		arr := []int{1, 2, 3}
		sum := Sum(arr)
		expected := 6
		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
} 


// slice統合演算子を使うことはできない。なのでreflect.DeepEqualを使う
// reflect.DeepEqualは、2つの変数が同じであるかを確認するために便利
// reflect.DeepEqualは型安全ではない。なのでコンパイルが通ってしまう。
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
func TestSumAllTails(t *testing.T) {
	// アサーションにerrorハンドリングを委譲する
    checkSums := func(t *testing.T, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("make the sums of tails of", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
        checkSums(t, got, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}
        checkSums(t, got, want)
    })

}
