package driver

type Orm struct {
}

var db = connect()

func (Orm) GetOne(m interface{}, where map[string]interface{}) (int64, error) {

	result := db.Where(where).Order("parent_id desc").First(&m)
	return result.RowsAffected, result.Error
}
