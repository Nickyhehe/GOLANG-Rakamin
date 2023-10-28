package app

type Photo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photoUrl"`
	UserID   string `json:"userId"`
}
