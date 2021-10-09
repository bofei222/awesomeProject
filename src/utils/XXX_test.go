package fileutil

import (
	"fmt"
	"testing"
)

func TestReadPropertiesFile(t *testing.T) {
	props, err := ReadPropertiesFile("sample_test.properties")
	if err != nil {
		t.Error("Error while reading properties file")
	}

	fmt.Println(props)
	if props["host"] != "localhost" || props["proxyHost"] != "test" || props["protocol"] != "https://" || props["chunk"] != "" {
		t.Error("Error properties not loaded correctly")
	}

}
