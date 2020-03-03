package ymir

type index struct{
	Message indexData `json:"message"`
}

type indexData struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Methods []string `json:"methods"`
}
