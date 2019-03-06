package helper

type Iterator interface{
	Next() interface{}
	HasNext() bool
}
