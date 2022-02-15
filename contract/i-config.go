package contract

type IConfig interface {
	GetStruct(res interface{}) error
}
