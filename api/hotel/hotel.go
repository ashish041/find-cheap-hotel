package hotel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

var (
	type1   = "TAXESANDFEES"
	type2   = "EXTRA_FEES"
	api_src = []string{/*"https://f704cb9e-bf27-440c-a927-4c8e57e3bad1.mock.pstmn.io/s1/availability",*/
		"https://f704cb9e-bf27-440c-a927-4c8e57e3bad1.mock.pstmn.io/s2/availability"}
)

var httpClient = &http.Client{Timeout: 5 * time.Second}

type Tax struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Type     string `json:"type,omitempty"`
}

type Room struct {
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	NetPrice   string `json:"net_price,omitempty"`
	NetRate    string `json:"net_rate,omitempty"`
	Taxes      []*Tax `json:"taxes,omitempty"`
	Total      string `json:"total,omitempty"`
	TotalPrice string `json:"totalPrice,omitempty"`
}

type CheapRoom struct {
	Code  string  `json:"code,omitempty"`
	Name  string  `json:"name,omitempty"`
	Total float64 `json:"total,omitempty"`
}

type Hotel struct {
	Name  string  `json:"name,omitempty"`
	Stars int  `json:"stars,omitempty"`
	Rooms []*Room `json:"rooms,omitempty"`
}

type Hotels struct {
	Hotels []*Hotel `json:"hotels"`
}

type SortedList []*CheapRoom

func (s SortedList) Len() int {
	return len(s)
}

func (s SortedList) Less(i, j int) bool {
	return s[i].Total < s[j].Total
}

func (s SortedList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func fetchData(url string) ([]byte, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		log.Printf("Error client.Get(%s): %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error ioutil.ReadAll(%v): %v", resp.Body, err)
		return nil, err
	}
	return body, nil
}

func calculateTotal(room *Room) (float64, error) {
	var total float64
	var err error

	if room.Total != "" {
		total, err = strconv.ParseFloat(room.Total, 8)
		if err != nil {
			return total, err
		}
	}
	if total == 0 && room.TotalPrice != "" {
		total, err = strconv.ParseFloat(room.TotalPrice, 8)
		if err != nil {
			return total, err
		}
	}
	if total == 0 {
		if room.NetRate != "" {
			total, err = strconv.ParseFloat(room.NetRate, 8)
			if err != nil {
				return total, err
			}
		}
		if total == 0 && room.NetPrice != "" {
			total, err = strconv.ParseFloat(room.NetPrice, 8)
			if err != nil {
				return total, err
			}
		}
		for _, tax := range room.Taxes {
			var amount float64
			if tax.Amount != "" {
				amount, err = strconv.ParseFloat(tax.Amount, 8)
				if err != nil {
					return total, err
				}
				total = total + amount
			}
		}
	}
	return total, nil
}

func roomToStruct(room *Room, total float64) *CheapRoom {
	return &CheapRoom{
		Code:  room.Code,
		Name:  room.Name,
		Total: total,
	}
}

func getAllHotels() ([]*Hotels, error) {
	var list []*Hotels
	for _, url := range api_src {
		var hotels *Hotels
		body, err := fetchData(url)
		if err != nil {
			return list, err
		}
		err = json.Unmarshal(body, &hotels)
		if err != nil {
			return list, err
		}
		list = append(list, hotels)
	}
	return list, nil
}

func uniqueHotels(allHotels []*Hotels) (map[string]*CheapRoom, error) {
	unique := map[string]*CheapRoom{}

	for _, hotels := range allHotels {
		for _, hotel := range hotels.Hotels {
			for _, room := range hotel.Rooms {
				total, err := calculateTotal(room)
				if err != nil {
					return unique, err
				}
				if val, ok := unique[room.Code]; ok {
					if val.Total > total {
						unique[room.Code] = roomToStruct(room, total)
					}
					continue
				}
				unique[room.Code] = roomToStruct(room, total)
			}
		}
	}
	return unique, nil
}

func GetHotelLists() ([]*CheapRoom, error) {
	var list []*CheapRoom

	allHotels, err := getAllHotels()
	if err != nil {
		return list, err
	}
	unique, err := uniqueHotels(allHotels)
	if err != nil {
		return list, err
	}
	for _, r := range unique {
		list = append(list, &CheapRoom{
			Code:  r.Code,
			Name:  r.Name,
			Total: r.Total,
		})
	}
	sort.Sort(SortedList(list))
	return list, nil
}
