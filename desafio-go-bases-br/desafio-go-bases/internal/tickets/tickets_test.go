package tickets

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"reflect"
	"testing"
)

func ExampleContains() {
	bolean := Contains([]string{"asdf", "ghjkl"}, "asdf")
	fmt.Println(bolean)
	//Output:true
}

func TestContainsPositivo(t *testing.T) {
	teste := Contains([]string{"teste", "teste2"}, "teste")
	esperado := true
	if teste != esperado {
		t.Error("Expected:", esperado, "Got:", teste)
	}
}
func TestContainsNegativo(t *testing.T) {
	teste := Contains([]string{"teste", "teste2"}, "teste3")
	esperado := false
	if teste != esperado {
		t.Error("Expected:", esperado, "Got:", teste)
	}
}

func TestOpenCSVPositivo(t *testing.T) {
	teste, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		t.Error("Expected:", reflect.TypeOf([]Ticket{}), "Got:", err)
	}
	if teste == nil {
		t.Error("Expected:", reflect.TypeOf([]Ticket{}), "Got:", nil)
	}
}

func ExampleOpenCSV() {
	list, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list[0])
	//Output:{1 Tait Mc Caughan tmc0@scribd.com Finland 17:11 785}
}

func TestOpenCSV2Negativo(t *testing.T) {
	teste, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets2.csv")
	if err == nil {
		t.Error("Expected:", reflect.TypeOf(os.PathError{}), "Got:", err)
	}
	if teste != nil {
		t.Error("Expected:", reflect.TypeOf(os.PathError{}), "Got:", reflect.TypeOf([]Ticket{}))
	}
}

type TesteTurno struct {
	data     turno
	esperado string
}

func ExampleTurno_Num() {
	turno := Manha.Num()
	fmt.Println(turno)
	//Output:1
}

func TestTurno_Num(t *testing.T) {
	testes := []TesteTurno{
		{data: Madrugada, esperado: "0"},
		{data: Manha, esperado: "1"},
		{data: Tarde, esperado: "2"},
		{data: Noite, esperado: "3"},
	}
	for _, teste := range testes {
		resultado := teste.data.Num()
		if resultado != teste.esperado {
			t.Error("("+teste.data.String()+") - "+"Expected:", teste.esperado, "Got:", resultado)
		}
	}
}

func ExampleTurno_String() {
	turno := Madrugada.String()
	fmt.Println(turno)
	//Output:madrugada
}

func TestTurno_String(t *testing.T) {
	testes := []TesteTurno{
		{data: Madrugada, esperado: "madrugada"},
		{data: Manha, esperado: "manha"},
		{data: Tarde, esperado: "tarde"},
		{data: Noite, esperado: "noite"},
	}
	for _, teste := range testes {
		resultado := teste.data.String()
		if resultado != teste.esperado {
			t.Error("Expected:", teste.esperado, "Got:", resultado)
		}
	}
}

var ciaAerea CiaAerea = CiaAerea{
	ID:   uuid.New().String(),
	Nome: "LATAM Airlines Brasil",
	Cnpj: "02.012.862/0001-60",
}

func ExampleCiaAerea_BuscarTotalDeDestinos() {

	var ciaAerea CiaAerea = CiaAerea{
		ID:   uuid.New().String(),
		Nome: "LATAM Airlines Brasil",
		Cnpj: "02.012.862/0001-60",
	}

	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	resultado, err := ciaAerea.BuscarTotalDeDestinos()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)
	//Output:125
}

func TestCiaAerea_BuscarTotalDeDestinos(t *testing.T) {
	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	resultado, err := ciaAerea.BuscarTotalDeDestinos()
	esperado := 125
	if err != nil {
		t.Error("Expected:", esperado, "Got:", err)
	}
	if resultado != esperado {
		t.Error("Expected:", esperado, "Got:", resultado)
	}
}

