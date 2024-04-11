package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func resolve(domain string, qtype uint16) {
	m := new(dns.Msg)                      // Crea un nuevo mensaje DNS.
	m.SetQuestion(dns.Fqdn(domain), qtype) // Configura la consulta DNS con el dominio y tipo de registro
	//m.RecursionDesired = true

	// Realiza la consulta a tu servidor DNS local.
	// Asegúrate de reemplazar "127.0.0.1:53" con la dirección y puerto de tu servidor DNS si es diferente.
	c := new(dns.Client)
	in, _, err := c.Exchange(m, "127.0.0.1:53")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ans := range in.Answer {
		fmt.Println(ans)
	}
}

func main() {
	resolve("example.com", dns.TypeA)
}
