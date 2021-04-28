package cgminer

type API struct {
	Suburl string //子路径
}

var SubUrl string

func init() {
	SubUrl = "/api"
}
