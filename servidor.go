package main

import (
	//"errors"
	"fmt"
	"net"
	"net/rpc"
)
type Alumno struct{
	nameStudent string
	Materia string
	Calificacion float64
}

type Server struct{
	Materias map[string]map[string]float64
	Alumnos map[string]map[string]float64
}

func (this *Server) AgregarA(al Alumno,  answ *string) error {
	
	if _, ok := this.Alumnos[al.nameStudent]; !ok {//condicion checa si el alumno no existe en la mapa
	
		cali := make(map[string]float64)
		cali[al.Materia] = al.Calificacion
		this.Alumnos[al.nameStudent] = cali
		*answ = "Estudiante guardado"
			
		
	}else{//si alumno es registrado por primer vez
		if _, ok := this.Alumnos[al.nameStudent][al.Materia]; ok { //si existe alguien que con la materia este igual
			*answ = "Error, El alumno con la materia ya existe"
		}
		*answ = "alumno existe"
	}	
	return nil
}

func (this *Server) CalPA(nameStudent string, answ *float64) error {
	var promedio float64
	var suma float64

	
	for i:= range this.Alumnos[nameStudent]{
		suma+= this.Alumnos[nameStudent][i]
	}
	promedio = suma/float64(len(this.Alumnos[nameStudent]))
	*answ = promedio
	return nil
}
func (this *Server) CalPG( al Alumno,  answ *float64) error {//esta bien
	var sumInd float64
	var sumT float64
	var promedioT float64
	
	

	for i:= range this.Alumnos{
		sumInd =0
		for j:= range this.Alumnos[i]{
			sumInd+= this.Alumnos[i][j]
			fmt.Println(sumInd)
			
		}
		sumT+=sumInd/float64(len(this.Alumnos[i]))
	}


	promedioT= sumT/float64(len(this.Alumnos))
	fmt.Println(promedioT)
	*answ =  promedioT
	return nil
}

func (this *Server) CalPM(mat string,  answ *float64) error {//calcular promedio de materia
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
	fmt.Println(promedioT)
	*answ =  promedioT
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