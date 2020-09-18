package main

import "fmt"

// func main() {
// 	hosts := findHosts()
// 	hosts.print()
// 	fmt.Println(hosts.getHost(1))

// 	fmt.Println(hosts.toString())

// 	hosts.saveToFile("test")

// 	hosts2 := getHostsFromFile("test")
// 	fmt.Println("Hosts2:")
// 	hosts2.print()
// }

func main() {
	host := host{name: "host1", mac: "aa:aa:aa:aa:aa:", ipv4: "10.0.0.1"}
	fmt.Println(host)
	fmt.Printf("%+v\n", host)
	host.mac = "bb:bb:bb:bb:bb"
	fmt.Printf("%+v\n", host)
	host.print()
	host.name = "hostx"
	host.print()
	host.updateName("hosty")
	host.print()

	emptyMap := make(map[string]string)
	emptyMap["white"] = "sdhjsdfljksdhf"
	delete(emptyMap, "white")

	fmt.Println(emptyMap)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#bf745",
	}

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hexcode for", colour, "is", hex)
	}
}
