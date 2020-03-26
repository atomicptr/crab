package crawl

import (
	"strings"
	"testing"
)

func TestValidateKeyValueStringsWithValidInputs(t *testing.T) {
	kvSet := []string{
		"a=b",
		"b=5",
		"test_name=value-that-is-kinda-long",
	}

	err := validateKeyValueStrings("test", kvSet)

	if err != nil {
		t.Error(err)
	}
}

func TestValidateKeyValueStringsWithInvalidInputs(t *testing.T) {
	kvSet := []string{
		"a=b",
		"b=5",
		"lorem ipsum dolor sit amet",
	}

	err := validateKeyValueStrings("test", kvSet)

	if err == nil {
		t.Fail()
	}
}

func TestCreateMapFromKeyValueStrings(t *testing.T) {
	kvSet := []string{
		"a=b",
		"b=5",
		"test_name=value-that-is-kinda-long",
	}

	kvMap := createMapFromKeyValueStrings(kvSet)

	for _, kvPair := range kvSet {
		parts := strings.Split(kvPair, "=")

		if kvMap[parts[0]] != parts[1] {
			t.Errorf("Expected '%s' to be '%s'", parts[0], parts[1])
		}
	}
}
