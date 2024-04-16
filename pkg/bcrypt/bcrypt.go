package bcrypt

import lib_bcrypt "golang.org/x/crypto/bcrypt"

type Interface interface {
	GenerateFromPassword(password string) (string, error)
	CompareAndHashPassword(hashPassword string, password string) error
}

type bcrypt struct {
	cost int
}

func Init() Interface {
	return &bcrypt{
		cost: 10,
	}
}

func (b *bcrypt) GenerateFromPassword(password string) (string, error) {
	bytePassword, err := lib_bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func (b *bcrypt) CompareAndHashPassword(hashPassword string, password string) error {
	if err := lib_bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}
