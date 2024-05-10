package gitTypes

type RefInfoPart struct {
	Id         string
	Location   RefLocation
	FullName   string
	ShortName  string
	RemoteName string
	SiblingId  string
	RefType    RefType
	Head       bool
}

type RefType int

const (
	Branch RefType = iota
	Tag
	Stash
)

type RefLocation int

const (
	Local RefLocation = iota
	Remote
)

type Commit struct {
	Author    string
	Email     string
	Date      DateResult
	Id        string
	Index     int
	ParentIds []string // Ordered by date?
	IsMerge   bool
	Message   string
	StashId   string
	Ref       []RefInfo

	// Filtered means not included. Relating to Find feature?
	Filtered   bool
	NumSkipped int
}

type DateResult struct {
	Ms         int
	Adjustment int
}

type RefInfo struct {
	Id         string
	Location   RefLocation
	FullName   string
	ShortName  string
	RemoteName string
	SiblingId  string
	RefType    RefType
	Head       bool
	CommitId   string
	Time       int // in milliseconds (ms)
}
