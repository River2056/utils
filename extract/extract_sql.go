package extract

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"utils/common"

	_ "github.com/go-sql-driver/mysql"
)

func ExtractSQL() {
	var s string
	s = os.Args[1]
	s = strings.ReplaceAll(s, ":campaignId", "6")
	s = strings.ReplaceAll(s, ":startDate", "'2023-08-01'")
	s = strings.ReplaceAll(s, ":endDate", "'2023-08-31'")
	s = strings.ReplaceAll(s, ":countryCode", "7043")

	fmt.Println(s)
	db, err := sql.Open("mysql", "au-transfer-mysql8:7BGW044kBbCnY4bI6pf5UW@tcp(18.235.112.133:49998)/vfsc_business")
	common.CheckError(err)
	defer db.Close()

	row := db.QueryRow(s)
	var data interface{}
	err = row.Scan(&data)
	common.CheckError(err)
	fmt.Println(data)
}
