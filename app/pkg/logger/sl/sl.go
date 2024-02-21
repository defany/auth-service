package sl

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/defany/auth-service/app/pkg/logger/handlers/slogpretty"
)

type Slog struct {
	Level     slog.Level              `json:"level" env-default:"debug"`
	AddSource bool                    `json:"add_source" env-default:"false"`
	Format    slogpretty.FieldsFormat `json:"format" env-default:"pretty"` // json, text or pretty
}

func NewSlogLogger(c Slog) *slog.Logger {
	o := &slog.HandlerOptions{Level: c.Level, AddSource: c.AddSource}
	w := os.Stdout
	var h slog.Handler

	switch c.Format {
	case "pretty":
		h = slogpretty.NewHandler().
			WithAddSource(c.AddSource).
			WithLevel(c.Level).
			WithLevelEmoji(true).
			WithFieldsFormat("json")
	case "json":
		h = slog.NewJSONHandler(w, o)
	case "text":
		h = slog.NewTextHandler(w, o)
	}

	return slog.New(h)
}

func Err(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}

func ErrAttr(err error) slog.Attr {
	return slog.Attr{
		Key:   "response",
		Value: slog.StringValue(err.Error()),
	}
}

func OpErrAttr(op string, err error) slog.Attr {
	return ErrAttr(Err(op, err))
}
