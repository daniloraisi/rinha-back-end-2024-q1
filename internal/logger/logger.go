package logger

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
)

const (
	messageKey = "msg"
	timeKey    = "timestamp"
	iso        = "2006-01-02T15:04:05.000Z0300"

	prodEnv = "prod"
)

type KeyValue struct {
	Key   string
	Value interface{}
}

type Logger struct {
	stdOut *zerolog.Logger
	stdErr *zerolog.Logger
}

func (l *Logger) Debug(msg string, keyValues ...KeyValue) {
	constructFields(l.stdOut.Debug(), keyValues...).Msg(msg)
}

func (l *Logger) Info(msg string, keyValues ...KeyValue) {
	constructFields(l.stdOut.Info(), keyValues...).Msg(msg)
}

func (l *Logger) Warn(msg string, keyValues ...KeyValue) {
	constructFields(l.stdOut.Warn(), keyValues...).Msg(msg)
}

func (l *Logger) Error(msg string, keyValues ...KeyValue) {
	constructFields(l.stdErr.Error(), keyValues...).Msg(msg)
}

func (l *Logger) Panic(msg string, keyValues ...KeyValue) {
	constructFields(l.stdErr.Panic(), keyValues...).Msg(msg)
}

func (l *Logger) Fatal(msg string, keyValues ...KeyValue) {
	constructFields(l.stdErr.Fatal(), keyValues...).Msg(msg)
}

func (l Logger) ErrProp(err error) KeyValue {
	return KeyValue{Key: "error", Value: err.Error()}
}

func (l Logger) Prop(key string, value interface{}) KeyValue {
	return KeyValue{Key: key, Value: value}
}

func constructFields(event *zerolog.Event, keyValues ...KeyValue) *zerolog.Event {
	for _, v := range keyValues {
		event.Interface(v.Key, v.Value)
	}

	return event
}

func shortenCallerField(pc uintptr, file string, line int) string {
	short := file
	found := false

	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			if found {
				short = file[i+1:]
				break
			}

			found = true
		}
	}

	file = short

	return file + ":" + strconv.Itoa(line)
}

func capitalizeLevelField(l zerolog.Level) string {
	return strings.ToUpper(l.String())
}

func setEncoderFields() {
	zerolog.CallerMarshalFunc = shortenCallerField
	zerolog.LevelFieldMarshalFunc = capitalizeLevelField
	zerolog.MessageFieldName = messageKey
	zerolog.TimeFieldFormat = iso
	zerolog.TimestampFieldName = timeKey
}

func newZeroLogger(w io.Writer) *zerolog.Logger {
	l := zerolog.New(w).With().
		Timestamp().
		CallerWithSkipFrameCount(3).
		Logger()

	return &l
}

func New(envName string) *Logger {
	logger := &Logger{
		stdOut: newZeroLogger(os.Stdout),
		stdErr: newZeroLogger(os.Stderr),
	}

	if strings.Contains(envName, prodEnv) {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	setEncoderFields()

	return logger
}
