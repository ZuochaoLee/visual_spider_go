package models

type Conf struct {
	ID         int
	TaskName   string
	Theardnum  int
	Cron       string
	Des        string
	Dbtype     string
	Dbhost     string
	Dbport     string
	Dbname     string
	Dbuser     string
	Dbpasswd   string
	ReqType    string
	RootUrl    string
	Cookie     string
	HeaderFile string
	UseProxy   string
	TextType   string
	PostData   string
	Status     int
	PagePre    string
	PageRule   string
	PageFun    string
	PageFour   string
	PageThree  string
	PageTwo    string
	PageOne    string
}
type Ruler struct {
	ID     int
	TaskId int
	Name   string
	Rule   string
	Fun    string
	Num    string
}
