package orcid

import (
	"encoding/json"
	"os"
	"testing"
)

func TestUnmarshalOrcidRecord(t *testing.T) {
	// Read the sample JSON file
	data, err := os.ReadFile("../orcid-record.json")
	if err != nil {
		t.Fatalf("Failed to read sample JSON file: %v", err)
	}

	// Unmarshal the JSON
	var record Record
	if err := json.Unmarshal(data, &record); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Basic validation
	if record.OrcidIdentifier == nil {
		t.Error("Expected OrcidIdentifier to be populated")
	} else {
		if record.OrcidIdentifier.Path != "0000-0003-1401-2056" {
			t.Errorf("Expected ORCID path to be 0000-0003-1401-2056, got %s", record.OrcidIdentifier.Path)
		}
	}

	if record.Person == nil {
		t.Error("Expected Person to be populated")
	} else {
		if record.Person.Name == nil {
			t.Error("Expected Person.Name to be populated")
		} else {
			if record.Person.Name.GivenNames == nil || record.Person.Name.GivenNames.Value != "Michael" {
				t.Error("Expected given name to be Michael")
			}
			if record.Person.Name.FamilyName == nil || record.Person.Name.FamilyName.Value != "Thicke" {
				t.Error("Expected family name to be Thicke")
			}
		}
	}

	if record.History == nil {
		t.Error("Expected History to be populated")
	} else {
		if record.History.CreationMethod != "MEMBER_REFERRED" {
			t.Errorf("Expected creation method to be MEMBER_REFERRED, got %s", record.History.CreationMethod)
		}
	}
}
