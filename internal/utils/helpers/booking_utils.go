package helpers

import (
	"fmt"
	"strings"
	"time"
)

func GenerateBookingCode() string {
	// Format: BKG + timestamp + random
	// contoh: BKG20250327A1B2
	timestamp := time.Now().Format("20060102")
	random := strings.ToUpper(RandomString(4))
	return fmt.Sprintf("BKG%s%s", timestamp, random)
}
