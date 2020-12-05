package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func setUpS2F64FlagSet(s2sp *map[string]float64) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToFloat64Var(s2sp, "s2f64", map[string]float64{}, "Command separated ls2f64t!")
	return f
}

func setUpS2F64FlagSetWithDefault(s2sp *map[string]float64) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToFloat64VarP(s2sp, "s2f64", "f", map[string]float64{"da": 1, "db": 2, "de": 5}, "Command separated ls2f64t!")
	return f
}

func createS2F64Flag(vals map[string]float64) string {
	records := make([]string, 0, len(vals)>>1)
	for k, v := range vals {
		records = append(records, fmt.Sprintf("%s=%f",k, v))
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return strings.TrimSpace(buf.String())
}

func TestEmptyS2F64(t *testing.T) {
	var s2f64 map[string]float64
	f := setUpS2F64FlagSet(&s2f64)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getS2F64, err := f.GetStringToFloat64("s2f64")
	if err != nil {
		t.Fatal("got an error from GetStringToFloat64():", err)
	}
	if len(getS2F64) != 0 {
		t.Fatalf("got s2f64 %v with len=%d but expected length=0", getS2F64, len(getS2F64))
	}
}

func TestS2F64(t *testing.T) {
	var s2f64 map[string]float64
	f := setUpS2F64FlagSet(&s2f64)

	vals := map[string]float64{"a": 1, "b": 2, "d": 4, "c": 3, "e": 5}
	arg := fmt.Sprintf("--s2f64=%s", createS2F64Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range s2f64 {
		if vals[k] != v {
			t.Fatalf("expected s2f64[%s] to be %f but got: %f", k, vals[k], v)
		}
	}
	getS2F64, err := f.GetStringToFloat64("s2f64")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	for k, v := range getS2F64 {
		if vals[k] != v {
			t.Fatalf("expected s2f64[%s] to be %f but got: %f from GetStringToFloat64", k, vals[k], v)
		}
	}
}

func TestS2F64Default(t *testing.T) {
	var s2f64 map[string]float64
	f := setUpS2F64FlagSetWithDefault(&s2f64)

	vals := map[string]float64{"da": 1, "db": 2, "de": 5}

	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range s2f64 {
		if vals[k] != v {
			t.Fatalf("expected s2f64[%s] to be %f but got: %f", k, vals[k], v)
		}
	}

	getS2F64, err := f.GetStringToFloat64("s2f64")
	if err != nil {
		t.Fatal("got an error from GetStringToString():", err)
	}
	for k, v := range getS2F64 {
		if vals[k] != v {
			t.Fatalf("expected s2f64[%s] to be %f from GetStringToFloat64 but got: %f", k, vals[k], v)
		}
	}
}

func TestS2F64WithDefault(t *testing.T) {
	var s2f64 *map[string]float64
	f := NewFlagSet("test", ContinueOnError)
	s2f64 = f.StringToFloat64P( "s2f64", "f", map[string]float64{"da": 1, "db": 2, "de": 5}, "Command separated ls2f64t!")

	vals := map[string]float64{"a": 1, "b": 2, "e": 5}
	arg := fmt.Sprintf("-f %s", createS2F64Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range *s2f64 {
		if vals[k] != v {
			t.Fatalf("expected s2f64[%s] to be %f but got: %f", k, vals[k], v)
		}
	}

	getS2F64, err := f.GetStringToFloat64("s2f64")
	if err != nil {
		t.Fatal("got an error from GetStringToFloat64():", err)
	}
	for k, v := range getS2F64 {
		if vals[k] != v {
			t.Fatalf("expected s2f64[%s] to be %f from GetStringToFloat64 but got: %f", k, vals[k], v)
		}
	}
}

func TestS2F64CalledTwice(t *testing.T) {
	var s2f64 *map[string]float64
	f := NewFlagSet("test", ContinueOnError)
	s2f64 = f.StringToFloat64("s2f64", map[string]float64{}, "Command separated ls2f64t!")

	in := []string{"a=1,b=2", "b=3", `"e=5"`, `f=7`}
	expected := map[string]float64{"a": 1, "b": 3, "e": 5, "f": 7}
	argfmt := "--s2f64=%s"
	arg0 := fmt.Sprintf(argfmt, in[0])
	arg1 := fmt.Sprintf(argfmt, in[1])
	arg2 := fmt.Sprintf(argfmt, in[2])
	arg3 := fmt.Sprintf(argfmt, in[3])
	err := f.fs.Parse([]string{arg0, arg1, arg2, arg3})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	if len(*s2f64) != len(expected) {
		t.Fatalf("expected %d flags; got %d flags", len(expected), len(*s2f64))
	}
	for i, v := range *s2f64 {
		if expected[i] != v {
			t.Fatalf("expected s2f64[%s] to be %f but got: %f", i, expected[i], v)
		}
	}
}