package gin_binding

type Person struct {
	Name   string `form:"name"`
	Addres string `form:"address"`
}
