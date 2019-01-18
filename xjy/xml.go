package xjy

import (
	"errors"

	u "github.com/cdutwhu/go-util"
)

// XMLScanObjects is ( ONLY LIKE  <SchoolInfo RefId="D3F5B90C-D85D-4728-8C6F-0D606070606C"> )
func XMLScanObjects(xml, idmark string) (ids, objtags []string, posarr []int) {
	idmark = u.Str(idmark).MakePrefix(" ")
	idmark = u.Str(idmark).MakeSuffix("=")
	lengthID := len(idmark)
	pLastAbs := 0
LOOKFOROBJ:
	if p := sI(xml[pLastAbs:], idmark); p > 0 {
		if op := sLI(xml[pLastAbs:pLastAbs+p], "<"); op >= 0 {
			obj := xml[pLastAbs : pLastAbs+p][op+1:]
			objtags = append(objtags, obj)
			ps := pLastAbs + op
			posarr = append(posarr, ps)
		}
		if ip := sI(xml[pLastAbs+p:], ">"); ip > 0 {
			id := xml[pLastAbs+p+lengthID : pLastAbs+p+ip]
			id = sT(id, "\"")
			ids = append(ids, id)
		}
		pLastAbs += (p + lengthID)
		goto LOOKFOROBJ
	}
	return
}

// XMLObjStrByID is
func XMLObjStrByID(xml, idmark, rid string) string {
	ids, objtags, posarr := XMLScanObjects(xml, idmark)
	for i, id := range ids {
		if id == rid {
			if i != len(ids)-1 {
				return sTR(xml[posarr[i]:posarr[i+1]], " \t\r\n")
			}
			/* last object */
			endtag := "</" + objtags[i] + ">"
			if end := sI(xml[posarr[i]:], endtag); end > 0 {
				return sTR(xml[posarr[i]:posarr[i]+end+len(endtag)], " \t\r\n")
			}
		}
	}
	return ""
}

// XMLEleStrByTag is (should only be used in one object string)
// func XMLEleStrByTag(xml, tag string) string {
// 	s, s1 := sI(xml, fSpf("<%s>", tag)), sI(xml, fSpf("<%s ", tag))
// 	if s1 > s {
// 		s = s1
// 	}
// 	if s >= 0 {
// 		if e := sI(xml[s:], fSpf("</%s>", tag)); e > 0 {
// 			return xml[s : s+e+len(tag)+3]
// 		}
// 		PE(errors.New("Not a valid XML"))
// 	}
// 	return ""
// }

// XMLEleStrByTag : (index from 1)
func XMLEleStrByTag(xml, tag string, idx int) string {
	startNext, cnt := 0, 0
AGAIN:
	xml = xml[startNext:]
	s, s1 := sI(xml, fSpf("<%s>", tag)), sI(xml, fSpf("<%s ", tag))
	if s1 > s {
		s = s1
	}
	if s >= 0 {
		if peR := sI(xml[s:], fSpf("</%s>", tag)); peR > 0 {
			startNext = s + peR + len(tag) + 3
			cnt++
			if idx == cnt {
				return xml[s:startNext]
			}
			goto AGAIN
		}
		PE(errors.New("Invalid XML"))
	}
	return ""
}

// XMLFindAttributes is (ONLY LIKE  <SchoolInfo RefId="D3F5B90C-D85D-4728-8C6F-0D606070606C" Type="LGL">)
func XMLFindAttributes(xmlele string) (attributes, attriValues []string, attributeList string) { /* 'map' may cause mis-order, so use slice */
	l := len(xmlele)
	if l == 0 || xmlele[0] != '<' || xmlele[l-1] != '>' {
		PE(fEpf("Not a valid XML section"))
		return nil, nil, ""
	}

	tag := xmlele[sLI(xmlele, "</")+2 : l-1]
	if eol := sI(xmlele, "\">") + 1; xmlele[len(tag)+1] == ' ' && eol > len(tag) { /* has attributes */
		kvs := sFF(xmlele[len(tag)+2:eol], func(c rune) bool { return c == ' ' })
		for _, kv := range kvs {
			kvstrs := sFF(kv, func(c rune) bool { return c == '=' })
			attributes = append(attributes, ("-" + kvstrs[0])) /* mark '-' before attribute for differentiating child */
			attriValues = append(attriValues, u.Str(kvstrs[1]).RemoveQuotes())
		}
	}
	return attributes, attriValues, sJ(attributes, " + ")
}

