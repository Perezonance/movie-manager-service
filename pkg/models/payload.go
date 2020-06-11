package models

type (
	TestFruitPayload struct {
		Fruits		[]string	`json:"fruits"`
	}

	RequestUserPayload struct {
		Payload 	[]User		`json:"users"`
	}

	RequestPostPayload struct {
		Payload 	[]Post	 	`json:"posts"`
	}

	User struct {
		Id			float64		`json:"id"`
		Name 		string		`json:"name"`
		Email 		string		`json:"email"`
		Expertise 	string		`json:"expertise"`
	}

	Post struct {
		Id 			string		`json:"id"`
		UserId 		string 		`json:"userId"`
		Title 		string 		`json:"title"`
		Body 		string 		`json:"body"`
	}
)


