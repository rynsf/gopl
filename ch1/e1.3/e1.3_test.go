package main

import (
	"testing"
)

var args = []string{"This", "is", "a", "test", "string"}

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo(args)
	}
}

func BenchmarkEchoWithStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithStringJoin(args)
	}
}
