package checker

type VersionComparer interface {
	Compare(v, w string) (n int, err error)
	IsValid(v string) bool
}

func Compare(c VersionComparer, v string, w string) (n int, err error) {
	return c.Compare(v, w)
}
