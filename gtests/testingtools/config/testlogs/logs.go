// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testlogs

import "github.com/andrerrcosta2/gtools/gtests/testingtools/gtests"

const (
	Default gtests.LoggerStrategy = iota
	ShowAfterTests
	OnErrors
	OnFailure
)
