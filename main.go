package main

import "fmt"

func main(){
	var conferenceName string="Go Conference"
	const conferenceTickets int=50
	var remainingTickets int =50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of",conferenceTickets,"tickets and",remainingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")

	
}
