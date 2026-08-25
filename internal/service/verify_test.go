package service

import "testing"

func TestRecoveryLoadsTripLatchBeforeRodControl(t *testing.T) {
	o := NewRecoveryOrder()
	got := o.Recover()
	if !o.Valid() || len(got) != 3 || got[0] != "trip-latch" { t.Fatalf("recovery order = %#v", got) }
}
