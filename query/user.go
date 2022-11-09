package yfantasy

type UserQuery struct {
	query
}

func Users() *UserQuery {
	return &UserQuery{
		query{
			resource: "users",
			params:   []string{"use_login=1"},
		},
	}
}

func (u *UserQuery) Games() *GameQuery {
	return &GameQuery{
		query{
			base:         u.ToString(),
			resource:     "games",
			isCollection: true,
		},
	}
}

func (u *UserQuery) Leagues() *LeagueQuery {
	return &LeagueQuery{
		query{
			base:         u.ToString(),
			resource:     "leagues",
			isCollection: true,
		},
	}
}

func (u *UserQuery) Teams() *TeamQuery {
	return &TeamQuery{
		query{
			base:         u.ToString(),
			resource:     "teams",
			isCollection: true,
		},
	}
}
