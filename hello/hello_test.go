package main

import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t testing.TB, got, want string) {
        t.Helper() // テスト失敗時に行番号を報告するために必要
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }
    // test内容をグループ化する
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello world'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, world"

        assertCorrectMessage(t, got, want)
    })
}
