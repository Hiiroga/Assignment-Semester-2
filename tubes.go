package main

import (
	"fmt"
)

type dataMahasiswa struct{
	nama,jurusan,status string
	tes,umur int
}
const NMAX = 1000
type tabMABA[NMAX] dataMahasiswa

var jumlahMaba int
var maba tabMABA
var kkm int = 75

func main(){
	var user int
	var keluar bool = false

	for !keluar{
		fmt.Println("---------------------------------------------")
		fmt.Println("HELLO SELAMAT DATANG  DI UNIVERSITAS TELKOM")
		fmt.Println("---------------------------------------------")
		fmt.Println("LOGIN SEBAGAI:")
		fmt.Println("1.ADMIN")
		fmt.Println("2.MABA")
		fmt.Println("3.EXIT")
		fmt.Print("pilihan: ")
		fmt.Scan(&user)
		if user == 1{
			admin()
		}else if user == 2 {
			daftar()
		}else if user == 3{
			fmt.Println("TERIMA KASIH")
			keluar = true
		}else{
			fmt.Println("MASUKKAN SESUAI PILIHAN !")
		}
	}
}

func daftar(){
	var user int
	var keluar bool = false
	for !keluar{
		fmt.Println("---------------------------------------------")
		fmt.Println("MENU")
		fmt.Println("1.DAFTAR")
		fmt.Println("2.LIHAT HASIL")
		fmt.Println("3.KEMBALI")
		fmt.Print("pilihan: ")
		fmt.Scan(&user)
		if user == 1{
			pendaftaraan()
		}else if user == 2 {
			lihatHasil()
		}else if user == 3{
			keluar = true
		}else{
			fmt.Println("MASUKKAN SESUAI PILIHAN !")
		}
	}
}

func pendaftaraan(){
	var n int
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELAMAT DATANG DI PENDAFTARAAN MAHASISWA BARU")
	fmt.Println("-----------------------------------------------")
	jumlahMaba++
	n = jumlahMaba - 1
	fmt.Printf("SILAHKAN ISI NAMA DAN JURUSAN PENDAFTAR KE %d\n",n+1)
	fmt.Print("NAMA = ")
	fmt.Scan(&maba[n].nama)
	fmt.Print("JURUSAN = ")
	fmt.Scan(&maba[n].jurusan)
	fmt.Print("UMUR = ")
	fmt.Scan(&maba[n].umur)
	fmt.Println()
}

func lihatHasil(){
	var nama,jurusan string
	var hasil bool = false
	var temp,umur int
	fmt.Println("SILAHKAN KETIKAN NAMA DAN JURUSAN")
	fmt.Print("NAMA = ")
	fmt.Scan(&nama)
	fmt.Print("JURUSAN = ")
	fmt.Scan(&jurusan)
	fmt.Print("UMUR = ")
	fmt.Scan(&umur)
	for i:=0;i<jumlahMaba ;i++{
		if maba[i].nama == nama && maba[i].jurusan==jurusan && maba[i].umur == umur{
			hasil = true
			temp = i
		}
	}
	if hasil == true{
		if maba[temp].tes != 0{
			fmt.Printf("HASIL TES ANDA %d dan anda dinyatakan %s\n",maba[temp].tes,maba[temp].status)
		}else{
			fmt.Println("TUNGGU SAMPAI HASIL KELUAR")
		}
	}else{
		fmt.Println("DATA TIDAK ADA, MASUKKAN NAMA DAN JURUSAN DENGAN BENAR")
	}
}


func admin(){
	var menu int
	var keluar bool = false

	for !keluar{
		fmt.Println("------")
		fmt.Println(" MENU ")
		fmt.Println("------")
		fmt.Println("1.EDIT DATA")
		fmt.Println("2.EDIT NILAI")
		fmt.Println("3.DATA MAHASISWA")
		fmt.Println("4.KEMBALI")
		fmt.Print("pilih menu: ")
		fmt.Scan(&menu)
		if menu == 1{
			editData(&maba,&jumlahMaba)
		}else if menu == 2{
			editNilai()
		}else if menu == 3{
			data()
		}else if menu == 4 {
			keluar = true
		}else{
			fmt.Println("MASUKKAN SESUAI PILIHAN !")
		}
	}
}

// ----------------------- EDIT DATA ----------------------------------
func editData(maba *tabMABA, jumlahMaba *int){
	var keluar bool = false
	var n int
	for !keluar{
		fmt.Println("--------------------")
		fmt.Println("1.TAMBAH DATA")
		fmt.Println("2.UBAH DATA")
		fmt.Println("3.HAPUS DATA")
		fmt.Println("4.KEMBALI")
		fmt.Println("--------------------")
		fmt.Print("pilih: ")
		fmt.Scan(&n)
		if n == 1{
			tambahData(maba,jumlahMaba)
		}else if n == 2{
			ubahData()
		}else if n == 3{
			hapusData()
		}else if n == 4{
			keluar = true
		}else{
			fmt.Println("MASUKKAN SESUAI PILIHAN !")
		}
	}
}

