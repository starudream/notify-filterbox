package server

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"
)

func Start() error {
	ln, err := net.Listen("tcp", config.Get("addr").String())
	if err != nil {
		return err
	}

	slog.Info("server started at %s", ln.Addr().String())

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	hs := &http.Server{Handler: mux}

	go func() {
		err = hs.Serve(ln)
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				slog.Error("server stopped with error: %v", err)
			} else {
				err = nil
			}
		}
	}()

	<-signalutil.Defer(func() { _ = hs.Shutdown(context.Background()) }).Done()

	return err
}
