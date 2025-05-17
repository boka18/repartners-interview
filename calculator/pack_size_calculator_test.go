package calculator_test

import (
	"testing"

	"reflect"

	"github.com/boka18/repartners-interview/calculator"
)

func TestCalculate(t *testing.T) {
	pc := calculator.NewPackSize()

	tests := []struct {
		name      string
		packSizes []int
		order     int
		wantTotal int
		wantPacks map[int]int
	}{
		{
			name:      "exact match with one pack",
			packSizes: []int{250, 500, 1000},
			order:     1000,
			wantTotal: 1000,
			wantPacks: map[int]int{1000: 1},
		},
		{
			name:      "small overage with fewer packs",
			packSizes: []int{250, 500},
			order:     600,
			wantTotal: 750,
			wantPacks: map[int]int{500: 1, 250: 1},
		},
		{
			name:      "non-trivial combination",
			packSizes: []int{23, 31, 53},
			order:     76,
			wantTotal: 76,
			wantPacks: map[int]int{23: 1, 53: 1},
		},
		{
			name:      "large amount",
			packSizes: []int{23, 31, 53},
			order:     500000,
			wantTotal: 500000,
			wantPacks: map[int]int{23: 2, 31: 7, 53: 9429},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pc.Calculate(tt.packSizes, tt.order)

			if got.TotalItems < tt.order {
				t.Errorf("TotalItems = %d, expected >= %d", got.TotalItems, tt.order)
			}

			if tt.wantPacks != nil && !reflect.DeepEqual(got.PacksUsed, tt.wantPacks) {
				t.Errorf("PacksUsed = %+v, want %+v", got.PacksUsed, tt.wantPacks)
			}
		})
	}
}
