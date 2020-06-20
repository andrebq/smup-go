package helper

type (
	// DirtyFlag is used to track if some information is stale and should be updated
	DirtyFlag bool
)

// Refresh returns true if DirtyFlag is true and changes its value to false
func (d *DirtyFlag) Refresh() bool {
	v := bool(*d)
	if v {
		*d = false
	}
	return v
}

// Do executes fn if DirtyFlag is true and the dirty flag is cleared.
//
// It returns true only if fn was actually executed
func (d *DirtyFlag) Do(fn func()) bool {
	if d.Refresh() {
		fn()
		return true
	}
	return false
}