func tambahData(maba *tabMABA,jumlahMaba *int){
	var n,total int
	fmt.Printf("TOTAL DATA SEKARANG %d\n",*jumlahMaba)
	fmt.Print("JUMLAH DATA YANG INGIN DI TAMBAH = ")
	fmt.Scan(&n)
	fmt.Println("")
	total = *jumlahMaba + n
	for i:=*jumlahMaba ; i < total ; i++{
		fmt.Printf(" DATA KE %d\n",i+1)
		fmt.Print("NAMA = ")
		fmt.Scan(&maba[i].nama)
		fmt.Print("JURUSAN = ")
		fmt.Scan(&maba[i].jurusan)
		fmt.Print("UMUR = ")
		fmt.Scan(&maba[i].umur)
		fmt.Println()
	}
	*jumlahMaba = total
}

func ubahData(){
	var i,n,pilih int
	var keluar bool = false
	for !keluar{
		cetakData()
		fmt.Printf("PILIH DATA KE BERAPA YANG INGIN DI UBAH : ")
		fmt.Scan(&i)
		if i <= jumlahMaba{
			fmt.Println()
			fmt.Println("APA YANG INGIN DI UBAH : ")
			fmt.Println("1. NAMA")
			fmt.Println("2. JURUSAN")
			fmt.Println("3. NILAI TEST ")
			fmt.Println("4. STATUS ")
			fmt.Println("5. UMUR ")
			fmt.Print("1/2/3/4/5: ")
			fmt.Scan(&pilih)
			fmt.Println("")
			fmt.Printf("DATA KE %d\n",i)
			i = i - 1
			if pilih == 1{
				fmt.Print("NAMA = ")
				fmt.Scan(&maba[i].nama)
			}else if pilih == 2{
				fmt.Print("JURUSAN = ")
				fmt.Scan(&maba[i].jurusan)
			}else if pilih == 3{
				fmt.Print("NILAI TES = ")
				fmt.Scan(&maba[i].tes)
			}else if pilih == 4{
				fmt.Print("STATUS = ")
				fmt.Scan(&maba[i].status)
			}else if pilih == 5{
				fmt.Print("UMUR = ")
				fmt.Scan(&maba[i].umur)
			} else {
				fmt.Println("DATA TIDAK ADA ")
			}
			fmt.Print("UBAH LAGI(1) / TIDAK(2): ")
			fmt.Scan(&n)
			fmt.Println()
			if n == 2{
				keluar = true
			}
			if n == 1{
				keluar = false
			}else{
				fmt.Println("MASUKKAN SESUAI PILIHAN !")
			}
		}else{
			fmt.Println("DATA TIDAK ADA")
		}
	}
	
}

func hapusData(){
	var i int
	cetakData()
	fmt.Print("PILIH DATA KE BERAPA YANG DIHAPUS: ")
	fmt.Scan(&i)
	if i <= jumlahMaba && i >=1{
		i = i - 1
		maba[i].nama = ""
		maba[i].jurusan = ""
		maba[i].tes = 0
		for i < jumlahMaba-1{
			maba[i] = maba[i+1]
			i++
		}
		jumlahMaba--
	}else{
		fmt.Println("DATA TIDAK ADA")
	}
}


// ----------------------- EDIT NILAI ----------------------------------

func editNilai(){
	var keluar bool = false
	var n int
	for !keluar{
		fmt.Println("--------------------")
		fmt.Println("1.UBAH KKM")
		fmt.Println("2.INPUT NILAI TEST")
		fmt.Println("3.PERBARUI STATUS")
		fmt.Println("4.KEMBALI")
		fmt.Println("--------------------")
		fmt.Print("pilih: ")
		fmt.Scan(&n)
		if n == 1{
			fmt.Print("MASUKKAN NILAI KKM = ")
			fmt.Scan(&kkm)
			fmt.Println()
		}else if n == 2{
			inputNilai()	
		}else if n == 3{
			cekStatus()
			fmt.Println("STATUS SUDAH DI PERBARUI")	
		}else if n == 4{
			keluar = true
		}else{
			fmt.Println("MASUKKAN SESUAI PILIHAN !")
		}
	}
}

func inputNilai(){
	var i,n,nilai int
	var keluar bool = false
	for !keluar{
		cetakData()
		fmt.Printf("PILIH DATA KE BERAPA YANG INGIN DI INPUT : ")
		fmt.Scan(&i)
		if i <= jumlahMaba{
			fmt.Printf("DATA KE %d\n",i)
			i = i - 1
			fmt.Print("NILAI = ")
			fmt.Scan(&nilai)
			maba[i].tes = nilai 
			fmt.Print("INPUT LAGI(1) / TIDAK(2): ")
			fmt.Scan(&n)
			if n == 2{
				keluar = true
			}
			if n == 1{
				keluar = false
			}else{
				fmt.Println("MASUKKAN SESUAI PILIHAN !")
			}
			fmt.Println("--------------------------------------")
		}else{
			fmt.Println("DATA TIDAK ADA")
		}
	}
}

