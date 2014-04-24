package models

type Setting struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

func GetStringSetting(key string) (value string, err error) {
	db := GetDbSession()
	var setting *Setting

	obj, err := db.Get(&Setting{}, key)
	if obj == nil {
		return "", err
	}
	setting = obj.(*Setting)
	return setting.Value, err
}

func SetStringSetting(key, value string) (err error) {
	db := GetDbSession()
	result, err := db.Exec("UPDATE settings SET value=$1 WHERE key=$2", value, key)

	rows, err := result.RowsAffected()
	if rows == 0 {
		_, err = db.Exec("INSERT INTO settings (key, value) VALUES($1, $2)", key, value)
	}

	return err
}
