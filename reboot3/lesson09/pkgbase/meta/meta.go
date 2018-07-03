package meta

type Data struct {
	x int32 // private
	Y int32 // public
}

func New() *Data {
	return new(Data)
}
