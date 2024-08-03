package main

import "testing"

func TestDeserailize(t *testing.T) {
	type args struct {
		raw  string
		dest any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    any
	}{
		{
			name: "int",
			args: args{
				raw:  "1",
				dest: new(int),
			},
			wantErr: false,
			want:    1,
		},
		{
			name: "float",
			args: args{
				raw:  "1.5",
				dest: new(float32),
			},
			wantErr: false,
			want:    1.5,
		},
		{
			name: "string",
			args: args{
				raw:  "\"aaa\"",
				dest: new(string),
			},
			wantErr: false,
			want:    "aaa",
		},
		{
			name: "boolean",
			args: args{
				raw:  "true",
				dest: new(bool),
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "array",
			args: args{
				raw:  "[1,2,3]",
				dest: new([]int),
			},
			wantErr: false,
			want:    []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Deserailize(tt.args.raw, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Deserailize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
