package selltech

// SDNEntry структура для разбора XML
type SDNEntry struct {
	Uid       string `xml:"uid"`
	FirstName string `xml:"firstName"`
	LastName  string `xml:"lastName"`
	SDNType   string `xml:"sdnType"`
}

// SDNList структура для разбора XML
type SDNList struct {
	Entries []SDNEntry `xml:"sdnEntry"`
}

// NameResult представляет результат запроса имен
type NameResult struct {
	UID       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
