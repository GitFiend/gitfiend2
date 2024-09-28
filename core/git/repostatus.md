# Repo Status Algorithm

Need to check for
- Whether we have modified files
  - Working dir and staging
  - Merge or Rebase conflict
- Current branch commit id and remote branch commit id
- Whether config has changed? What does this effect?
- All branch names so we know which branch across all repos we can switch to

## Working Dir
Probably can't avoid running git status.

## Push/Pull status
Currently using countCommitsBetweenFallback to count num of commits different. We do this always if
we find the head and remote and refs dir with no memo.

## Config
We could load this just once and not bother again until a repo is selected? 
We could check if the file has changed (os.Stat) rather than reloading and parsing if we do need to 
keeping checking for changes.
The current parser does extra work than needed, so could be sped up a little.

## All branches
Tricky problem, though not sure how fast the current thing is. 
We could run it always and use pack-refs to speed it up if there are tons of files.
We could check if packed refs file has been modified? Not sure how fast os.Stat is vs just reading a smallish file is.
