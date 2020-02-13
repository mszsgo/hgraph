package hgraph

import (
	"reflect"
	"strings"
	"testing"
)

func TestGateway(t *testing.T) {
	graphql := `
	{
		"requestId":"uuid",
		"token":"t",
		"query":"
			query {
			  captcha {
				number(len: 6) {
				  captchaId
				  base64Image
				}
			  }
			}
		"
	}
	`
	graphql = strings.ReplaceAll(graphql, "\n", "")
	graphql = strings.ReplaceAll(graphql, "\t", "")
	type args struct {
		body []byte
	}
	tests := []struct {
		name      string
		args      args
		wantBytes []byte
	}{
		// TODO: Add test cases.
		{name: "TGateway", args: args{body: []byte(graphql)}, wantBytes: []byte("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes := Gateway(tt.args.body)
			if !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Logf("Gateway() gotString = %s ", string(gotBytes))
			}
		})
	}
}
