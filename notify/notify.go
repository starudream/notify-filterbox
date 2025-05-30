package notify

import (
	"os/exec"
	"strings"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
)

type Params struct {
	Id       string
	Title    string
	SubTitle string
	Message  string
	GroupId  string
}

type Result struct {
	Id  string
	Out string
}

func Notify(p Params) (Result, error) {
	res := Result{Id: p.Id}

	args := []string{"--message", p.Message, "--timeout", "30"}
	if p.Title != "" {
		args = append(args, "--title", p.Title)
	}
	if p.SubTitle != "" {
		args = append(args, "--subtitle", p.SubTitle)
	}

	path := "npx node-notifier-cli notify"
	if config.Exists("notify.path") {
		path = config.Get("notify.path").String()
	}

	ss := strings.Split(path, " ")
	if len(ss) > 1 {
		path = ss[0]
		args = append(ss[1:], args...)
	}

	slog.Debug("exec: %s %v", path, strings.Join(args, " "))

	bs, err := exec.Command(path, args...).CombinedOutput()
	res.Out = string(bs)

	return res, err
}
