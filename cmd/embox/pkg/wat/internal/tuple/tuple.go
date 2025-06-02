// While go don't have variadic type parameters ( https://github.com/golang/go/issues/66651 )
package tuple

type Tuple interface {
	Size() int
}

func New2[T1, T2 any](m1 T1, m2 T2) Of2[T1, T2] {
	return Of2[T1, T2]{m1, m2}
}

func Default2[T1, T2 any]() Of2[T1, T2] {
	return Of2[T1, T2]{}
}

type Of2[T1, T2 any] struct {
	M1 T1
	M2 T2
}

func (t *Of2[T1, T2]) Size() int {
	return 2
}

func New3[T1, T2, T3 any](m1 T1, m2 T2, m3 T3) Of3[T1, T2, T3] {
	return Of3[T1, T2, T3]{m1, m2, m3}
}

func Default3[T1, T2, T3 any]() Of3[T1, T2, T3] {
	return Of3[T1, T2, T3]{}
}

type Of3[T1, T2, T3 any] struct {
	M1 T1
	M2 T2
	M3 T3
}

func (t *Of3[T1, T2, T3]) Size() int {
	return 3
}

func New4[T1, T2, T3, T4 any](m1 T1, m2 T2, m3 T3, m4 T4) Of4[T1, T2, T3, T4] {
	return Of4[T1, T2, T3, T4]{m1, m2, m3, m4}
}

func Default4[T1, T2, T3, T4 any]() Of4[T1, T2, T3, T4] {
	return Of4[T1, T2, T3, T4]{}
}

type Of4[T1, T2, T3, T4 any] struct {
	M1 T1
	M2 T2
	M3 T3
	M4 T4
}

func (t *Of4[T1, T2, T3, T4]) Size() int {
	return 4
}

func New5[T1, T2, T3, T4, T5 any](m1 T1, m2 T2, m3 T3, m4 T4, m5 T5) Of5[T1, T2, T3, T4, T5] {
	return Of5[T1, T2, T3, T4, T5]{m1, m2, m3, m4, m5}
}

func Default5[T1, T2, T3, T4, T5 any]() Of5[T1, T2, T3, T4, T5] {
	return Of5[T1, T2, T3, T4, T5]{}
}

type Of5[T1, T2, T3, T4, T5 any] struct {
	M1 T1
	M2 T2
	M3 T3
	M4 T4
	M5 T5
}

func (t *Of5[T1, T2, T3, T4, T5]) Size() int {
	return 5
}

func New6[T1, T2, T3, T4, T5, T6 any](m1 T1, m2 T2, m3 T3, m4 T4, m5 T5, m6 T6) Of6[T1, T2, T3, T4, T5, T6] {
	return Of6[T1, T2, T3, T4, T5, T6]{m1, m2, m3, m4, m5, m6}
}

func Default6[T1, T2, T3, T4, T5, T6 any]() Of6[T1, T2, T3, T4, T5, T6] {
	return Of6[T1, T2, T3, T4, T5, T6]{}
}

type Of6[T1, T2, T3, T4, T5, T6 any] struct {
	M1 T1
	M2 T2
	M3 T3
	M4 T4
	M5 T5
	M6 T6
}

func (t *Of6[T1, T2, T3, T4, T5, T6]) Size() int {
	return 6
}

func New7[T1, T2, T3, T4, T5, T6, T7 any](m1 T1, m2 T2, m3 T3, m4 T4, m5 T5, m6 T6, m7 T7) Of7[T1, T2, T3, T4, T5, T6, T7] {
	return Of7[T1, T2, T3, T4, T5, T6, T7]{m1, m2, m3, m4, m5, m6, m7}
}

func Default7[T1, T2, T3, T4, T5, T6, T7 any]() Of7[T1, T2, T3, T4, T5, T6, T7] {
	return Of7[T1, T2, T3, T4, T5, T6, T7]{}
}

type Of7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	M1 T1
	M2 T2
	M3 T3
	M4 T4
	M5 T5
	M6 T6
	M7 T7
}

func (t *Of7[T1, T2, T3, T4, T5, T6, T7]) Size() int {
	return 7
}

func New8[T1, T2, T3, T4, T5, T6, T7, T8 any](m1 T1, m2 T2, m3 T3, m4 T4, m5 T5, m6 T6, m7 T7, m8 T8) Of8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Of8[T1, T2, T3, T4, T5, T6, T7, T8]{m1, m2, m3, m4, m5, m6, m7, m8}
}

func Default8[T1, T2, T3, T4, T5, T6, T7, T8 any]() Of8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Of8[T1, T2, T3, T4, T5, T6, T7, T8]{}
}

type Of8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	M1 T1
	M2 T2
	M3 T3
	M4 T4
	M5 T5
	M6 T6
	M7 T7
	M8 T8
}

func (t *Of8[T1, T2, T3, T4, T5, T6, T7, T8]) Size() int {
	return 8
}
