package slices

import "testing"

func TestC128SliceString(t *testing.T) {
	ConfirmString := func(s *C128Slice, r string) {
		if x := s.String(); x != r {
			t.Fatalf("%v erroneously serialised as '%v'", r, x)
		}
	}

	ConfirmString(C128List(), "()")
	ConfirmString(C128List(0), "((0+0i))")
	ConfirmString(C128List(0, 1), "((0+0i) (1+0i))")
	ConfirmString(C128List(0, 1i), "((0+0i) (0+1i))")
}

func TestC128SliceLen(t *testing.T) {
	ConfirmLength := func(s *C128Slice, i int) {
		if x := s.Len(); x != i {
			t.Fatalf("%v.Len() should be %v but is %v", s, i, x)
		}
	}
	
	ConfirmLength(C128List(0), 1)
	ConfirmLength(C128List(0, 1), 2)
}

func TestC128SliceSwap(t *testing.T) {
	ConfirmSwap := func(s *C128Slice, i, j int, r *C128Slice) {
		if s.Swap(i, j); !r.Equal(s) {
			t.Fatalf("Swap(%v, %v) should be %v but is %v", i, j, r, s)
		}
	}
	ConfirmSwap(C128List(0, 1, 2), 0, 1, C128List(1, 0, 2))
	ConfirmSwap(C128List(0, 1, 2), 0, 2, C128List(2, 1, 0))
}

func TestC128SliceCut(t *testing.T) {
	ConfirmCut := func(s *C128Slice, start, end int, r *C128Slice) {
		if s.Cut(start, end); !r.Equal(s) {
			t.Fatalf("Cut(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 0, 1, C128List(1, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 1, 2, C128List(0, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 2, 3, C128List(0, 1, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 3, 4, C128List(0, 1, 2, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 4, 5, C128List(0, 1, 2, 3, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 5, 6, C128List(0, 1, 2, 3, 4))

	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), -1, 1, C128List(1, 2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 0, 2, C128List(2, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 1, 3, C128List(0, 3, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 2, 4, C128List(0, 1, 4, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 3, 5, C128List(0, 1, 2, 5))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 4, 6, C128List(0, 1, 2, 3))
	ConfirmCut(C128List(0, 1, 2, 3, 4, 5), 5, 7, C128List(0, 1, 2, 3, 4))
}

func TestC128SliceTrim(t *testing.T) {
	ConfirmTrim := func(s *C128Slice, start, end int, r *C128Slice) {
		if s.Trim(start, end); !r.Equal(s) {
			t.Fatalf("Trim(%v, %v) should be %v but is %v", start, end, r, s)
		}
	}

	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 0, 1, C128List(0))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 1, 2, C128List(1))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 2, 3, C128List(2))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 3, 4, C128List(3))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 4, 5, C128List(4))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 5, 6, C128List(5))

	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), -1, 1, C128List(0))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 0, 2, C128List(0, 1))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 1, 3, C128List(1, 2))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 2, 4, C128List(2, 3))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 3, 5, C128List(3, 4))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 4, 6, C128List(4, 5))
	ConfirmTrim(C128List(0, 1, 2, 3, 4, 5), 5, 7, C128List(5))
}

func TestC128SliceDelete(t *testing.T) {
	ConfirmDelete := func(s *C128Slice, index int, r *C128Slice) {
		if s.Delete(index); !r.Equal(s) {
			t.Fatalf("Delete(%v) should be %v but is %v", index, r, s)
		}
	}

	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), -1, C128List(0, 1, 2, 3, 4, 5))
	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), 0, C128List(1, 2, 3, 4, 5))
	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), 1, C128List(0, 2, 3, 4, 5))
	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), 2, C128List(0, 1, 3, 4, 5))
	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), 3, C128List(0, 1, 2, 4, 5))
	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), 4, C128List(0, 1, 2, 3, 5))
	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), 5, C128List(0, 1, 2, 3, 4))
	ConfirmDelete(C128List(0, 1, 2, 3, 4, 5), 6, C128List(0, 1, 2, 3, 4, 5))
}

