package geoF

type figuraGeo interface {
	Area() float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Cuadrado struct {
	Lado float64
}

type Pentagono struct {
	Lado    float64
	Apotema float64
}

func (t *Triangulo) Area() float64 {
	return t.Altura * t.Base
}

func (c *Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (p *Pentagono) Area() float64 {
	per := p.Lado * 5
	return per * p.Apotema / 2
}

func Figure(lados string) string {
	if lados == "3" {
		return "Es un triángulo"
	} else if lados == "4" {
		return "Es un cuadrado"
	} else if lados == "5" {
		return "Es un pentágono"
	} else {
		return "Valor no válido"
	}
}
