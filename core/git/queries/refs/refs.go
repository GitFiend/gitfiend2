package refs

import (
	. "gitfiend2/core/git"
	. "gitfiend2/core/parser"
	"regexp"
	"strings"
)

var refNameParser = Regex(regexp.MustCompile(`[^,\s()]+`))

var PRefName = Map(
	refNameParser, func(result string) RefInfoPart {
		cleaned := strings.Replace(result, "^{}", "", -1)
		parts := strings.Split(cleaned, "/")
		refType := getTypeFromName(parts[1])

		return RefInfoPart{
			Id:         cleaned,
			RefType:    refType,
			Location:   getRefLocation(parts),
			ShortName:  getShortName(parts),
			FullName:   cleaned,
			RemoteName: getRemoteName(parts),
		}
	},
)

func getTypeFromName(second string) RefType {
	switch second {
	case "tags":
		return Tag
	case "stash":
		return Stash
	default:
		return Branch
	}
}

func getRefLocation(refParts []string) RefLocation {
	if len(refParts) >= 3 {
		if refParts[1] == "heads" {
			return Local
		}
		return Remote
	}
	return Local
}

func getShortName(refParts []string) string {
	if refParts[1] == "remotes" {
		return strings.Join(refParts[3:], "/")
	}
	return strings.Join(refParts[2:], "/")
}

func getRemoteName(refParts []string) string {
	if len(refParts) > 3 && refParts[1] == "remotes" {
		return refParts[2]
	}
	return ""
}

var PTagRef = Map(
	And2(Word(`tag: `), PRefName), func(result And2Result[string, RefInfoPart]) RefInfoPart {
		return result.B
	},
)

var PHeadRef = Map(
	And2(Word(`HEAD -> `), PRefName), func(result And2Result[string, RefInfoPart]) RefInfoPart {
		return result.B
	},
)

var PCommitRef = Or(PHeadRef, PTagRef, PRefName)

var PCommitRefs = Map(
	And3(Char('('), RepSep(PCommitRef, ","), Char(')')),
	func(result And3Result[rune, []RefInfoPart, rune]) []RefInfoPart {
		return result.B
	},
)

var POptionalRefs = Or(
	PCommitRefs, OptionalWhiteSpace[[]RefInfoPart](),
)

func MakeRefInfo(part RefInfoPart) RefInfo {
	ref := RefInfo{
		Id:         part.Id,
		Location:   part.Location,
		FullName:   part.FullName,
		ShortName:  part.ShortName,
		RemoteName: part.RemoteName,
		SiblingId:  part.SiblingId,
		RefType:    part.RefType,
		Head:       part.Head,
	}

	return ref
}
