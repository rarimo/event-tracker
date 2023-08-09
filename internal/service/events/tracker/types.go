package tracker

import "context"

// EventTracker is an interface for event tracker
type EventTracker interface {
	// Run starts the event tracker to track contract events
	Run(ctx context.Context)
}
