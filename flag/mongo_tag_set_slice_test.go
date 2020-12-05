package flag

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/tag"
	"strings"
	"testing"
)

func setUpMongoTagSetSFlagSet(mtss *[]tag.Set) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.MongoTagSetSliceVar(mtss, "mtss", []tag.Set{}, "Command separated list!")
	return f
}

func setUpMongoTagSetSWithDefault(mtss *[]tag.Set) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.MongoTagSetSliceVar(mtss, "mtss", []tag.Set{
		{
			tag.Tag{
				Name:  "n1",
				Value: "v1",
			},
			tag.Tag{
				Name:  "n2",
				Value: "v2",
			},
		},
		{
			tag.Tag{
				Name:  "n3",
				Value: "v3",
			},
			tag.Tag{
				Name:  "n4",
				Value: "v4",
			},
			tag.Tag{
				Name:  "n5",
				Value: "v5",
			},
		},
	},
	"Command separated list!")

	return f
}

func TestEmptyMongoTagSetSlice(t *testing.T) {
	var mtss []tag.Set
	f := setUpMongoTagSetSFlagSet(&mtss)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getMTSS, err := f.GetMongoTagSetSlice("mtss")
	if err != nil {
		t.Fatal("got an error from GetMongoTagSetSlice():", err)
	}

	if len(getMTSS) != 0 {
		t.Fatalf("got mtss %v with len=%d but expected length=0", getMTSS, len(getMTSS))
	}
}

func TestMongoTagSetSlice(t *testing.T) {
	var mtss []tag.Set
	f := setUpMongoTagSetSFlagSet(&mtss)

	vals := []string{"[n1=v1 n2=v2]", "[n3=v3 n4=v4 n5=v5]"}
	arg := fmt.Sprintf("--mtss=%s", strings.Join(vals, ","))

	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	for i, v := range mtss {
		tv := "[" + strings.ReplaceAll(v.String(), "," ," ") + "]"

		if vals[i] != tv {
			t.Fatalf("expected mtss[%d] to be %s but got: %s", i, vals[i], tv)
		}
	}


	getMTSS, err := f.GetMongoTagSetSlice("mtss")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	for i, v := range getMTSS {
		tv := fmt.Sprintf("[%s]", strings.ReplaceAll(v.String(), ",", " "))

		if vals[i] != tv {
			t.Fatalf("expected mtss[%d] to be %s but got: %s", i, vals[i], tv)
		}
	}
}

func TestMongoTagSetSliceDefault(t *testing.T) {
	var mtss []tag.Set
	f := setUpMongoTagSetSWithDefault(&mtss)

	vals := []string{"[n1=v1 n2=v2]", "[n3=v3 n4=v4 n5=v5]"}

	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range mtss {
		tv := "[" + strings.ReplaceAll(v.String(), "," ," ") + "]"

		if vals[i] != tv {
			t.Fatalf("expected mtss[%d] to be %s but got: %s", i, vals[i], tv)
		}
	}

	getMTSS, err := f.GetMongoTagSetSlice("mtss")
	if err != nil {
		t.Fatal("got an error from GetMongoTagSetSlice():", err)
	}

	for i, v := range getMTSS {
		tv := fmt.Sprintf("[%s]", strings.ReplaceAll(v.String(), ",", " "))

		if vals[i] != tv {
			t.Fatalf("expected mtss[%d] to be %s but got: %s", i, vals[i], tv)
		}
	}
}

func TestMongoTagSetSliceWithDefault(t *testing.T) {
	var mtss []tag.Set
	f := setUpMongoTagSetSWithDefault(&mtss)

	vals := []string{"[n1=v1 n2=v2 n3=v3]", "[n4=v4 n5=v5 n6=v6]"}
	arg := fmt.Sprintf("--mtss=%s", strings.Join(vals, ","))
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range mtss {

		tv := "[" + strings.ReplaceAll(v.String(), "," ," ") + "]"

		if vals[i] != tv {
			t.Fatalf("expected mtss[%d] to be %s but got: %s", i, vals[i], tv)
		}
	}

	getMTSS, err := f.GetMongoTagSetSlice("mtss")
	if err != nil {
		t.Fatal("got an error from GetMongoTagSetSlice():", err)
	}
	for i, v := range getMTSS {
		tv := fmt.Sprintf("[%s]", strings.ReplaceAll(v.String(), ",", " "))

		if vals[i] != tv {
			t.Fatalf("expected mtss[%d] to be %s but got: %s", i, vals[i], tv)
		}
	}
}