func ExampleCiaAerea_GetTotalTickets() {

	var ciaAerea CiaAerea = CiaAerea{
		ID:   uuid.New().String(),
		Nome: "LATAM Airlines Brasil",
		Cnpj: "02.012.862/0001-60",
	}

	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	resultado, err := ciaAerea.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)
	//Output:45
}

func TestCiaAerea_GetTotalTickets(t *testing.T) {
	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	resultado, err := ciaAerea.GetTotalTickets("United States")
	esperado := 25
	if err != nil {
		t.Error("Expected:", esperado, "Got:", err)
	}
	if resultado != esperado {
		t.Error("Expected:", esperado, "Got:", resultado)
	}
}

func ExampleCiaAerea_GetMornings() {

	var ciaAerea CiaAerea = CiaAerea{
		ID:   uuid.New().String(),
		Nome: "LATAM Airlines Brasil",
		Cnpj: "02.012.862/0001-60",
	}

	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	resultado, err := ciaAerea.GetMornings("madrugada")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)
	//Output:304
}

type TesteCiaAerea struct {
	data     string
	esperado int
}

type TesteCiaAereaError struct {
	data     string
	esperado error
}

func TestCiaAerea_GetMorningsPositivo(t *testing.T) {

	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	testes := []TesteCiaAerea{
		{data: "00:30", esperado: 304},
		{data: "madrugada", esperado: 304},
		{data: "0", esperado: 304},
		{data: "08:00", esperado: 256},
		{data: "manha", esperado: 256},
		{data: "1", esperado: 256},
		{data: "tarde", esperado: 289},
		{data: "17:23", esperado: 289},
		{data: "2", esperado: 289},
		{data: "noite", esperado: 151},
		{data: "22:37", esperado: 151},
		{data: "3", esperado: 151},
	}

	for _, teste := range testes {
		resultado, err := ciaAerea.GetMornings(teste.data)
		if err != nil {
			t.Error("("+teste.data+") - "+"Expected:", teste.esperado, "Got:", err)
		}
		if resultado != teste.esperado {
			t.Error("("+teste.data+") - "+"Expected:", teste.esperado, "Got:", resultado)
		}
	}
}

func TestCiaAerea_GetMorningsNegativo(t *testing.T) {
	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	testes := []TesteCiaAereaError{
		{data: "00:60", esperado: ErrTimeFormat},
		{data: "24:00", esperado: ErrTimeFormat},
		{data: "6:", esperado: ErrTimeFormat},
		{data: "23:", esperado: ErrTimeFormat},
		{data: ":4", esperado: ErrTimeFormat},
		{data: "4", esperado: ErrTimeFormat},
		{data: "shdh:", esperado: ErrTimeFormat},
		{data: "shd", esperado: ErrTimeFormat},
	}

	for _, teste := range testes {
		resultado, err := ciaAerea.GetMornings(teste.data)
		if err == nil {
			t.Error("("+teste.data+") - "+"Expected:", teste.esperado, "Got:", resultado)
		}
	}
}

func ExampleCiaAerea_AverageDestination() {

	var ciaAerea CiaAerea = CiaAerea{
		ID:   uuid.New().String(),
		Nome: "LATAM Airlines Brasil",
		Cnpj: "02.012.862/0001-60",
	}

	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	resultado, err := ciaAerea.AverageDestination(len(ciaAerea.Tickets24Horas))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)
	//Output:8
}

func TestCiaAerea_AverageDestination(t *testing.T) {
	ticketsList, err := OpenCSV("C:\\Users\\jarde\\Documents\\curso_ctd\\ctd-ano2-jardel-silva\\bim3\\especializacao-backed03\\checkpoint01\\desafio-go-bases-br\\desafio-go-bases\\tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	ciaAerea.Tickets24Horas = ticketsList

	resultado, err := ciaAerea.AverageDestination(len(ciaAerea.Tickets24Horas))
	esperado := 8.0
	if err != nil {
		t.Error("Expected:", esperado, "Got:", err)
	}
	if resultado != esperado {
		t.Error("Expected:", esperado, "Got:", resultado)
	}
}
