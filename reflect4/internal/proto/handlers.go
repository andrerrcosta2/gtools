// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package proto

//// Deep performs a deep extraction of a given data value, returning its Prototype representation.
//// It classifies the data type and processes it accordingly to extract detailed metadata.
//func Deep(data any) (*reflective.Prototype, error) {
//	return gdp(reflectlite.TypeOf(data).Prototype())
//}
//
//// Shallow performs a shallow extraction of a given data value, returning its Prototype representation.
//// It classifies the data type and processes it accordingly to extract metadata.
//func Shallow(data any) *reflective.Prototype {
//	return reflectlite.TypeOf(data).Prototype()
//}
//
//// gdp generates a deep prototype representation of a given reflective type.
//// It handles various kinds of types, including primitives, pointers, iterables, maps, and structs,
//// recursively extracting detailed metadata for each type class.
//// If the type is invalid or unrecognized, it returns an error.
//func gdp(t *reflective.Prototype) (proto *reflective.Prototype, err error) {
//	switch {
//	case kinds.IsPrimitive(t.Kind):
//		return t, nil
//	case kinds.IsPointer(t.Kind):
//		return gdp(t.Value)
//	case kinds.IsMap(t.Kind):
//		return deepMap(t)
//	case kinds.IsIterable(t.Kind):
//		return deepIterable(t)
//	case kinds.IsStruct(t.Kind):
//		return deepStruct(t)
//	default:
//		return nil, fmt.Errorf("unsupported kind: %v", proto.Kind)
//	}
//}
//
//// deepMap returns a deep prototype representation of a given map type.
//// It returns a prototype with Name and OfValue fields set to the corresponding
//// prototypes of the map's key and element types.
//// The prototypes of the key and element types are recursively extracted
//// using the Deep function.
//func deepMap(t *reflective.Prototype) (proto *reflective.Prototype, err error) {
//	// Recursively extract the prototype of the key type
//	var keyPrototype *reflective.Prototype
//	keyPrototype, err = gdp(t.Key)
//
//	if err != nil {
//		return
//	}
//
//	// Recursively extract the prototype of the element type
//	var valuePrototype *reflective.Prototype
//	valuePrototype, err = gdp(t.Value)
//
//	if err != nil {
//		return
//	}
//
//	proto = &reflective.Prototype{
//		Kind:  reflect.Map,
//		Key:   keyPrototype,
//		Value: valuePrototype,
//	}
//	return
//}
//
//// deepIterable returns a deep prototype representation of a given iterable type (slice or array).
//// It extracts the prototype of the element type recursively using the Deep function.
//func deepIterable(t *reflective.Prototype) (proto *reflective.Prototype, err error) {
//	// Recursively extract the prototype of the element type
//	var value *reflective.Prototype
//	value, err = gdp(t.Value)
//	if err != nil {
//		return
//	}
//
//	// ToSet the extracted element prototype
//	return &reflective.Prototype{
//		Value: value,
//	}, err
//}
//
//// deepStruct returns a deep prototype representation of a given struct type.
//// It extracts metadata such as the type name, kind, and field names of the struct
//// and its fields, and recursively extracts the prototypes of the fields using the Deep function.
//func deepStruct(t *reflective.Prototype) (proto *reflective.Prototype, err error) {
//	// Iterate through each field of the struct
//	for i := 0; i < len(t.Fields); i++ {
//		t.Fields[i], err = gdp(t.Fields[i])
//		if err != nil {
//			return
//		}
//	}
//	return
//}
