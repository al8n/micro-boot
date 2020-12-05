package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func setUpS2F32FlagSet(s2sp *map[string]float32) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToFloat32Var(s2sp, "s2f32", map[string]float32{}, "Command separated ls2f32t!")
	return f
}

func setUpS2F32FlagSetWithDefault(s2sp *map[string]float32) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToFloat32VarP(s2sp, "s2f32", "f", map[string]float32{"da": 1, "db": 2, "de": 5}, "Command separated ls2f32t!")
	return f
}

func createS2F32Flag(vals map[string]float32) string {
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

func TestEmptyS2F32(t *testing.T) {
	var s2f32 map[string]float32
	f := setUpS2F32FlagSet(&s2f32)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getS2F32, err := f.GetStringToFloat32("s2f32")
	if err != nil {
		t.Fatal("got an error from GetStringToFloat32():", err)
	}
	if len(getS2F32) != 0 {
		t.Fatalf("got s2f32 %v with len=%d but expected length=0", getS2F32, len(getS2F32))
	}
}

func TestS2F32(t *testing.T) {
	var s2f32 map[string]float32
	f := setUpS2F32FlagSet(&s2f32)

	vals := map[string]float32{"a": 1, "b": 2, "d": 4, "c": 3, "e": 5}
	arg := fmt.Sprintf("--s2f32=%s", createS2F32Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range s2f32 {
		if vals[k] != v {
			t.Fatalf("expected s2f32[%s] to be %f but got: %f", k, vals[k], v)
		}
	}
	getS2F32, err := f.GetStringToFloat32("s2f32")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	for k, v := range getS2F32 {
		if vals[k] != v {
			t.Fatalf("expected s2f32[%s] to be %f but got: %f from GetStringToFloat32", k, vals[k], v)
		}
	}
}

func TestS2F32Default(t *testing.T) {
	var s2f32 map[string]float32
	f := setUpS2F32FlagSetWithDefault(&s2f32)

	vals := map[string]float32{"da": 1, "db": 2, "de": 5}

	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range s2f32 {
		if vals[k] != v {
			t.Fatalf("expected s2f32[%s] to be %f but got: %f", k, vals[k], v)
		}
	}

	getS2F32, err := f.GetStringToFloat32("s2f32")
	if err != nil {
		t.Fatal("got an error from GetStringToString():", err)
	}
	for k, v := range getS2F32 {
		if vals[k] != v {
			t.Fatalf("expected s2f32[%s] to be %f from GetStringToFloat32 but got: %f", k, vals[k], v)
		}
	}
}

func TestS2F32WithDefault(t *testing.T) {
	var s2f32 *map[string]float32
	f := NewFlagSet("test", ContinueOnError)
	s2f32 = f.StringToFloat32P( "s2f32", "f", map[string]float32{"da": 1, "db": 2, "de": 5}, "Command separated ls2i32t!")

	vals := map[string]float32{"a": 1, "b": 2, "e": 5}
	arg := fmt.Sprintf("-f %s", createS2F32Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range *s2f32 {
		if vals[k] != v {
			t.Fatalf("expected s2f32[%s] to be %f but got: %f", k, vals[k], v)
		}
	}

	getS2F32, err := f.GetStringToFloat32("s2f32")
	if err != nil {
		t.Fatal("got an error from GetStringToFloat32():", err)
	}
	for k, v := range getS2F32 {
		if vals[k] != v {
			t.Fatalf("expected s2f32[%s] to be %f from GetStringToFloat32 but got: %f", k, vals[k], v)
		}
	}
}

func TestS2F32CalledTwice(t *testing.T) {
	var s2f32 *map[string]float32
	f := NewFlagSet("test", ContinueOnError)
	s2f32 = f.StringToFloat32("s2f32", map[string]float32{}, "Command separated ls2i32t!")

	in := []string{"a=1,b=2", "b=3", `"e=5"`, `f=7`}
	expected := map[string]float32{"a": 1, "b": 3, "e": 5, "f": 7}
	argfmt := "--s2f32=%s"
	arg0 := fmt.Sprintf(argfmt, in[0])
	arg1 := fmt.Sprintf(argfmt, in[1])
	arg2 := fmt.Sprintf(argfmt, in[2])
	arg3 := fmt.Sprintf(argfmt, in[3])
	err := f.fs.Parse([]string{arg0, arg1, arg2, arg3})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	if len(*s2f32) != len(expected) {
		t.Fatalf("expected %d flags; got %d flags", len(expected), len(*s2f32))
	}
	for i, v := range *s2f32 {
		if expected[i] != v {
			t.Fatalf("expected s2f32[%s] to be %f but got: %f", i, expected[i], v)
		}
	}
}