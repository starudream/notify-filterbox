package notify

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"

	"github.com/starudream/notify-filterbox/util"
)

func TestNotify(t *testing.T) {
	res, err := Notify(Params{
		Id:       util.UUID(),
		Title:    "foo",
		SubTitle: "bar",
		Message:  "hello",
		GroupId:  "baz",
	})
	testutil.LogNoErr(t, err, res)
}
