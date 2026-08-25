package irradiation

import (
	"sync"
	"testing"
)

func TestConcurrentClaimHasSingleOwner(t *testing.T) {
	const workers = 64
	book := NewLeaseBook()
	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		winners  []string
	)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			owner := "exp-" + string(rune('A'+i))
			if book.Claim(owner) {
				mu.Lock()
				winners = append(winners, owner)
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	if len(winners) != 1 {
		t.Fatalf("expected exactly one owner to win the position, got %d (%v)", len(winners), winners)
	}
	if got := book.Owner(); got != winners[0] {
		t.Fatalf("position owner mismatch: book has %q, winner was %q", got, winners[0])
	}
}

func TestReleaseAllowsNextOwner(t *testing.T) {
	book := NewLeaseBook()
	if !book.Claim("exp-A") {
		t.Fatal("first claim should succeed")
	}
	if book.Claim("exp-B") {
		t.Fatal("second claim while position held should fail")
	}
	if !book.Release("exp-A") {
		t.Fatal("release by owner should succeed")
	}
	if !book.Claim("exp-B") {
		t.Fatal("claim after release should succeed")
	}
}
