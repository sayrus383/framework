package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
)

type orderedFields struct {
	keys   []string
	values map[string]interface{}
}

func (f *orderedFields) Set(key string, value interface{}) {
	f.keys = append(f.keys, key)
	f.values[key] = value
}

type SlogHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type SlogHandler struct {
	slog.Handler
	l     *log.Logger
	attrs []slog.Attr
}

func NewSlogHandler(out io.Writer, opts SlogHandlerOptions) *SlogHandler {
	h := &SlogHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

func (h *SlogHandler) Handle(ctx context.Context, r slog.Record) error {
	var (
		level  = r.Level.String()
		fields = &orderedFields{
			values: make(map[string]interface{}),
		}
	)

	fields.Set(KeyTime, r.Time.Format(fmt.Sprintf("%s", TimeFormat)))
	fields.Set(KeyLevel, r.Level.String())
	fields.Set(KeyMessage, r.Message)

	r.Attrs(func(a slog.Attr) bool {
		if err, ok := a.Value.Any().(error); ok {
			fields.Set(a.Key, err.Error())
			return true
		}

		fields.Set(a.Key, a.Value.Any())
		return true
	})

	for _, a := range h.attrs {
		fields.Set(a.Key, a.Value.Any())
	}

	var jsonFields []byte
	jsonFields = append(jsonFields, '{')
	for i, key := range fields.keys {
		val, _ := json.Marshal(fields.values[key])
		jsonFields = append(jsonFields, fmt.Sprintf(`"%s":%s`, key, val)...)
		if i < len(fields.keys)-1 {
			jsonFields = append(jsonFields, ',')
		}
	}
	jsonFields = append(jsonFields, '}')

	/* Приводим к единому формату, потому что банковский стандарт
	[2024-01-04 14:22:41.860][INFO]: {
		"timestamp":"2024-01-04 14:22:41.860",
		"request_id":"",
		"log_type":"INFO",
		"module":"dbp-order-adapter",
		"version":"development",
		"file":"",
		"message":"initialized"
	}
	*/
	h.l.Println(
		fmt.Sprintf("%s[%s]:", r.Time.Format(fmt.Sprintf("[%s]", TimeFormat)), level),
		string(jsonFields),
	)

	return nil
}

func (h *SlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &SlogHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}
