// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package format

import "github.com/andrerrcosta2/gtools/core/format/code/indent"

// Ser defines the type of serializable data
type Ser string

// Sprintable defines a type that can sprint itself to a string using an indent.Tab
// for recursive formatting
type Sprintable interface {
	Sprint(indentor indent.Indentor) string
}

// Decoder defines a type that can decode a byte array to a data
type Decoder[T any] interface {
	Decode(b []byte) (T, error)
}

// Encoder defines a type that can encode a data to a byte array
type Encoder[T any] interface {
	Encode(t T) ([]byte, error)
}

// WithContentType defines a type that can return its content type
type WithContentType interface {
	ContentType() string
}
