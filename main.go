package main

import (
	"log"

	"github.com/jpillora/cloud-torrent/server"
	"github.com/jpillora/opts"
)

var VERSION = "1.1.3-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "NKR's Torrent Leecher",
		Port:       3000,
		ConfigPath: "cloud-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
