package entity

type Todos struct {
	Todos []TODO `json:"todos"`
}

type TODO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
}
