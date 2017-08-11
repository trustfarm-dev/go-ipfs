package gc

import (
	"math/rand"
	"testing"
	"unsafe"
)

func TestTrielementSize(t *testing.T) {
	t.Log("trielement size is:", unsafe.Sizeof(trielement{}))
}

func BenchmarkMapInserts(b *testing.B) {
	b.N = 10e6
	keys := make([]string, b.N)
	buf := make([]byte, 64)
	for i := 0; i < b.N; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
		keys[i] = string(buf)
	}

	set := make(map[string]trielement)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set[keys[i]] = trielement{1}
	}
}
func BenchmarkMapUpdate(b *testing.B) {
	b.N = 10e6
	keys := make([]string, b.N)
	buf := make([]byte, 64)
	for i := 0; i < b.N; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
		keys[i] = string(buf)
	}

	set := make(map[string]trielement)
	for i := 0; i < b.N; i++ {
		set[keys[i]] = trielement{1}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set[keys[i]] = trielement{2}
	}
}
