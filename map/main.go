package main

import "fmt"
// map is reference type, struct is value type
func main()  {
	colors := map[string]string{"red": "#fff234", "white": "#ffffff", "green": "#d3d4d5"}
	colors["yellow"] = "#ggeee"
	delete(colors, "yellow")
	printMap(colors)
}

func printMap(c map[string]string)  {
	for key, value :=range c{
		fmt.Println("The color of ", key, "of hex code is", value)
	}
}