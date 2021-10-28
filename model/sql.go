package model

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "database/sql"
    "fmt"
    "log"
    "strings"
    "os"
    "errors"
)

type Query struct {
    where string
    column []string
    limit int
    offset int
    table string
    order []string
    connection string
}

type Example struct {
    Idx int `db:"id"`
    Msg string `db:"msg"`
}

var DB map[string]*sqlx.DB

/*
 * DB 연결
 *
 * @param string $connection 연결DB 이름
 *
 * @return *sqlx.DB 연결된 DB
*/
func Conn() *sqlx.DB {
    connection := "default"
    if DB == nil{
        DB = make(map[string]*sqlx.DB)
    }
    db := DB[connection]
    if db != nil {
        return db
    }
    db, err := sqlx.Connect("mysql", os.Getenv("MYSQL_CONNECT_STRING"))
    if err != nil {
        log.Print(err)
    }
    row, _ := db.Queryx("SET NAMES utf8mb4")
    defer row.Close()

    DB[connection] = db
    return db
}

/*
 * Select
 *
 * @param interface $dest 할당받을 변수
 * @param string $sql 실행할 쿼리
 * @param string $conn 커넥션 이름
 *
 * @return err
*/
func Get(dest interface{}, sql string, params ...string) error {
    db := Conn()
    return db.Select(dest, sql)
}

func Scan(sql string, params ...string) string {
    db := Conn()
    row := db.QueryRowx(sql)
    if row == nil {
        return ""
    }
    var result string
    row.Scan(&result)
    return result
}
/*
 * Update
 *
 * @param map[string]interface{} $data 저장할 데이터
 * @param string $table 테이블 이름
 * @param string $where 조건절
 *
 * @return void
*/
func Update(data map[string]interface{}, table string, where string, param ...interface{}) (sql.Result, error) {
    db := Conn()

    set := []string{}
    for column, value := range data {
        set = append(set, fmt.Sprintf("`%s` = '%s'", column, value))
    }
    sql := fmt.Sprintf("UPDATE `%s` SET %s WHERE %s", table, strings.Join(set, ", "), where)
    if len(param) > 0 {
        sql = fmt.Sprintf("%s LIMIT %d", sql, param[0])
    }
    return db.Exec(sql)
}

func Insert(data interface{}, table string) (result sql.Result, err error) {
    db := Conn()
    sql, err := InsertSql(data, table)
    if err != nil {
        return
    }
    return db.Exec(sql)
}

func InsertSql(data interface{}, table string) (string, error) {
    sql := ""

    if row, ck := data.(map[string]interface{}); ck {
        // 한개
        columns := []string{}
        values := []string{}
        for column, value := range row {
            columns = append(columns, column)
            values = append(values, value.(string))
        }
        sql = fmt.Sprintf("INSERT INTO `%s` (`%s`) VALUES ('%s')", table, strings.Join(columns, "`, `"), strings.Join(values, "', '"))
    } else if rows, ck := data.([]map[string]interface{}); ck {
        // 여러개
        columns := []string{}
        values2 := []string{}
        for i, row := range rows {
            if i == 0 {
                for column, _ := range row {
                    columns = append(columns, column)
                }
            }
            values := []string{}
            for _, column := range columns {
                if value, ck := row[column]; ck {
                    values = append(values, value.(string))
                } else {
                    return sql, errors.New("invalid data format")
                }
            }
            values2 = append(values2, strings.Join(values, "', '"))
        }
        sql = fmt.Sprintf("INSERT INTO `%s` (`%s`) VALUES\n('%s')", table, strings.Join(columns, "`, `"), strings.Join(values2, "'),\n('"))
    }
    return sql, nil
}

func Upsert(data interface{}, table string, option map[string]string) (sql.Result, error) {
    db := Conn()

    sql, err := InsertSql(data, table)
    if err != nil {
        return nil, err
    }
    upsert := []string{}
    for column, value := range option {
        upsert = append(upsert, fmt.Sprintf("`%s` = %s", column, value))
    }
    sql = fmt.Sprintf("%s ON DUPLICATE KEY UPDATE %s", sql, strings.Join(upsert, ", "))

    return db.Exec(sql)
}
