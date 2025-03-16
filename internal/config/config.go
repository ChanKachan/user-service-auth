package config

var JWTSecretKey = &[]byte{}

func Init(JWTKey *[]byte) {
	JWTSecretKey = JWTKey

	if len(*JWTSecretKey) == 0 {
		panic("JWT secret key can't be empty")
	}
}
