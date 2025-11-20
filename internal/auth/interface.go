package auth

type PassHasher interface {
	Hash(password string) (string, error)
	//Verify(password, hashedPassword string) (bool, error)
}
