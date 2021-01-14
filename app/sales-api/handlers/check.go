package handlers

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/tleuzhan13/service/foundation/web"
	"log"
	"math/rand"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if n := rand.Intn(100); n%2 == 0 {
		return errors.New("untrusted error")
	}

	status := struct {
		Status string
	}{
		Status: "service ready",
	}

	c.log.Println(r, status)
	return web.Respond(ctx, w, status, http.StatusOK)
}
