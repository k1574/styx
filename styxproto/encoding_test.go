package styxproto

import "testing"

func TestEncode(t *testing.T) {
	var (
		qbuf    = make([]byte, 13)
		buf     = make([]byte, MinBufSize)
		statbuf = make([]byte, maxStatLen)
	)
	encode := func(v interface{}, _ []byte, err error) interface{} {
		if err != nil {
			t.Fatalf("- %T %s", v, err)
		} else {
			t.Logf("+ %s", v)
		}
		return v
	}

	qid := encode(NewQid(qbuf, 1, 203, 0x83208)).(Qid)
	stat := encode(NewStat(statbuf, "georgia", "gopher", "gopher", "")).(Stat)
	stat.SetLength(492)
	stat.SetMode(02775)
	stat.SetQid(qid)

	encode(NewTversion(buf, 1<<12, "9P2000"))
	encode(NewRversion(buf, 1<<11, "9P2000"))
	encode(NewTauth(buf, 1, 1, "gopher", ""))
	encode(NewRauth(buf, 1, qid))
	encode(NewTattach(buf, 2, 2, 1, "gopher", ""))
	encode(NewRattach(buf, 2, qid))
	encode(NewRerror(buf, 0, "some error"))
	encode(NewTflush(buf, 3, 2))
	encode(NewRflush(buf, 3))
	encode(NewTwalk(buf, 4, 4, 4, "var", "log", "messages"))
	encode(NewRwalk(buf, 4, qid))
	encode(NewTopen(buf, 0, 1, 1))
	encode(NewRopen(buf, 0, qid, 300))
	encode(NewTcreate(buf, 1, 4, "frogs.txt", 0755, 3))
	encode(NewRcreate(buf, 1, qid, 1200))
	encode(NewTread(buf, 0, 32, 803280, 5308))
	encode(NewRread(buf, 0, 0, nil))
	encode(NewTwrite(buf, 1, 4, 10, 0, nil))
	encode(NewRwrite(buf, 1, 0))
	encode(NewTclunk(buf, 5, 4))
	encode(NewRclunk(buf, 5))
	encode(NewTremove(buf, 18, 9))
	encode(NewRremove(buf, 18))
	encode(NewTstat(buf, 6, 13))
	encode(NewRstat(buf, 6, stat))
	encode(NewTwstat(buf, 7, 3, stat))
	encode(NewRwstat(buf, 7))
}