// ----------------------- DATA MAHASISWA ----------------------------------

func data() {
	var menu int
	var keluar bool = false

	for !keluar {
		fmt.Println("------")
		fmt.Println(" DATA ")
		fmt.Println("------")
		fmt.Println("1.TAMPILKAN DATA")
		fmt.Println("2.URUTKAN DATA DARI NILAI")
		fmt.Println("3.URUTKAN DATA DARI UMUR")
		fmt.Println("4.CARI DATA BERDASARKAN UMUR")
		fmt.Println("5.KEMBALI")
		fmt.Print("pilih menu: ")
		fmt.Scan(&menu)
		if menu == 1 {
			cetakData()
		} else if menu == 2 {
			urutData()
		} else if menu == 3 {
			urutDataUmur()
		} else if menu == 4 {
			cariDataUmur()
		} else if menu == 5 {
			keluar = true
		} else {
			fmt.Println("MASUKKAN SESUAI PILIHAN !")
		}
	}
}

func cetakData(){
	fmt.Println("----------------------------------------")
	fmt.Println("DATA MABA : ")
	for i:=0 ; i<jumlahMaba;i++{
		fmt.Printf("%d) %s %s %d %d %s\n",i+1,maba[i].nama,maba[i].jurusan,maba[i].tes,maba[i].umur,maba[i].status)
	}
	fmt.Println("----------------------------------------")
}

func cekStatus(){
	for i:=0;i<jumlahMaba;i++{
		if maba[i].tes >= kkm{
			maba[i].status = "LULUS"
		}else{
			maba[i].status = "TIDAK LULUS"
		}
	}
}


func urutData(){
	var pass,i,idx,temp int
	var temp1,temp2,temp3 string
	pass = 1
	for pass <= jumlahMaba-1{
		idx = pass - 1
		i = pass
		for i<jumlahMaba {
			if maba[idx].tes < maba[i].tes{
				idx = i
			}
			i++
		}
		temp = maba[pass-1].tes
		temp1 = maba[pass-1].nama
		temp2 = maba[pass-1].jurusan
		temp3 = maba[pass-1].status
		maba[pass-1].tes=maba[idx].tes
		maba[pass-1].nama=maba[idx].nama
		maba[pass-1].jurusan=maba[idx].jurusan
		maba[pass-1].status=maba[idx].status
		maba[idx].tes = temp
		maba[idx].nama = temp1
		maba[idx].jurusan = temp2
		maba[idx].status = temp3
		pass++
	}
	cetakData()
}

func urutDataUmur() {
    var i, j, temp, tempTes int
    var tempNama, tempJurusan, tempStatus string

    for i = 1; i < jumlahMaba; i++ {
        temp = maba[i].umur
        tempNama = maba[i].nama
        tempJurusan = maba[i].jurusan
		tempTes = maba[i].tes
        tempStatus = maba[i].status

        j = i - 1
        for j >= 0 && maba[j].umur > temp {
            maba[j+1].umur = maba[j].umur
            maba[j+1].nama = maba[j].nama
            maba[j+1].jurusan = maba[j].jurusan
			maba[j+1].tes = maba[j].tes
            maba[j+1].status = maba[j].status
            j = j - 1
        }
        maba[j+1].umur = temp
        maba[j+1].nama = tempNama
        maba[j+1].jurusan = tempJurusan
		maba[j+1].tes = tempTes
        maba[j+1].status = tempStatus
    }
    cetakData()
}

func binarySearchUmur(x int) int {
	urutDataUmur()
	left, right := 0, jumlahMaba-1
	for left <= right {
		mid := (left + right) / 2
		if maba[mid].umur == x {
			return mid
		}
		if maba[mid].umur < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func cariDataUmur() {
	var umur int
	fmt.Print("Masukkan umur yang dicari: ")
	fmt.Scan(&umur)
	idx := binarySearchUmur(umur)
	if idx != -1 {
		fmt.Printf("Mahasiswa dengan umur %d ditemukan:\n", umur)
		count := 1 

		i := idx
		for i >= 0 && maba[i].umur == umur {
			fmt.Printf("%d) %s %s %d %d %s\n", count, maba[i].nama, maba[i].jurusan, maba[i].tes, maba[i].umur, maba[i].status)
			i--
			count++
		}

		i = idx + 1
		for i < jumlahMaba && maba[i].umur == umur {
			fmt.Printf("%d) %s %s %d %d %s\n", count, maba[i].nama, maba[i].jurusan, maba[i].tes, maba[i].umur, maba[i].status)
			i++
			count++
		}
	} else {
		fmt.Println("Tidak ada mahasiswa dengan umur tersebut.")
	}
}