package utils

func StrPtr(a string) *string {
	return &a
}
func StrVal(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}

func SafeStrPtr(a string) *string {
	if a == "" {
		return nil
	}
	return &a
}

func BoolPtr(a bool) *bool {
	return &a
}
func BoolVal(a *bool) bool {
	if a == nil {
		return false
	}
	return *a
}

func SafeBoolPtr(a bool) *bool {
	if a {
		return nil
	}
	return &a
}

func IntPtr(a int) *int {
	return &a
}
func IntVal(a *int) int {
	if a == nil {
		return 0
	}
	return *a
}
func Int32Ptr(a int32) *int32 {
	return &a
}

func Int32Val(a *int32) int32 {
	if a == nil {
		return 0
	}
	return *a
}

func Int64Ptr(a int64) *int64 {
	return &a
}

func Int64Val(a *int64) int64 {
	if a == nil {
		return 0
	}
	return *a
}

func Float32Val(a *float32) float32 {
	if a == nil {
		return 0
	}
	return *a
}

func Float32Ptr(a float32) *float32 {
	return &a
}

func Float64Val(a *float64) float64 {
	if a == nil {
		return 0
	}
	return *a
}

func Float64Ptf(a float64) *float64 {
	return &a
}
