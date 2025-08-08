package spec

import (
	"bytes"
	"testing"
)

func TestGenerateHTMLSpec(t *testing.T) {
	htmlDoc := `
<html><head></head><body><script></script><p></p><h2 id="skipme"></h2><h2 id=""></h2></body></html>
`

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := GenerateHTMLSpec(tt.args.url, w)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHTMLSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("GenerateHTMLSpec() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
