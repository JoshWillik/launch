package config

import "testing"

// TestRead checks to see if project config is read correctly in all cases
func TestRead(t *testing.T) {
	goodConfig := Read("./test/good-project")
	if goodConfig.ContextRoot != "something/else" {
		t.Log("good project config is incorrect: " + goodConfig.ContextRoot)
		t.Fail()
	}
	badConfig := Read("./test/bad-project")
	if badConfig.ContextRoot != "./test/bad-project" {
		t.Log("bad project config is incorrect: " + badConfig.ContextRoot)
		t.Fail()
	}
	missingConfig := Read("./test/missing-project")
	if missingConfig.ContextRoot != "./test/missing-project" {
		t.Log("missing project config is incorrect: " + missingConfig.ContextRoot)
		t.Fail()
	}
}
