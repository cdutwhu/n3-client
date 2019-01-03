package xjy

import (
	u "github.com/cdutwhu/util"
)

const (
	lenGUID = 36 // GUID 36 chars length
)

// JSONScanObjects is (must have top-level "id" attribute like `"id": "6690e6c9-3ef0-4ed3-8b37-7f3964730bee",` )
func JSONScanObjects(json, idmark string) (ids, objstrs []string, posarr []int) {
	idmark = u.Str(idmark).MakeQuotes(u.QDouble)
	idmark = u.Str(idmark).MakeSuffix(":")

	level, arrLevel, done1, done2 := 0, 0, false, false
	for i, c := range json {
		switch c {
		case '{':
			level++
		case '}':
			level--
		case '[':
			arrLevel++
		case ']':
			arrLevel--
		}

		/* single object */
		if level == 1 && arrLevel == 0 && done1 {
			continue
		}
		if level == 1 && arrLevel == 0 && !done1 {
			if p, pe, parr := sI(json[i:], idmark), sI(json[i:], "\","), sI(json[i:], "["); p > 0 && pe > 0 && parr < 0 {
				posarr = append(posarr, i)
				ids = append(ids, json[i+pe-lenGUID:i+pe])
				done1 = true
			}
		}
		if level == 0 && arrLevel == 0 && done1 {
			objstrs = append(objstrs, json[posarr[0]:i+1])
			break
		}

		/* object array */
		if level == 2 && arrLevel == 1 && done2 {
			continue
		}
		if level == 2 && arrLevel == 1 && !done2 {
			if p, pe := sI(json[i:], idmark), sI(json[i:], "\","); p > 0 && pe > 0 {
				posarr = append(posarr, i)
				ids = append(ids, json[i+pe-lenGUID:i+pe])
				done2 = true
			}
		}
		if level == 1 && arrLevel == 1 && done2 {
			objstrs = append(objstrs, json[posarr[len(posarr)-1]:i+1])
			done2 = false
		}
	}
	return
}

// JSONObjStrByID is
func JSONObjStrByID(json, idmark, ID string) string {
	ids, objstrs, _ := JSONScanObjects(json, idmark)
	for i, id := range ids {
		if id == ID {
			return objstrs[i]
		}
	}
	return ""
}

// JSONEleStrByTag is
func JSONEleStrByTag(json, tag string) string {
	tag = u.Str(tag).MakeQuotes(u.QDouble)
	tag = u.Str(tag).MakeSuffix(":")

	level, arrLevel := 0, 0
	for _, c := range json {
		switch c {
		case '{':
			level++
		case '}':
			level--
		case '[':
			arrLevel++
		case ']':
			arrLevel--
		}

		if p := sI(json, tag); p >= 0 {
			if level == 1 {
				peR := sI(json[p:], "\",")
				hasDepth, _, _ := u.Str(json[p:p+peR+1]).HasAny('{', '}')

				if peR > 0 && !hasDepth { /* not last one, plain one */
					return u.Str(json[p : p+peR+1]).MakeBrackets(u.BCurly)
				}

				if peR < 0 && !hasDepth {
					peR = sLI(json[p:], "\"")
					return u.Str(json[p : p+peR+1]).MakeBrackets(u.BCurly)
				}

				if hasDepth { /* complex one */
					_, rR := u.Str(json[p:]).BracketsPos(u.BCurly, 1, 1)
					return u.Str(json[p : p+rR+1]).MakeBrackets(u.BCurly)
				}
			}
		}
	}
	return u.Str("").MakeBrackets(u.BCurly)
}
