package database

type CRUDInterface interface {
	Create()
	Remove()
	Update()
	Delete()
}
