package entity

type Item struct {
  Id      int    `json:"id"`
  Name    string `json:"name"`
  Price   int    `json:"price"`
  Remarks string `json:"remarks"`
}
