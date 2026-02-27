package main

import "fmt"

type Repository[T any] interface {
	Save(item T)
	Get() T
}
type MemoryStorage[T any] struct {
	data T
}

func (ms *MemoryStorage[T]) Save(item T) {
	ms.data = item
	fmt.Println("Данные сохранены")
}
func (ms *MemoryStorage[T]) Get() T {
	return ms.data
}

func main() {
	strStorage := &MemoryStorage[string]{}
	useRepository(strStorage, "Привет, Go!")
	intStorage := &MemoryStorage[int]{}
	useRepository(intStorage, 42)
}

func useRepository[T any](repo Repository[T], val T) {
	repo.Save(val)
	fmt.Printf("Из хранилища получено: %v\n", repo.Get())
}
