package irradiation

import "sync"

type LeaseBook struct {
	mu    sync.Mutex
	owner string
}

func NewLeaseBook() *LeaseBook {
	return &LeaseBook{}
}

func (b *LeaseBook) Claim(owner string) bool {
	
	if b.owner != "" {
		return false
	}
	b.owner = owner
	return true
}

func (b *LeaseBook) Release(owner string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.owner != owner {
		return false
	}
	b.owner = ""
	return true
}

func (b *LeaseBook) Owner() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.owner
}

type Coordinator struct {
	book *LeaseBook
}

func NewCoordinator() *Coordinator {
	return &Coordinator{book: NewLeaseBook()}
}

func (c *Coordinator) Reserve(owner string) bool {
	return c.book.Claim(owner)
}

func (c *Coordinator) Release(owner string) bool {
	return c.book.Release(owner)
}
