package jsm

import (
	"reflect"
	"strings"
	"testing"
)

func TestDbSpecificQuery(t *testing.T) {
	tests := []struct {
		name string
		want map[dataBase]func(sb *strings.Builder)
	}{
		{
			"postgres",
			map[dataBase]func(sb *strings.Builder){
				postgres: func(sb *strings.Builder) {
					sb.WriteString(" WHERE")
				},
				mysql: func(sb *strings.Builder) {
					sb.WriteString(" WHERE")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DbSpecificQuery(); !reflect.DeepEqual(len(got), len(tt.want)) {
				t.Errorf("DbSpecificQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
