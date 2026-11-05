// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chrono

import "strings"

// Cron builds a 5-field cron expression:
// minute hour day-of-month month day-of-week
//
// It does not validate field contents; callers are responsible for passing
// values compatible with their scheduler.
func Cron(minute, hour, dayOfMonth, month, dayOfWeek string) string {
	return strings.Join([]string{minute, hour, dayOfMonth, month, dayOfWeek}, " ")
}

// CronWithSeconds builds a 6-field cron expression:
// second minute hour day-of-month month day-of-week
//
// It does not validate field contents; callers are responsible for passing
// values compatible with their scheduler.
func CronWithSeconds(second, minute, hour, dayOfMonth, month, dayOfWeek string) string {
	return strings.Join([]string{second, minute, hour, dayOfMonth, month, dayOfWeek}, " ")
}
