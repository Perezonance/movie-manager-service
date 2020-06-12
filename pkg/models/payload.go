package models

type (
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
		Id 			float64		`json:"id"`
		UserId 		float64 	`json:"userId"`
		Title 		string 		`json:"title"`
		Body 		string 		`json:"body"`
	}
)


