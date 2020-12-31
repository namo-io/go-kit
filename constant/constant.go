package constant

// ContextKey not to use basic type string
type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

var (
	// ContextApplication application name
	ContextApplication = ContextKey("application")

	// ContextRequestID contextRequestID for tracing
	ContextRequestID = ContextKey("requestID")
)
