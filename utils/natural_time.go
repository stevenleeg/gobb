package utils

import (
	"fmt"
	"time"
)

func TimeRelativeToNow(in time.Time) string {
	diff := time.Since(in)

	if diff.Hours()/24 > 7 || diff.Hours() < 0 {
		return in.Format("Mon Jan 2 2006")
	} else if int(diff.Hours()/24) == 1 {
		return "1 day ago"
	} else if int(diff.Hours()/24) > 1 {
		return fmt.Sprintf("%d days ago", int(diff.Hours()/24))
	} else if diff.Seconds() <= 1 {
		return "just now"
	} else if diff.Seconds() < 60 {
		return fmt.Sprintf("%d seconds ago", int(diff.Seconds()))
	} else if diff.Seconds() < 120 {
		return "1 minute ago"
	} else if diff.Seconds() < 3600 {
		return fmt.Sprintf("%d minutes ago", int(diff.Seconds()/60))
	} else if diff.Seconds() < 7200 {
		return "1 hour ago"
	} else {
		return fmt.Sprintf("%d hours ago", int(diff.Seconds()/60/60))
	}
}
