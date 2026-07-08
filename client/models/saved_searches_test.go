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

func TestSavedSearchObjectUnmarshalKeepsStringAndNumericStringFields(t *testing.T) {
	// Splunk returns everything as strings, including numeric fields without a
	// ",string" json option (e.g. alert.severity). The decoder must still
	// populate the rest of the struct instead of bailing out on those.
	payload := `{
		"search":"index=main",
		"disabled":"1",
		"action.httpalert.param.verify_ssl_certificate":"1",
		"alert_threshold":"1",
		"alert.severity":"4",
		"action.email.sendcsv":"1"
	}`

	var obj SavedSearchObject
	if err := json.Unmarshal([]byte(payload), &obj); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	// String fields must be preserved (regression: whole struct came back empty).
	if obj.Search != "index=main" {
		t.Fatalf("expected Search to be preserved, got %q", obj.Search)
	}

	// Non-bool field with "1" stays a string, untouched by bool normalization.
	if obj.AlertThreshold != "1" {
		t.Fatalf("expected AlertThreshold to stay \"1\", got %q", obj.AlertThreshold)
	}

	// Bool fields are normalized.
	if !obj.Disabled {
		t.Fatalf("expected Disabled to be true")
	}
	if !obj.ActionHttpalertParamVerifySslCertificate {
		t.Fatalf("expected ActionHttpalertParamVerifySslCertificate to be true")
	}

	// Int field with ",string" option still decodes correctly.
	if obj.ActionEmailSendCSV != 1 {
		t.Fatalf("expected ActionEmailSendCSV to be 1, got %d", obj.ActionEmailSendCSV)
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
