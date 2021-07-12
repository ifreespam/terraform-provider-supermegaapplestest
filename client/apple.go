package client

type Apple struct {
	Name  string `json:"name"`
	Sort  string `json:"sort"`
	Owner string `json:"owner"`
	Count int    `json:"count"`
}

// HashApple is doing something
// Deprecated: should not be used
func HashApple(apple *Apple) string {
	return "[" + apple.Name + "|" + apple.Sort + "]"
}

// HashAppleV2 recommended for usage
// Deprecated: not recommended
func HashAppleV2(apple *Apple) string {
	return "[<<" + apple.Name + "|" + apple.Sort + ">>]"
}

func HashAppleV3(apple *Apple) string {
	return "[[[" + apple.Name + "|||" + apple.Sort + "]]]"
}
