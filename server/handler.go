package server

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/gh"
	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/notify-filterbox/filterbox"
	"github.com/starudream/notify-filterbox/notify"
	"github.com/starudream/notify-filterbox/util"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("read body error: %v", err)
		return
	}

	body, err = json.Compact(body)
	if err != nil {
		slog.Error("compact body error: %v", err)
		return
	}

	slog.Debug("req: %s", body)

	msg, err := json.UnmarshalTo[*filterbox.Message](body)
	if err != nil {
		slog.Error("unmarshal body error: %v, raw: %s", err, base64.StdEncoding.EncodeToString(body))
		return
	}

	if msg.Ongoing {
		return
	}

	if msg.Title == "{android.title}" {
		msg.Title = ""
	}
	if msg.Text == "{android.text}" {
		msg.Text = ""
	}
	if msg.BigText == "{android.bigText}" {
		msg.BigText = ""
	}

	if msg.Text == "" && msg.BigText == "" {
		return
	}

	go send(util.UUID(), msg)

	w.WriteHeader(200)
}

func send(id string, msg *filterbox.Message) {
	res, err := notify.Notify(notify.Params{
		Id:       id,
		Title:    msg.Title,
		SubTitle: msg.AppName + "(" + msg.PackageName + ")",
		Message:  gh.Ternary(msg.BigText != "", msg.BigText, msg.Text),
	})
	if err != nil {
		slog.Error("notify error: %v, out: %s", err, res.Out)
	}
}
