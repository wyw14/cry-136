package permit

func CanIssuePowerPermit(channels []string) bool {
	if len(channels) < 2 {
		return false
	}
	return channels[0] != channels[1]
}

type PowerPermit struct {
	NeutronRevision uint64
	CoolantRevision uint64
	RadiationOK     bool
}

func (p PowerPermit) Valid() bool {
	return p.NeutronRevision > 0 && p.CoolantRevision > 0 && p.RadiationOK
}

func (p PowerPermit) Revision() uint64 {
	if p.NeutronRevision > p.CoolantRevision {
		return p.NeutronRevision
	}
	return p.CoolantRevision
}
