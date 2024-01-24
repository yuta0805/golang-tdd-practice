package main

import "testing"

func TestHello(t *testing.T) {
    // test内容をグループ化する
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("say 'Hello world'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, world"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
}
