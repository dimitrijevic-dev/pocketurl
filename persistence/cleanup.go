package persistence

import (
	"time"
)

func startCleanupScheduler() {
	ticker := time.NewTicker(1 * time.Hour)
	for range ticker.C { CleanExpiredLinks() }
}