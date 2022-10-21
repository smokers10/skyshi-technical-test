package requestbody

type GetOne struct {
	ID int `json:"id"`
}

type Create struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

type Delete struct {
	ID int `json:"id"`
}
