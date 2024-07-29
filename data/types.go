package data

import (
	//"encoding/json"
	"fmt"
)

/* ---------------------------------------- 
type BotRequest struct {
	Type BotRequestType
	Data json.RawMessage
}

func (b BotRequest) GetUserData() (UserData, error) {
	var x UserData
	err := json.Unmarshal(b.Data, &x)
	return x, err
}
func (b BotRequest) GetSecretData() (SecretData, error) {
	var x SecretData
	err := json.Unmarshal(b.Data, &x)
	return x, err
}
func (b BotRequest) GetPermissionData() (PermissionData, error) {
	var x PermissionData
	err := json.Unmarshal(b.Data, &x)
	return x, err
}

 ---------------------------------------- */

/* ---------------------------------------- */
type SecretData struct {
	Id                int64
	Name              string
	Value             string
	Description       string
	Author            *UserData `pg:"rel:has-one"`
	EncryptedPassword string
	AllowedUses       int8
	UseCount          int8
	CreatedInstant    int64
	Lifetime          int64
}

func (s SecretData) String() string {
	return fmt.Sprintf(
		"SecretData<%d %s %s %s %s %s %d %d %d %d>",
		s.Id, s.Name, s.Value, s.Description, s.Author,
		s.EncryptedPassword, s.AllowedUses, s.UseCount,
		s.CreatedInstant, s.Lifetime)
}

/* ---------------------------------------- */

/* ---------------------------------------- */
type UserData struct {
	Id              int64
	Name            string
	Email           string
	PermissionLevel UserPermission
	DiscordId       int64
	SlackId         int64
	TeamsId         int64
	CreatedInstant  int64
}

func (u UserData) String() string {
	return fmt.Sprintf(
		"UserData<%d %s %s %s %d %d %d %d>",
		u.Id, u.Name, u.Email, u.PermissionLevel,
		u.DiscordId, u.SlackId, u.TeamsId, u.CreatedInstant,
	)
}

/* ---------------------------------------- */

/* ---------------------------------------- */
type PermissionData struct {
	Id                int64
	User              *UserData   `pg:"rel:has-one"`
	Secret            *SecretData `pg:"rel:has-one"`
	EncryptedPassword string
	AllowedUses       int8
	UseCount          int8
	CreatedInstant    int64
	Lifetime          int64
}

func (p PermissionData) String() string {
	return fmt.Sprintf(
		"PermissionData<%d %s %s %s %d %d %d %d>",
		p.Id, p.User, p.Secret, p.EncryptedPassword,
		p.AllowedUses, p.UseCount, p.CreatedInstant,
		p.Lifetime,
	)
}

/* ---------------------------------------- */
