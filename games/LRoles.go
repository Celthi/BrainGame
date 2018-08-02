package Games


type LRoles struct {
	Name string
	Skill string
}

func (role LRoles) Play() string {
	return role.Name + "is playing"
}