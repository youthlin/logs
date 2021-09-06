package trie_test

import (
	"testing"

	"github.com/youthlin/logs/pkg/trie"
)

func TestTire(t *testing.T) {
	tree := trie.NewTire(10)
	tree.Insert("github.com", 20)
	tree.Insert("github.com/youthlin", 30)
	tree.Insert("github.com/youthlin/中文", 35)
	tree.Insert("github.com/youthlin/logs", 40)
	t.Logf("Dump trie: %v", tree.Dump())
	for _, tCase := range []struct {
		path string
		want int
	}{
		{"", 10},
		{"abc", 10},
		{"中文", 10},
		{"github", 10},
		{"github.com", 20},
		{"github.com/abc", 20},
		{"github.com/youthlin", 30},
		{"github.com/youthlin/t", 30},
		{"github.com/youthlin/中文", 35},
		{"github.com/youthlin/中文子串", 35},
		{"github.com/youthlin/中文/子串", 35},
		{"github.com/youthlin/logs", 40},
		{"github.com/youthlin/logs/util", 40},
	} {
		if got := tree.Search(tCase.path).(int); got != tCase.want {
			t.Errorf("search %q fail. got=%v, want=%v", tCase.path, got, tCase.want)
		}
	}
}
