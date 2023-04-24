package db

import (
	"fmt"
)

func (m *Mysql) GetBTCHeight(name string) (int, error) {
	height := 0
	sql := fmt.Sprintf("select * from t_task where f_name = \"%s\" ;", name)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&height)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, nil
	}
	return height, err
}
func (m *Mysql) GetBSCHeight(name string) (int, error) {
	height := 0
	sql := fmt.Sprintf("select * from t_mechanism where f_name = \"%s\" ;", name)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&height)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, nil
	}

	return height, err
}
func (m *Mysql) GetETHHeight(name string) (int, error) {
	height := 0
	sql := fmt.Sprintf("select * from t_mechanism where f_name = \"%s\" ;", name)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&height)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, nil
	}

	return height, err
}
func (m *Mysql) GetHUIHeight(name string) (int, error) {
	height := 0
	sql := fmt.Sprintf("select * from t_mechanism where f_name = \"%s\" ;", name)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&height)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, nil
	}

	return height, err
}
func (m *Mysql) GetTRONHeight(name string) (int, error) {
	height := 0
	sql := fmt.Sprintf("select * from f_task where f_name = \"%s\" ;", name)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&height)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, nil
	}

	return height, err
}
