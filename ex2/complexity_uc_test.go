package main

import (
	"reflect"
	"testing"
)

func TestComplexityUC_calculate(t *testing.T) {
	type fields struct {
		tags TagsConfig
	}
	type args struct {
		tag       int
		iteration int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "case 1",
			fields: fields{
				tags: TagsConfig{},
			},
			args: args{
				tag:       4,
				iteration: 15,
			},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "case 2",
			fields: fields{
				tags: TagsConfig{},
			},
			args: args{
				tag:       3,
				iteration: 15,
			},
			want: []int{3, 2, 1},
		},
		{
			name: "case 3",
			fields: fields{
				tags: TagsConfig{},
			},
			args: args{
				tag:       4,
				iteration: 5,
			},
			want: []int{4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ComplexityUC{
				tags: tt.fields.tags,
			}
			if got := c.calculate(tt.args.tag, tt.args.iteration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewComplexityUC(t *testing.T) {
	type args struct {
		tags TagsConfig
	}
	tests := []struct {
		name string
		args args
		want *ComplexityUC
	}{
		{
			name: "success",
			args: args{
				tags: TagsConfig{},
			},
			want: &ComplexityUC{
				tags: TagsConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComplexityUC(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComplexityUC() = %v, want %v", got, tt.want)
			}
		})
	}
}
