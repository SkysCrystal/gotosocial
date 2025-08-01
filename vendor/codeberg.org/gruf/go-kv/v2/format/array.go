package format

import "unsafe"

// iterArrayType returns a FormatFunc capable of iterating
// and formatting the given array type currently in typenode{}.
// note this will fetch a sub-FormatFunc for the array element
// type, and also handle special cases of [n]byte, [n]rune arrays.
func (fmt *Formatter) iterArrayType(t typenode) FormatFunc {

	// Array element type.
	elem := t.rtype.Elem()

	// Get nested elem typenode with appropriate flags.
	flags := reflect_array_elem_flags(t.flags, elem)
	et := t.next(elem, flags)

	// Get elem format func.
	fn := fmt.loadOrGet(et)
	if fn == nil {
		panic("unreachable")
	}

	// Handle possible sizes.
	switch t.rtype.Len() {
	case 0:
		return emptyArrayType(t)
	case 1:
		return iterSingleArrayType(t, fn)
	default:
		return iterMultiArrayType(t, fn)
	}
}

func emptyArrayType(t typenode) FormatFunc {
	if !t.needs_typestr() {
		// Simply append empty.
		return func(s *State) {
			s.B = append(s.B, "[]"...)
		}
	}

	// Array type string with refs.
	typestr := t.typestr_with_refs()

	// Append empty with type.
	return func(s *State) {
		if s.A.WithType() {
			s.B = append(s.B, typestr...)
			s.B = append(s.B, "{}"...)
		} else {
			s.B = append(s.B, "[]"...)
		}
	}
}

func iterSingleArrayType(t typenode, fn FormatFunc) FormatFunc {
	if !t.needs_typestr() {
		return func(s *State) {
			// Wrap 'fn' in braces.
			s.B = append(s.B, '[')
			fn(s)
			s.B = append(s.B, ']')
		}
	}

	// Array type string with refs.
	typestr := t.typestr_with_refs()

	// Wrap in type+braces.
	return func(s *State) {

		// Open / close braces.
		var open, close uint8
		open, close = '[', ']'

		// Include type info.
		if s.A.WithType() {
			s.B = append(s.B, typestr...)
			open, close = '{', '}'
		}

		// Wrap 'fn' in braces.
		s.B = append(s.B, open)
		fn(s)
		s.B = append(s.B, close)
	}
}

func iterMultiArrayType(t typenode, fn FormatFunc) FormatFunc {
	// Array element in-memory size.
	esz := t.rtype.Elem().Size()

	// Number of elements.
	n := t.rtype.Len()

	if !t.needs_typestr() {
		// Wrap elems in braces.
		return func(s *State) {
			ptr := s.P

			// Prepend array brace.
			s.B = append(s.B, '[')

			for i := 0; i < n; i++ {
				// Format at array index.
				offset := esz * uintptr(i)
				s.P = add(ptr, offset)
				fn(s)

				// Append separator.
				s.B = append(s.B, ',')
			}

			// Drop final space.
			s.B = s.B[:len(s.B)-1]

			// Prepend array brace.
			s.B = append(s.B, ']')
		}
	}

	// Array type string with refs.
	typestr := t.typestr_with_refs()

	// Wrap in type+braces.
	return func(s *State) {
		ptr := s.P

		// Open / close braces.
		var open, close uint8
		open, close = '[', ']'

		// Include type info.
		if s.A.WithType() {
			s.B = append(s.B, typestr...)
			open, close = '{', '}'
		}

		// Prepend array brace.
		s.B = append(s.B, open)

		for i := 0; i < n; i++ {
			// Format at array index.
			offset := esz * uintptr(i)
			s.P = add(ptr, offset)
			fn(s)

			// Append separator.
			s.B = append(s.B, ',')
		}

		// Drop final comma.
		s.B = s.B[:len(s.B)-1]

		// Prepend array brace.
		s.B = append(s.B, close)
	}
}

func wrapByteArray(t typenode, fn FormatFunc) FormatFunc {
	n := t.rtype.Len()
	if !t.needs_typestr() {
		return func(s *State) {
			if s.A.AsText() || s.A.AsQuotedText() || s.A.AsQuotedASCII() {
				var v string
				p := (*unsafeheader_String)(unsafe.Pointer(&v))
				p.Len = n
				p.Data = s.P
				appendString(s, v)
			} else {
				fn(s)
			}
		}
	}
	typestr := t.typestr_with_ptrs()
	return func(s *State) {
		if s.A.AsText() || s.A.AsQuotedText() || s.A.AsQuotedASCII() {
			var v string
			p := (*unsafeheader_String)(unsafe.Pointer(&v))
			p.Len = n
			p.Data = s.P
			if s.A.WithType() {
				s.B = append(s.B, "("+typestr+")("...)
				appendString(s, v)
				s.B = append(s.B, ")"...)
			} else {
				appendString(s, v)
			}
		} else {
			fn(s)
		}
	}
}

func wrapRuneArray(t typenode, fn FormatFunc) FormatFunc {
	n := t.rtype.Len()
	if !t.needs_typestr() {
		return func(s *State) {
			if s.A.AsText() || s.A.AsQuotedText() || s.A.AsQuotedASCII() {
				var v []rune
				p := (*unsafeheader_Slice)(unsafe.Pointer(&v))
				p.Cap = n
				p.Len = n
				p.Data = s.P
				appendString(s, string(v))
			} else {
				fn(s)
			}
		}
	}
	typestr := t.typestr_with_ptrs()
	return func(s *State) {
		if s.A.AsText() || s.A.AsQuotedText() || s.A.AsQuotedASCII() {
			var v []rune
			p := (*unsafeheader_Slice)(unsafe.Pointer(&v))
			p.Cap = n
			p.Len = n
			p.Data = s.P
			if s.A.WithType() {
				s.B = append(s.B, "("+typestr+")("...)
				appendString(s, string(v))
				s.B = append(s.B, ")"...)
			} else {
				appendString(s, string(v))
			}
		} else {
			fn(s)
		}
	}
}
