package gin_binding

type Person struct {
	Name   string `form:"name"`
	Addres string `form:"address"`
}

type BindinPerson struct {	
	Name      string    `form:"name,default=ilham"`
	Age       int       `form:"age,default=20"`
	Friends   []string  `form:"friends,default=ilham;ramadan"`
	Addresses [2]string `form:"addresses,default=indah safitri" collection_format:"ssv"`
	LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

type ApiResponse struct {
	Code    int
	Message string
	Status  string
	Data    any
}
