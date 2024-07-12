package computer

type Spec struct { //exported struct
	Maker string //exported field
	Price int32  //exported field
	model string //unexported field
}
