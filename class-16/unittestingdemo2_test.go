package main

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	mock "github.com/shahincsejnu/demo-golang/class-16/mock"
)

type testStruct struct {
}

func (ts testStruct) Sum(x, y int) (int, error) {
	return 0, fmt.Errorf("dummy error!")
}

func Test_process(t *testing.T) {
	t.Skip()
	type args struct {
		d Demo
		x int
		y int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				d: Temp{},
				x: 10,
				y: 20,
			},
			want:    30,
			wantErr: false,
		},
		{
			name: "second",
			args: args{
				d: testStruct{},
				x: 10,
				y: 30,
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := process(tt.args.d, tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessMock(t *testing.T) {
	mockController := gomock.NewController(t)

	mockStruct := mock.NewMockDemo(mockController)

	mockStruct.EXPECT().Sum(1, 2).Return(4, nil)

	res, err := process(mockStruct, 1, 2)
	fmt.Println(res)
	if err != nil {
		t.Fatal(err)
	}

	if res != 4 {
		t.Fatalf("Expected 3 but found: %v", res)
	}
}
