# Учебный проект: Реализация связного списка на Golang / Educational Project: Implementing a Linked List in Golang

> **Замечание:** Эта реализация связного списка произведена в учебных целях. Связный список уже реализован в Go.

> **Note:** This implementation of the linked list is made for educational purposes. The linked list is already implemented in Go.

[![RUS](https://img.shields.io/badge/lang-RUS-blue)](#русская-версия) | [![ENG](https://img.shields.io/badge/lang-ENG-red)](#english-version)

---

## RUS

Свя́зный спи́сок — базовая динамическая структура данных в информатике, состоящая из узлов, содержащих данные и ссылки («связки») на следующий и/или предыдущий узел списка. Принципиальным преимуществом перед массивом является структурная гибкость: порядок элементов связного списка может не совпадать с порядком расположения элементов данных в памяти компьютера, а порядок обхода списка всегда явно задаётся его внутренними связями. Прежде чем приступить к написанию кода, необходимо определить, какими функциями (иногда буду употреблять слово методы) будет обладать список.

### Мой связный список будет обладать:

- Удаление/добавление элементов смещает остальные при их наличии.
- Добавление элемента в начало.
- Добавление элемента в конец.
- Добавление элемента по индексу.
- Удаление элемента по индексу.
- Удаление первого элемента.
- Удаление последнего элемента.
- Получение длины списка.
- Переворачивание списка (реверсирование).
- Получение элемента.

Связный список реализую двумя способами:

1. **Реализация поверх слайса** (расширяемый массив). Это не совсем честно, но быстро и просто в создании. Вдобавок, мы сможем использовать уже существующие функции для создания новых или сами в других целях.
2. **Реализация на основе новых типов(структур)**. Я создам два типа: «элемент», содержащий ссылку на следующий элемент и значение, и «связный список», который хранит ссылку на первый элемент. С таким способом придется реализовывать методы с нуля.

---

## ENG

A linked list is a basic dynamic data structure in computer science, consisting of nodes containing data and links ("links") to the next and/or previous node in the list. The principal advantage over an array is structural flexibility: the order of elements in a linked list may not match the order of data elements in the computer's memory, and the traversal order of the list is always explicitly defined by its internal links. Before starting to write code, it is necessary to determine what functions (sometimes I will use the word methods) the list will have.

### My linked list will have:

- Deleting/adding elements shifts the others if present.
- Adding an element to the beginning.
- Adding an element to the end.
- Adding an element by index.
- Deleting an element by index.
- Deleting the first element.
- Deleting the last element.
- Getting the length of the list.
- Reversing the list.
- Getting an element.

I will implement the linked list in two ways:

1. **Implementation on top of a slice** (dynamic array). This is not entirely fair, but it is quick and easy to create. Additionally, we can use existing functions to create new ones or use them on their own.
2. **Implementation based on new types**. I will create two types: "element", containing a reference to the next element and value, and "linked list", which stores a reference to the first element. With this method, you will have to implement the methods from scratch.
