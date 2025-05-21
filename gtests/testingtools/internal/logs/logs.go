// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package logs

type LoggerLevel int

const (
	LOG_ON_CALL LoggerLevel = iota
	LOG_ON_ERRORS
	LOG_ON_FAILURE
)
