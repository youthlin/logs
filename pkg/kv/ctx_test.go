package kv_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/youthlin/logs/pkg/kv"
)

func TestGet(t *testing.T) {
	bg := context.Background()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantKvs []interface{}
	}{
		{"empty", args{bg}, nil},
		{"Add1", args{kv.Add(bg, "one")}, nil},
		{"Add2", args{kv.Add(bg, "one", "two")}, []interface{}{"one", "two"}},
		{"Add3", args{kv.Add(bg, "one", "two", "three")}, []interface{}{"one", "two"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKvs := kv.Get(tt.args.ctx); !reflect.DeepEqual(gotKvs, tt.wantKvs) {
				t.Errorf("Get() = %v, want %v", gotKvs, tt.wantKvs)
			}
		})
	}
}
