package htemel

import "io"

type Node interface {
	Render(w io.Writer) error
}
