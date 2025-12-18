package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

/*
METHOD
Go tidak memiliki class. Namun, anda bisa mendefinisikan method pada tipe.
Sebuah method adalah sebuah fungsi dengan argumen khusus receiver.
receiver diletakkan pada bagian antara kata kunci func and nama method.

PENTING: METHOD HANYALAH SEBUAH FUNGSI DENGAN ARGUMEN
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ini penulisan fungsi versi 2
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method lanjutan: Dimana hanya bisa mendeklarasikan sebuah method dengan sebuah
// receiver yang tipenya didefinisikan di paket yang sama dengan method-nya.
func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

/*
POINTER RECEIVER (Dengan METHOD)
Dengan pointer-receiver method dapat mengubah nilai yang ditunjukan pada setiap receiver

pointer-receiver lebih umum digunakan daripada receiver dengan value.
*/
func (v *Vertex) ScaleMethod(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
POINTER DAN FUNGSI
*/
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// cara memanggil method
	fmt.Println(v.Abs())
	// cara memanggil fungsi dengan parameter
	fmt.Println(Abs2(v))
	// method lanjutan
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs3())

	/*
		Pada function jika terdapat pointer, maka saat memanggil function
		tersebut harus menyertakan tanda & pada variabel param pointernya.

		Namun pada method tanda (&) tidak perlu ditambahkan walaupun pada method terdapat pointer
	*/
	// POINTER - RECEIVER
	v.ScaleMethod(10)
	fmt.Println(v.Abs())

	// POINTER DAN FUNGSI
	ScaleFunc(&v, 10) // maka pada bagian scale param v tanda & juga dihapus
	fmt.Println(Abs2(v))

	// pada method
	vertex := Vertex{3, 4}
	vertex.ScaleMethod(2)
	ScaleFunc(&vertex, 10)

	// pada function
	p := &Vertex{4, 3}
	p.ScaleMethod(3)
	ScaleFunc(p, 8)

	fmt.Println(vertex, p)

	/*
		Begitu juga dengan penerimaan parameter:
	*/

}
