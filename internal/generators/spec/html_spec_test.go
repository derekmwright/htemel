package spec

import (
	"bytes"
	"io"
	"testing"
)

func TestGenerateHTMLSpec(t *testing.T) {
	htmlDoc := `
<html>
	<head></head>
	<body>
		<script></script>
		<p></p>
		<h2 id="skip-me"></h2>
		<h2 id="semantics"></h2><h4><p><code>tag</code></p></h4><p>Good description</p><p>I shouldn't be in output</p>
		<h2 id="parsing-should-stop"><h4><p><code>badtag</code></p></h4><p>Bad description</p></h2>
	</body>
</html>
`

	type args struct {
		rc io.ReadCloser
	}
	tests := []struct {
		name    string
		args    args
		want    *Spec
		wantErr bool
	}{
		{
			name: "basic parse",
			args: args{
				rc: io.NopCloser(bytes.NewBufferString(htmlDoc)),
			},
			want: &Spec{
				Name: "HTML",
				Elements: []*Element{
					{Tag: "tag", Description: "Good description"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateHTMLSpec(tt.args.rc)

			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHTMLSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got.Elements) != 1 {
				t.Errorf("len(gotArray) = %d, want 1", len(got.Elements))
				t.FailNow()
			}

			gotName := got.Name
			wantName := tt.want.Name
			if gotName != wantName {
				t.Errorf("GenerateHTMLSpec().Name = %v, want %v", gotName, wantName)
			}

			gotTag := got.Elements[0].Tag
			wantTag := tt.want.Elements[0].Tag
			if gotTag != wantTag {
				t.Errorf("GenerateHTMLSpec() Element.Tag got = %v, want %v", gotTag, wantTag)
			}

			gotDescription := got.Elements[0].Description
			wantDescription := tt.want.Elements[0].Description
			if gotDescription != wantDescription {
				t.Errorf("GenerateHTMLSpec() Element.Description got = %v, want %v", gotDescription, wantDescription)
			}
		})
	}
}
