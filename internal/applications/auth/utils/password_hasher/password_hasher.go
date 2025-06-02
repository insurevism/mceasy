package passwordhasher

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(password, encodedHash string) (bool, error)
}
