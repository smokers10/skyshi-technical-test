package todo

type GetOne struct {
	ID int `json:"id"`
}

type Create struct {
	ActivityGroupID int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        string `json:"is_active"`
	Priority        string `json:"priority"`
}

type Delete struct {
	ID int `json:"id"`
}

type Update struct {
	ID              int    `json:"id"`
	ActivityGroupID int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        string `json:"is_active"`
	Priority        string `json:"priority"`
	UpdatedAt       string
}
