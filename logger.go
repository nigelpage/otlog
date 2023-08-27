package otlog

import (
	"encoding/json"
	"fmt"

	"github.com/gookit/slog"
)

/*
 The data-model for an OpenTelemetry log record can be found at
 https://opentelemetry.io/docs/reference/specification/logs/data-model/
*/

// N.B. slog uses the term 'Level' to describe the importance of a log message, whilst OpenTelemetry uses the term 'Severity'
type Severity slog.Level

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

// OpenTelemetry logger

type OTLogger struct {
	logger slog.Logger
}

const (
	TraceSeverity  Severity = 1
	Trace2Severity Severity = 2
	Trace3Severity Severity = 3
	Trace4Severity Severity = 4
	DebugSeverity  Severity = 5
	Debug2Severity Severity = 6
	Debug3Severity Severity = 7
	Debug4Severity Severity = 8
	InfoSeverity   Severity = 9
	Info2Severity  Severity = 10
	Info3Severity  Severity = 11
	Info4Severity  Severity = 12
	WarnSeverity   Severity = 13
	Warn2Severity  Severity = 14
	Warn3Severity  Severity = 15
	Warn4Severity  Severity = 16
	ErrorSeverity  Severity = 17
	Error2Severity Severity = 18
	Error3Severity Severity = 19
	Error4Severity Severity = 20
	FatalSeverity  Severity = 21
	Fatal2Severity Severity = 22
	Fatal3Severity Severity = 23
	Fatal4Severity Severity = 24
)

var severityToStr = map[Severity]string{
	TraceSeverity:  "TRACE",
	Trace2Severity: "TRACE2",
	Trace3Severity: "TRACE3",
	Trace4Severity: "TRACE4",
	DebugSeverity:  "DEBUG",
	Debug2Severity: "DEBUG2",
	Debug3Severity: "DEBUG3",
	Debug4Severity: "DEBUG4",
	InfoSeverity:   "INFO",
	Info2Severity:  "INFO2",
	Info3Severity:  "INFO3",
	Info4Severity:  "INFO4",
	WarnSeverity:   "WARN",
	Warn2Severity:  "WARN2",
	Warn3Severity:  "WARN3",
	Warn4Severity:  "WARN4",
	ErrorSeverity:  "ERROR",
	Error2Severity: "ERROR2",
	Error3Severity: "ERROR3",
	Error4Severity: "ERROR4",
	FatalSeverity:  "FATAL",
	Fatal2Severity: "FATAL2",
	Fatal3Severity: "FATAL3",
	Fatal4Severity: "FATAL4",
}

var strToSeverity = map[string]Severity{
	`"TRACE"`:  TraceSeverity,
	`"TRACE2"`: Trace2Severity,
	`"TRACE3"`: Trace3Severity,
	`"TRACE4"`: Trace4Severity,
	`"DEBUG"`:  DebugSeverity,
	`"DEBUG2"`: Debug2Severity,
	`"DEBUG3"`: Debug3Severity,
	`"DEBUG4"`: Debug4Severity,
	`"INFO"`:   InfoSeverity,
	`"INFO2"`:  Info2Severity,
	`"INFO3"`:  Info3Severity,
	`"INFO4"`:  Info4Severity,
	`"WARN"`:   WarnSeverity,
	`"WARN2"`:  Warn2Severity,
	`"WARN3"`:  Warn3Severity,
	`"WARN4"`:  Warn4Severity,
	`"ERROR"`:  ErrorSeverity,
	`"ERROR2"`: Error2Severity,
	`"ERROR3"`: Error3Severity,
	`"ERROR4"`: Error4Severity,
	`"FATAL"`:  FatalSeverity,
	`"FATAL2"`: Fatal2Severity,
	`"FATAL3"`: Fatal3Severity,
	`"FATAL4"`: Fatal4Severity,
}

/*
	Convenience functions for manipulating Severity strings and values
*/

