package main

import (
	"fmt"
	"math"
	"time"
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

	/*
			PENEGASAN TIPE
			penegasan tipe menyediakan akses ke isi interface di balik nilai konkritnya.

		CONTOH: t := i.(T)
		Perintah di atas menegaskan bahwa isi interface i menyimpan tipe konkrit T
		dan memberikan nilai T ke variabel t. Dan apabila interface tidak mengandung tipe T, akan memicu panic.

		Untuk memeriksa apakah sebuah isi interface benar mengandung tipe tertentu, dapat menggunakan
		t, ok := i.(T)
		dan akan mereturn nilai boolean.
	*/
	var interf2 interface{} = "hello"

	s := interf2.(string)
	fmt.Println(s)

	s, ok := interf2.(string)
	fmt.Println(s, ok)

	r, ok := interf2.(float64)
	fmt.Println(r, ok)

	//r = interf2.(float64) // panic
	//fmt.Println(r)

	/*
			SWITCH UNTUK PENEGASAN TIPE
			tipe switch adalah bentukan yang membolehkan beberapa penegasan tipe secara serial.

		Tipe switch sama seperti perintah switch biasa, tapi dengan nilai case mengacu pada tipe (bukan nilai),
		dan nilai case tersebut dibandingkan dengan tipe yang dikandung oleh isi interface yang diberikan.
	*/
	do(21)
	do("hello")
	do(true)

	/*
		STRINGER
		stringer adalah Interface yang ada dimanapun, dan didefinisikan oleh paket fmt.

		! STRINGER BERUPA METHOD STRING DIBAWAH !

		Sebuah Stringer adalah suatu tipe yang mendeskripsikan dirinya sendiri sebagai string.
		Paket fmt (dan banyak lainnya) menggunakan interface ini untuk mencetak nilai.
	*/
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	/*
		LATIHAN STRINGER
	*/
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	/*
			ERROR
			Program Go mengekspresikan keadaan error dengan nilai error.

			Tipe error adalah interface buatan mirip dengan fmt.Stringer.
			Seperti dengan fmt.Stringer, paket fmt mencari interface error saat mencetak nilai.

			Fungsi terkadang mengembalikan nilai error, dan kode yang memanggilnya harus menangani error
			dengan memeriksa apakah error bernilai nil.

		i, err := strconv.Atoi("42")
		if err != nil {
		    fmt.Printf("couldn't convert number: %v\n", err)
		    return
		}
		fmt.Println("Converted integer:", i)

			error yang nil menandakan sukses; error yang bukan-nil menandakan adanya kesalahan.
	*/
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// LATIHAN ERROR
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// SWITCH TIPE
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Dua kali %v adalah %v\n", v, v*2)
	case string:
		fmt.Printf("%q adalah %v bytes panjangnya\n", v, len(v))
	default:
		fmt.Printf("Saya tidak kenal dengan tipe %T!\n", v)
	}
}

// STRINGER
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Latihan Stinger
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// ERROR
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"tidak bekerja!",
	}
}

// LATIHAN ERROR
type ErrNegativeSqrt float64

// implementasi error
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Tidak bisa dipangkatkan (Sqrt) angka negatif : %v", float64(e))
}

// fungsi Sqrt dengan error
func Sqrt(x float64) (float64, error) {

	// Jika x negatif → return error
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	// Newton method
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}
