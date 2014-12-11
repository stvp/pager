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

func (p *Pager) trigger(description, key string, details map[string]interface{}) (incidentKey string, err error) {
	if len(p.ServiceKey) == 0 {
		return "", fmt.Errorf("ServiceKey is not set")
	}

	payload := map[string]interface{}{
		"service_key": p.ServiceKey,
		"event_type":  "trigger",
		"description": description,
		"details":     details,
	}
	if len(key) > 0 {
		payload["incident_key"] = key
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(p.Endpoint, "application/json", bytes.NewReader(jsonPayload))
	defer resp.Body.Close()

	if err = errorFromResponse(resp); err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("PagerDuty request was successful but an error occurred while reading the response body: %s", err.Error())
	}

	respBody := map[string]string{}
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return "", fmt.Errorf("PagerDuty request was successful but an error occurred while parsing the response body JSON: %s", err.Error())
	}

	return respBody["incident_key"], nil
}

// errorFromResponse returns an error with a helpful message if the given
// PagerDuty response is an error response.
func errorFromResponse(resp *http.Response) (err error) {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("PagerDuty request failed (%s) and an error occurred while reading the response body: %s", resp.Status, err.Error())
	}

	return fmt.Errorf("PagerDuty request failed (%s): %s", resp.Status, string(bodyBytes))
}
