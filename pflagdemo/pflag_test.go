package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/spf13/pflag"
)

var (
	testBool                     = pflag.Bool("test_bool", false, "bool value")
	testInt                      = pflag.Int("test_int", 0, "int value")
	testInt64                    = pflag.Int64("test_int64", 0, "int64 value")
	testUint                     = pflag.Uint("test_uint", 0, "uint value")
	testUint64                   = pflag.Uint64("test_uint64", 0, "uint64 value")
	testString                   = pflag.String("test_string", "0", "string value")
	testFloat                    = pflag.Float64("test_float64", 0, "float64 value")
	testDuration                 = pflag.Duration("test_duration", 0, "time.Duration value")
	testOptionalInt              = pflag.Int("test_optional_int", 0, "optional int value")
	normalizeFlagNameInvocations = 0
)

func boolString(s string) string {
	if s == "0" {
		return "false"
	}
	return "true"
}

func TestEverything(t *testing.T) {
	pflag.Parse()
	m := make(map[string]*pflag.Flag)
	desired := "0"
	visitor := func(f *pflag.Flag) {
		if len(f.Name) > 5 && f.Name[0:5] == "test_" {
			m[f.Name] = f
			ok := false
			switch {
			case f.Value.String() == desired:
				ok = true
			case f.Name == "test_bool" && f.Value.String() == boolString(desired):
				ok = true
			case f.Name == "test_duration" && f.Value.String() == desired+"s":
				ok = true
			}
			if !ok {
				t.Error("Visit: bad value", f.Value.String(), "for", f.Name)
			}
		}
	}
	pflag.VisitAll(visitor)
	if len(m) != 9 {
		t.Error("VisitAll misses some flags")
		for k, v := range m {
			t.Log(k, *v)
		}
	}
	m = make(map[string]*pflag.Flag)
	pflag.Visit(visitor)
	if len(m) != 0 {
		t.Errorf("Visit sees unset flags")
		for k, v := range m {
			t.Log(k, *v)
		}
	}
	// Now set all flags
	pflag.Set("test_bool", "true")
	pflag.Set("test_int", "1")
	pflag.Set("test_int64", "1")
	pflag.Set("test_uint", "1")
	pflag.Set("test_uint64", "1")
	pflag.Set("test_string", "1")
	pflag.Set("test_float64", "1")
	pflag.Set("test_duration", "1s")
	pflag.Set("test_optional_int", "1")
	desired = "1"
	pflag.Visit(visitor)
	if len(m) != 9 {
		t.Error("Visit fails after set")
		for k, v := range m {
			t.Log(k, *v)
		}
	}
	// Now test they're visited in sort order.
	var flagNames []string
	pflag.Visit(func(f *pflag.Flag) { flagNames = append(flagNames, f.Name) })
	if !sort.StringsAreSorted(flagNames) {
		t.Errorf("flag names not sorted: %v", flagNames)
	}
}

func TestAddFlagSet(t *testing.T) {
	oldSet := pflag.NewFlagSet("old", pflag.ContinueOnError)
	newSet := pflag.NewFlagSet("new", pflag.ContinueOnError)

	oldSet.String("flag1", "flag1", "flag1")
	oldSet.String("flag2", "flag2", "flag2")

	newSet.String("flag2", "flag2", "flag2")
	newSet.String("flag3", "flag3", "flag3")

	oldSet.AddFlagSet(newSet)

	// if len(oldSet.formal) != 3 {
	// 	t.Errorf("Unexpected result adding a FlagSet to a FlagSet %v", oldSet)
	// }
	fmt.Printf("%+v", oldSet)
}
