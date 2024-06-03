package git

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

// TODO: Pull out refs
type Commit struct {
	Author    string     `json:"author"`
	Email     string     `json:"email"`
	Date      DateResult `json:"date"`
	Id        string     `json:"id"`
	Index     int        `json:"index"`
	ParentIds []string   `json:"parentIds"` // Ordered by date?
	IsMerge   bool       `json:"isMerge"`
	Message   string     `json:"message"`
	StashId   string     `json:"stashId"`
	Ref       []RefInfo  `json:"refs"`

	// Filtered means not included. Relating to Find feature?
	// TODO: Is this still used?
	Filtered   bool `json:"filtered"`
	NumSkipped int  `json:"numSkipped"`
}

// TODO: Adjustment doesn't seem to be used in the frontend.
type DateResult struct {
	Ms         int `json:"ms"`
	Adjustment int `json:"adjustment"`
}

type RefInfo struct {
	Id         string      `json:"id"`
	Location   RefLocation `json:"location"`
	FullName   string      `json:"fullName"`
	ShortName  string      `json:"shortName"`
	RemoteName string      `json:"remoteName"`
	SiblingId  string      `json:"siblingId"`
	RefType    RefType     `json:"refType"`
	Head       bool        `json:"head"`
	CommitId   string      `json:"commitId"`
	Time       int         `json:"time"` // in milliseconds (ms)
}
