package commander

import "fmt"

type Flag struct {
	Name        string
	Description string
}

func (f *Flag) String() string {
	return fmt.Sprintf("--%v = %v", f.Name, f.Description)
}