func TestC128SliceDeleteIf(t *testing.T) {
	ConfirmDeleteIf := func(s *C128Slice, f interface{}, r *C128Slice) {
		if s.DeleteIf(f); !r.Equal(s) {
			t.Fatalf("DeleteIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), complex128(0), C128List(1, 3, 5))
	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), complex128(1), C128List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), complex128(6), C128List(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(0) }, C128List(1, 3, 5))
	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(1) }, C128List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(6) }, C128List(0, 1, 0, 3, 0, 5))

	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(0) }, C128List(1, 3, 5))
	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(1) }, C128List(0, 0, 3, 0, 5))
	ConfirmDeleteIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(6) }, C128List(0, 1, 0, 3, 0, 5))
}

func TestC128SliceEach(t *testing.T) {
	var count	complex128
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(index int, i interface{}) {
		if index != int(real(i.(complex128))) {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(key, i interface{}) {
		if complex(float64(key.(int)), 0) != i {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})

	count = 0
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(i complex128) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(index int, i complex128) {
		if int(real(i)) != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).Each(func(key interface{}, i complex128) {
		if key.(int) != int(real(i)) {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestC128SliceBlockCopy(t *testing.T) {
	ConfirmBlockCopy := func(s *C128Slice, destination, source, count int, r *C128Slice) {
		s.BlockCopy(destination, source, count)
		if !r.Equal(s) {
			t.Fatalf("BlockCopy(%v, %v, %v) should be %v but is %v", destination, source, count, r, s)
		}
	}

	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 0, 4, C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 9, 9, 4, C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 2, 4, C128List(0, 1, 2, 3, 4, 2, 3, 4, 5, 9))
	ConfirmBlockCopy(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 2, 5, 4, C128List(0, 1, 5, 6, 7, 8, 6, 7, 8, 9))
}

func TestC128SliceBlockClear(t *testing.T) {
	ConfirmBlockClear := func(s *C128Slice, start, count int, r *C128Slice) {
		s.BlockClear(start, count)
		if !r.Equal(s) {
			t.Fatalf("BlockClear(%v, %v) should be %v but is %v", start, count, r, s)
		}
	}

	ConfirmBlockClear(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, 4, C128List(0, 0, 0, 0, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 4, C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmBlockClear(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 4, C128List(0, 1, 2, 3, 4, 0, 0, 0, 0, 9))
}

func TestC128SliceOverwrite(t *testing.T) {
	ConfirmOverwrite := func(s *C128Slice, offset int, v, r *C128Slice) {
		s.Overwrite(offset, *v)
		if !r.Equal(s) {
			t.Fatalf("Overwrite(%v, %v) should be %v but is %v", offset, v, r, s)
		}
	}

	ConfirmOverwrite(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 0, C128List(10, 9, 8, 7), C128List(10, 9, 8, 7, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, C128List(10, 9, 8, 7), C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	ConfirmOverwrite(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, C128List(11, 12, 13, 14), C128List(0, 1, 2, 3, 4, 11, 12, 13, 14, 9))
}

func TestC128SliceReallocate(t *testing.T) {
	ConfirmReallocate := func(s *C128Slice, l, c int, r *C128Slice) {
		o := s.String()
		el := l
		if el > c {
			el = c
		}
		switch s.Reallocate(l, c); {
		case s == nil:				t.Fatalf("%v.Reallocate(%v, %v) created a nil value for Slice", o, l, c)
		case s.Cap() != c:			t.Fatalf("%v.Reallocate(%v, %v) capacity should be %v but is %v", o, l, c, c, s.Cap())
		case s.Len() != el:			t.Fatalf("%v.Reallocate(%v, %v) length should be %v but is %v", o, l, c, el, s.Len())
		case !r.Equal(s):			t.Fatalf("%v.Reallocate(%v, %v) should be %v but is %v", o, l, c, r, s)
		}
	}

	c := make(C128Slice, 0, 10)
	ConfirmReallocate(C128List(), 0, 10, &c)
	ConfirmReallocate(C128List(0, 1, 2, 3, 4), 3, 10, C128List(0, 1, 2))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4), 5, 10, C128List(0, 1, 2, 3, 4))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4), 10, 10, C128List(0, 1, 2, 3, 4, 0, 0, 0, 0, 0))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 1, 5, C128List(0))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 5, 5, C128List(0, 1, 2, 3, 4))
	ConfirmReallocate(C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), 10, 5, C128List(0, 1, 2, 3, 4))
}

func TestC128SliceExtend(t *testing.T) {
	ConfirmExtend := func(s *C128Slice, n int, r *C128Slice) {
		c := s.Cap()
		s.Extend(n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Extend(%v) len should be %v but is %v", n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Extend(%v) cap should be %v but is %v", n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Extend(%v) should be %v but is %v", n, r, s)
		}
	}

	ConfirmExtend(C128List(), 1, C128List(0))
	ConfirmExtend(C128List(), 2, C128List(0, 0))
}

func TestC128SliceExpand(t *testing.T) {
	ConfirmExpand := func(s *C128Slice, i, n int, r *C128Slice) {
		c := s.Cap()
		s.Expand(i, n)
		switch {
		case s.Len() != r.Len():	t.Fatalf("Expand(%v, %v) len should be %v but is %v", i, n, r.Len(), s.Len())
		case s.Cap() != c + n:		t.Fatalf("Expand(%v, %v) cap should be %v but is %v", i, n, c + n, s.Cap())
		case !r.Equal(s):			t.Fatalf("Expand(%v, %v) should be %v but is %v", i, n, r, s)
		}
	}

	ConfirmExpand(C128List(), -1, 1, C128List(0))
	ConfirmExpand(C128List(), 0, 1, C128List(0))
	ConfirmExpand(C128List(), 1, 1, C128List(0))
	ConfirmExpand(C128List(), 0, 2, C128List(0, 0))

	ConfirmExpand(C128List(0, 1, 2), -1, 2, C128List(0, 0, 0, 1, 2))
	ConfirmExpand(C128List(0, 1, 2), 0, 2, C128List(0, 0, 0, 1, 2))
	ConfirmExpand(C128List(0, 1, 2), 1, 2, C128List(0, 0, 0, 1, 2))
	ConfirmExpand(C128List(0, 1, 2), 2, 2, C128List(0, 1, 0, 0, 2))
	ConfirmExpand(C128List(0, 1, 2), 3, 2, C128List(0, 1, 2, 0, 0))
	ConfirmExpand(C128List(0, 1, 2), 4, 2, C128List(0, 1, 2, 0, 0))
}

func TestC128SliceReverse(t *testing.T) {
	sxp := C128List(1, 2, 3, 4, 5)
	rxp := C128List(5, 4, 3, 2, 1)
	sxp.Reverse()
	if !rxp.Equal(sxp) {
		t.Fatalf("Reversal failed: %v", sxp)
	}
}

func TestC128SliceCar(t *testing.T) {
	ConfirmCar := func(s *C128Slice, r complex128) {
		n := s.Car().(complex128)
		if ok := n == r; !ok {
			t.Fatalf("head should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCar(C128List(1, 2, 3), 1)
}

func TestC128SliceCdr(t *testing.T) {
	ConfirmCdr := func(s, r *C128Slice) {
		if n := s.Cdr(); !n.Equal(r) {
			t.Fatalf("tail should be '%v' but is '%v'", r, n)
		}
	}
	ConfirmCdr(C128List(1, 2, 3), C128List(2, 3))
}

func TestC128SliceRplaca(t *testing.T) {
	ConfirmRplaca := func(s *C128Slice, v interface{}, r *C128Slice) {
		if s.Rplaca(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplaca(C128List(1, 2, 3, 4, 5), complex128(0), C128List(0, 2, 3, 4, 5))
}

func TestC128SliceRplacd(t *testing.T) {
	ConfirmRplacd := func(s *C128Slice, v interface{}, r *C128Slice) {
		if s.Rplacd(v); !s.Equal(r) {
			t.Fatalf("slice should be '%v' but is '%v'", r, s)
		}
	}
	ConfirmRplacd(C128List(1, 2, 3, 4, 5), nil, C128List(1))
	ConfirmRplacd(C128List(1, 2, 3, 4, 5), complex128(10), C128List(1, 10))
	ConfirmRplacd(C128List(1, 2, 3, 4, 5), C128List(5, 4, 3, 2), C128List(1, 5, 4, 3, 2))
	ConfirmRplacd(C128List(1, 2, 3, 4, 5, 6), C128List(2, 4, 8, 16), C128List(1, 2, 4, 8, 16))
}

func TestC128SliceSetIntersection(t *testing.T) {
	ConfirmSetIntersection := func(s, o, r *C128Slice) {
		x := s.SetIntersection(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetIntersection(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetIntersection(C128List(1, 2, 3), C128List(), C128List())
	ConfirmSetIntersection(C128List(1, 2, 3), C128List(1), C128List(1))
	ConfirmSetIntersection(C128List(1, 2, 3), C128List(1, 1), C128List(1))
	ConfirmSetIntersection(C128List(1, 2, 3), C128List(1, 2, 1), C128List(1, 2))
}

func TestC128SliceSetUnion(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *C128Slice) {
		x := s.SetUnion(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(C128List(1, 2, 3), C128List(), C128List(1, 2, 3))
	ConfirmSetUnion(C128List(1, 2, 3), C128List(1), C128List(1, 2, 3))
	ConfirmSetUnion(C128List(1, 2, 3), C128List(1, 1), C128List(1, 2, 3))
	ConfirmSetUnion(C128List(1, 2, 3), C128List(1, 2, 1), C128List(1, 2, 3))
}

func TestC128SliceSetDifference(t *testing.T) {
	ConfirmSetUnion := func(s, o, r *C128Slice) {
		x := s.SetDifference(*o)
		x.Sort()
		if !r.Equal(x) {
			t.Fatalf("%v.SetUnion(%v) should be %v but is %v", s, o, r, x)
		}
	}

	ConfirmSetUnion(C128List(1, 2, 3), C128List(), C128List(1, 2, 3))
	ConfirmSetUnion(C128List(1, 2, 3), C128List(1), C128List(2, 3))
	ConfirmSetUnion(C128List(1, 2, 3), C128List(1, 1), C128List(2, 3))
	ConfirmSetUnion(C128List(1, 2, 3), C128List(1, 2, 1), C128List(3))
}

func TestC128SliceFind(t *testing.T) {
	ConfirmFind := func(s *C128Slice, v complex128, i int) {
		if x, ok := s.Find(v); !ok || x != i {
			t.Fatalf("%v.Find(%v) should be %v but is %v", s, v, i, x)
		}
	}

	ConfirmFind(C128List(0, 1, 2, 3, 4), 0, 0)
	ConfirmFind(C128List(0, 1, 2, 3, 4), 1, 1)
	ConfirmFind(C128List(0, 1, 2, 4, 3), 2, 2)
	ConfirmFind(C128List(0, 1, 2, 4, 3), 3, 4)
	ConfirmFind(C128List(0, 1, 2, 4, 3), 4, 3)
}

func TestC128SliceFindN(t *testing.T) {
	ConfirmFindN := func(s *C128Slice, v complex128, n int, i *ISlice) {
		if x := s.FindN(v, n); !x.Equal(i) {
			t.Fatalf("%v.Find(%v, %v) should be %v but is %v", s, v, n, i, x)
		}
	}

	ConfirmFindN(C128List(1, 0, 1, 0, 1), 2, 3, IList())
	ConfirmFindN(C128List(1, 0, 1, 0, 1), 1, 0, IList(0, 2, 4))
	ConfirmFindN(C128List(1, 0, 1, 0, 1), 1, 1, IList(0))
	ConfirmFindN(C128List(1, 0, 1, 0, 1), 1, 2, IList(0, 2))
	ConfirmFindN(C128List(1, 0, 1, 0, 1), 1, 3, IList(0, 2, 4))
	ConfirmFindN(C128List(1, 0, 1, 0, 1), 1, 4, IList(0, 2, 4))
}

func TestC128SliceKeepIf(t *testing.T) {
	ConfirmKeepIf := func(s *C128Slice, f interface{}, r *C128Slice) {
		if s.KeepIf(f); !r.Equal(s) {
			t.Fatalf("KeepIf(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), complex128(0), C128List(0, 0, 0))
	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), complex128(1), C128List(1))
	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), complex128(6), C128List())

	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(0) }, C128List(0, 0, 0))
	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(1) }, C128List(1))
	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(6) }, C128List())

	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(0) }, C128List(0, 0, 0))
	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(1) }, C128List(1))
	ConfirmKeepIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(6) }, C128List())
}

