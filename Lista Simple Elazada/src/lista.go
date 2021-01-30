package main
import "fmt"

func main() {
	
	lista := Lista{}
	lista.Agregar("1","Mario")
	lista.Agregar("2","Maria")
	lista.Agregar("3","Pepe")
	lista.Agregar("4","Chepe")
	lista.Agregar("5","Chepe2")
	
	lista.Buscar("Pablo")
}

//--------------------------------Clase Nodo----------------------------------------

type Nodo struct {
	carnet string
	nombre string
	siguiente *Nodo
}

//Get para atributos
func (c *Nodo) getCarnet() string {
	return c.carnet
}

func (c *Nodo) getNombre() string {
	return c.nombre
}

func (c *Nodo) getSiguiente() *Nodo {
	return c.siguiente
}

//Set para atributos
func (c *Nodo) SetCarnet(ncarnet string) {
	c.carnet = ncarnet
}

func (c *Nodo) SetNombre(nnombre string) {
	c.nombre =	nnombre
}

func (c *Nodo) SetSiguiente(sn *Nodo) {
	c.siguiente = sn
}

//_______________________________________________________________________________________

//---------------------------------Clase Lista-------------------------------------------

type Lista struct {
	primero *Nodo
	ultimo *Nodo
}

//Metodo Agregar
func (l *Lista) Agregar(c string, n string) {
	if (l.primero == nil) {
		nuevo := Nodo{
			carnet : c,
			nombre : n,
			siguiente : nil,
		}
		l.primero = &nuevo
		l.ultimo = &nuevo
	}else{
		nuevo := Nodo{
			carnet : c,
			nombre : n,
			siguiente : nil,
		}
		l.ultimo.SetSiguiente(&nuevo)
		l.ultimo = &nuevo
	}
}

//Mostrar Lista
func (l *Lista) MostrarLista() {
	var aux *Nodo = l.primero
	for {
		fmt.Println("Nombre: " + aux.nombre + "--" + "Carnet: " + aux.carnet)
		if (aux == l.ultimo){
			return
		}
		aux = aux.getSiguiente()
	}
}


//Metodo Buscar
func (l *Lista) Buscar(nom string){
	var aux *Nodo = l.primero

	for {
		if(aux.nombre == nom){
			fmt.Println("Si se encontro: " + nom)
			return
		}else if (aux.getSiguiente() == nil){
			fmt.Println(nom + " No existe en esta lista")
			return
		}
		aux = aux.getSiguiente()
	}

}


//Eliminar Nodo
func (l *Lista) Eliminar(c string) {
	var nodoeliminar *Nodo = l.primero
	var aux *Nodo = l.primero

	if (l.primero.nombre == c) {
		l.primero = aux.getSiguiente()
		aux = nil
	}else if(l.ultimo.nombre == c){
		for {
			if (aux.getSiguiente().nombre == c) {
				l.ultimo = aux
				aux = nil
				return
			}
			aux = aux.getSiguiente()
		}
	}else{
		for{
			if (nodoeliminar.nombre == c) {
				for {
					if (aux.getSiguiente().nombre == c) {
						aux.SetSiguiente(nodoeliminar.getSiguiente())
						nodoeliminar = nil
						return
					}
					aux = aux.getSiguiente()
				}
			}
			nodoeliminar = nodoeliminar.getSiguiente()
		}
	}
}

//_______________________________________________________________________________________