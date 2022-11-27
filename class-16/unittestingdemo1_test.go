package main

// import (
// 	"testing"
// )

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "first",
// 		},
// 		{
// 			name: "second",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

// func TestSum(t *testing.T) {
// 	type args struct {
// 		x int
// 		y int
// 	}

// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "first",
// 			args: args{
// 				x: 10,
// 				y: 20,
// 			},
// 			want: 30,
// 		},
// 		{
// 			name: "second",
// 			args: args{
// 				x: -10,
// 				y: 10,
// 			},
// 			want: 0,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
// 				t.Errorf("Sum() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestAbs(t *testing.T) {
// 	type args struct {
// 		x int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "first",
// 			args: args{
// 				x: -10,
// 			},
// 			want: 10,
// 		},
// 		{
// 			name: "second",
// 			args: args{
// 				x: 10,
// 			},
// 			want: 10,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Abs(tt.args.x); got != tt.want {
// 				t.Errorf("Abs() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
