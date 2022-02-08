package model

type Author struct {
	Name     string `json:"name,omitempty"`
	Age      int32  `json:"age,string,omitempty"`
	Address  string `json:"address,omitempty"`
	PostCode string `json:"post_code,omitempty"`
}
