package main

import "fmt"

//Первая заглавная буква в названии типов, разрешает использовать их в другом файле, в другом пакете. В этом же пакете необязательно заглавную.

//Мы создаем новый кастомный тип, который строится на структуре "связный список", которая содержит слайс (расширяемый массив) с данными.
//Именно этому типу мы навязываем новые методы работы, поэтому можно вызывать методы через точку.
type LinkedList struct {
	data []int
}

//Некоторые методы полностью строятся на встроенных для работы со слайсами.
func (slice *LinkedList) AppEnd(num int) {
	slice.data = append(slice.data, num)
}

//Удаление последнего элемента слайса обеспечивается приравниванием этого же слайса, без последнего элемента. Некоторые методы строятся похожим способом.
func (slice *LinkedList) DeleteEnd() {
	slice.data = slice.data[:len(slice.data)-1]
}
func (slice *LinkedList) DeletFirst() {
	slice.data = slice.data[1:]
}
func (slice *LinkedList) AppFirst(num int) {
	var sliceCopy []int
	sliceCopy = append(sliceCopy, num)
	sliceCopy = append(sliceCopy, slice.data...)
	slice.data = sliceCopy
}
func (slice *LinkedList) InsertAt(num int, index int) {
	if index == 0 {
		slice.AppFirst(num)
		return
	}
	if index == len(slice.data) {
		slice.AppEnd(num)
		return
	}
	var sliceCopy []int
	sliceCopy = append(sliceCopy, slice.data[0:index]...)
	sliceCopy = append(sliceCopy, num)
	sliceCopy = append(sliceCopy, slice.data[index:]...)
	slice.data = sliceCopy
}
func (slice *LinkedList) DeleteElement(index int) {
	if index == 0 {
		slice.DeletFirst()
		return
	}
	if index == len(slice.data)-1 {
		slice.DeleteEnd()
		return
	}
	sliceCopy := slice.data[0:index]
	sliceCopy = slice.data[index:]
	slice.data = sliceCopy
}

func (slice *LinkedList) Len() int {
	return len(slice.data)
}
func (slice *LinkedList) ReverseList() {
	var sliceCopy []int
	for i := len(slice.data) - 1; i >= 0; i-- {
		sliceCopy = append(sliceCopy, slice.data[i])
	}
	slice.data = sliceCopy
}

//Получени элемента по индексу возможно за счет прямого обращения к слайсу.
func (slice *LinkedList) getElement(index int) int {
	return slice.data[index]

}
func main() {
	var myList LinkedList
	fmt.Println(`Создание пустого списка`)
	fmt.Println(myList)

	fmt.Println(`Добавление в начало с помощью appFirst`)
	myList.AppFirst(1)
	fmt.Println(myList)

	fmt.Println(`Добавление в конец с помощью appEnd`)
	myList.AppEnd(3)
	fmt.Println(myList)

	fmt.Println(`Добавление в середину с помошью InsertAt`)
	myList.InsertAt(2, 1)
	fmt.Println(myList)

	fmt.Println(`Обращение к элементу по индексу`)
	fmt.Println(myList.getElement(1))

	fmt.Println(`Реверсирование с помошью ReverseList`)
	myList.ReverseList()
	fmt.Println(myList)

	fmt.Println(`Удаление середины (по индексу) с помощью DeleteElement`)
	myList.DeleteElement(1)
	fmt.Println(myList)

	fmt.Println(`Удаление из начала с помошью DeletFirst`)
	myList.DeletFirst()
	fmt.Println(myList)

	fmt.Println(`Удаление из конца с помошью DeleteEnd`)
	myList.DeleteEnd()
	fmt.Println(myList)

	fmt.Println(`Получение длины`)
	fmt.Println(myList.Len())
}
