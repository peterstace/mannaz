package chess

import "testing"

func TestFENInitialPosition(t *testing.T) {
	pos := InitialPosition()
	got := pos.FEN()
	const want = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	expectEq(t, got, want)
}

func expectEq[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Logf("want: %v", want)
		t.Logf("got:  %v", got)
		t.Log("mismatch")
	}
}
