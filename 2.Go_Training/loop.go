package main

import (
	"fmt"
)

func main(){
	employeeName:="Pooja"
	employeeDesignation:=ifEmployee(employeeName)
	if employeeName== employeeDesignation{
		fmt.Printf("If loop-Employee name is %s and employee has no designation specified", employeeName)
	}else{
	fmt.Printf("If loop-Designation of employee %s is %s", employeeName,employeeDesignation)
	}
	employeeDesignation=switchEmployee(employeeName)
	if employeeName== employeeDesignation{
		fmt.Printf("\nSwitch loop-Employee name is %s and employee has no designation specified", employeeName)
	}else{
		fmt.Printf("\nSwitch loop-Designation of employee %s is %s", employeeName,employeeDesignation)
	}	
}

 func ifEmployee(employeename string)string{
	var designation string
	if employeename == "Suresh" {
		designation = "Manager"
	} else if employeename == "Asutosh" {
		designation = "Automation Engineer"
	} else {
		designation = employeename
	}
	return designation
 }
 func switchEmployee(employeename string)string{
	var designation string
	switch employeename {
	case "Suresh":
		designation = "Manager"
	case "Asutosh":
		designation = "Automation Engineer"
	default:
		designation = employeename
	}
	return designation
}