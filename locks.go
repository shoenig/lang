package lang

// Mutex is anything that implements Lock and Unlock.
type Mutex interface {
	Lock()
	Unlock()
}

// Critical executes f while holding the mutex.
func Critical(mutex Mutex, f func()) {
	mutex.Lock()
	f()
	mutex.Unlock()
}

// CriticalGet executes f while holding the mutex, and returns the return value
// of f.
func CriticalGet[T any](mutex Mutex, f func() T) T {
	mutex.Lock()
	result := f()
	mutex.Unlock()
	return result
}
