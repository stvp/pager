package pager

import (
	"os"
	"testing"
)

func TestServiceKey(t *testing.T) {
	ServiceKey = ""

	_, err := Trigger("TestServiceKey")
	if err == nil {
		t.Errorf("didn't get an error")
	}
}

func TestTrigger(t *testing.T) {
	ServiceKey = os.Getenv("SERVICE_KEY")

	key, err := Trigger("TestTrigger")
	if err != nil {
		t.Error(err)
	}
	if len(key) == 0 {
		t.Errorf("didn't get an incident key back")
	}

	originalKey := key
	key, err = TriggerIncidentKey("TestTriggerIncidentKey", originalKey)
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if key != originalKey {
		t.Errorf("expected %#v, got %#v", originalKey, key)
	}
}

func TestTriggerWithDetails(t *testing.T) {
	ServiceKey = os.Getenv("SERVICE_KEY")

	details := map[string]interface{}{"testing": true}
	key, err := TriggerWithDetails("TestTriggerWithDetails", details)
	if err != nil {
		t.Error(err)
	}
	if len(key) == 0 {
		t.Errorf("didn't get an incident key back")
	}

	originalKey := key
	key, err = TriggerIncidentKeyWithDetails("TestTriggerIncidentKeyWithDetails", originalKey, details)
	if err != nil {
		t.Error(err)
	}
	if key != originalKey {
		t.Errorf("expected %#v, got %#v", originalKey, key)
	}
}

func TestTriggerIncidentKeyWithDetails(t *testing.T) {
	ServiceKey = os.Getenv("SERVICE_KEY")

	details := map[string]interface{}{"testing": true}
	key, err := TriggerIncidentKeyWithDetails("TestTriggerIncidentKeyWithDetails", "TestKey", details)
	if err != nil {
		t.Error(err)
	}
	if len(key) == 0 {
		t.Errorf("didn't get an incident key back")
	}

	originalKey := key
	key, err = TriggerIncidentKeyWithDetails("TestTriggerIncidentKeyWithDetails", originalKey, details)
	if err != nil {
		t.Error(err)
	}
	if key != originalKey {
		t.Errorf("expected %#v, got %#v", originalKey, key)
	}
}

func TestResolveIncidentKey(t *testing.T) {
	ServiceKey = os.Getenv("SERVICE_KEY")

	key, err := Trigger("TestResolveIncidentKey")
	if err != nil {
		t.Error(err)
	}
	if len(key) == 0 {
		t.Errorf("didn't get an incident key back")
	}

	err = ResolveIncidentKey(key)
	if err != nil {
		t.Error(err)
	}
}
