package flag

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

func setUpReadPreferenceModeFlagSet(mode *readpref.Mode) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.MongoReadPreferenceModeVar(mode, "mode", 0, "Command read preference mode!")
	return f
}

func setUpReadPreferenceModeWithDefaultNearestMode(mode *readpref.Mode) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.MongoReadPreferenceModeVar(mode, "mode", readpref.NearestMode, "Command read preference mode!")
	return f
}

func TestEmptyReadPreferenceMode(t *testing.T) {
	var mode readpref.Mode
	f := setUpReadPreferenceModeFlagSet(&mode)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	_, err = f.GetMongoReadPreferenceMode("mode")
	if assert.Error(t, err) {
		assert.Equal(t, errMongoReadPreferenceMode, err)
	}
}

func TestMongoReadPreferenceMode(t *testing.T) {
	var mode readpref.Mode
	f := setUpReadPreferenceModeFlagSet(&mode)

	arg := fmt.Sprintf("--mode=primary")
	err := f.fs.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getMode, err := f.GetMongoReadPreferenceMode("mode")
	assert.NoError(t, err)
	assert.Equal(t, getMode, readpref.PrimaryMode)
}

func TestMongoReadPreferenceModeWithDefault(t *testing.T) {
	var mode readpref.Mode
	f := setUpReadPreferenceModeWithDefaultNearestMode(&mode)

	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getMode, err := f.GetMongoReadPreferenceMode("mode")
	assert.NoError(t, err)
	assert.Equal(t, getMode, readpref.NearestMode)
}

func TestMongoReadPreferenceModePrimaryPreferred(t *testing.T) {
	var mode *readpref.Mode
	f := NewFlagSet("test", ContinueOnError)
	mode = f.MongoReadPreferenceMode("mode", readpref.PrimaryPreferredMode, "")

	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getMode, err := f.GetMongoReadPreferenceMode("mode")
	assert.NoError(t, err)
	assert.Equal(t, getMode, *mode)
}

func TestMongoReadPreferenceModeSecondary(t *testing.T) {
	var mode *readpref.Mode
	f := NewFlagSet("test", ContinueOnError)
	mode = f.MongoReadPreferenceModeP("mode", "m", readpref.PrimaryPreferredMode, "")

	err := f.fs.Parse([]string{"-m secondary"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getMode, err := f.GetMongoReadPreferenceMode("mode")
	assert.NoError(t, err)
	assert.Equal(t, getMode, *mode)
}

func TestMongoReadPreferenceModeSecondaryPreferredMode(t *testing.T) {
	var mode *readpref.Mode
	f := NewFlagSet("test", ContinueOnError)
	mode = f.MongoReadPreferenceModeP("mode", "m", readpref.PrimaryPreferredMode, "")

	err := f.fs.Parse([]string{"-m secondary-preferred"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getMode, err := f.GetMongoReadPreferenceMode("mode")
	assert.NoError(t, err)
	assert.Equal(t, getMode, *mode)
}

func TestMongoReadPreferenceModeErrorMode(t *testing.T)  {
	var mode readpref.Mode
	f := NewFlagSet("test", ContinueOnError)
	f.MongoReadPreferenceModeVarP(&mode,"mode", "m", readpref.PrimaryPreferredMode, "")

	err := f.fs.Parse([]string{"-m error"})
	assert.Error(t, err)
}