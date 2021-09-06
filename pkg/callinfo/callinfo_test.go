package callinfo_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/youthlin/logs/pkg/callinfo"
)

func TestCallInfo_ShortFileName(t *testing.T) {
	type fields struct {
		PkgName   string
		FileName  string
		ShortFile string
		FuncName  string
		Line      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"not-empty", fields{ShortFile: "abc"}, "abc"},
		{"lazy", fields{FileName: "abc"}, "abc"},
		{"lazy-split", fields{FileName: "dir/filename"}, "filename"},
		{"root", fields{FileName: "/filename"}, "filename"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &callinfo.CallInfo{
				PkgName:   tt.fields.PkgName,
				FileName:  tt.fields.FileName,
				ShortFile: tt.fields.ShortFile,
				FuncName:  tt.fields.FuncName,
				Line:      tt.fields.Line,
			}
			if got := c.ShortFileName(); got != tt.want {
				t.Errorf("CallInfo.ShortFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	ci := callinfo.Get()
	dir, _ := filepath.Abs(".")
	want := &callinfo.CallInfo{
		PkgName:   "github.com/youthlin/logs/pkg/callinfo_test",
		FileName:  filepath.Join(dir, "callinfo_test.go"),
		ShortFile: "", // lazy get
		FuncName:  "TestGet",
		Line:      46,
	}
	if !reflect.DeepEqual(ci, want) {
		t.Errorf("callinfo.Get() got = %#v want: %#v", ci, want)
	}
}
