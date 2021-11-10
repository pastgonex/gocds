package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
)

func init() {
	RootCmd.AddCommand(dbCommand)
}

//map for converting mysql type to golang types
var typeForMysqlToGo = map[string]string{
	"int":                "int64",
	"integer":            "int64",
	"tinyint":            "int64",
	"smallint":           "int64",
	"mediumint":          "int64",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int64",
	"smallint unsigned":  "int64",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func deal(dblink string, modelName string) (err error) {
	//fmt.Println(dblink)
	dbName := getDbName(dblink)
	//fmt.Println(dbName)

	// Init DB
	DB, err := initDb(dblink)
	if err != nil {
		return err
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(DB)

	cmd := "select table_name from information_schema.tables " +
		"where table_schema='" + dbName + "';"
	tables, err := DB.Query(cmd)
	if err != nil {
		return err
	}

	// 如果不存在model目录， 则创建一个
	if Exists(modelName) == false {
		err = os.Mkdir(modelName, os.ModePerm)
		if err != nil {
			return err
		}
	}

	err = dealEachTable(tables, cmd, dbName, DB, modelName)

	return err
}

func getDbName(dblink string) string {
	var dbName string // 数据库名称
	for i := len(dblink) - 1; dblink[i] != '/'; i-- {
		dbName += string(dblink[i])
	}
	dbName = reverse(dbName)
	return dbName
}

func reverse(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func dealEachTable(tables *sql.Rows, cmd string, dbName string, DB *sql.DB, modelName string) (err error) {
	for tables.Next() {
		var tableName string
		if err = tables.Scan(&tableName); err != nil {
			fmt.Println("scan table name error")
			return err
		}
		//fmt.Println(tableName)
		// 开始处理表，对于每个表，查询出所有字段名和字段类型，并生成对应的struct
		cmd = "select column_name,column_type from information_schema.columns " +
			"where table_schema='" + dbName + "' and table_name='" + tableName + "';"
		lists, err := DB.Query(cmd)
		if err != nil {
			fmt.Println("query column error")
			return err
		}

		f, err := initGoFile(modelName, tableName, err)
		if err != nil {
			return err
		}

		err = dealEachLine(lists, f)
		if err != nil {
			return err
		}
	}
	return err
}

func dealEachLine(lists *sql.Rows, f *os.File) (err error) {
	for lists.Next() {
		var columnName, columnType string
		if err = lists.Scan(&columnName, &columnType); err != nil {
			fmt.Println("scan column error")
			return err
		}
		//fmt.Println(columnName, columnType)
		columnType = handlingTypeSuffix(columnType)

		_, err := io.WriteString(f, "\t"+columnName+" "+typeForMysqlToGo[columnType]+"\n")
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(f, "}\n")
	err = f.Close()
	if err != nil {
		return err
	}
	return err
}

func initGoFile(modelName string, tableName string, err error) (*os.File, error) {
	targetFile := "./" + modelName + "/" + tableName + ".go"
	err = ioutil.WriteFile(targetFile, []byte("package "+modelName+"\n\n"), 0644)
	f, _ := os.OpenFile(targetFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	_, err = io.WriteString(f, "type "+tableName+" struct {\n")
	if err != nil {
		return nil, err
	}
	return f, err
}

func initDb(dblink string) (*sql.DB, error) {
	DB, e := sql.Open("mysql", dblink)
	if e != nil {
		return nil, e
	}
	// 判断是否已经连接上数据库
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail...")
		return nil, e
	}
	return DB, e
}

func handlingTypeSuffix(columnType string) string {
	if string(columnType[len(columnType)-1]) == ")" {
		for string(columnType[len(columnType)-1]) != "(" {
			columnType = columnType[:len(columnType)-1]
		}
		columnType = columnType[:len(columnType)-1]
	}
	return columnType
}

var dbCommand = &(cobra.Command{
	Use: "db",
	Short: "Generate structure files (go files) " +
		"corresponding to each table based on the database links",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case len(args) == 0 || len(args) == 1:
			fmt.Println("Too few arguments.")
			os.Exit(1)
		case len(args) == 2:
			if err := deal(args[0], args[1]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case len(args) > 2:
			fmt.Println("Too many arguments.")
			os.Exit(1)
		}
	},
	Example: "\"cds db root:rootroot@tcp(127.0.0.1:3306)/gocds\" model",
})
