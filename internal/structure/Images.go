package structure

type Image struct {
	Name   string  `json:"name"`
	Width  int32   `json:"width"`
	Height int32   `json:"height"`
	Size   float64 `json:"size"`
	Owner  int     `json:"owner"`
}
