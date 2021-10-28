package fibonacci

func Fibonacci() func() int {
	a,b:=1,1
	return func() int {
		defer func() {
			a,b=b,a+b
		}()
		return a
	}
}
