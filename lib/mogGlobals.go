package lib

// Some parameters may vary per connection
// Recommended to create new libGlobals instead of reusing
type libGlobals struct {
	// TODO fill with parameters
	Hotbar []HotbarRec
}

func NewGlobals() *libGlobals {
	return &libGlobals{
		Hotbar: make([]HotbarRec, GlobalHotbarAmount),
	}
}

const (
	GlobalHotbarAmount = 13
)
