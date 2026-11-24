package main

import (
	"fmt"
	"time"
)

func ParseAndExtractTime(raw string, targetLocation string) (string, error) {
	// 1. Strict parsing: Reject any string that omits offset/Z
	t, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return "", fmt.Errorf("invalid RFC3339 timestamp: %w", err)
	}

	// 2. Load the IANA Location (handles DST automatically)
	loc, err := time.LoadLocation(targetLocation)
	if err != nil {
		return "", fmt.Errorf("unknown location: %w", err)
	}

	// 3. Shift the representation to the user's location
	localTime := t.In(loc)

	// 4. Extract wall-clock time string (HH:MM:SS)
	return localTime.Format("15:04:05"), nil
}
