package data 

/* ---------------------------------------- 
type BotRequestType int8
const (
    InsertUser BotRequestType = iota + 1
    RemoveUser
    UpdateUser
    InsertSecret
    RemoveSecret
    UpdateSecret
    InsertPermission
    RemovePermission
    UpdatePermission
)

func (r BotRequestType) String() string {
    return [...]string{
        "InsertUser", 
        "RemoveUser", 
        "UpdateUser", 
        "InsertSecret", 
        "RemoveSecret", 
        "UpdateSecret",
        "InsertPermission",
        "RemovePermission",
        "UpdatePermission",
    }[r-1]
}
func (r BotRequestType) EnumIndex() int {
    return int(r)
}
 ---------------------------------------- */

/* ---------------------------------------- */
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
/* ---------------------------------------- */
