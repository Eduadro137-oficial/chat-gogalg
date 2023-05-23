package portscanner

import (
	"reflect"
	"testing"
	"time"
)

func TestNewPortScanner(t *testing.T) {
	type args struct {
		host    string
		timeout time.Duration
		threads int
	}
	tests := []struct {
		name string
		args args
		want *PortScanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPortScanner(tt.args.host, tt.args.timeout, tt.args.threads); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPortScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}
