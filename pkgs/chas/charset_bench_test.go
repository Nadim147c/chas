package chas

import (
	"bytes"
	"testing"
)

func writeBuf(b *testing.B, w *bytes.Buffer) {
	b.Helper()
	for range 10000 {
		w.WriteString("the quick brown fox jumps over the lazy dog 🚀\n")
		w.WriteString("no match here\n")
		w.WriteString("another line with a partial match: quick brown\n")
	}
}

func BenchmarkSearch(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	go writeBuf(b, buf)

	charset := "abcdefghijklmnopqrstuvwxyz🚀"
	b.ResetTimer()

	for b.Loop() {
		_ = Search(charset, buf, b.Output())
	}
}
