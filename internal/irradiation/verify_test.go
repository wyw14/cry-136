package irradiation

import (
	"sync"
	"testing"
)

func TestConcurrentExperimentsKeepIrradiationPositionExclusive(t *testing.T) {
	b := NewLeaseBook()
	var wg sync.WaitGroup
	results := make(chan bool, 2)
	for _, owner := range []string{"alpha", "beta"} {
		wg.Add(1)
		go func(value string) { defer wg.Done(); results <- b.Claim(value) }(owner)
	}
	wg.Wait()
	close(results)
	count := 0
	for result := range results { if result { count++ } }
	if count != 1 { t.Fatalf("claim count = %d", count) }
}
