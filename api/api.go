package api

type Services struct {
}

func Start(port string) {
	services := Services{}
	r := routes(&services)
	server := newServer(port, r)

	server.Start()
}
