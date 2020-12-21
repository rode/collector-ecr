package listener

import (
	"net/http"

	"go.uber.org/zap"

	pb "github.com/liatrio/rode-api/proto/v1alpha1"
)

type listener struct {
	rodeClient pb.RodeClient
	logger     *zap.Logger
}

type Listener interface {
	ProcessEvent(http.ResponseWriter, *http.Request)
}

func NewListener(logger *zap.Logger, client pb.RodeClient) Listener {
	return &listener{
		rodeClient: client,
		logger:     logger,
	}
}

// ProcessEvent handles incoming webhook events
func (l *listener) ProcessEvent(w http.ResponseWriter, request *http.Request) {
	log := l.logger.Named("ProcessEvent")
	// Add authorization/authentication logic

	// Process Event

	// Send to Rode API
	log.Info("Received event for processing")
	w.WriteHeader(200)
}
