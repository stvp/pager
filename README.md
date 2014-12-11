pager
=====

`pager` is a Golang package for triggering [PagerDuty][pagerduty] incidents.

API
---

[Go API Documentation][godocs]

    Trigger(serviceKey string, description string) (incidentKey string, err error)
    TriggerIncidentKey(serviceKey string, description, key string) (incidentKey string, err error)
    TriggerWithDetails(serviceKey string, description string, details map[string]interface{}) (incidentKey string, err error)
    TriggerIncidentKeyWithDetails(serviceKey string, description, key string, details map[string]interface{}) (incidentKey string, err error)

Example
-------

```go
package main

import (
  "github.com/stvp/pager"
)

func main() {
  // ...
  incidentKey, err := pager.Trigger("a0d9345d0b041d12d702fa8c0cfe6516", "Everything is on fire.")
}
```

[pagerduty]: http://pagerduty.com
[godocs]: http://godoc.org/github.com/stvp/pager

