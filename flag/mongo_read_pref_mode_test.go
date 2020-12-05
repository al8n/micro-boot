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
