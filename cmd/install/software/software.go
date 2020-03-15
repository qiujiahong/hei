package software

type Item struct {
	Tag             string
	Url             string
	FileName        string
	InstallFileName string
	Default         bool
}

type SoftWare struct {
	Name     string
	Versions []Item
}

func (soft *SoftWare) GetItemByTag(tag string) (Item, bool) {
	for _, item := range soft.Versions {
		if tag == item.Tag {
			return item, false
		} else if (tag == "") && item.Default {
			return item, false
		}
	}
	return Item{}, true
}
