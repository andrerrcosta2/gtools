// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package op

type (
	// Clone is a bitmask type used to configure the clone on reflection operations
	Clone uint16
	// Compare is a bitmask type used to configure the conditions of equality on reflection operations.
	Compare uint16
	// Read is a bitmask type used to configure the behavior of reflect read operations.
	// It controls how unexported field4, functions, channels, and interface values are handled,
	// enabling customization from lenient to strict - reflect.DeepEqual-like - semantics.
	Read uint8
	// Write is a bitmask type used to configure the behavior of reflect write operations.
	// It controls how unexported field4, functions, channels, and interface values are handled,
	// enabling customization from lenient to strict - reflect.DeepEqual-like - semantics.
	Write uint8
)

func (Read) Op() string    { return "Reading" }
func (Write) Op() string   { return "Writing" }
func (Compare) Op() string { return "Comparing" }
func (Clone) Op() string   { return "Cloning" }

//func Resolve(op ...reflect4.Option) (read Read, write Write, comparing Compare, clone Clone) {
//	for _, o := range op {
//		switch o.Op() {
//		case Reading:
//			read |= o.(Read)
//		case Writing:
//			write |= o.(Write)
//		case Comparing:
//			comparing |= o.(Compare)
//		case Cloning:
//			clone |= o.(Clone)
//		}
//	}
//	return
//}
