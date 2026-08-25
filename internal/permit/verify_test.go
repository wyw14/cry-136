package permit

import "testing"

func TestPowerPermitRequiresIndependentNeutronChains(t *testing.T) {
	if CanIssuePowerPermit([]string{"shared", "shared"}) { t.Fatal("correlated channels counted twice") }
	if !CanIssuePowerPermit([]string{"a", "b"}) { t.Fatal("independent channels rejected") }
}
