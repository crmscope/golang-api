package library

type Data struct {
	Name  string
	Value string
}

func GetVal(name string, data []Data) string {

	for i, _ := range data {
		if data[i].Name == name {
			return data[i].Value
		}
	}
	return ""
}
