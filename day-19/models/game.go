package Game

type State struct {
	Patterns []string
}

func NewState() *State {
	return &State{Patterns: []string{}}
}

func equal(p1, p2 []string) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i, v := range p1 {
		if v != p2[i] {
			return false
		}
	}
	return true
}

func isValid(patternSet map[string]bool, s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 1; i <= len(s); i++ {
		if patternSet[s[:i]] && isValid(patternSet, s[i:]) {
			return true
		}
	}
	return false
}

func (s *State) IsValidDesign(design string) bool {
	// Create a map for quick pattern lookup
	patternSet := make(map[string]bool)
	for _, pattern := range s.Patterns {
		patternSet[pattern] = true
	}

	// DP array to store validity of substrings
	n := len(design)
	dp := make([]bool, n+1)
	dp[0] = true // Empty substring is valid

	// fill the DP table
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && patternSet[design[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func (s *State) TotalSolutions(design string) int {
	// Create a map for quick pattern lookup
	patternSet := make(map[string]bool)
	for _, pattern := range s.Patterns {
		patternSet[pattern] = true
	}

	// DP array to store validity of substrings
	n := len(design)
	dp := make([]int, n+1)
	dp[0] = 1 // Base case: there's one way to decompose an empty string

	// fill the DP table
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if patternSet[design[j:i]] {
				dp[i] += dp[j]
			}
		}
	}

	return dp[n]
}
