package monitoring

import "strings"
import "go.opencensus.io/stats/view"

// StackdriverPrefix ...
const StackdriverPrefix = "stackdriver_"

// UseStackdriverStandardMetrics ...
func UseStackdriverStandardMetrics() bool {
	// this should be based on some env var or flag or whatever...
	return true
}

// GetMetricType ...
func GetMetricType(view *view.View) string {
	if UseStackdriverStandardMetrics() {
		if strings.HasPrefix(view.Name, StackdriverPrefix) {
			return "istio.io/control/" + strings.TrimPrefix(view.Name, StackdriverPrefix)
		}
	}
	return "custom.googleapis.com/istio_control/" + view.Name
}
