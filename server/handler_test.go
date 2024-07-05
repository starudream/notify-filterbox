package server

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/config"

	"github.com/starudream/notify-filterbox/filterbox"
)

func TestFilter(t *testing.T) {
	config.Set("script", `
function filter(data) {
    if (data.AppName === "哔哩哔哩") {
        return false
    }
    return true
}
`)
	t.Log(filter(&filterbox.Message{
		Title:       "标题标题",
		Text:        "文本文本",
		AppName:     "哔哩哔哩",
		PackageName: "tv.danmaku.bili",
	}))
}
