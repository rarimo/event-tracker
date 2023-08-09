package service

import (
	"context"
	"gitlab.com/rarimo/identity/event-tracker/internal/config"
	"gitlab.com/rarimo/identity/event-tracker/internal/database/postgres"
	"gitlab.com/rarimo/identity/event-tracker/internal/service/events/handler"
	"gitlab.com/rarimo/identity/event-tracker/internal/service/events/tracker"
	"gitlab.com/rarimo/identity/event-tracker/internal/service/galxe"
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
