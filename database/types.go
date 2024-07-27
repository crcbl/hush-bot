package main

type Secret struct {
	Id                int64
	Name              string
	Value             string
	Description       string
	Author            *User `pg:"rel:has-one"`
	EncryptedPassword string
	AllowedUses       int8
	UseCount          int8
	CreatedInstant    int64
	Lifetime          int64
}

type User struct {
	Id              int64
	Name            string
	Email           string
	PermissionLevel UserPermission
	DiscordId       int64
	SlackId         int64
	TeamsId         int64
	CreatedInstant  int64
}

type Permission struct {
	Id                int64
	User              *User   `pg:"rel:has-one"`
	Secret            *Secret `pg:"rel:has-one"`
	EncryptedPassword string
	AllowedUses       int8
	UseCount          int8
	CreatedInstant    int64
	Lifetime          int64
}

// User permission enum
type UserPermission int8

const (
	NONE       UserPermission = iota // 0
	READ                             // 1
	WRITE                            // 2
	READ_WRITE                       // 3
)

func (p UserPermission) String() string {
	return [...]string{"NONE", "READ", "WRITE", "READ_WRITE"}[p]
}
func (p UserPermission) EnumIndex() int {
	return int(p)
}
