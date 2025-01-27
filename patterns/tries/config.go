// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import "runtime"

var defaultConcurrentInternalOperations = max(6, runtime.NumCPU())
