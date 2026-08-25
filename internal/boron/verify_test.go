package boron

import "testing"

func TestScramResultPreservesBoronInjectionFailure(t *testing.T) {
	result := MergeScram(true, "injection timeout")
	if result.Status != "failed" || result.Error == "" { t.Fatalf("scram result = %#v", result) }
}
