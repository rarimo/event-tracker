package service

import (
	"context"
	"github.com/rarimo/event-tracker/internal/config"
	"github.com/rarimo/event-tracker/internal/database/postgres"
	"github.com/rarimo/event-tracker/internal/service/events/handler"
	"github.com/rarimo/event-tracker/internal/service/events/tracker"
	"github.com/rarimo/event-tracker/internal/service/galxe"
)

// Run runs the service
func Run(ctx context.Context, cfg config.Config) {
	tracker.NewEventTracker(
		cfg.EventTrackerSettings(),
		cfg.Log(),
		postgres.NewServiceDB(cfg.DB()),
		handler.MustNewEventHandler(
			cfg.Log(),
			galxe.NewSubmitter(cfg.GalxeSubmitterSettings()),
		),
	).Run(ctx)
}
