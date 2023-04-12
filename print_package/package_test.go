package print_package

import (
	"testing"
)

func TestPrintSomething(t *testing.T) {
	type args struct {
		kata string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive case",
			args: args{
				kata: "TEST",
			},
			want: "TESTTEST",
		},
		{
			name: "Positive case 2",
			args: args{
				kata: "ABC",
			},
			want: "HALO ABC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintSomething(tt.args.kata); got != tt.want {
				t.Errorf("PrintSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintSomething2(t *testing.T) {
	type args struct {
		kata  string
		kata2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive case",
			args: args{
				kata:  "TEST",
				kata2: "HALO",
			},
			want: "TEST HALO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintSomething2(tt.args.kata, tt.args.kata2); got != tt.want {
				t.Errorf("PrintSomething2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintSomething3(t *testing.T) {
	type args struct {
		kata  string
		kata2 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Positive case",
			args: args{
				kata:  "TEST",
				kata2: "HALO",
			},
			want:    "TEST HALO",
			wantErr: false,
		},
		{
			name: "Negative case kata",
			args: args{
				kata:  "",
				kata2: "HALO",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Negative case kata2",
			args: args{
				kata:  "HALO",
				kata2: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrintSomething3(tt.args.kata, tt.args.kata2)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintSomething3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrintSomething3() = %v, want %v", got, tt.want)
			}
		})
	}
}
