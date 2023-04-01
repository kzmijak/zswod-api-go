package userQuery

func (uq UserQuery) QueryFullUser() UserQuery {
	query := uq.WithAvatar()
	return FromQuery(query)
}