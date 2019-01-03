package xjy

import (
	"strings"

	u "github.com/cdutwhu/util"
)

// var mapidll = make(map[string][]int)

// InitLineLevelBuf is (deprecated)
// func InitLineLevelBuf(rid string, nline int) {
// 	if _, ok := mapidll[rid]; !ok {
// 		mapidll[rid] = make([]int, nline)
// 	}
// }

// IsYAMLPath is
func IsYAMLPath(line string) bool {
	return line[len(line)-1] == ':' && !strings.Contains(line, ": ")
}

// IsYAMLValueLine is
func IsYAMLValueLine(line string) bool {
	/* Not a Path line */
	return !IsYAMLPath(line)
}

// YAMLTag is
func YAMLTag(line string) string {
	// fmt.Println("---", line)
	if IsYAMLValueLine(line) {
		if pos := sI(line, ": "); pos >= 0 { /* Normal 'Tag: Value' line */
			if pos1 := sI(line, "- "); pos1 >= 0 {
				return u.Str(sTL(line[pos1+2:pos], " ")).RemoveQuotes()
			}
			return u.Str(sTL(line[:pos], " ")).RemoveQuotes()
		}
		if pos := sI(line, "- "); pos >= 0 { /* Array Element '- Value' line */
			return "" /* array element obj */
		}
	}
	return u.Str(sTL(line[:len(line)-1], " ")).RemoveQuotes() /* Pure One Path Section */
}

// YAMLValue is
func YAMLValue(line string) (value string, arrEleValue bool) {
	if IsYAMLValueLine(line) {
		if pos := sI(line, ": "); pos >= 0 { /* Normal 'Sub: Obj' line */
			return sT(line[pos+2:len(line)], "\""), false
		}
		if pos := sI(line, "- "); pos >= 0 { /* Array Element '- Obj' line */
			return sT(line[pos+2:len(line)], "\""), true
		}
	}
	return "", false /* Pure One Path Section */
}

// YAMLLevel is
func YAMLLevel(line string) int {
	for i, c := range line {
		if c != ' ' && line[i+1] != ' ' {
			// PC(i%2 == 1, fmt.Errorf("error yaml format %s: in YAMLLevel", line))
			return i / 2
		}
	}
	return -1
}

// UpperLevelLine is (deprecated) (using line index is to avoid identical lines in yaml file)
func UpperLevelLine(idx int, lines []string) (int, string) {
	thislevel := YAMLLevel(lines[idx])
	// mapidll["rid"][idx] = thislevel
	if thislevel == 0 {
		return -1, ""
	}
	for i := idx - 1; i >= 0; i-- {
		//level := mapidll["rid"][i] /* much faster than YAMLLevel again */
		level := YAMLLevel(lines[i]) /* much slower than map */
		if thislevel-level == 1 {
			return i, lines[i]
		}
	}
	return -1, ""
}

// UpperLevelLines is (deprecated) (Too slow even use map, avoid using this function)
func UpperLevelLines(idx int, lines []string, self, up2low bool) (idxes []int, strs []string) {
	if self {
		idxes, strs = append(idxes, idx), append(strs, lines[idx])
	}
	idx, uline := UpperLevelLine(idx, lines)
	if idx >= 0 {
		//for mapidll["rid"][idx] > 0 {
		for YAMLLevel(lines[idx]) > 0 {
			idxes, strs = append(idxes, idx), append(strs, uline)
			idx, uline = UpperLevelLine(idx, lines)
		}
		idxes, strs = append(idxes, idx), append(strs, uline)
	}
	if up2low {
		for l, r := 0, len(idxes)-1; l < r; l, r = l+1, r-1 {
			idxes[l], idxes[r] = idxes[r], idxes[l]
			strs[l], strs[r] = strs[r], strs[l]
		}
	}
	return
}

/*******************************************************************/

// YAMLLines2Nodes is ,
func YAMLLines2Nodes(lines []string, idmark string) *[]Node {
	if !strings.HasPrefix(idmark, "-") {
		idmark = "-" + idmark
	}

	nodes := make([]Node, len(lines))
	objID := ""
	pn0 := &nodes[0]
	pn0.tag, pn0.value, pn0.path, pn0.aevalue, pn0.level, pn0.levelXPath, pn0.id = YAMLTag(lines[0]), "", YAMLTag(lines[0]), false, 0, []int{0}, objID

	for i, l := range lines[1:] {
		i++

		pn, pnlast := &nodes[i], &nodes[i-1]
		pn.tag = YAMLTag(l)
		pn.value, pn.aevalue = YAMLValue(l) /* pn.path will be filled below from levelXPath */
		pn.level = YAMLLevel(l)
		pn.levelXPath = make([]int, pn.level+1)
		copy(pn.levelXPath, pnlast.levelXPath)

		if sI(l, idmark) >= 0 {
			objID = pn.value
		}
		pn.id = objID
		if (pnlast.level == pn.level || pnlast.level == pn.level-1) && pnlast.id == "" {
			pnlast.id = pn.id
		}

		switch {
		case pn.level == pnlast.level+1: /*jump into*/
			pn.levelXPath[pn.level-1] = i - 1
		case pn.level == pnlast.level: /*next sibling*/
			//copy(pn.levelFootPath, pnlast.levelFootPath)
		case pn.level == pnlast.level-1: /*jump 1 out */
		case pn.level == pnlast.level-2: /*jump 2 out */
		case pn.level == pnlast.level-3: /*jump 3 out */
		default:
			/* incorrect file */
		}
		pn.levelXPath[pn.level] = i

		for _, p := range pn.levelXPath {
			tag := YAMLTag(lines[p])
			if strings.HasPrefix(tag, "- ") { /* remove YAML array symbol '- ' */
				tag = tag[2:]
			}
			if len(tag) == 0 {
				continue
			}
			pn.path += (tag + ".")
		}
		pn.path = pn.path[:len(pn.path)-1] /* remove last '.' */
	}

	/* remove 'RefId' nodes */
	nodesNoID := []Node{}
	for _, n := range nodes {
		if n.tag != idmark {
			nodesNoID = append(nodesNoID, n)
		}
	}

	return &nodesNoID
}

// YAMLAllValuesAsync is
func YAMLAllValuesAsync(yamlstr, objIDMark string, skipDir bool, OnOneValueFetch func(path, value, id string), done chan<- int) {
	lines := strings.FieldsFunc(yamlstr, func(c rune) bool { return c == '\n' })
	nodes := YAMLLines2Nodes(lines, objIDMark)
	for _, n := range *nodes {
		if skipDir {
			if len(sT(n.value, " ")) != 0 {
				OnOneValueFetch(n.path, n.value, n.id)
			}
		} else {
			OnOneValueFetch(n.path, n.value, n.id)
		}
	}
	done <- 1
}
