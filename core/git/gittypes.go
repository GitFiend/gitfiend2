package git

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

// CommitInfo
// TODO: Should probably be private
type CommitInfo struct {
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
	// TODO: Is this still used?
	Filtered   bool
	NumSkipped int
}

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
	Ref       []string   `json:"refs"`

	// Filtered means not included. Relating to Find feature?
	// TODO: Is this still used?
	Filtered   bool `json:"filtered"`
	NumSkipped int  `json:"numSkipped"`
}

// DateResult
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

type PatchType rune

const (
	A PatchType = 'A'
	C           = 'C'
	B           = 'B'
	D           = 'D'
	M           = 'M'
	R           = 'R'
	T           = 'T'
	U           = 'U'
	X           = 'X'
)

type Patch struct {
	CommitId  string    `json:"commitId"`
	OldFile   string    `json:"oldFile"`
	NewFile   string    `json:"newFile"`
	PatchType PatchType `json:"patchType"`
	Id        string    `json:"id"`
	IsImage   bool      `json:"isImage"`
}
