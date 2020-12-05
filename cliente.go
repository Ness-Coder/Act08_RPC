package main

import (
	"fmt"
	"net/rpc"
)
type Alumno struct{
	NombreEstudiante string
	Materia string
	Calificacion float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	var al Alumno

	for {
		fmt.Println("\n1) Agregar calificacion")
		fmt.Println("2) Imprimir promedio de Alumno")
		fmt.Println("3) Imprimir promedio general")
		fmt.Println("4) Imprimir promedio de una materia")
		fmt.Println("5) Salir")
		fmt.Print("Opcion:")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var result string
			var nombre string
			var materia string
			var calificacion float64
			fmt.Println("Agregar")
			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)
			al.NombreEstudiante = nombre

			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			al.Materia = materia

			fmt.Print("Calificacion: ")
			fmt.Scanln(&calificacion)
			al.Calificacion =calificacion	
			
			
			err = c.Call("Server.AgregarA", al, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 2:
			var nombre string
			var result float64
			fmt.Println("Promedio de Alumno")
			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)

			err = c.Call("Server.CalPA", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Print("Promedio:") 
				fmt.Println(result)
			}
		case 3:
			var result float64
			fmt.Println("Promedio General")
			err = c.Call("Server.CalPG",  al, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Print("Promedio General:") 
				fmt.Println(result)
			}
		case 4:
			var materia string
			var result float64
			fmt.Println("Promedio Materia")
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)

			err = c.Call("Server.CalPM",  materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 5:
			return
		}
	}
}

func main() {
	client()
}