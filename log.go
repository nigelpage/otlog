package otlog

import (
	"fmt"
	"github.com/gookit/slog"
)

/*
 The data-model for an OpenTelemetry log record can be found at
 https://opentelemetry.io/docs/reference/specification/logs/data-model/
*/

type Level slog.Level

const (
	TraceLevel  Level = 1
	Trace2Level Level = 2
	Trace3Level Level = 3
	Trace4Level Level = 4
	DebugLevel  Level = 5
	Debug2Level Level = 6
	Debug3Level Level = 7
	Debug4Level Level = 8
	InfoLevel   Level = 9
	Info2Level  Level = 10
	Info3Level  Level = 11
	Info4Level  Level = 12
	WarnLevel   Level = 13
	Warn2Level  Level = 14
	Warn3Level  Level = 15
	Warn4Level  Level = 16
	ErrorLevel  Level = 17
	Error2Level Level = 18
	Error3Level Level = 19
	Error4Level Level = 20
	FatalLevel  Level = 21
	Fatal2Level Level = 22
	Fatal3Level Level = 23
	Fatal4Level Level = 24
)

var levelToStr = map[Level]string{
	TraceLevel:  "TRACE",
	Trace2Level: "TRACE2",
	Trace3Level: "TRACE3",
	Trace4Level: "TRACE4",
	DebugLevel:  "DEBUG",
	Debug2Level: "DEBUG2",
	Debug3Level: "DEBUG3",
	Debug4Level: "DEBUG4",
	InfoLevel:   "INFO",
	Info2Level:  "INFO2",
	Info3Level:  "INFO3",
	Info4Level:  "INFO4",
	WarnLevel:   "WARN",
	Warn2Level:  "WARN2",
	Warn3Level:  "WARN3",
	Warn4Level:  "WARN4",
	ErrorLevel:  "ERROR",
	Error2Level: "ERROR2",
	Error3Level: "ERROR3",
	Error4Level: "ERROR4",
	FatalLevel:  "FATAL",
	Fatal2Level: "FATAL2",
	Fatal3Level: "FATAL3",
	Fatal4Level: "FATAL4",
}

var strToLevel = map[string]Level{
	`"TRACE"`:  TraceLevel,
	`"TRACE2"`: Trace2Level,
	`"TRACE3"`: Trace3Level,
	`"TRACE4"`: Trace4Level,
	`"DEBUG"`:  DebugLevel,
	`"DEBUG2"`: Debug2Level,
	`"DEBUG3"`: Debug3Level,
	`"DEBUG4"`: Debug4Level,
	`"INFO"`:   InfoLevel,
	`"INFO2"`:  Info2Level,
	`"INFO3"`:  Info3Level,
	`"INFO4"`:  Info4Level,
	`"WARN"`:   WarnLevel,
	`"WARN2"`:  Warn2Level,
	`"WARN3"`:  Warn3Level,
	`"WARN4"`:  Warn4Level,
	`"ERROR"`:  ErrorLevel,
	`"ERROR2"`: Error2Level,
	`"ERROR3"`: Error3Level,
	`"ERROR4"`: Error4Level,
	`"FATAL"`:  FatalLevel,
	`"FATAL2"`: Fatal2Level,
	`"FATAL3"`: Fatal3Level,
	`"FATAL4"`: Fatal4Level,
}

// OpenTelemetry logger

type OTLogger struct {
	logger slog.Logger
}

/*
	Convenience functions for manipulating Level strings and values
*/

func LevelValue(str string) (Level, error) {
	if val, ok := strToLevel[str]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("invalid OpenTelemetry log level name - %s", str)
}

func (l Level) String() string {
	return levelToStr[Level(Level(l))]
}

func (l Level) IsTraceLevel() bool {
	if l >= TraceLevel && l <= Trace4Level {
		return true
	}

	return false
}

func (l Level) IsDebugLevel() bool {
	if l >= DebugLevel && l < Debug4Level {
		return true
	}

	return false
}
func (l Level) IsInfoLevel() bool {
	if l >= InfoLevel && l <= Info4Level {
		return true
	}

	return false
}
func (l Level) IsWarnLevel() bool {
	if l >= WarnLevel && l <= Warn4Level {
		return true
	}

	return false
}

func (l Level) IsErrorLevel() bool {
	if l >= ErrorLevel && l <= Error4Level {
		return true
	}

	return false
}
func (l Level) IsFatalLevel() bool {
	if l >= FatalLevel && l <= Fatal4Level {
		return true
	}

	return false
}

// Convenience functions for logging with OpenTelemetry log levels
// N.B. When using ensure that the message is the last argument

func (l *OTLogger) Trace(args ...any) { l.logger.Log(slog.Level(TraceLevel), args)}
func (l *OTLogger) Trace2(args ...any) { l.logger.Log(slog.Level(Trace2Level), args) }
func (l *OTLogger) Trace3(args ...any) { l.logger.Log(slog.Level(Trace3Level), args) }
func (l *OTLogger) Trace4(args ...any) { l.logger.Log(slog.Level(Trace4Level), args) }
func (l *OTLogger) Tracef(format string, args ...any) { l.logger.Logf(slog.Level(TraceLevel), format, args)}
func (l *OTLogger) Tracef2(format string, args ...any) { l.logger.Logf(slog.Level(Trace2Level), format, args) }
func (l *OTLogger) Tracef3(format string, args ...any) { l.logger.Logf(slog.Level(Trace3Level), format, args) }
func (l *OTLogger) Tracef4(format string, args ...any) { l.logger.Logf(slog.Level(Trace4Level), format, args) }

