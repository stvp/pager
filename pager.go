package pager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	Endpoint   = "https://events.pagerduty.com/generic/2010-04-15/create_event.json"
	ServiceKey = ""
)

type Pager struct {
	Endpoint   string
	ServiceKey string
}

func New(serviceKey string) *Pager {
	return &Pager{Endpoint, serviceKey}
}

func Trigger(description string) (incidentKey string, err error) {
	return trigger(description, "", map[string]interface{}{})
}

func TriggerIncidentKey(description, key string) (incidentKey string, err error) {
	return trigger(description, key, map[string]interface{}{})
}

func TriggerWithDetails(description string, details map[string]interface{}) (incidentKey string, err error) {
	return trigger(description, "", details)
}

func TriggerIncidentKeyWithDetails(description, key string, details map[string]interface{}) (incidentKey string, err error) {
	return trigger(description, key, details)
}

func trigger(description, key string, details map[string]interface{}) (incidentKey string, err error) {
	p := Pager{Endpoint, ServiceKey}
	return p.trigger(description, key, details)
}

func (p *Pager) Trigger(description string) (incidentKey string, err error) {
	return p.trigger(description, "", map[string]interface{}{})
}

func (p *Pager) TriggerIncidentKey(description, key string) (incidentKey string, err error) {
	return p.trigger(description, key, map[string]interface{}{})
}

func (p *Pager) TriggerWithDetails(description string, details map[string]interface{}) (incidentKey string, err error) {
	return p.trigger(description, "", details)
}

func (p *Pager) TriggerIncidentKeyWithDetails(description, key string, details map[string]interface{}) (incidentKey string, err error) {
	return p.trigger(description, key, details)
}

func (p *Pager) trigger(description, incidentKey string, details map[string]interface{}) (newIncidentKey string, err error) {
	payload, err := triggerIncidentJSON(p.ServiceKey, description, incidentKey, details)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(p.Endpoint, "application/json", bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = responseError(resp)
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	respBody := map[string]string{}
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return "", err
	}

	return respBody["incident_key"], nil
}

// -- Utils

// triggerIncidentJSON builds the JSON payload for triggering a new PagerDuty
// incident.
func triggerIncidentJSON(serviceKey, description, incidentKey string, details map[string]interface{}) ([]byte, error) {
	payload := map[string]interface{}{
		"service_key": serviceKey,
		"event_type":  "trigger",
		"description": description,
		"details":     details,
	}
	if len(incidentKey) > 0 {
		payload["incident_key"] = incidentKey
	}
	return json.Marshal(payload)
}

// responseError returns an error if the given PagerDuty response is an error
// response.
func responseError(resp *http.Response) (err error) {
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusNoContent {
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf("%s: %s", resp.Status, string(body))
}
