package main

import "testing"

func Test_validateSort(t *testing.T) {
	type args struct {
		slc *[]float64
	}
	unsorted := []float64{3.0, 2.0, 4.0, 1.0}
	sorted := []float64{1.0, 2.0, 3.0, 4.0}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Non-Sorted", args: args{&unsorted}},
		{name: "Sorted", args: args{&sorted}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateSort(tt.args.slc); got != tt.want {
				t.Errorf("validateSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
