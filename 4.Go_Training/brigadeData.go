package main

import(
	"fmt"
)
func main(){
	employees:=[]string{
		"Suresh","Pooja","Sowmya","Biswa","Banu","Selva","Geo","Megha","Manjunath","Giri","Uthra","Muthya","Anurag","Kuldeep","Nikhila",
	}
	platformBrigade:=employees[0:5]
	sustainingBrigade:=employees[5:10]
	devicesBrigade:=employees[10:15]
	fmt.Printf("\nPlatform brigade members are: \n")
	for _,v:=range platformBrigade {
		fmt.Printf("%v\t", v)
	}
	fmt.Printf("\nSustaining Team brigade members are: \n")
	for _,v:=range sustainingBrigade {
		fmt.Printf("%v\t", v)
	}
	fmt.Printf("\nDevices brigade members are: \n")
	for _,v:=range devicesBrigade {
		fmt.Printf("%v\t", v)
	}
}

