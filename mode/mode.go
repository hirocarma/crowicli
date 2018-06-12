package mode

// Mode is interface that has Api function.
type Mode interface {
    API() (string, error)
}
