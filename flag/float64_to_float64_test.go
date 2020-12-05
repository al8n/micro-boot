// Copyright 2009 The Go Authors. All rights reserved.
// Use of ths2s source code s2s governed by a BSD-style
// license that can be found in the LICENSE file.

package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func setUpF642F64FlagSet(f2fp *map[float64]float64) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.Float64ToFloat64Var(f2fp, "f642f64", map[float64]float64{}, "Command separated lf642f64t!")
	return f
}

func setUpF642F64FlagSetWithDefault(s2sp *map[float64]float64) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.Float64ToFloat64Var(s2sp, "f642f64", map[float64]float64{1.3: 1.3, 1.4: 1.4, 1.5: 1.5}, "Command separated lf642f64t!")
	return f
}

func createF642F64Flag(vals map[float64]float64) string {
	records := make([]string, 0, len(vals)>>1)
	for k, v := range vals {
		records = append(records, fmt.Sprintf("%f=%f", k, v))
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return strings.TrimSpace(buf.String())
}

func TestEmptyF642F64(t *testing.T) {
	var f642f64 map[float64]float64
	f := setUpF642F64FlagSet(&f642f64)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getF642F64, err := f.GetFloat64ToFloat64("f642f64")
	if err != nil {
		t.Fatal("got an error from GetFloat64ToFloat64():", err)
	}
	if len(getF642F64) != 0 {
		t.Fatalf("got f642f64 %v with len=%d but expected length=0", getF642F64, len(getF642F64))
	}
}

func TestF642F64(t *testing.T) {
	var f642f64 map[float64]float64
	f := NewFlagSet("test", ContinueOnError)
	f.Float64ToFloat64VarP(&f642f64, "f642f64", "f", map[float64]float64{}, "Command separated lf642f64t!")

	vals := map[float64]float64{1.4: 1.4, 1.5: 1.5, 1.6: 1.6, 1.7: 1.7, 1.8: 1.8}
	arg := fmt.Sprintf("--f642f64=%s", createF642F64Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range f642f64 {
		if vals[k] != v {
			t.Fatalf("expected f642f64[%f] to be %f but got: %f", k, vals[k], v)
		}
	}
	getF642F64, err := f.GetFloat64ToFloat64("f642f64")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	for k, v := range getF642F64 {
		if vals[k] != v {
			t.Fatalf("expected f642f64[%f] to be %f but got: %f from GetFloat64ToFloat64", k, vals[k], v)
		}
	}
}

func TestF642F64Default(t *testing.T) {
	var f642f64 *map[float64]float64
	f := NewFlagSet("test", ContinueOnError)
	f642f64 = f.Float64ToFloat64( "f642f64", map[float64]float64{1.3: 1.3, 1.4: 1.4, 1.5: 1.5}, "Command separated lf642f64t!")

	vals := map[float64]float64{1.1: 1.1, 1.2: 1.2, 1.3: 1.3, 1.4: 1.4, 1.5: 1.5}

	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range *f642f64 {
		if vals[k] != v {
			t.Fatalf("expected f642f64[%f] to be %f but got: %f", k, vals[k], v)
		}
	}

	getF642F64, err := f.GetFloat64ToFloat64("f642f64")
	if err != nil {
		t.Fatal("got an error from GetFloat64ToFloat64():", err)
	}
	for k, v := range getF642F64 {
		if vals[k] != v {
			t.Fatalf("expected f642f64[%f] to be %f from GetStringToString but got: %f", k, vals[k], v)
		}
	}
}

func TestF642F64WithDefault(t *testing.T) {
	var f642f64 *map[float64]float64
	f := NewFlagSet("test", ContinueOnError)
	f642f64 = f.Float64ToFloat64P( "f642f64", "f", map[float64]float64{1.3: 1.3, 1.4: 1.4, 1.5: 1.5}, "Command separated lf642f64t!")

	vals := map[float64]float64{1.1: 1.1, 1.2: 1.2, 1.3: 1.3, 1.4: 1.4, 1.5: 1.5}
	arg := fmt.Sprintf("--f642f64=%s", createF642F64Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range *f642f64 {
		if vals[k] != v {
			t.Fatalf("expected f642f64[%f] to be %f but got: %f", k, vals[k], v)
		}
	}

	getF642F64, err := f.GetFloat64ToFloat64("f642f64")
	if err != nil {
		t.Fatal("got an error from GetStringToString():", err)
	}
	for k, v := range getF642F64 {
		if vals[k] != v {
			t.Fatalf("expected s2s[%f] to be %f from GetFloat64ToFloat64 but got: %f", k, vals[k], v)
		}
	}
}

func TestF642F64CalledTwice(t *testing.T) {
	var f642f64 map[float64]float64
	f := setUpF642F64FlagSet(&f642f64)

	in := []string{"1.3=1.3,1.4=1.4", "1.5=1.5", `"1.6=1.6"`, `1.7=1.7`}
	expected := map[float64]float64{1.3: 1.3, 1.4:1.4, 1.5: 1.5, 1.6: 1.6, 1.7: 1.7}
	argfmt := "--f642f64=%s"
	arg0 := fmt.Sprintf(argfmt, in[0])
	arg1 := fmt.Sprintf(argfmt, in[1])
	arg2 := fmt.Sprintf(argfmt, in[2])
	arg3 := fmt.Sprintf(argfmt, in[3])
	err := f.fs.Parse([]string{arg0, arg1, arg2, arg3})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	if len(f642f64) != len(expected) {
		t.Fatalf("expected %d flags; got %d flags", len(expected), len(f642f64))
	}
	for i, v := range f642f64 {
		if expected[i] != v {
			t.Fatalf("expected f642f64[%f] to be %f but got: %f", i, expected[i], v)
		}
	}
}
