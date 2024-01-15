package src

import "testing"

func TestCompute_AddTwo(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		c    *Compute
		args args
		want int
	}{

		{
			name: "test1",
			c: &Compute{
				D: 10,
			},
			args: args{
				a: 1,
				b: 2,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.AddTwo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compute.AddTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompute_Divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		c       *Compute
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test1",
			c: &Compute{
				D: 10,
			},
			args: args{
				a: 10,
				b: 2,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "test2",
			c: &Compute{
				D: 10,
			},
			args: args{
				a: 10,
				b: -2,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Divide(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compute.Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compute.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