func TestC128SliceReverseEach(t *testing.T) {
	var count	complex128
	count = 9
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(i interface{}) {
		if i != count {
			t.Fatalf("0: element %v erroneously reported as %v", count, i)
		}
		count--
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(index int, i interface{}) {
		if index != int(real(i.(complex128))) {
			t.Fatalf("1: element %v erroneously reported as %v", index, i)
		}
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(key, i interface{}) {
		if complex(float64(key.(int)), 0) != i {
			t.Fatalf("2: element %v erroneously reported as %v", key, i)
		}
	})

	count = 9
	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(i complex128) {
		if i != count {
			t.Fatalf("3: element %v erroneously reported as %v", count, i)
		}
		count--
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(index int, i complex128) {
		if int(real(i)) != index {
			t.Fatalf("4: element %v erroneously reported as %v", index, i)
		}
	})

	C128List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9).ReverseEach(func(key interface{}, i complex128) {
		if key.(int) != int(real(i)) {
			t.Fatalf("5: element %v erroneously reported as %v", key, i)
		}
	})
}

func TestC128SliceReplaceIf(t *testing.T) {
	ConfirmReplaceIf := func(s *C128Slice, f, v interface{}, r *C128Slice) {
		if s.ReplaceIf(f, v); !r.Equal(s) {
			t.Fatalf("ReplaceIf(%v, %v) should be %v but is %v", f, v, r, s)
		}
	}

	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), complex128(0), complex128(1), C128List(1, 1, 1, 3, 1, 5))
	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), complex128(1), complex128(0), C128List(0, 0, 0, 3, 0, 5))
	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), complex128(6), complex128(0), C128List(0, 1, 0, 3, 0, 5))

	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(0) }, complex128(1), C128List(1, 1, 1, 3, 1, 5))
	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(1) }, complex128(0), C128List(0, 0, 0, 3, 0, 5))
	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(6) }, complex128(0), C128List(0, 1, 0, 3, 0, 5))

	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(0) }, complex128(1), C128List(1, 1, 1, 3, 1, 5))
	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(1) }, complex128(0), C128List(0, 0, 0, 3, 0, 5))
	ConfirmReplaceIf(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(6) }, complex128(0), C128List(0, 1, 0, 3, 0, 5))
}

