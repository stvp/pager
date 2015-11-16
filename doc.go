// Package pager provides PagerDuty incident triggers.
//
// Global usage:
//
//     pager.ServiceKey = "3961B1F4AD08424C9DA704DEBCBBF8F3"
//     incidentKey, err := pager.Trigger("Everything is on fire.")
//
// Individual services:
//
//     opsPager := pager.New("09D0A4B9B3F54047BCD7B65704A58333")
//     incidentKey, err = opsPager.Trigger("Server out of memory.")
//
// Including extra details:
//
//     pager.TriggerWithDetails("Oh no", map[string]interface{}{
//         "cause": "it's a mystery",
//         "responsible": "not me!",
//    })
//
//
package pager
