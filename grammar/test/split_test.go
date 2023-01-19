package test

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			var got = Split(test.input, test.sep)
			var want = test.want
			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected:%#v, got:%#v", want, got)
			}
		})

	}

}
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

//const n = 30

// func BenchmarkFib(b *testing.B) {
//
//		for i := 0; i < b.N; i++ {
//			Fib(n)
//		}
//	}
//
//	func BenchmarkFib2(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			 Fib2(n)
//		}
//	}
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib2(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
