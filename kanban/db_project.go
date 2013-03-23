package main

//setter les admins multiples
func changeAdminProject(p *ConnectionPoolWrapper, id uint32, state bool) error {
	db := p.GetConnection()
	_, err := db.Exec("update projects set admin = $1 where id = $2", state, id)
	p.ReleaseConnection(db)
	return err
}

func GetNbProjects(p *ConnectionPoolWrapper) (int, error) {
	var num int
	db := p.GetConnection()
	row := db.QueryRow("select count(*) from projects")
	err := row.Scan(&num)
	p.ReleaseConnection(db)
	return num, err
}

func (u *Project) Add(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
    _, err := db.Exec("INSERT INTO projects(name, content, admins_id, read) VALUES($1, $2, $3, $4);",
		u.Name, u.admins_id, u.Read, u.Content)
	return err
}

func (u *Project) Del(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("delete from projects where id = $1", u.Id)
	return err
}

func (u *Project) Update(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
    _, err := db.Exec("update projects set name = $1, admins_id = $2, read = $3, content,  where id = $4",
	u.Name, u.admins_id, u.Read, u.Content, u.Id)
	return err
}

func (u *Project) ChangeAdmin(p *ConnectionPoolWrapper, admins []int) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	_, err := db.Exec("update projects set admins_id = $1 where id = $2", admins[0], u.Id)
	return err
}

func (u *Project) GetById(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select * from projects where id = $1", u.Id)
	err := row.Scan(&u.Id, &u.Name, &u.admins_id, &u.Read, &u.Content)
	return err
}

func (u *Project) GetByName(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select * from projects where name = $1", u.Name)
	err := row.Scan(&u.Id, &u.Name, &u.admins_id, &u.Read, &u.Content)
	return err
}

func (u *Project) Get(p *ConnectionPoolWrapper) error {
	db := p.GetConnection()
	defer p.ReleaseConnection(db)
	row := db.QueryRow("select * from projects where name = $1", u.Name)
	err := row.Scan(&u.Id, &u.Name, &u.admins_id, &u.Read, &u.Content)
	return err
}

func (u *Project) PutAdmin(p *ConnectionPoolWrapper) error {
	return changeAdminProject(p, u.Id, true)
}

func (u *Project) Unadmin(p *ConnectionPoolWrapper) error {
	return changeAdminProject(p, u.Id, false)
}