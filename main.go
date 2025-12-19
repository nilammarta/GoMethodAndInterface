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
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ini penulisan fungsi versi 2
func AbsFunc(v Vertex) float64 {
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

/*
INTERFACE
*/
type Abser interface {
	Abs() float64
}

type I interface {
	Method()
}

type T struct {
	S string
}

// Method berikut berarti type T mengimplementasikan interface I,
// tapi kita tidak perlu secara eksplisit mendeklarasikannya.
func (t T) Method() {
	fmt.Println(t.S)
}

// Nilai Interface
type Interf interface {
	M()
}

func (t *T) M() {
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

type F float64

func describe(i Interf) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
INTERFACE NILAI NIL
*/
type nilIntrf interface {
	nilM()
}

type T2 struct {
	S string
}

func (t *T2) nilM() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe2(i nilIntrf) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
INTERFACE YG NIL
*/
type InterNil interface {
	MNil()
}

func describe3(i InterNil) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	/*
		METHOD
	*/
	v := Vertex{3, 4}
	// cara memanggil method
	fmt.Println(v.AbsMethod())
	// cara memanggil fungsi dengan parameter
	fmt.Println(AbsFunc(v))
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
	fmt.Println(v.AbsMethod())

	// POINTER DAN FUNGSI
	ScaleFunc(&v, 10) // maka pada bagian scale param v tanda & juga dihapus
	fmt.Println(AbsFunc(v))

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
		* Fungsi yang menerima argumen nilai harus mendapatkan sebuah nnilai dari tipe yang spesifik.
		* Sedangkan pada method dengan value-receiver bisa menerima sebuah nilai atau sebuha pointer sebagai receivernya
	*/
	// pada function:
	fmt.Println(AbsFunc(v)) // OK
	//fmt.Println(AbsFunc(&v)) // Compile error!

	// pada method:
	fmt.Println(v.AbsMethod()) // OK
	pointer := &v
	fmt.Println(pointer.AbsMethod()) // OK

	/*
			POINTER / RECEIVER :
		Alasan menggunakan pointer: Pertama agar method dapat mengubah nilai yang ditunjuk oleh receiver.
		dan Kedua untuk menghindari menyalin nilai setiap kali method dipanggil.
	*/
	receiver := &Vertex{3, 4}
	fmt.Printf("Sebelum scaling: %+v, Abs: %v\n", receiver, receiver.AbsMethod())
	receiver.ScaleMethod(5)
	fmt.Printf("Setelah scaling: %+v, Abs: %v\n", receiver, receiver.AbsMethod())

	/*
		INTERFACE
		Sebuah type interface didefinisikan sebagai sebuah kumpulan method penanda.
		Nilai dari tipe interface dapat menyimpan nilai apapun yang mengimplementasikan method tersebut.

		Interface tidak peduli struktur data, hanya peduli:
		method apa yang ada
		signature method tersebut
	*/
	// inisiasi function
	//var a Abser
	//varF := MyFloat(-math.Sqrt2)
	//varV := Vertex{3, 4}

	//a = varF  // a MyFloat mengimplementasikan Abser
	//a = &varV // a *Vertex mengimplementasikan Abser
	//
	//// Pada baris berikut, v adalah sebuah Vertex (bukan *Vertex)
	//// dan TIDAK mengimplementasikan Abser.
	//a = varV
	//
	//fmt.Println(a.Abs())

	/*
			Interface dipenuhi secara IMPLISIT dan BUKAN eksplisit, dimana tidak ada perintah implements.

		Melainkan, mengimplementasikan sebuah interface dengan mengimplementasikan method-methodnya.
	*/
	var i I = T{"hello"}
	i.Method()

	/*
			NILAI INTERFACE
		Isi interface dapat dibayangkan sebagai sebuah pasangan nilai dan sebuah tipe (nilai, tipe)
		Memanggil suatu method terhadap suatu interface akan mengeksekusi method
		dengan nama yang sama pada tipe yang dipegangnya
	*/
	var i2 Interf

	i2 = &T{"Hello"}
	describe(i2)
	i2.M()

	i2 = F(math.Pi)
	describe(i2)
	i2.M()

	/*
		INTERFACE NILAI NIL
		Jika nilai sebenarnya dari interface itu sendiri adalah nil,
		maka method akan dipanggil dengan receiver bernilai nil.
	*/
	var i3 nilIntrf

	var t *T2
	i3 = t
	describe2(i3)
	i3.nilM()

	i3 = &T2{"hello"}
	describe2(i3)
	i3.nilM()

	// JIka sebuah method pada interface yang nill maka akan terjadi error
	var inter InterNil
	describe3(inter)
	// inter.MNil() // cb running ini untuk mengetahui errornya

	/*
		INTERFACE KOSONG
		Interface Kosong merupakan tipe interface yang tidak memiliki method
		interface{}

		Sebuah interface kosong bisa menyimpan nilai dari tipe apapun.
	*/
	var interf interface{}
	describe4(interf)

	interf = 42
	describe4(interf)

	interf = "hello"
	describe4(interf)
}

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
