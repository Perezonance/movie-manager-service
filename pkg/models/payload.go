package models

type (
	RequestUserPayload struct {
		Payload 	[]User		`json:"payload"`
	}

	RequestPostPayload struct {
		Payload 	[]Post 		`json:"payload"`
	}

	User struct {
		Id			string		`json:"id"`
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


