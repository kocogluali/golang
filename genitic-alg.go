package main

import (
	"os"
	"log"
	"bufio"
	"math/rand"

)

func kontrol(g int, slice []int) bool  {
	for _, v := range slice {
		if v == g {
			return true
		}
	}
	return false
}

func random(min,max,baslangic,son int) int {

	slice := make([]int, 0)
	for i:=baslangic;i<son; i++{
		deger :=rand.Intn(max-min)
		if b:=kontrol(deger,slice) ; b == false {
			slice=append(slice,deger)
		}else{
			i = i-1
		}

	}
	for _ , deger:=range slice{
		log.Println(deger)
	}
	return 0
}

func main()  {
	dosya, err :=os.Open("data.txt")
	if err !=nil{
		konum , _ :=os.Getwd() //konumu alıyoruz.
		log.Println("Burada böyle bir dosya yok ", konum)
	}
	defer dosya.Close()
	scanner := bufio.NewScanner(dosya)
	for scanner.Scan() {
		data:=scanner.Text()
		log.Println(data)

	}
	random(1,30,0,10)

}