func TestC128SliceReplace(t *testing.T) {
	ConfirmReplace := func(s *C128Slice, v interface{}) {
		if s.Replace(v); !s.Equal(v) {
			t.Fatalf("Replace() should be %v but is %v", s, v)
		}
	}

	ConfirmReplace(C128List(0, 1, 2, 3, 4, 5), C128List(9, 8, 7, 6, 5))
	ConfirmReplace(C128List(0, 1, 2, 3, 4, 5), C128Slice{ 9, 8, 7, 6, 5 })
	ConfirmReplace(C128List(0, 1, 2, 3, 4, 5), &[]complex128{ 9, 8, 7, 6, 5 })
	ConfirmReplace(C128List(0, 1, 2, 3, 4, 5), []complex128{ 9, 8, 7, 6, 5 })
}

func TestC128SliceSelect(t *testing.T) {
	ConfirmSelect := func(s *C128Slice, f interface{}, r *C128Slice) {
		if x := s.Select(f); !r.Equal(x) {
			t.Fatalf("Select(%v) should be %v but is %v", f, r, s)
		}
	}

	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), complex128(0), C128List(0, 0, 0))
	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), complex128(1), C128List(1))
	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), complex128(6), C128List())

	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(0) }, C128List(0, 0, 0))
	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(1) }, C128List(1))
	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), func(x interface{}) bool { return x == complex128(6) }, C128List())

	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(0) }, C128List(0, 0, 0))
	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(1) }, C128List(1))
	ConfirmSelect(C128List(0, 1, 0, 3, 0, 5), func(x complex128) bool { return x == complex128(6) }, C128List())
}

