package geoF

type figuraGeo interface {
	area() float64
}

type Triangulo struct {
	base   float64
	altura float64
}

type Cuadrado struct {
	lado float64
}

type Pentagono struct {
	lado    float64
	apotema float64
}

func (t *Triangulo) area() float64 {
	return t.altura * t.base
}

func (c *Cuadrado) area() float64 {
	return c.lado * c.lado
}

func (p *Pentagono) area() float64 {
	per := p.lado * 5
	return per * p.apotema / 2
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
