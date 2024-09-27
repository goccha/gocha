package refs

func P[T any](t T) *T {
	return &t
}

func String(v *string, defaultValue ...string) string {
	if v == nil {
		return ""
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Int(v *int, defaultValue ...int) int {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Int32(v *int32, defaultValue ...int32) int32 {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Int64(v *int64, defaultValue ...int64) int64 {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Uint(v *uint, defaultValue ...uint) uint {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Uint32(v *uint32, defaultValue ...uint32) uint32 {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Uint64(v *uint64, defaultValue ...uint64) uint64 {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Float32(v *float32, defaultValue ...float32) float32 {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Float64(v *float64, defaultValue ...float64) float64 {
	if v == nil {
		return 0
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}

func Bool(v *bool, defaultValue ...bool) bool {
	if v == nil {
		return false
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return *v
}
