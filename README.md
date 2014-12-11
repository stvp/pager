pager
=====

`pager` is a Golang package for triggering [PagerDuty][pagerduty] incidents.

[Go API Documentation][godocs]

Example
-------

```go
package main

import (
  "github.com/stvp/pager"
)

func main() {
  // Global useage
  pager.ServiceKey = "3961B1F4AD08424C9DA704DEBCBBF8F3"
  incidentKey, err := pager.Trigger("Everything is on fire.")

  // Individual endpoints
  opsPager := pager.New("09D0A4B9B3F54047BCD7B65704A58333")
  incidentKey, err = opsPager.Trigger("Server out of memory.")
}
```

[pagerduty]: http://pagerduty.com
[godocs]: http://godoc.org/github.com/stvp/pager

