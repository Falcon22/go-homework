package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/pkg/errors"
)

type myFiles []os.FileInfo

func (files myFiles) Len() int {
	return len(files)
}

func (files myFiles) Less(i, j int) bool {
	return files[i].Name() < files[j].Name()
}

func (files myFiles) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

func (files myFiles) getLastInd(withFiles bool) int {
	if withFiles {
		return len(files) - 1
	}
	lastDirInd := 0
	for i := range files {
		if files[i].IsDir() {
			lastDirInd = i
		}
	}
	return lastDirInd
}

func (t *tree) printFile(prefix string, fi os.FileInfo, last bool) string {
	var newPrefix, symbol, postfix string
	if last {
		postfix = "└───"
		symbol = ""
	} else {
		postfix = "├───"
		symbol = "│"
	}
	fmt.Fprint(t.out, prefix, postfix, fi.Name())
	newPrefix = prefix + symbol + "\t"
	if !fi.IsDir() {
		if fi.Size() == 0 {
			fmt.Fprint(t.out, " (empty)")
		} else {
			fmt.Fprintf(t.out, " (%db)", fi.Size())
		}
	}
	fmt.Fprintln(t.out)
	return newPrefix
}

func (t *tree) printDir(path string, prefix string) error {
	dir, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "can't open dir %s", path)
	}

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return errors.Wrapf(err, "can't read dir %s", path)
	}

	if err = dir.Close(); err != nil {
		return errors.Wrap(err, "can't close dir")
	}

	files := myFiles(fileInfos)
	sort.Sort(files)
	last := files.getLastInd(t.printFiles)
	newPrefix := prefix
	for _, fi := range files {
		if t.printFiles || fi.IsDir() {
			newPrefix = t.printFile(prefix, fi, fi == fileInfos[last])
		}
		if fi.IsDir() {
			newPath := path + "/" + fi.Name()
			err = t.printDir(newPath, newPrefix)
			if err != nil {
				return errors.Wrap(err, "can't print dir")
			}
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	t := tree{
		out:        out,
		printFiles: printFiles,
	}
	return t.printDir(path, "")
}

type tree struct {
	out        io.Writer
	printFiles bool
}