func SeverityValue(str string) (Severity, error) {
	if val, ok := strToSeverity[str]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("invalid OpenTelemetry log Severity name - %s", str)
}

func (s Severity) String() string {
	if s < TraceSeverity || s > Fatal4Severity {
		return fmt.Sprintf("**INVALID OpenTelemetry log Severity value - %d**", s)
	}
	return severityToStr[Severity(Severity(s))]
}

func (s Severity) IsTraceSeverity() bool {
	if s >= TraceSeverity && s <= Trace4Severity {
		return true
	}

	return false
}

func (s Severity) IsDebugSeverity() bool {
	if s >= DebugSeverity && s < Debug4Severity {
		return true
	}

	return false
}
func (s Severity) IsInfoSeverity() bool {
	if s >= InfoSeverity && s <= Info4Severity {
		return true
	}

	return false
}
func (s Severity) IsWarnSeverity() bool {
	if s >= WarnSeverity && s <= Warn4Severity {
		return true
	}

	return false
}

// Any Severity of 17 or above indicates an erroneous situation and should be handled
func (s Severity) ShouldNotIgnore() bool {
	return s >= ErrorSeverity
}

func (s Severity) IsErrorSeverity() bool {
	if s >= ErrorSeverity && s <= Error4Severity {
		return true
	}

	return false
}
func (s Severity) IsFatalSeverity() bool {
	if s >= FatalSeverity && s <= Fatal4Severity {
		return true
	}

	return false
}

// Convenience functions for logging with OpenTelemetry log Levels
// N.B. When using ensure that the message is the last argument

func (l *OTLogger) Trace(args ...any)  { l.logger.Log(slog.Level(TraceSeverity), args) }
func (l *OTLogger) Trace2(args ...any) { l.logger.Log(slog.Level(Trace2Severity), args) }
func (l *OTLogger) Trace3(args ...any) { l.logger.Log(slog.Level(Trace3Severity), args) }
func (l *OTLogger) Trace4(args ...any) { l.logger.Log(slog.Level(Trace4Severity), args) }
func (l *OTLogger) Tracef(format string, args ...any) {
	l.logger.Logf(slog.Level(TraceSeverity), format, args)
}
func (l *OTLogger) Tracef2(format string, args ...any) {
	l.logger.Logf(slog.Level(Trace2Severity), format, args)
}
func (l *OTLogger) Tracef3(format string, args ...any) {
	l.logger.Logf(slog.Level(Trace3Severity), format, args)
}
func (l *OTLogger) Tracef4(format string, args ...any) {
	l.logger.Logf(slog.Level(Trace4Severity), format, args)
}

func (l *OTLogger) Debug(args ...any)  { l.logger.Log(slog.Level(DebugSeverity), args) }
func (l *OTLogger) Debug2(args ...any) { l.logger.Log(slog.Level(Debug2Severity), args) }
func (l *OTLogger) Debug3(args ...any) { l.logger.Log(slog.Level(Debug3Severity), args) }
func (l *OTLogger) Debug4(args ...any) { l.logger.Log(slog.Level(Debug4Severity), args) }
func (l *OTLogger) Debugf(format string, args ...any) {
	l.logger.Logf(slog.Level(DebugSeverity), format, args)
}
func (l *OTLogger) Debugf2(format string, args ...any) {
	l.logger.Logf(slog.Level(Debug2Severity), format, args)
}
func (l *OTLogger) Debugf3(format string, args ...any) {
	l.logger.Logf(slog.Level(Debug3Severity), format, args)
}
func (l *OTLogger) Debugf4(format string, args ...any) {
	l.logger.Logf(slog.Level(Debug4Severity), format, args)
}