// XMLFindChildren : (NOT search grandchildren)
func XMLFindChildren(xmlele, id string) (uids, children []string, childList string, arrCnt int) {
	l := len(xmlele)
	if l == 0 || xmlele[0] != '<' || xmlele[l-1] != '>' {
		fPln(xmlele)
		PE(fEpf("Not a valid XML section"))
		return nil, nil, "nil", -1
	}

	skip, childpos, level, inflag := false, []int{}, 0, false
	for i, c := range xmlele[1:] { // skip the first '<'
		i++

		if c == '<' && xmlele[i:i+4] == "<!--" {
			skip = true
		}
		if c == '>' && xmlele[i-2:i+1] == "-->" {
			skip = false
		}
		if skip {
			continue
		}

		if c == '<' && xmlele[i+1] != '/' {
			level++
		}
		if c == '<' && xmlele[i+1] == '/' {
			level--
			if level == 0 {
				inflag = false
			}
		}

		if level == 1 {
			if !inflag {
				childpos = append(childpos, i+1)
				inflag = true
			}
		}
	}

	for _, p := range childpos {
		pe, peA := sI(xmlele[p:], ">"), sI(xmlele[p:], " ")
		if peA > 0 && peA < pe {
			pe = peA
		}
		child := xmlele[p : p+pe]
		children = append(children, child)
		uids = append(uids, id)
	}

	if len(children) > 1 && u.AllAreIdentical(children...) {
		return uids, children, fSpf("[%d]%s", len(children), children[0]), len(children)
	}

	return uids, children, sJ(children, " + "), 0
}

// XMLYieldFamilyTree is (We pack attributes in return map, value like '-...')
func XMLYieldFamilyTree(xmlstr string, ids, objs []string, skipNoChild bool, mapkeyprefix string,
	mapEleChildlist *map[string]string, mapEleObjIDArrcnt *map[string]objIDArrcnt) {

	if len(mapkeyprefix) > 0 {
		mapkeyprefix += "."
	}
	for i, obj := range objs {
		curPath := mapkeyprefix + obj

		if _, ok := (*mapEleChildlist)[curPath]; ok {
			continue
		}

		xmlele := XMLEleStrByTag(xmlstr, obj, 1)
		uids, children, childlist, arrCnt := XMLFindChildren(xmlele, ids[i]) /* uniform ids, children */
		attributes, _, attributeList := XMLFindAttributes(xmlele)            /* attributes */

		/* array children info */
		if arrCnt > 0 {
			(*mapEleObjIDArrcnt)[curPath] = objIDArrcnt{
				objID:  ids[i],
				arrCnt: arrCnt,
			}
		}

		/* attributes */
		if len(attributes) > 0 {
			(*mapEleChildlist)[curPath] = attributeList + " + "
			if len(children) == 0 {
				(*mapEleChildlist)[curPath] += "#content"
			}
		}

		/* children */
		if skipNoChild {
			if len(children) > 0 {
				(*mapEleChildlist)[curPath] += childlist
			}
		} else {
			(*mapEleChildlist)[curPath] += childlist
		}

		if len(children) == 0 && len(attributeList) == 0 { /* attributes */
			continue
		} else {
			/* recursive */
			XMLYieldFamilyTree(xmlele, uids, children, skipNoChild, curPath, mapEleChildlist, mapEleObjIDArrcnt)
		}
	}
}

// ObjIDArrcnt :
type objIDArrcnt struct {
	objID  string
	arrCnt int
}

// XMLStructAsync is
func XMLStructAsync(xmlstr, ObjIDMark string, skipNoChild bool, OnStruFetch func(string, string), OnArrFetch func(string, string, int), done chan<- int) {
	ids, objs, _ := XMLScanObjects(xmlstr, ObjIDMark)
	mapEleChildlist, mapEleObjIDArrcnt := &map[string]string{}, &map[string]objIDArrcnt{}
	XMLYieldFamilyTree(xmlstr, ids, objs, skipNoChild, "", mapEleChildlist, mapEleObjIDArrcnt)
	// for k := range *mapEleChildlist {
	// 	(*mapEleChildlist)[k] = sTR((*mapEleChildlist)[k], "+ ")
	// }
	for k, v := range *mapEleChildlist {
		OnStruFetch(k, sTR(v, "+ "))
	}
	for k, v := range *mapEleObjIDArrcnt {
		OnArrFetch(k, v.objID, v.arrCnt)
	}
	done <- 1
}

