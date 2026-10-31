package server

import (
	"gotinyns/records"
	"net"
	"strconv"
	"strings"

	"github.com/miekg/dns"
)

func toRR(qname string, rec records.Record) dns.RR {
	owner := string(rec.Name)
	if owner == "" {
		owner = qname
	}

	if !strings.HasSuffix(owner, ".") {
		owner += "."
	}

	hdr := dns.RR_Header{
		Name:   owner,
		Rrtype: rec.Type,
		Class:  dns.ClassINET,
		Ttl:    rec.TTL,
	}

	switch rec.Type {

	case dns.TypeA:
		if len(rec.Value) != 4 {
			return nil
		}

		return &dns.A{Hdr: hdr, A: net.IP(rec.Value)}

	case dns.TypeAAAA:
		if len(rec.Value) != 16 {
			return nil
		}

		return &dns.AAAA{Hdr: hdr, AAAA: net.IP(rec.Value)}

	case dns.TypeCNAME:
		return &dns.CNAME{Hdr: hdr, Target: ensureDot(string(rec.Value))}

	case dns.TypeNS:
		return &dns.NS{Hdr: hdr, Ns: ensureDot(string(rec.Value))}

	case dns.TypePTR:
		return &dns.PTR{Hdr: hdr, Ptr: ensureDot(string(rec.Value))}

	case dns.TypeMX:
		parts := strings.Fields(string(rec.Value))
		if len(parts) != 2 {
			return nil
		}

		pref, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil
		}

		return &dns.MX{Hdr: hdr, Preference: uint16(pref), Mx: ensureDot(parts[1])}

	case dns.TypeTXT:
		return &dns.TXT{Hdr: hdr, Txt: []string{string(rec.Value)}}

	default:
		return nil

	}
}

func ensureDot(s string) string {
	if strings.HasSuffix(s, ".") {
		return s
	}

	return s + "."
}
