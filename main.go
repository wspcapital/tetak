	package main

	import "fmt"

	func fact(f int, n int) int {
		if n==f+1 {
			return 1
		}
		
		ping := n%3
		pong := n%4
		if ping > 0 && pong > 0{
			fmt.Println("PingPong:",n)
		} else if ping > 0{
			fmt.Println("Ping:",n)	
		} else if pong > 0{
			fmt.Println("Pong:",n) 
		}
		
		return fact(f,n+1)
	}

	func main(){
		fact(200,1)
	}
