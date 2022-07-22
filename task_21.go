package main

type DataService interface { //сервис данных который принимает в формате XML
	SendXmlData()
}

type XmlDoc struct { //некая структура xml документа
}

func (x XmlDoc) SendXmlData() {
	println("Getting xml doc")
}

type JsonDoc struct { //необходимая нам структура json
}

func (j JsonDoc) ConvertedToXml() string {
	return "<xml></xml>"
}

type Adapter struct { //сам адаптер
	jsonDoc *JsonDoc
}

func (a Adapter) SendXmlData() { //реализация полей адаптера
	a.jsonDoc.ConvertedToXml()
	println("Getting xml data")
}

func main() {
	//плюс паттерна: сокрытие от клиента преобразований интефресов
	//минус: усложнение кода из-за добавления структуры и методов
}

