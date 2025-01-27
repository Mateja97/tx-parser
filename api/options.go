package api

func WithPort(port string) func(*handler) {
	return func(h *handler) {
		h.port = port
	}
}
