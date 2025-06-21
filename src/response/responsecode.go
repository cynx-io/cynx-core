package response

type Code string

func (r Code) String() string {
	return string(r)
}