func (l *OTLogger) Debug(args ...any) { l.logger.Log(slog.Level(DebugLevel), args)}
func (l *OTLogger) Debug2(args ...any) { l.logger.Log(slog.Level(Debug2Level), args) }
func (l *OTLogger) Debug3(args ...any) { l.logger.Log(slog.Level(Debug3Level), args) }
func (l *OTLogger) Debug4(args ...any) { l.logger.Log(slog.Level(Debug4Level), args) }
func (l *OTLogger) Debugf(format string, args ...any) { l.logger.Logf(slog.Level(DebugLevel), format, args)}
func (l *OTLogger) Debugf2(format string, args ...any) { l.logger.Logf(slog.Level(Debug2Level), format, args) }
func (l *OTLogger) Debugf3(format string, args ...any) { l.logger.Logf(slog.Level(Debug3Level), format, args) }
func (l *OTLogger) Debugf4(format string, args ...any) { l.logger.Logf(slog.Level(Debug4Level), format, args) }

func (l *OTLogger) Info(args ...any) { l.logger.Log(slog.Level(InfoLevel), args)}
func (l *OTLogger) Info2(args ...any) { l.logger.Log(slog.Level(Info2Level), args) }
func (l *OTLogger) Info3(args ...any) { l.logger.Log(slog.Level(Info3Level), args) }
func (l *OTLogger) Info4(args ...any) { l.logger.Log(slog.Level(Info4Level), args) }
func (l *OTLogger) Infof(format string, args ...any) { l.logger.Logf(slog.Level(InfoLevel), format, args)}
func (l *OTLogger) Infof2(format string, args ...any) { l.logger.Logf(slog.Level(Info2Level), format, args) }
func (l *OTLogger) Infof3(format string, args ...any) { l.logger.Logf(slog.Level(Info3Level), format, args) }
func (l *OTLogger) Infof4(format string, args ...any) { l.logger.Logf(slog.Level(Info4Level), format, args) }

func (l *OTLogger) Warn(args ...any) { l.logger.Log(slog.Level(WarnLevel), args)}
func (l *OTLogger) Warn2(args ...any) { l.logger.Log(slog.Level(Warn2Level), args) }
func (l *OTLogger) Warn3(args ...any) { l.logger.Log(slog.Level(Warn3Level), args) }
func (l *OTLogger) Warn4(args ...any) { l.logger.Log(slog.Level(Warn4Level), args) }
func (l *OTLogger) Warnf(format string, args ...any) { l.logger.Logf(slog.Level(WarnLevel), format, args)}
func (l *OTLogger) Warnf2(format string, args ...any) { l.logger.Logf(slog.Level(Warn2Level), format, args) }
func (l *OTLogger) Warnf3(format string, args ...any) { l.logger.Logf(slog.Level(Warn3Level), format, args) }
func (l *OTLogger) Warnf4(format string, args ...any) { l.logger.Logf(slog.Level(Warn4Level), format, args) }

func (l *OTLogger) Error(args ...any) { l.logger.Log(slog.Level(ErrorLevel), args)}
func (l *OTLogger) Error2(args ...any) { l.logger.Log(slog.Level(Error2Level), args) }
func (l *OTLogger) Error3(args ...any) { l.logger.Log(slog.Level(Error3Level), args) }
func (l *OTLogger) Error4(args ...any) { l.logger.Log(slog.Level(Error4Level), args) }
func (l *OTLogger) Errorf(format string, args ...any) { l.logger.Logf(slog.Level(ErrorLevel), format, args)}
func (l *OTLogger) Errorf2(format string, args ...any) { l.logger.Logf(slog.Level(Error2Level), format, args) }
func (l *OTLogger) Errorf3(format string, args ...any) { l.logger.Logf(slog.Level(Error3Level), format, args) }
func (l *OTLogger) Errorf4(format string, args ...any) { l.logger.Logf(slog.Level(Error4Level), format, args) }

func (l *OTLogger) Fatal(args ...any) { l.logger.Log(slog.Level(FatalLevel), args)}
func (l *OTLogger) Fatal2(args ...any) { l.logger.Log(slog.Level(Fatal2Level), args) }
func (l *OTLogger) Fatal3(args ...any) { l.logger.Log(slog.Level(Fatal3Level), args) }
func (l *OTLogger) Fatal4(args ...any) { l.logger.Log(slog.Level(Fatal4Level), args) }
func (l *OTLogger) Fatalf(format string, args ...any) { l.logger.Logf(slog.Level(FatalLevel), format, args)}
func (l *OTLogger) Fatalf2(format string, args ...any) { l.logger.Logf(slog.Level(Fatal2Level), format, args) }
func (l *OTLogger) Fatalf3(format string, args ...any) { l.logger.Logf(slog.Level(Fatal3Level), format, args) }
func (l *OTLogger) Fatalf4(format string, args ...any) { l.logger.Logf(slog.Level(Fatal4Level), format, args) }