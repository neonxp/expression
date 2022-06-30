package expression

import (
	"reflect"
	"testing"
)

func TestEvaluator_Eval(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "simple math",
			args: args{
				expression: "2 + 2 * 2 + max(4,9)",
			},
			want:    2 + 2*2 + 9,
			wantErr: false,
		},
		{
			name: "simple math 2",
			args: args{
				expression: "10 % 5",
			},
			want:    10 % 5,
			wantErr: false,
		},
		{
			name: "simple math 3",
			args: args{
				expression: "10 / 5",
			},
			want:    10 / 5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New()
			got, err := e.Eval(tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluator.Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluator.Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