/*********************************************************************/

// /* Only work on well-formated xml file like below */
// /* If NOT, use X(ml)file2Y(aml) */
// // <Name Type="LGL">
// // --><FamilyName>Smith</FamilyName>
// // --><GivenName>Fred</GivenName>
// // --><FullName>Fred Smith</FullName>
// // </Name>

// // IsXMLPath is
// func IsXMLPath(line string) bool {
// 	return sI(line, "</") == -1
// }

// // IsXMLEndTag is
// func IsXMLEndTag(line string) bool {
// 	// return !IsXMLPath(line) && (line[:2] == "</" || sI(line, "\t</") >= 0)
// 	return sHP(sTL(line, "\t"), "</")
// }

// // IsXMLValue is
// func IsXMLValue(line string) bool {
// 	return !IsXMLPath(line) && !IsXMLEndTag(line)
// }

// // XMLTag is
// func XMLTag(line string) string {
// 	l, r := sI(line, "<"), sI(line, ">")
// 	return sFF(line[l+1:r], func(c rune) bool { return c == ' ' })[0]
// }

// // XMLValue is
// func XMLValue(line string) string {
// 	if IsXMLValue(line) {
// 		l, r := sI(line, ">"), sI(line, "</")
// 		return line[l+1 : r]
// 	}
// 	return ""
// }

// // XMLObjRefId is
// // func XMLObjRefId(line string) string {

// // }

// // XMLAttr is
// func XMLAttr(line string) (tags []string, values []string) {
// 	// if sI(line, " RefId=") >= 0 {
// 	// 	return nil, nil
// 	// }
// 	if IsXMLEndTag(line) {
// 		return nil, nil
// 	}
// 	l, r := sI(line, "<"), sI(line, ">")
// 	if line[l+1:r] == XMLTag(line) {
// 		return nil, nil
// 	}

// 	attrs := sFF(line[l+1:r], func(c rune) bool { return c == ' ' })
// 	for _, attr := range attrs {
// 		av := sFF(attr, func(c rune) bool { return c == '=' })
// 		tags = append(tags, av[0])
// 		values = append(values, sT(av[1], "\""))
// 	}
// 	return
// }

// // XMLLevel is
// func XMLLevel(line string) int {
// 	if IsXMLEndTag(line) {
// 		return -1
// 	}
// 	for i, c := range line {
// 		if c != '\t' {
// 			return i
// 		}
// 	}
// 	return -1
// }

// // XMLLines2Nodes is
// // func XMLLines2Nodes(lines []string, idmark string) *[]Node {
// // 	nodes := []Node{}
// // 	ignore := false
// // 	objID := ""

// // 	for i, l := range lines {
// // 		if sI(l, "<!--") >= 0 {
// // 			ignore = true
// // 			continue
// // 		}
// // 		if sI(l, "-->") >= 0 {
// // 			ignore = false
// // 			continue
// // 		}
// // 		if ignore {
// // 			continue
// // 		}

// // 		// nodes = append(nodes, Node{})
// // 		// pn, pnlast := &nodes[len(nodes)-1], &nodes[len(nodes)-1]
// // 		// if len(nodes) > 1 {
// // 		// 	pnlast = &nodes[len(nodes)-2]
// // 		// }

// // 		// pn.tag = XMLTag(l)
// // 		// pn.value = XMLValue(l)
// // 		// pn.level = XMLLevel(l)
// // 		// pn.levelXPath = make([]int, pn.level+1)
// // 		// copy(pn.levelXPath, pnlast.levelXPath)

// // 		// if ts, vs := XMLAttr(l); ts != nil && len(ts) > 0 { /* with attributes */
// // 		// 	if ts[0] == idmark {
// // 		// 		objID = vs[0]
// // 		// 	}

// // 		// }

// // 		// pn.id = objID
// // 	}

// // 	return &nodes
// // }
