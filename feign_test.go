package hgraph

import (
	"reflect"
	"testing"
)

func TestGraphql(t *testing.T) {
	type args struct {
		hql string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Graphql(tt.args.hql); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graphql() = %v, want %v", got, tt.want)
			}
		})
	}
}
