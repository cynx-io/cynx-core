package usertype

type UserType int

const (
	Guest  = iota
	Normal = iota
)

func New(usertype string) UserType {
	switch usertype {
	case "GUEST":
		return Guest
	case "NORMAL":
		return Normal
	default:
		return Guest
	}
}

func (t UserType) String() string {
	switch t {
	case Guest:
		return "GUEST"
	case Normal:
		return "NORMAL"
	default:
		return "UNKNOWN"
	}
}
