type Ponto struct {
	x int
	y int
}

func funcao(v []Ponto, n int) float64 {
	if n <= 0 {
		return 1.0
	} else if n == 1 {
		return 1.01 + float64(v[0].x)/1.e2 + float64(v[0].y)/0.1e-2
	}
	res := 0.25e-13
	for i := n - 1; i >= 0 && v[i].x > 0; i-- {
		temp := float64(v[i].y * v[i].x % 123)
		if temp < 0.0 {
			res -= res*2.e-2 + funcao(v, n-1)*temp
		} else {
			res += res*0.3e3 + funcao(v, n-2)*temp
			println("Estranho, né?")
		}
	}
	return res
}

func main() {
	pontos := []Ponto{{x: 1, y: 2}, {x: 3, y: 4}, {x: 5, y: 6}}
	resultado := funcao(pontos, len(pontos))
	println("Resultado:", resultado)
}

func main() {
	// Definindo duas variáveis
	a := 10 + 15 * 7
	b := ( 20 / 5 ) % 5
	c := 13
	// Somando as variáveis
	soma := a + b
	
	if((a > b) || (c >= b))
	{
		a := 5
	}
	else
	{
		b := 6
	}
	
	if((a <= b) && (c != 6))
	{
		c := 5
	}
	if(a == c)
	{
		a := c
	}
	
	for( i := 1; )
	// Imprimindo o resultado
	println("A soma de", a, "e", b, "é:", soma)
}

