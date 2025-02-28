package main

import "fmt"

func main() {
	Patinhos(10)
}

func Patinhos(patinhos int){
	
	frase1 := " patinhos foram passear,\n além das montanhas para brincar.\n A mamãe gritou 'Quack Quack Quack Quack!',\n mas só "
	frase2 := " patinhos voltaram de lá.\n\n\n"
	
	for ; patinhos != 0; patinhos--{
		
		var patinhosQueVoltaramDeLa int = patinhos - 1; 
		
		if patinhos != 1{
			fmt.Printf("%d%s%d%s", patinhos, frase1, patinhosQueVoltaramDeLa, frase2)
		}

		if patinhosQueVoltaramDeLa == 1{
			frase2 = " patinho voltou de lá.\n\n\n"
			var patinhoQueVoltouDeLa string;
			frase1 = " patinho foi passear,\n além das montanhas para brincar.\n A mamãe gritou 'Quack Quack Quack Quack!',\n mas nenhum"
			fmt.Printf("%d%s%s%s", patinhos, frase1, patinhoQueVoltouDeLa, frase2)
		}
	} 
}