package main

import (
	//"errors"
	"fmt"
	"net"
	"net/rpc"
)
type Alumno struct{
	NombreEstudiante string
	Materia string
	Calificacion float64
}

type Server struct{
	Materias map[string]map[string]float64
	Alumnos map[string]map[string]float64
}

func (this *Server) AgregarA(al Alumno,  reply *string) error {
	
	if _, ok := this.Alumnos[al.NombreEstudiante]; !ok {//condicion checa si el alumno no existe en la mapa
	
		cali := make(map[string]float64)
		cali[al.Materia] = al.Calificacion
		this.Alumnos[al.NombreEstudiante] = cali
		*reply = "Alumno guardado"
		
	}else{//Setencias por si ocurre una modificacion o añadimiento y/o por si agrega alguien que ya este 
		if _, ok := this.Alumnos[al.NombreEstudiante][al.Materia]; ok { 
			*reply = "Error, El alumno con la materia ya existe"
		} else {
			this.Alumnos[al.NombreEstudiante][al.Materia] = al.Calificacion 
			*reply = "Se modifico"

		}
	}	
	return nil
}

func (this *Server) CalPA(nombre string, reply *float64) error {
	var promedio float64
	var suma float64

	
	for i:= range this.Alumnos[nombre]{
		suma+= this.Alumnos[nombre][i]
	}
	promedio = suma/float64(len(this.Alumnos[nombre]))
	*reply = promedio
	return nil
}
func (this *Server) CalPG( al Alumno,  reply *float64) error {//esta bien
	var sumInd float64
	var sumT float64
	var promedioT float64
	
	

	for i:= range this.Alumnos{
		sumInd =0
		for j:= range this.Alumnos[i]{
			sumInd+= this.Alumnos[i][j]
			//fmt.Println(sumInd)
			
		}
		sumT+=sumInd/float64(len(this.Alumnos[i]))
	}


	promedioT= sumT/float64(len(this.Alumnos))
	//fmt.Println(promedioT)
	*reply =  promedioT
	return nil
}

func (this *Server) CalPM(mat string,  reply *float64) error {//calcular promedio de materia
	var sumInd float64
	var sumT float64
	var promedioT float64
	
	

	for i:= range this.Alumnos{
		for  range this.Alumnos[i]{
			if _, ok := this.Alumnos[i][mat]; ok {
				sumInd += this.Alumnos[i][mat]
				sumT++
				break
			}
		}
	}


	promedioT= sumInd/sumT
	//fmt.Println(promedioT)
	*reply =  promedioT
	return nil
}


func server() {
	serv:= new(Server)
	serv.Alumnos = make(map[string]map[string]float64)
	serv.Materias = make(map[string]map[string]float64)
	rpc.Register(serv)
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()
	var input string
	fmt.Scanln(&input)
}