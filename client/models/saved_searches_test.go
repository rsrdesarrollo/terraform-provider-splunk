package models

import (
	"encoding/json"
	"testing"
)

func TestSavedSearchObjectUnmarshalNormalizesBoolStrings(t *testing.T) {
	payload := `{"action.httpalert.param.verify_ssl_certificate":"1","disabled":"0"}`

	var obj SavedSearchObject
	if err := json.Unmarshal([]byte(payload), &obj); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if !obj.ActionHttpalertParamVerifySslCertificate {
		t.Fatalf("expected ActionHttpalertParamVerifySslCertificate to be true")
	}

	if obj.Disabled {
		t.Fatalf("expected Disabled to be false")
	}
}

func TestSavedSearchObjectUnmarshalNormalizesBoolNumbers(t *testing.T) {
	payload := `{"action.httpalert.param.verify_ssl_certificate":1,"disabled":0}`

	var obj SavedSearchObject
	if err := json.Unmarshal([]byte(payload), &obj); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if !obj.ActionHttpalertParamVerifySslCertificate {
		t.Fatalf("expected ActionHttpalertParamVerifySslCertificate to be true")
	}

	if obj.Disabled {
		t.Fatalf("expected Disabled to be false")
	}
}