func TestC128SliceUniq(t *testing.T) {
	ConfirmUniq := func(s, r *C128Slice) {
		if s.Uniq(); !r.Equal(s) {
			t.Fatalf("Uniq() should be %v but is %v", r, s)
		}
	}

	ConfirmUniq(C128List(0, 0, 0, 0, 0, 0), C128List(0))
	ConfirmUniq(C128List(0, 1, 0, 3, 0, 5), C128List(0, 1, 3, 5))
}

func TestC128SliceShuffle(t *testing.T) {
	ConfirmShuffle := func(s, r *C128Slice) {
		if s.Shuffle(); s.Equal(r) {
			t.Fatalf("%v.Shuffle() should change order of elements", s)
		}
		if s.Sort(); !s.Equal(r) {
			t.Fatalf("Shuffle() when sorted should be %v but is %v", r, s)
		}
	}

	ConfirmShuffle(C128List(0, 1, 2, 3, 4, 5), C128List(0, 1, 2, 3, 4, 5))
}

func TestC128SliceValuesAt(t *testing.T) {
	ConfirmValuesAt := func(s *C128Slice, i []int, r *C128Slice) {
		if x := s.ValuesAt(i...); !r.Equal(x) {
			t.Fatalf("%v.ValuesAt(%v) should be %v but is %v", s, i, r, x)
		}
	}

	ConfirmValuesAt(C128List(0, 1, 2, 3, 4, 5), []int{}, C128List())
	ConfirmValuesAt(C128List(0, 1, 2, 3, 4, 5), []int{ 0, 1 }, C128List(0, 1))
	ConfirmValuesAt(C128List(0, 1, 2, 3, 4, 5), []int{ 0, 3 }, C128List(0, 3))
	ConfirmValuesAt(C128List(0, 1, 2, 3, 4, 5), []int{ 0, 3, 4, 3 }, C128List(0, 3, 4, 3))
}

