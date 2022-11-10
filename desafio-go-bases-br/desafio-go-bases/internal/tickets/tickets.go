package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID            int     `json:"ID"`
	Nome          string  `json:"Nome"`
	Email         string  `json:"Email"`
	PaisDeDestino string  `json:"PaisDeDestino"`
	HoraDoVoo     string  `json:"HoraDoVoo"`
	Preco         float64 `json:"Preco"`
}

type CiaAerea struct {
	ID             string   `json:"ID"`
	Nome           string   `json:"Nome"`
	Cnpj           string   `json:"Cnpj"`
	Tickets24Horas []Ticket `json:"Tickets24Horas"`
}

var ErrNoTickets = errors.New("A lista de tickets está vazia")

func (c CiaAerea) BuscarTotalDeDestinos() (int, error) {
	var sliceAux []string
	if len(c.Tickets24Horas) > 0 {
		for _, ticket := range c.Tickets24Horas {
			if !Contains(sliceAux, ticket.PaisDeDestino) {
				sliceAux = append(sliceAux, ticket.PaisDeDestino)
			}
		}
	} else {
		return 0, ErrNoTickets
	}
	fmt.Errorf("%v", sliceAux)
	return len(sliceAux), nil
}

var ErrTimeFormat = errors.New(`insira um formato de horário válido.
	São aceitos os valores para:
		Madrugada: "madrugada", 1, 00:00 a 06:00
		Manhã: "manha", 2, 07:00 a 12:00
		Tarde: "tarde", 3, 13:00 a 19:00
		Noite: "noite", 4, 20:00 a 23:00
	Ex: 14:00 ou tarde ou 3`)

// GetTotalTickets Busca total de tickets de um determinado país em um dia
func (c CiaAerea) GetTotalTickets(destination string) (int, error) {
	// deve retornar quantas pessoas vão para o destino
	//tiketsList, err := c.CarregarTickets24Horas()
	//if err != nil {
	//	return 0, err
	//}
	contador := 0
	for _, ticket := range c.Tickets24Horas {
		if strings.ToUpper(ticket.PaisDeDestino) == strings.ToUpper(destination) {
			contador++
		}
	}
	return contador, nil
}

// GetMornings Busca total de tickets dado detrminado horário em um dia
func (c CiaAerea) GetMornings(time string) (int, error) {
	var turno string

	// identificar e validar parâmetro recebido "time"
	if time != Madrugada.String() && time != Manha.String() && time != Tarde.String() && time != Noite.String() && time != Madrugada.Num() && time != Manha.Num() && time != Tarde.Num() && time != Noite.Num() {
		if strings.Contains(time, ":") {
			partes := strings.SplitN(time, ":", 2)
			horas, err := strconv.Atoi(partes[0])
			if err != nil {
				return 0, err
			}
			if horas < 0 || horas > 23 {
				return 0, ErrTimeFormat
			}
			if horas >= 0 && horas <= 6 {
				turno = "madrugada"
			} else if horas >= 7 && horas <= 12 {
				turno = "manha"
			} else if horas >= 13 && horas <= 19 {
				turno = "tarde"
			} else {
				turno = "noite"
			}
		}
	} else {
		turno = time
	}

	contador := 0
	for _, v := range c.Tickets24Horas {
		partesAux := strings.SplitN(v.HoraDoVoo, ":", 2)
		horasAux, err := strconv.Atoi(partesAux[0])
		if err != nil {
			return 0, err
		}
		if (horasAux >= 0 && horasAux <= 6) && (turno == "madrugada" || time == Madrugada.Num()) {
			contador++
		} else if (horasAux >= 7 && horasAux <= 12) && (turno == "manha" || time == Manha.Num()) {
			contador++
		} else if (horasAux >= 13 && horasAux <= 19) && (turno == "tarde" || time == Tarde.Num()) {
			contador++
		} else if (horasAux >= 20 && horasAux <= 23) && (turno == "noite" || time == Noite.Num()) {
			contador++
		}
	}

	return contador, nil
}

// AverageDestination Calcula média de ticktes para um determinado país em um dia
func (c CiaAerea) AverageDestination(destination string, totalDeTickets int) (float64, error) {
	// total viagens/total de paises
	totalDePaises, err := c.BuscarTotalDeDestinos()
	if err != nil {
		return 0, err
	}
	media := float64(totalDeTickets) / float64(totalDePaises)
	return media, nil
}

type turno uint

var turnos = [...]string{"madrugada", "manha", "tarde", "noite"}

const (
	Madrugada turno = iota
	Manha
	Tarde
	Noite
)

func (t turno) String() string {
	return turnos[t]
}

func (t turno) Num() string {
	switch t {
	case Madrugada:
		return "0"
	case Manha:
		return "1"
	case Tarde:
		return "2"
	case Noite:
		return "3"
	default:
		return ""
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func OpenCSV(path string) ([]Ticket, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	var tiketsList []Ticket
	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Errorf("erro: %w", err)
			break
		}
		id, err := strconv.Atoi(record[0])
		preco, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, err
		}
		tiketsList = append(tiketsList, Ticket{
			ID:            id,
			Nome:          record[1],
			Email:         record[2],
			PaisDeDestino: record[3],
			HoraDoVoo:     record[4],
			Preco:         preco,
		})
	}
	return tiketsList, nil
}
