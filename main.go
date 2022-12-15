package main

import (
	harvester "network/Harvester"
)

func main() {

	harvester := harvester.NewHarvester("https://google.com", "GET")
	harvester.Hijack()
}
