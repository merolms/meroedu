package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validConfig = []byte(`
debug: true
server:
  address: ":9090"
context:
  timeout: 2
database:
  params:
  parseTime: "true"
  host: localhost
`)

var invalidConfig = []byte(`
debug: true
server: address: ":9090"
context:
  timeout: 2
database:
  params:
  parseTime: "true"
  host: localhost
`)

func createTemporaryConfig(t *testing.T) *os.File {
	f, err := ioutil.TempFile("", "meroedu.yml")
	if err != nil {
		t.Fatalf("unable to create temporary config: %v", err)
	}
	return f
}
func removeTemporaryConfig(t *testing.T, f *os.File) {
	err := f.Close()
	if err != nil {
		t.Fatalf("unable to remove temporary config: %v", err)
	}
}
func TestLoadConfigSuccess(t *testing.T) {

	f := createTemporaryConfig(t)
	defer removeTemporaryConfig(t, f)
	_, err := f.Write(validConfig)
	if err != nil {
		t.Fatalf("error writing config to temporary file: %v", err)
	}
	// Load the valid config
	err = ReadConfig(f.Name())
	if err != nil {
		t.Fatalf("error loading config from temporary file: %v", err)
	}

	// Load an invalid config
	err = ReadConfig("bogusfile")
	if err == nil {
		t.Fatalf("expected error when loading invalid config, but got %v", err)
	}
}

func TestLoadConfigError(t *testing.T) {

	f := createTemporaryConfig(t)
	defer removeTemporaryConfig(t, f)
	_, err := f.Write(invalidConfig)
	if err != nil {
		t.Fatalf("error writing config to temporary file: %v", err)
	}
	// Load the valid config
	err = ReadConfig(f.Name())

	assert.Error(t, err)

	// Load an invalid config
	err = ReadConfig("bogusfile")
	if err == nil {
		t.Fatalf("expected error when loading invalid config, but got %v", err)
	}
}
