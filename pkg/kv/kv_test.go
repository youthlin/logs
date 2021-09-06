package kv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrimOdd(t *testing.T) {
	type args struct {
		kvs []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"empty", args{nil}, nil},
		{"len-1", args{[]interface{}{1}}, nil},
		{"len-2", args{[]interface{}{1, 2}}, []interface{}{1, 2}},
		{"len-3", args{[]interface{}{1, 2, 3}}, []interface{}{1, 2}},
		{"len-4", args{[]interface{}{1, 2, 3, 4}}, []interface{}{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimOdd(tt.args.kvs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type S struct {
		A int
	}
	type args struct {
		kvs []interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"empty", args{nil}, [][]string{}},
		{"len=1", args{[]interface{}{1}}, [][]string{}},
		{"len=2", args{[]interface{}{1, 2}}, [][]string{{"1", "2"}}},
		{"len=3", args{[]interface{}{1, 2, "a"}}, [][]string{{"1", "2"}}},
		{"len=4", args{[]interface{}{1, 2, "s", S{100}}}, [][]string{{"1", "2"}, {"s", `{"A":100}`}}},
		{"len=6", args{[]interface{}{false, true, 1.1, []byte(`ok`), "err", fmt.Errorf("error")}},
			[][]string{{"false", "true"}, {"1.100", "ok"}, {"err", "error"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.kvs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToString() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
