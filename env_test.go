package env

import "testing"

type TestEnv = map[string]string

const randName = "RAND1610118849722"

func TestLoadEnv(t *testing.T) {
	got := LoadEnv()

	if len(got) < 10 {
		t.Errorf("Expect to load os environment variables (they all have 10 variables at least)")
	}

	if got["PKG_NAME"] != "test" {
		t.Errorf("Expect .env file to be loaded")
	}

	if got[randName] != "" {
		t.Errorf("Expect not found env to be empty string")
	}
}

func TestHasEnv(t *testing.T) {
	if HasEnv("PKG_NAME") != true {
		t.Errorf("Expect env to have \"PKG_NAME\" key")
	}

	if HasEnv("PKG_name") != true {
		t.Errorf("Expect env to read uppercase of \"PKG_name\" key")
	}

	if HasEnv(randName) != false {
		t.Errorf("Expect env to not have a random key")
	}
}

func TestGetEnv(t *testing.T) {
	if GetEnv("PKG_NAME") != "test" {
		t.Errorf("Expect env to have \"PKG_NAME\" key set with \"test\"")
	}

	if GetEnv(randName) != "" {
		t.Errorf("Expect GetEnv to return empty string for a random key")
	}
}
