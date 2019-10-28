package goaext

type middlewareKey int

const (
	RealIPKey middlewareKey = iota + 1
	ReqIDKey
)
