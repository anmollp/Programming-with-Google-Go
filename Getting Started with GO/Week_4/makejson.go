package main

import (
"fmt"
"encoding/json"
"bufio"
"os"
"strings"
)

func main() {
	person := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name:")
	name,err := reader.ReadString('\n')
	name = strings.Trim(name," \n")
	if err != nil {
		fmt.Printf("Could not read string")
	}
	fmt.Println("Enter address:")
	address, err := reader.ReadString('\n')
	address = strings.Trim(address, " \n")
	if err != nil {
                fmt.Printf("Could not read string")
        }
	
	person["name"] = name
	person["address"] = address

	j, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Error Marshaling")
	}
	fmt.Println(string(j))
}
