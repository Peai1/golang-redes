package main

// go mod init servidor -> crear modulo
// go get github.com/miekg/dns -> instalar libreria

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

var records = map[string]map[uint16]string{
	"example.com.": {
		dns.TypeA:     "192.168.0.2",
		dns.TypeAAAA:  "::1",
		dns.TypeCNAME: "cname.example.com.",
		dns.TypeNS:    "ns.example.com.",
		dns.TypeMX:    "mx.example.com.",
	},
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)      // Crea un nuevo mensaje DNS.
	msg.SetReply(r)          // Configura el mensaje de respuesta con la consulta original.
	msg.Authoritative = true // Indica que el servidor DNS es autoritativo para el dominio.

	if rec, ok := records[r.Question[0].Name]; ok {
		switch r.Question[0].Qtype { // Verifica el tipo de consulta DNS.
		case dns.TypeA: // Ipv4
			// Agrega un registro de respuesta de tipo A con la dirección IP
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(rec[dns.TypeA]),
			})
		case dns.TypeAAAA: // Ipv6
			// Agrega un registro de respuesta de tipo AAAA con la dirección IP
			msg.Answer = append(msg.Answer, &dns.AAAA{
				Hdr:  dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 60},
				AAAA: net.ParseIP(rec[dns.TypeAAAA]),
			})
		case dns.TypeCNAME:
			//	Agrega un registro de respuesta de tipo CNAME con el nombre del dominio al que apunta.
			msg.Answer = append(msg.Answer, &dns.CNAME{
				Hdr:    dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 60},
				Target: rec[dns.TypeCNAME],
			})
		case dns.TypeNS:
			// Agrega un registro de respuesta de tipo NS con el nombre del servidor de nombres.
			msg.Answer = append(msg.Answer, &dns.NS{
				Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: 60},
				Ns:  rec[dns.TypeNS],
			})
		case dns.TypeMX:
			// Agrega un registro de respuesta de tipo MX con el nombre del servidor de correo.
			msg.Answer = append(msg.Answer, &dns.MX{
				Hdr: dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeMX, Class: dns.ClassINET, Ttl: 60},
				Mx:  rec[dns.TypeMX],
			})
		}
	} else {
		msg.SetRcode(r, dns.RcodeNameError)
	}
	w.WriteMsg(msg)
}

func main() {
	// Configura el servidor DNS.
	dns.HandleFunc(".", handleDNSRequest)
	server := &dns.Server{
		Addr: ":53",
		Net:  "udp",
	}
	log.Printf("Servidor DNS iniciado en puerto 53")
	err := server.ListenAndServe()
	defer server.Shutdown()

	if err != nil {
		log.Fatalf("Fallo al iniciar servidor: %s\n ", err.Error())
	}
}
