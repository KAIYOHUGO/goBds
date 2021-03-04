package api

// export

// Sstatus ...
// return bool,true mean something is right
type Sstatus struct {
	Status bool
}

// Sinfo ...
type Sinfo struct {
	Status bool
	Info   int8
}

// Slist ...
// return name + status array
type Slist []struct {
	Name   string
	Status bool
}

// Sterminal ...
// return line
type Sterminal struct {
	line string
}
