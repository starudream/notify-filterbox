package notify

import (
	"os/exec"
	"strings"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
)

func notify(p Params) (Result, error) {
	res := Result{Id: p.Id}

	args := []string{"-json", "-timeout", "30", "-actions", "Close", "-message", p.Message}
	if p.Title != "" {
		args = append(args, "-title", p.Title)
	}
	if p.SubTitle != "" {
		args = append(args, "-subtitle", p.SubTitle)
	}

	slog.Debug("alerter %v", strings.Join(args, " "))

	path := "/opt/homebrew/bin/alerter"
	if config.Exists("alerter.path") {
		path = config.Get("alerter.path").String()
	}
	bs, err := exec.Command(path, args...).Output()
	res.Out = string(bs)

	return res, err
}