func TestC128SliceInsert(t *testing.T) {
	ConfirmInsert := func(s *C128Slice, n int, v interface{}, r *C128Slice) {
		if s.Insert(n, v); !r.Equal(s) {
			t.Fatalf("Insert(%v, %v) should be %v but is %v", n, v, r, s)
		}
	}

	ConfirmInsert(C128List(), 0, complex128(0), C128List(0))
	ConfirmInsert(C128List(), 0, C128List(0), C128List(0))
	ConfirmInsert(C128List(), 0, C128List(0, 1), C128List(0, 1))

	ConfirmInsert(C128List(0), 0, complex128(1), C128List(1, 0))
	ConfirmInsert(C128List(0), 0, C128List(1), C128List(1, 0))
	ConfirmInsert(C128List(0), 1, complex128(1), C128List(0, 1))
	ConfirmInsert(C128List(0), 1, C128List(1), C128List(0, 1))

	ConfirmInsert(C128List(0, 1, 2), 0, complex128(3), C128List(3, 0, 1, 2))
	ConfirmInsert(C128List(0, 1, 2), 1, complex128(3), C128List(0, 3, 1, 2))
	ConfirmInsert(C128List(0, 1, 2), 2, complex128(3), C128List(0, 1, 3, 2))
	ConfirmInsert(C128List(0, 1, 2), 3, complex128(3), C128List(0, 1, 2, 3))

	ConfirmInsert(C128List(0, 1, 2), 0, C128List(3, 4), C128List(3, 4, 0, 1, 2))
	ConfirmInsert(C128List(0, 1, 2), 1, C128List(3, 4), C128List(0, 3, 4, 1, 2))
	ConfirmInsert(C128List(0, 1, 2), 2, C128List(3, 4), C128List(0, 1, 3, 4, 2))
	ConfirmInsert(C128List(0, 1, 2), 3, C128List(3, 4), C128List(0, 1, 2, 3, 4))
}