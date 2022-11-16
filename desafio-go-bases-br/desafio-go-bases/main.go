package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jardel197/checkpoint-bkd3/desafio-go-bases-br/desafio-go-bases/internal/tickets"
)

var CiaAerea tickets.CiaAerea

func main() {
	ticketsList, err := tickets.OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	CiaAerea = tickets.CiaAerea{
		ID:             uuid.New().String(),
		Nome:           "LATAM Airlines Brasil",
		Cnpj:           "02.012.862/0001-60",
		Tickets24Horas: ticketsList,
	}

	//Requisito 1
	totalBrasil, err := CiaAerea.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	totalUSA, err := CiaAerea.GetTotalTickets("United States")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total de tikets para o Brasil:", totalBrasil)
	fmt.Println("Total de tikets para o United States:", totalUSA)

	// Requisito 2
	ticketsMadrugada, err := CiaAerea.GetMornings("0")
	if err != nil {
		fmt.Println(err)
	}
	ticketsManha, err := CiaAerea.GetMornings("manha")
	if err != nil {
		fmt.Println(err)
	}
	ticketsTarde, err := CiaAerea.GetMornings("14:30")
	if err != nil {
		fmt.Println(err)
	}
	ticketsNoite, err := CiaAerea.GetMornings("3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Quantidades de pessoas que viajam pela magrugada:", ticketsMadrugada)
	fmt.Println("Quantidades de pessoas que viajam pela manhã:", ticketsManha)
	fmt.Println("Quantidades de pessoas que viajam pela tarde:", ticketsTarde)
	fmt.Println("Quantidades de pessoas que viajam pela noite:", ticketsNoite)
	totalTickets24h := ticketsMadrugada + ticketsManha + ticketsTarde + ticketsNoite
	fmt.Println("Total de tickets nas últimas 24h:", totalTickets24h)

	//	Requisito 3
	mediaTicketsPorDestino, err := CiaAerea.AverageDestination(len(CiaAerea.Tickets24Horas))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Média de tickets por destinos em um dia: %v\n", mediaTicketsPorDestino)
}
