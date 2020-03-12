package software

type Item struct {
	Tag string
	Url string
	FileName string
}

type SoftWare struct {
	Name     string
	Versions []Item
}
