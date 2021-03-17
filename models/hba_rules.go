package models

type Rules struct {
	LineNumber int32  `gorm:"column:line_number;"`
	Type       string `gorm:"column:type;"`
	database   string `gorm:"column:database;"`
	UserName   string `gorm:"column:user_name;"`
	Address    string `gorm:"column:address;"`
	Netmask    string `gorm:"column:netmask;"`
	AuthMethod string `gorm:"column:auth_method;"`
	Options    string `gorm:"column:options;"`
	Error      string `gorm:"column:error;"`
}
