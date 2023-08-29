package otlog

// Log fields required by OpenTelemetry that don't exist in slog

type OTRecord struct {
	SpanId [16]byte		`json:"spanId,omitempty"`
	TraceId [16]byte	`json:"traceId,omitempty"`
	Service string		`json:"service.name"`
	telemetrySDK TelemetrySDK
}

type TelemetrySDK struct {
	TelemetrySDKname string		`json:"telemetry.language.name"`
	TelemetrySDKlanguage string	`json:"telemetry.language.language"`
	TelemetrySDKversion string	`json:"telemetry.language.version"`
}

func newTelemetrySDK() *TelemetrySDK {
	tsdk := TelemetrySDK{
		TelemetrySDKname: "otlog",
		TelemetrySDKlanguage: "go",
		TelemetrySDKversion: moduleVersion,
	}

	return &tsdk
}