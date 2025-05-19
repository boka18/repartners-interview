package calculator

import (
	"math"
)

type Pack struct {
	ID   int `json:"id"`
	Size int `json:"size"`
}

type Result struct {
	TotalItems int
	PacksUsed  map[int]int
}

type PackSize struct {
}

func NewPackSize() *PackSize {
	return &PackSize{}
}

// Calculate determines the optimal combination of pack sizes needed to fulfill a given order.
// The algorithm ensures:
//  1. The total number of items shipped is greater than or equal to the order amount.
//  2. Among all valid combinations, the one with the fewest total packs is chosen.
//
// This is solved using a bottom-up dynamic programming (DP) approach.
// Each dp[i] represents the best way to ship exactly `i` items:
//   - `totalPacks`: the minimal number of packs needed to reach `i`
//   - `packsUsed`: a breakdown of how many of each pack size was used
//
// The function builds up solutions from dp[0] up to dp[order + max(packSizes)],
// checking at each step whether we can build the current amount `i` by adding
// one more pack to a previously solved subproblem at dp[i - size].
//
// After filling the DP table, the function finds the first valid solution with
// a total >= order and returns its breakdown.
//
// Parameters:
//   - packSizes: a slice of available pack sizes (e.g., [250, 500, 1000])
//   - order: the target number of items to ship
//
// Returns:
//   - Result{TotalItems, PacksUsed} with the minimal overage and fewest packs possible
func (pc *PackSize) Calculate(packSizes []int, order int) Result {
	maxExtra := max(packSizes...) // to allow minimal overage
	limit := order + maxExtra

	// DP state: for each amount, store total packs used and pack breakdown
	type state struct {
		totalPacks int
		packsUsed  map[int]int
	}

	dp := make([]*state, limit+1)
	dp[0] = &state{totalPacks: 0, packsUsed: map[int]int{}}

	for i := 1; i <= limit; i++ {
		for _, size := range packSizes {
			if i >= size && dp[i-size] != nil {
				prev := dp[i-size]
				totalPacks := prev.totalPacks + 1

				// copy previous packs and add current
				newPacks := make(map[int]int)
				for k, v := range prev.packsUsed {
					newPacks[k] = v
				}
				newPacks[size]++

				// use this if dp[i] is nil or this solution uses fewer packs
				if dp[i] == nil || totalPacks <= dp[i].totalPacks {
					dp[i] = &state{totalPacks: totalPacks, packsUsed: newPacks}
				}
			}
		}
	}

	// find the best solution at or above the target order
	for i := order; i <= limit; i++ {
		if dp[i] != nil {
			return Result{
				TotalItems: i,
				PacksUsed:  dp[i].packsUsed,
			}
		}
	}

	// fallback (shouldn't happen with valid input)
	return Result{TotalItems: -1, PacksUsed: map[int]int{}}
}

func max(nums ...int) int {
	m := math.MinInt
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}
