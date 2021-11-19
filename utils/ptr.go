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

func BoolPtr(a bool) *bool {
	return &a
}
func BoolVal(a *bool) bool {
	if a == nil {
		return false
	}
	return *a
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

func Int32Var(a *int32) int32 {
	if a == nil {
		return 0
	}
	return *a
}

func Int64Ptr(a int64) *int64 {
	return &a
}

func Int64Var(a *int64) int64 {
	if a == nil {
		return 0
	}
	return *a
}
