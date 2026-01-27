//go:build encore_app

package reqtrack

import (
	"encore.dev/appruntime/shared/appconf"
	"encore.dev/appruntime/shared/logging"
	"encore.dev/appruntime/shared/platform"
	"encore.dev/appruntime/shared/traceprovider"
)

var Singleton *RequestTracker

func init() {
	var traceFactory traceprovider.Factory
	tracingEnabled := appconf.Runtime.TraceEndpoint != ""
	if tracingEnabled {
		logging.RootLogger.Info().Str("endpoint", appconf.Runtime.TraceEndpoint).Msg("tracing enabled")
		traceFactory = &traceprovider.DefaultFactory{
			SampleRate: appconf.Runtime.TraceSamplingRate,
		}
	}

	Singleton = New(logging.RootLogger, platform.Singleton, traceFactory)
}
