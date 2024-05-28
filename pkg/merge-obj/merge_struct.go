package merge_obj

import "reflect"

type Config struct {
	Transformers                 Transformers
	Overwrite                    bool
	ShouldNotDereference         bool
	AppendSlice                  bool
	TypeCheck                    bool
	overwriteWithEmptyValue      bool
	overwriteSliceWithEmptyValue bool
	sliceDeepCopy                bool
	debug                        bool
}

type Transformers interface {
	Transformer(reflect.Type) func(dst, src reflect.Value) error
}

type visit struct {
	typ  reflect.Type
	next *visit
	ptr  uintptr
}
