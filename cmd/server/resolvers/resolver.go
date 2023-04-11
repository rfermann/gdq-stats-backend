package resolvers

import "github.com/rfermann/gdq-stats-backend/internal/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services *services.Services
}