func (l *OTLogger) Info(args ...any)  { l.logger.Log(slog.Level(InfoSeverity), args) }
func (l *OTLogger) Info2(args ...any) { l.logger.Log(slog.Level(Info2Severity), args) }
func (l *OTLogger) Info3(args ...any) { l.logger.Log(slog.Level(Info3Severity), args) }
func (l *OTLogger) Info4(args ...any) { l.logger.Log(slog.Level(Info4Severity), args) }
func (l *OTLogger) Infof(format string, args ...any) {
	l.logger.Logf(slog.Level(InfoSeverity), format, args)
}
func (l *OTLogger) Infof2(format string, args ...any) {
	l.logger.Logf(slog.Level(Info2Severity), format, args)
}
func (l *OTLogger) Infof3(format string, args ...any) {
	l.logger.Logf(slog.Level(Info3Severity), format, args)
}
func (l *OTLogger) Infof4(format string, args ...any) {
	l.logger.Logf(slog.Level(Info4Severity), format, args)
}

func (l *OTLogger) Warn(args ...any)  { l.logger.Log(slog.Level(WarnSeverity), args) }
func (l *OTLogger) Warn2(args ...any) { l.logger.Log(slog.Level(Warn2Severity), args) }
func (l *OTLogger) Warn3(args ...any) { l.logger.Log(slog.Level(Warn3Severity), args) }
func (l *OTLogger) Warn4(args ...any) { l.logger.Log(slog.Level(Warn4Severity), args) }
func (l *OTLogger) Warnf(format string, args ...any) {
	l.logger.Logf(slog.Level(WarnSeverity), format, args)
}
func (l *OTLogger) Warnf2(format string, args ...any) {
	l.logger.Logf(slog.Level(Warn2Severity), format, args)
}
func (l *OTLogger) Warnf3(format string, args ...any) {
	l.logger.Logf(slog.Level(Warn3Severity), format, args)
}
func (l *OTLogger) Warnf4(format string, args ...any) {
	l.logger.Logf(slog.Level(Warn4Severity), format, args)
}

func (l *OTLogger) Error(args ...any)  { l.logger.Log(slog.Level(ErrorSeverity), args) }
func (l *OTLogger) Error2(args ...any) { l.logger.Log(slog.Level(Error2Severity), args) }
func (l *OTLogger) Error3(args ...any) { l.logger.Log(slog.Level(Error3Severity), args) }
func (l *OTLogger) Error4(args ...any) { l.logger.Log(slog.Level(Error4Severity), args) }
func (l *OTLogger) Errorf(format string, args ...any) {
	l.logger.Logf(slog.Level(ErrorSeverity), format, args)
}
func (l *OTLogger) Errorf2(format string, args ...any) {
	l.logger.Logf(slog.Level(Error2Severity), format, args)
}
func (l *OTLogger) Errorf3(format string, args ...any) {
	l.logger.Logf(slog.Level(Error3Severity), format, args)
}
func (l *OTLogger) Errorf4(format string, args ...any) {
	l.logger.Logf(slog.Level(Error4Severity), format, args)
}

func (l *OTLogger) Fatal(args ...any)  { l.logger.Log(slog.Level(FatalSeverity), args) }
func (l *OTLogger) Fatal2(args ...any) { l.logger.Log(slog.Level(Fatal2Severity), args) }
func (l *OTLogger) Fatal3(args ...any) { l.logger.Log(slog.Level(Fatal3Severity), args) }
func (l *OTLogger) Fatal4(args ...any) { l.logger.Log(slog.Level(Fatal4Severity), args) }
func (l *OTLogger) Fatalf(format string, args ...any) {
	l.logger.Logf(slog.Level(FatalSeverity), format, args)
}
func (l *OTLogger) Fatalf2(format string, args ...any) {
	l.logger.Logf(slog.Level(Fatal2Severity), format, args)
}
func (l *OTLogger) Fatalf3(format string, args ...any) {
	l.logger.Logf(slog.Level(Fatal3Severity), format, args)
}
func (l *OTLogger) Fatalf4(format string, args ...any) {
	l.logger.Logf(slog.Level(Fatal4Severity), format, args)
}
