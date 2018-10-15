package example

type Params struct {
	Val int
}

func (p *Params) IsValid() bool {
	return p.Val > 0
}
