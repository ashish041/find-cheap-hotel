package hotel

import (
	"reflect"
	"testing"
)

func TestUniqueHotels_1(t *testing.T) {
	var allHotels []*Hotels

	hotel1 := &Hotel{
		Name: "a1",
		Rooms: []*Room{
			&Room{
				Code:       "B11",
				Total:      "",
				TotalPrice: "200",
			}, &Room{
				Code:       "A11",
				Total:      "105",
				TotalPrice: "",
			},
		},
	}
	hotel2 := &Hotel{
		Name: "b1",
		Rooms: []*Room{
			&Room{
				Code:       "B11",
				Total:      "",
				TotalPrice: "205",
			}, &Room{
				Code:       "A11",
				Total:      "100",
				TotalPrice: "",
			},
		},
	}
	hotels := []*Hotel{hotel1, hotel2}
	allHotels = append(allHotels, &Hotels{
		Hotels: hotels,
	})
	expected := map[string]*CheapRoom{}
	expected["A11"] = &CheapRoom{
		Code:  "A11",
		Name:  "",
		Total: 100,
	}
	expected["B11"] = &CheapRoom{
		Code:  "B11",
		Name:  "",
		Total: 200,
	}

	got, err := uniqueHotels(allHotels)
	if err != nil {
		t.Error(
			"For", allHotels,
			"expected", expected,
			"got", err,
		)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"For", allHotels,
			"expected", expected,
			"got", got,
		)
	}
}

func TestUniqueHotels_2(t *testing.T) {
	var allHotels []*Hotels

	hotel1 := &Hotel{
		Name: "a1",
		Rooms: []*Room{
			&Room{
				Code:       "D11",
				Total:      "",
				TotalPrice: "500",
			}, &Room{
				Code:       "C11",
				Total:      "",
				TotalPrice: "400",
			},
		},
	}
	hotel2 := &Hotel{
		Name: "b1",
		Rooms: []*Room{
			&Room{
				Code:       "C11",
				Total:      "380",
				TotalPrice: "",
			}, &Room{
				Code:       "D11",
				Total:      "480",
				TotalPrice: "",
			},
		},
	}
	hotels := []*Hotel{hotel1, hotel2}
	allHotels = append(allHotels, &Hotels{
		Hotels: hotels,
	})
	expected := map[string]*CheapRoom{}
	expected["C11"] = &CheapRoom{
		Code:  "C11",
		Name:  "",
		Total: 380,
	}
	expected["D11"] = &CheapRoom{
		Code:  "D11",
		Name:  "",
		Total: 480,
	}

	got, err := uniqueHotels(allHotels)
	if err != nil {
		t.Error(
			"For", allHotels,
			"expected", expected,
			"got", err,
		)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"For", allHotels,
			"expected", expected,
			"got", got,
		)
	}
}

func TestCalculateTotal_1(t *testing.T) {
	var expected float64

	room := &Room{
		Code:       "A11",
		Total:      "105",
		TotalPrice: "",
	}
	expected = 105

	got, err := calculateTotal(room)
	if err != nil {
		t.Error(
			"For", room,
			"expected", expected,
			"got", err,
		)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"For", room,
			"expected", expected,
			"got", got,
		)
	}
}

func TestCalculateTotal_2(t *testing.T) {
	var expected float64

	room := &Room{
		Code: "A11",
		Taxes: []*Tax{
			&Tax{
				Amount:   "8.00",
				Currency: "EUR",
				Type:     "TAXESANDFEES",
			},
			&Tax{
				Amount:   "4.00",
				Currency: "EUR",
				Type:     "EXTRA_FEES",
			},
		},
		Total:      "",
		TotalPrice: "147.00",
	}
	expected = 147

	got, err := calculateTotal(room)
	if err != nil {
		t.Error(
			"For", room,
			"expected", expected,
			"got", err,
		)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"For", room,
			"expected", expected,
			"got", got,
		)
	}
}

func TestCalculateTotal_3(t *testing.T) {
	var expected float64

	room := &Room{
		Code:    "A11",
		NetRate: "135.00",
		Taxes: []*Tax{
			&Tax{
				Amount:   "8.00",
				Currency: "EUR",
				Type:     "TAXESANDFEES",
			},
			&Tax{
				Amount:   "4.00",
				Currency: "EUR",
				Type:     "EXTRA_FEES",
			},
		},
	}
	expected = 147

	got, err := calculateTotal(room)
	if err != nil {
		t.Error(
			"For", room,
			"expected", expected,
			"got", err,
		)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Error(
			"For", room,
			"expected", expected,
			"got", got,
		)
	}
}
