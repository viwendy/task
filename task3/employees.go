package main

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func SelectEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", department)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func SelectHighestSalaryEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func main() {
	db := sqlx.MustConnect("mysql", "user:password@tcp(localhost:3306)/dbname")
	SelectEmployeesByDepartment(db, "技术部")
	SelectHighestSalaryEmployee(db)
}
