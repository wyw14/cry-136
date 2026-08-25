package cycle

import "github.com/google/uuid"

type PermitLink struct {
	CycleID uuid.UUID
	Epoch   uint64
}

func NewPermitLink(id uuid.UUID, epoch uint64) PermitLink {
	return PermitLink{CycleID: id, Epoch: epoch}
}

func (p PermitLink) Matches(id uuid.UUID, epoch uint64) bool {
	return p.CycleID == id && p.Epoch == epoch
}
