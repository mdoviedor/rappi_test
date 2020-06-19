package main

import (
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	type args struct {
		complexityUC ComplexityUCInterface
		tag          int
		iteration    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				complexityUC: NewComplexityUC(TagsConfig{}),
				tag:          4,
				iteration:    15,
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Handler(tt.args.complexityUC, tt.args.tag, tt.args.iteration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerWithMock(t *testing.T) {
	type args struct {
		complexityUC *ComplexityUCMock
		tag          int
		iteration    int
	}
	tests := []struct {
		name string
		args args
		want []int
		mock func(a args)
	}{
		{
			name: "case 1",
			args: args{
				complexityUC: &ComplexityUCMock{},
				tag:          4,
				iteration:    15,
			},
			want: []int{4, 3, 2, 1},
			mock: func(a args) {
				a.complexityUC.On("calculate", 4, 15).Once().Return([]int{4, 3, 2, 1})
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.args)

		t.Run(tt.name, func(t *testing.T) {
			if got := Handler(tt.args.complexityUC, tt.args.tag, tt.args.iteration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})

		tt.args.complexityUC.AssertExpectations(t)
	}
}
