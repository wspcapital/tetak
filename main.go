	package main

	import "fmt"
	import "flag"

	func fact(f int, n int) int {
		if n==f+1 {
			return 1
		}
		
		ping := n%3
		pong := n%4
		if ping == 0 && pong == 0{
			fmt.Println("PingPong:",n)
		} else if ping == 0{
			fmt.Println("Ping:",n)	
		} else if pong == 0{
			fmt.Println("Pong:",n) 
		} else{
			fmt.Println(":",n)
		}
		
		return fact(f,n+1)
	}

	func main(){
		fromPtr := flag.Int("from", 1, "an int")
		toPtr := flag.Int("to", 200, "an int")

		flag.Parse()
                if *toPtr <= *fromPtr{
			panic(fmt.Sprintf("Wrong params From=%v To=%v", *fromPtr, *toPtr))
		}
		fact(*toPtr,*fromPtr)
	}
