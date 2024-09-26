package refs

func P[T any](t T) *T {
	return &t
}

func String(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func Int(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func Int32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func Int64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func Uint(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

func Uint32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func Uint64(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func Float32(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

func Float64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func Bool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}
