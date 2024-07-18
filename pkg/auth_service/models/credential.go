package models

type Credential struct {
	FullName       string `json:"full_name,omitempty"`
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	HashedPassword []byte `json:"hashed_password,omitempty"`
}

func (c *Credential) ToJWTPayload() Credential {
	return Credential{
		Username: c.Username,
	}
}
