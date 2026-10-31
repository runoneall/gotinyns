package server

import (
	"gotinyns/records"
	"strings"

	"github.com/miekg/dns"
)

func handler(w dns.ResponseWriter, r *dns.Msg) {
	select {

	case sem <- struct{}{}:
		defer func() { <-sem }()

	default:
		m := new(dns.Msg)
		m.SetRcode(r, dns.RcodeServerFailure)

		w.WriteMsg(m)
		return

	}

	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	if len(r.Question) == 0 {
		m.SetRcode(r, dns.RcodeFormatError)

		w.WriteMsg(m)
		return
	}

	for _, q := range r.Question {
		name := []byte(strings.TrimSuffix(q.Name, "."))
		recs := records.Find(name, q.Qtype)

		for _, rec := range recs {
			rr := toRR(q.Name, rec)

			if rr != nil {
				m.Answer = append(m.Answer, rr)
			}
		}
	}

	w.WriteMsg(m)
}
