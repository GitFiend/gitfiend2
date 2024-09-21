package git

import (
	p "gitfiend2/core/parser"
	"log/slog"
	"os"
	"path"
	"strings"
)

// TODO: Does this always exist? It may be normal for this to fail?
func loadPackedRefs(repoPath string) ([]PackedRef, bool) {
	repo, ok := cache.GetRepoPath(repoPath)
	if !ok {
		slog.Error("repo missing, so couldn't load packed refs")
		return nil, false
	}
	packedDir := path.Join(repo.GitPath, "packed-refs")
	bytes, err := os.ReadFile(packedDir)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}
	text := string(bytes)

	res, ok := p.ParseAll(pLines, text)
	if ok {
		return res, true
	}

	return nil, false
}

type PackedRef struct {
	// RemoteName left empty if we have a local ref.
	// Some lines contain junk, so check for empty name.
	CommitId, RemoteName, Name string
}

var pLocalRef = p.Map(p.And4(p.AnyWord, p.Rune(' '), p.Word("refs/heads/"), p.UntilLineEnd), makeLocal)

func makeLocal(res p.And4Result[string, rune, string, string]) PackedRef {
	return PackedRef{
		CommitId: res.R1,
		Name:     res.R4,
	}
}

var pRemoteRef = p.Map(p.And4(p.AnyWord, p.Rune(' '), p.Word("refs/remotes/"), p.UntilLineEnd), makeRemote)

func makeRemote(res p.And4Result[string, rune, string, string]) PackedRef {
	remoteName, name := removeRemote(res.R4)

	return PackedRef{
		CommitId:   res.R1,
		RemoteName: remoteName,
		Name:       name,
	}
}

var pOtherRef = p.Map(p.UntilLineEnd, func(result string) PackedRef {
	return PackedRef{}
})

var pLines = p.Many(p.Or(pLocalRef, pRemoteRef, pOtherRef))

func removeRemote(refPart string) (string, string) {
	parts := strings.SplitN(refPart, "/", 2)
	return parts[0], parts[1]
}
