package cmd

import (
	_ "embed"
	"fmt"
	"strings"
)

func New(tableName string) *DataTable {
	return &DataTable{
		tableName: tableName,
		conn:      dbConn,
	}
}

func (dt *DataTable) CmdInsert(data map[string]interface{}) error {
	if len(data) == 0 {
		return fmt.Errorf("empty data")
	}
	columns := ""
	params := []interface{}{}
	placeholders := ""

	i := 1
	for key, value := range data {
		columns += key + ", "
		params = append(params, value)
		placeholders += fmt.Sprintf("$%d, ", i)
		i++
	}
	columns = strings.TrimSuffix(columns, ", ")
	placeholders = strings.TrimSuffix(placeholders, ", ")
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", dt.tableName, columns, placeholders)
	err := dt.conn.Conn()
	if err != nil {
		return err
	}
	_, err = dt.conn.Exec(sql, params...)
	if err != nil {
		return fmt.Errorf("unable to insert data: %v", err)
	}
	return nil
}

func (dt *DataTable) CmdUpdate(data map[string]interface{}, id int) error {
	if len(data) == 0 || id <= 0 {
		return fmt.Errorf("empty data")
	}
	set := ""
	params := []interface{}{}
	i := 1
	for key, value := range data {
		set += fmt.Sprintf("%s = $%d, ", key, i)
		params = append(params, value)
		i++
	}
	set = strings.TrimSuffix(set, ", ")
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", dt.tableName, set, i)
	params = append(params, id)
	err := dt.conn.Conn()
	if err != nil {
		return err
	}
	_, err = dt.conn.Exec(sql, params...)
	if err != nil {
		return fmt.Errorf("unable to update data: %v", err)
	}
	return nil
}

func (dt *DataTable) CmdDelete(field string, value interface{}) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", dt.tableName, field)
	err := dt.conn.Conn()
	if err != nil {
		return err
	}
	_, err = dt.conn.Exec(sql, value)
	if err != nil {
		return fmt.Errorf("unable to delete data: %v", err)
	}
	return nil
}

func (dt *DataTable) CmdRead(field string, value interface{}) ([]map[string]interface{}, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", dt.tableName, field)
	err := dt.conn.Conn()
	if err != nil {
		return nil, err
	}
	result, err := dt.conn.Exec(sql, value)
	if err != nil {
		return nil, fmt.Errorf("can't read data form database: %v", err)
	}
	return result, nil
}
