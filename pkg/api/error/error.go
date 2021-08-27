package error

type UserAlreadyExist struct{
	message string
}

func (m *UserAlreadyExist) Error() string {
	return m.message
}

func NewUserAlreadyExist(message string) *UserAlreadyExist{
	return &UserAlreadyExist{
		message: message,
	}
}

type UserNotFound struct{
	message string
}

func (m *UserNotFound) Error() string {
	return m.message
}

func NewUserNotFound(message string) *UserNotFound{
	return &UserNotFound{
		message: message,
	}
}

type InvalidUserPassword struct{
	message string
}

func (m *InvalidUserPassword) Error() string {
	return m.message
}

func NewInvalidUserPassword(message string) *InvalidUserPassword{
	return &InvalidUserPassword{
		message: message,
	}
}


type RedisClientError struct{
	message string
}

func (m *RedisClientError) Error() string {
	return m.message
}

func NewRedisClientError(message string) *RedisClientError{
	return &RedisClientError{
		message: message,
	}
}

type InvalidHeader struct{
	message string
}

func (m *InvalidHeader) Error() string {
	return m.message
}

func NewInvalidHeader(message string) *InvalidHeader{
	return &InvalidHeader{
		message: message,
	}
}

type ExpiredAccessToken struct{
	message string
}

func (m *ExpiredAccessToken) Error() string {
	return m.message
}

func NewExpiredAccessTokenError(message string) *ExpiredAccessToken{
	return &ExpiredAccessToken{
		message: message,
	}
}

type ExpiredRefreshToken struct{
	message string
}

func (m *ExpiredRefreshToken) Error() string {
	return m.message
}

func NewExpiredRefreshTokenError(message string) *ExpiredRefreshToken{
	return &ExpiredRefreshToken{
		message: message,
	}
}

type InvalidAccessToken struct{
	message string
}

func (m *InvalidAccessToken) Error() string {
	return m.message
}

func NewInvalidAccessTokenError(message string) *InvalidAccessToken{
	return &InvalidAccessToken{
		message: message,
	}
}

type InvalidRefreshToken struct{
	message string
}

func (m *InvalidRefreshToken) Error() string {
	return m.message
}

func NewInvalidRefreshTokenError(message string) *InvalidRefreshToken{
	return &InvalidRefreshToken{
		message: message,
	}
}

type TokenVerifyFail struct{
	message string
}

func (m *TokenVerifyFail) Error() string {
	return m.message
}

func NewTokenVerifyFailError(message string) *TokenVerifyFail{
	return &TokenVerifyFail{
		message: message,
	}
}