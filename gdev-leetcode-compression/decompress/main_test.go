package main

import "testing"

var table = map[string]string {
	"3[abc]4[ab]c": "abcabcabcababababc",
	"10[a]": "aaaaaaaaaa",
	"2[3[a]b]": "aaabaaab",


	"a" : "a",
	"3[a]" : "aaa",
	"1[a]c" : "ac",
	"ab2[ac]xy" : "abacacxy",
	"a2[2[b]2[c]]" : "abbccbbcc",
}

func TestDecompressRecursive(t *testing.T){
	for in, out := range table{
		w, _ := decompressRecursive([]rune(in))
		if string(w) != out {
			t.Errorf("in '%s' expected '%s' got '%s'", in, out, string(w))
		}
	}
}

func TestDecompressIterative(t *testing.T){
	for in, out := range table{
		w := decompressIterative([]rune(in))
		if string(w) != out {
			t.Errorf("in '%s' expected '%s' got '%s'", in, out, string(w))
		}
	}
}
