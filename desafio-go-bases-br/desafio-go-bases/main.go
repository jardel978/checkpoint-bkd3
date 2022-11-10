package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/google/uuid"
)

var ciaAerea tickets.CiaAerea

func main() {
	ticketsList, err := tickets.OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Errorf("erro: %w", err)
	}
	ciaAerea = tickets.CiaAerea{
		ID:             uuid.New().String(),
		Nome:           "LATAM Airlines Brasil",
		Cnpj:           "02.012.862/0001-60",
		Tickets24Horas: ticketsList,
	}

	//Requisito 1
	totalBrasil, err := ciaAerea.GetTotalTickets("Brazil")
	totalUSA, err := ciaAerea.GetTotalTickets("United States")
	if err != nil {
		fmt.Errorf("erro: %w", err)
	}
	fmt.Println("Total de tikets para o Brasil:", totalBrasil)
	fmt.Println("Total de tikets para o United States:", totalUSA)

	// Requisito 2
	ticketsMadrugada, err := ciaAerea.GetMornings("0")
	ticketsManha, err := ciaAerea.GetMornings("manha")
	ticketsTarde, err := ciaAerea.GetMornings("14:30")
	ticketsNoite, err := ciaAerea.GetMornings("3")
	if err != nil {
		fmt.Errorf("erro: %w", err)
	}
	fmt.Println("Quantidades de pessoas que viajam pela magrugada:", ticketsMadrugada)
	fmt.Println("Quantidades de pessoas que viajam pela manhã:", ticketsManha)
	fmt.Println("Quantidades de pessoas que viajam pela tarde:", ticketsTarde)
	fmt.Println("Quantidades de pessoas que viajam pela noite:", ticketsNoite)
	totalTickets24h := ticketsMadrugada + ticketsManha + ticketsTarde + ticketsNoite
	fmt.Println("Total de tickets nas últimas 24h:", totalTickets24h)

	//	Requisito 3
	mediaTicketsPorDestino, err := ciaAerea.AverageDestination("Brazil", len(ciaAerea.Tickets24Horas))
	if err != nil {
		fmt.Errorf("erro: %w", err)
	}
	fmt.Printf("Média de tickets para o Brasil em um dia dentre um total de %v: %v\n", len(ciaAerea.Tickets24Horas), mediaTicketsPorDestino)
}
