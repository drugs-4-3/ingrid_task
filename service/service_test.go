package service

import (
	"reflect"
	"testing"
)

func Test_sortByEffort(t *testing.T) {
	tests := []struct {
		name        string
		inputRoutes []Route
		want        []Route
	}{
		{
			name: "should sort by duration",
			inputRoutes: []Route{
				{
					Duration: 100,
					Distance: 200,
				},
				{
					Duration: 110,
					Distance: 190,
				},
				{
					Duration: 90,
					Distance: 210,
				},
			},
			want: []Route{
				{
					Duration: 90,
					Distance: 210,
				},
				{
					Duration: 100,
					Distance: 200,
				},
				{
					Duration: 110,
					Distance: 190,
				},
			},
		},
		{
			name: "should sort by distance if duration is equal",
			inputRoutes: []Route{
				{
					Duration: 100,
					Distance: 200,
				},
				{
					Duration: 100,
					Distance: 190,
				},
				{
					Duration: 110,
					Distance: 180,
				},
			},
			want: []Route{
				{
					Duration: 100,
					Distance: 190,
				},
				{
					Duration: 100,
					Distance: 200,
				},
				{
					Duration: 110,
					Distance: 180,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByEffort(tt.inputRoutes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByEffort() = %v, want %v", got, tt.want)
			}
		})
	}
}
