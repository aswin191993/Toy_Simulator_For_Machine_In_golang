package main

import "fmt"
import "strings"
import "./subpgm"

var reg [10] int

func subfun(a,b int){
	op:=a
	if(op==1||op==2||op==3){
		loadregister(a,b)
	}else if(op==5 || op==6 || op==7){
		operation(a,b)
	}	
}

func loadregister(optn,val int){
	fmt.Println("Register inputed")
	rval:= val%10000
	regnum:= rval/1000
	rvalu:= val%1000
	if (optn==3){
		reg[regnum]=rvalu
	}
}

func operation(obj,opval int){
	fmt.Println("Operations:\n")
	regtemp:=opval%10000
	reg1:=regtemp/1000
	reg2:=opval%10
	if(obj==5){
		reg[reg1]=reg[reg1]+reg[reg2]	
	}else if(obj==6){
		reg[reg1]=reg[reg1]-reg[reg2]	
	}else if(obj==7){
		reg[reg1]=reg[reg1]/reg[reg2]	
	}else if(obj==8){
		reg[reg1]=reg[reg1]*reg[reg2]	
	}
}

func display(ival,dval int){
	fmt.Printf("10%d:%d",ival,dval)
	for i:= 0;reg[i]!=0;i++{
		t:=(30+i)*1000+reg[i]
		fmt.Printf("\nRegister%d value:%d\n",i,t)
	}
}

func main(){
	k:=subpgm.Readfile()
	lines := strings.Split(string(k), "\n")
	var i int
	fmt.Println("###########################\n\tSIMULATOR\n###########################")

	for j:= 0;j<len(lines);j++{
		fmt.Println(lines[j])
		if _, err := fmt.Sscanf(lines[j], "%6d", &i); err == nil {
		}
		k:=i/10000
		subfun(k,i)
		display(j,i)
	}
}
/*
000000       Halt
01rmmm       Load register r with contents of address mmm.
02rmmm       Store the contents of register r at address mmm.
03rnnn       Load register r with the number nnn.
04r00s       Load register r with the memory word addressed by register s.
05r00s       Add contents of register s to register r
06r00s       Sub contents of register s from register r
07r00s       Mul contents of register r by register s
08r00s       Div contents of register r by register s
100mmm       Jump to location mmm
11rmmm       Jump to location mmm if register r is zero
*/