func TestMongoTagSetSAsSliceValue(t *testing.T) {
	var mtss []tag.Set
	f := setUpMongoTagSetSFlagSet(&mtss)

	in := []string{"[n1=v1 n2=v2 n3=v3]", "[n4=v4 n5=v5 n6=v6]"}
	argfmt := "--mtss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.fs.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	f.fs.VisitAll(func(f *pflag.Flag) {
		if val, ok := f.Value.(pflag.SliceValue); ok {
			_ = val.Replace([]string{"[n7=v7 n9=v9]", "[n10=v10 n11=v11]"})
		}
	})
	if len(mtss) != 2 || !assert.Equal(t, mtss, []tag.Set{
		{
			{
				Name: "n7",
				Value: "v7",
			},
			{
				Name: "n9",
				Value: "v9",
			},
		},
		{
			{
				Name: "n10",
				Value: "v10",
			},
			{
				Name: "n11",
				Value: "v11",
			},
		},
	}) {
		t.Fatalf("Expected ss to be overwritten with '[n7=v7 n9=v9], [n10=v10 n11=v11]', but got: %v", mtss)
	}
}

func TestMongoTagSetSCalledTwice(t *testing.T) {
	var mtss *[]tag.Set
	f := NewFlagSet("test", ContinueOnError)
	mtss = f.MongoTagSetSliceP( "mtss", "m", []tag.Set{}, "Command separated list!")

	in := []string{"[n7=v7 n9=v9],[n1=v1 n2=v2]", "[n3=v3]"}
	expected := []tag.Set{
		{
			tag.Tag{
				Name:  "n7",
				Value: "v7",
			},
			tag.Tag{
				Name:  "n9",
				Value: "v9",
			},
		},
		{
			tag.Tag{
				Name:  "n1",
				Value: "v1",
			},
			tag.Tag{
				Name:  "n2",
				Value: "v2",
			},
		},
		{
			tag.Tag{
				Name: "n3",
				Value: "v3",
			},
		},
	}
	argfmt := "--mtss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.fs.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range *mtss {
		if !assert.Equal(t, expected[i], v ) {
			t.Fatalf("expected mtss[%d] to be %v but got: %v", i, expected[i], v)
		}
	}
}

func TestMongoTagSetSAppend(t *testing.T)  {
	var mtss *[]tag.Set
	f := NewFlagSet("test", ContinueOnError)
	mtss = f.MongoTagSetSlice( "mtss", []tag.Set{}, "Command separated list!")

	in := []string{"[n1=v1 n2=v2 n3=v3]", "[n4=v4 n5=v5 n6=v6]"}
	argfmt := "--mtss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.fs.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	f.fs.VisitAll(func(f *pflag.Flag) {
		if val, ok := f.Value.(pflag.SliceValue); ok {
			_ = val.Append("[n7=v7 n9=v9], [n10=v10 n11=v11]")
		}
	})
	if len(*mtss) != 4 || !assert.Equal(t, *mtss, []tag.Set{
		{
			{
				Name: "n1",
				Value: "v1",
			},
			{
				Name: "n2",
				Value: "v2",
			},
			{
				Name: "n3",
				Value: "v3",
			},
		},
		{
			{
				Name: "n4",
				Value: "v4",
			},
			{
				Name: "n5",
				Value: "v5",
			},
			{
				Name: "n6",
				Value: "v6",
			},
		},
		{
			{
				Name: "n7",
				Value: "v7",
			},
			{
				Name: "n9",
				Value: "v9",
			},
		},
		{
			{
				Name: "n10",
				Value: "v10",
			},
			{
				Name: "n11",
				Value: "v11",
			},
		},
	}) {
		t.Fatalf("Expected ss to be append: '[n1=v1,n2=v2,n3=v3 n4=v4,n5=v5,n6=v6 n7=v7,n9=v9 n10=v10,n11=v11]', but got: %v", mtss)
	}
}

func TestMongoTagSetSGetSlice(t *testing.T)  {
	var mtss []tag.Set
	f := NewFlagSet("test", ContinueOnError)
	f.MongoTagSetSliceVarP(&mtss, "mtss", "m", []tag.Set{}, "Command separated list!")

	in := []string{"[n1=v1 n2=v2 n3=v3]", "[n4=v4 n5=v5 n6=v6]"}
	argfmt := "--mtss=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.fs.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	f.fs.VisitAll(func(f *pflag.Flag) {
		if val, ok := f.Value.(pflag.SliceValue); ok {
			slice := val.GetSlice()
			assert.Equal(t, []string{
				"[n1=v1 n2=v2 n3=v3]",
				"[n4=v4 n5=v5 n6=v6]",
			}, slice)
		}
	})
}