package game

type Data struct {
}

var data *Data

func GetData() *Data {
	if data == nil {
		data = &Data{}
	}

	return data
}
