pager
=====

`pager` is a Golang package for triggering [PagerDuty][pagerduty] incidents.

API
---

[Go API Documentation][godocs]

    Trigger(description string) (incidentKey string, err error)
    TriggerIncidentKey(description, key string) (incidentKey string, err error)
    TriggerWithDetails(description string, details map[string]interface{}) (incidentKey string, err error)
    TriggerIncidentKeyWithDetails(description string, details map[string]interface{}) (incidentKey string, err error)

Example
-------

```go
package main

import (
  "github.com/stvp/pager"
)

func main() {
  pager.ServiceKey = "a0d9345d0b041d12d702fa8c0cfe6516"
  // ...
  incidentKey, err := pager.Trigger("Everything is on fire.")
}
```

[pagerduty]: http://pagerduty.com
[godocs]: http://godoc.org/github.com/stvp/pager

