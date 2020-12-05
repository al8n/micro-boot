package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func setUpS2I32FlagSet(s2sp *map[string]int32) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToInt32Var(s2sp, "s2i32", map[string]int32{}, "Command separated ls2i32t!")
	return f
}

func setUpS2I32FlagSetWithDefault(s2sp *map[string]int32) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.StringToInt32VarP(s2sp, "s2i32", "i", map[string]int32{"da": 1, "db": 2, "de": 5}, "Command separated ls2i32t!")
	return f
}

func createS2I32Flag(vals map[string]int32) string {
	records := make([]string, 0, len(vals)>>1)
	for k, v := range vals {
		records = append(records, fmt.Sprintf("%s=%d",k, v))
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return strings.TrimSpace(buf.String())
}

func TestEmptyS2I32(t *testing.T) {
	var s2i32 map[string]int32
	f := setUpS2I32FlagSet(&s2i32)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getS2I32, err := f.GetStringToInt32("s2i32")
	if err != nil {
		t.Fatal("got an error from GetStringToString():", err)
	}
	if len(getS2I32) != 0 {
		t.Fatalf("got s2i32 %v with len=%d but expected length=0", getS2I32, len(getS2I32))
	}
}

func TestS2I32(t *testing.T) {
	var s2i32 map[string]int32
	f := setUpS2I32FlagSet(&s2i32)

	vals := map[string]int32{"a": 1, "b": 2, "d": 4, "c": 3, "e": 5}
	arg := fmt.Sprintf("--s2i32=%s", createS2I32Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range s2i32 {
		if vals[k] != v {
			t.Fatalf("expected s2i32[%s] to be %d but got: %d", k, vals[k], v)
		}
	}
	getS2I32, err := f.GetStringToInt32("s2i32")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	for k, v := range getS2I32 {
		if vals[k] != v {
			t.Fatalf("expected s2i32[%s] to be %d but got: %d from GetStringToInt32", k, vals[k], v)
		}
	}
}

func TestS2I32Default(t *testing.T) {
	var s2i32 map[string]int32
	f := setUpS2I32FlagSetWithDefault(&s2i32)

	vals := map[string]int32{"da": 1, "db": 2, "de": 5}

	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range s2i32 {
		if vals[k] != v {
			t.Fatalf("expected s2i32[%s] to be %d but got: %d", k, vals[k], v)
		}
	}

	getS2I32, err := f.GetStringToInt32("s2i32")
	if err != nil {
		t.Fatal("got an error from GetStringToString():", err)
	}
	for k, v := range getS2I32 {
		if vals[k] != v {
			t.Fatalf("expected s2i32[%s] to be %d from GetStringToInt32 but got: %d", k, vals[k], v)
		}
	}
}

func TestS2I32WithDefault(t *testing.T) {
	var s2i32 *map[string]int32
	f := NewFlagSet("test", ContinueOnError)
	s2i32 = f.StringToInt32P( "s2i32", "i", map[string]int32{"da": 1, "db": 2, "de": 5}, "Command separated ls2i32t!")

	vals := map[string]int32{"a": 1, "b": 2, "e": 5}
	arg := fmt.Sprintf("-i %s", createS2I32Flag(vals))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for k, v := range *s2i32 {
		if vals[k] != v {
			t.Fatalf("expected s2i32[%s] to be %d but got: %d", k, vals[k], v)
		}
	}

	getS2I32, err := f.GetStringToInt32("s2i32")
	if err != nil {
		t.Fatal("got an error from GetStringToString():", err)
	}
	for k, v := range getS2I32 {
		if vals[k] != v {
			t.Fatalf("expected s2i32[%s] to be %d from GetStringToInt32 but got: %d", k, vals[k], v)
		}
	}
}

func TestS2I32CalledTwice(t *testing.T) {
	var s2int32 *map[string]int32
	f := NewFlagSet("test", ContinueOnError)
	s2int32 = f.StringToInt32("s2i32", map[string]int32{}, "Command separated ls2i32t!")

	in := []string{"a=1,b=2", "b=3", `"e=5"`, `f=7`}
	expected := map[string]int32{"a": 1, "b": 3, "e": 5, "f": 7}
	argfmt := "--s2i32=%s"
	arg0 := fmt.Sprintf(argfmt, in[0])
	arg1 := fmt.Sprintf(argfmt, in[1])
	arg2 := fmt.Sprintf(argfmt, in[2])
	arg3 := fmt.Sprintf(argfmt, in[3])
	err := f.fs.Parse([]string{arg0, arg1, arg2, arg3})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	if len(*s2int32) != len(expected) {
		t.Fatalf("expected %d flags; got %d flags", len(expected), len(*s2int32))
	}
	for i, v := range *s2int32 {
		if expected[i] != v {
			t.Fatalf("expected s2i32[%s] to be %d but got: %d", i, expected[i], v)
		}
	}
}