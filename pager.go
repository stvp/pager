package pager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const endpoint = "https://events.pagerduty.com/generic/2010-04-15/create_event.json"

// ServiceKey is the integration key used for the global pager client. This
// comes from the "integration key" that is generated when you create a generic
// API integration in PagerDuty.
var ServiceKey = ""

// Pager is a PagerDuty client configured to trigger and resolve pager alerts
// for a service.
type Pager struct {
	ServiceKey string
}

// New returns a new PagerDuty client for the given integration.
func New(serviceKey string) *Pager {
	return &Pager{serviceKey}
}

// -- Trigger

// Trigger creates a new PagerDuty incident using the default client with the
// given description. The returned incident key can be used to resolve the
// incident. It can also be used by the "TriggerIncidentKey*" functions to
// trigger an incident only if that specific incident has been resolved.
func Trigger(description string) (incidentKey string, err error) {
	return trigger(description, "", map[string]interface{}{})
}

// TriggerIncidentKey triggers an incident using the default client with a
// given incident key only if that incident has been resolved or if that
// incident doesn't exist yet.
func TriggerIncidentKey(description string, key string) (incidentKey string, err error) {
	return trigger(description, key, map[string]interface{}{})
}

// TriggerWithDetails triggers an incident using the default client with a
// description string and a key-value map that will be saved as the incident's
// "details".
func TriggerWithDetails(description string, details map[string]interface{}) (incidentKey string, err error) {
	return trigger(description, "", details)
}

// TriggerIncidentKeyWithDetails triggers an incident using the default client
// with a given incident key only if that incident has been resolved or if that
// incident doesn't exist yet.
func TriggerIncidentKeyWithDetails(description string, key string, details map[string]interface{}) (incidentKey string, err error) {
	return trigger(description, key, details)
}

func trigger(description string, key string, details map[string]interface{}) (newIncidentKey string, err error) {
	p := Pager{ServiceKey}
	return p.trigger(description, key, details)
}

// Trigger creates a new PagerDuty incident with the given description. The
// returned incident key can be used to resolve the incident. It can also be
// used by the "TriggerIncidentKey*" functions to trigger an incident only if
// that specific incident has been resolved.
func (p *Pager) Trigger(description string) (incidentKey string, err error) {
	return p.trigger(description, "", map[string]interface{}{})
}

// TriggerIncidentKey triggers an incident with a given incident key only if
// that incident has been resolved or if that incident doesn't exist yet.
func (p *Pager) TriggerIncidentKey(description string, key string) (incidentKey string, err error) {
	return p.trigger(description, key, map[string]interface{}{})
}

// TriggerWithDetails triggers an incident with a description string and a
// key-value map that will be saved as the incident's "details".
func (p *Pager) TriggerWithDetails(description string, details map[string]interface{}) (incidentKey string, err error) {
	return p.trigger(description, "", details)
}

// TriggerIncidentKeyWithDetails triggers an incident with a given incident key
// only if that incident has been resolved or if that incident doesn't exist
// yet.
func (p *Pager) TriggerIncidentKeyWithDetails(description string, key string, details map[string]interface{}) (incidentKey string, err error) {
	return p.trigger(description, key, details)
}

func (p *Pager) trigger(description string, incidentKey string, details map[string]interface{}) (newIncidentKey string, err error) {
	payload := map[string]interface{}{
		"service_key": p.ServiceKey,
		"event_type":  "trigger",
		"description": description,
		"details":     details,
	}
	if len(incidentKey) > 0 {
		payload["incident_key"] = incidentKey
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	respBody, err := post(reqBody)
	if err != nil {
		return "", err
	}

	return respBody["incident_key"], nil
}

// -- Resolve

// ResolveIncidentKey resolves a triggered PagerDuty incident using the default
// client.
func ResolveIncidentKey(incidentKey string) error {
	return resolve(incidentKey)
}

func resolve(incidentKey string) error {
	p := Pager{ServiceKey}
	return p.resolve(incidentKey)
}

// ResolveIncidentKey resolves a triggered PagerDuty incident.
func (p *Pager) ResolveIncidentKey(incidentKey string) error {
	return p.resolve(incidentKey)
}

func (p *Pager) resolve(incidentKey string) error {
	payload := map[string]interface{}{
		"service_key":  p.ServiceKey,
		"event_type":   "resolve",
		"incident_key": incidentKey,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = post(reqBody)
	return err
}

// -- Utils

// post sends a POST request with the given request body to PagerDuty. See this
// URL for details about PagerDuty API response codes, etc:
// https://developer.pagerduty.com/documentation/integration/events
func post(reqBody []byte) (respBody map[string]string, err error) {
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, string(bodyBytes))
	}

	respBody = map[string]string{}
	err = json.Unmarshal(bodyBytes, &respBody)
	return respBody, err
}
