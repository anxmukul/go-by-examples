package main
import ("fmt"
		"encoding/json"
)

type Address struct {
	City string		`json:"city"`
	District string	`json:"district"`
	State string	`json:"state"`
	Postalcode int	`json:"postalcode"`
}

type Owner struct {
	Name string		`json:"name"`
	ContactNo string	`json:"constactno."`
	Email string	`json:"email"`
	Address `json:"address"`

}

type Flat struct {
	Flatno string	`json:"flatno"`
	Flattype string	`json:"flattype"`
	Owner	`json:"ownerDetails"`

}

func encodeIntoJson(allflats []Flat) {
	for _, flat := range allflats {
		// finalJson, err := json.Marshal(flat)
		finalJson, err := json.MarshalIndent(flat, "", "\t")	// MarshalIndent used to beutify json in console

		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Printf("%s\n", finalJson)
			// fmt.Println(finalJson)
		}
	}
	// fmt.Println(allflats)
	// finalJson, err := json.Marshal(allflats)
	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Printf("%s\n", finalJson)
	// }
}

func main(){
	flat1 := Flat{
		Flatno: "O-71",
		Flattype: "3BHK",
		Owner: Owner{
			Name: "AbhishekSah",
			ContactNo: "787658493",
			Email: "random@gmail.com",
			Address: Address{
				City: "Rampur",
				District: "Bhagwanpur",
				State: "Rajasthan",
				Postalcode: 853421,
			},
		},
	}
	flat2 := Flat{
		Flatno: "D-56",
		Flattype: "2BHK",
		Owner: Owner{
			Name: "Satya Bhia",
			ContactNo: "895432512",
			Email: "satyamalik@hotmail.com",
			Address: Address{
				City: "Ramgandh",
				District: "Lakhnow",
				State: "UP",
				Postalcode: 794321,
			},
		},
	}
	flat3 := Flat{
		Flatno: "Q-41",
		Flattype: "1BHK",
		Owner:Owner{
			Name: "Chota Don",
			ContactNo: "9999944444",
			Email: "chotadon@google.com",
			Address: Address{
				City: "Mumbai",
				District: "Lokhandwala",
				State: "Maharashtra",
				Postalcode: 123456,
			},
		},
	}
	flat4:= Flat{
		Flatno: "A-34",
		Flattype: "2BHK",
		Owner: Owner{
			Name: "Arundhatti Roy",
			ContactNo: "875432159",
			Email: "royarun@reddifmail.com",
			Address: Address{
				City: "Bnagaluru",
				District: "Bengalore",
				State: "Karnataka",
				Postalcode: 8004531,
			},
		},
	}
	// fmt.Println(flat1)
	// fmt.Println(flat2)
	// fmt.Println(flat3)
	// fmt.Println(flat4)

	flats := []Flat{flat1, flat2, flat3, flat4}
	encodeIntoJson(flats)
}

// Converting struct data in json using marshal
// Struct field should be capatalized because then it will be public otherwise it will be private