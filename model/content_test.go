package model

import (
	"testing"
)

func TestCheckJson(t *testing.T) {

	json_str := `{"Type":"map","Collection":"abc","Content":{"key1": "value", "key2": {"skey1": 1, "skey2": 2}}}`
	jbytes := []byte(json_str)
	resp := &ApiRequestResponse{}

	jsonObj := checkJson(jbytes, resp)

	if resp != nil && resp.Error {
		t.Fatal("checkJson failed to decode Json")
	}

	if jsonObj.Collection != "abc" {
		t.Fatal("Failed to read Content")
	}
}
