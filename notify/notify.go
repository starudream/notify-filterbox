package notify

type Params struct {
	Id       string
	Title    string
	SubTitle string
	Message  string
	GroupId  string
}

type Result struct {
	Id  string
	Out string
}

func Notify(p Params) (Result, error) {
	return notify(p)
}
