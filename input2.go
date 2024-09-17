func usoDeArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i < len(arr){
		i := 1 + arr[i]
		i := i + 1
	}
	return i
}

func main() {
	a := 15
	b := 5
	
	soma := a + b * a - b
	subtracao := a - b
	multiplicacao := a * b
	divisao := a / b
	modulo := a % b

	// Controle de fluxo
	if a >= 10 {
		a:=10
	} else if a == 10 {
		b := 5
	} else{
		c := 10
	}
	if b >= 5 {
		c := 10
	}
	if b != 6 {
		c := 10
	}
	if a <= 5  {
		c := 10
	}
	if (a <= 5) && (b > 6) {
		c := 10
	}
	if (a <= 5) || (b > (usoDeArray()+1) ) {
		c := 10
	}
	if !(a <= 5) || (b > 6) {
		c := 10
	}

	i := 0
	for i < 5 {
		i:= i + 1
	}

	usoDeArray()
}

