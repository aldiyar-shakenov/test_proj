package main

import "fmt"

func person_is_seller(name string) bool {
	return name[len(name)-1:] == "m"
}

func person_was_searched(person string, arr []string) bool {
	for _, names := range arr {
		if names == person {
			return true
		}
	}
	return false
}

func searchSeller(name string, nameGraph map[string][]string) {
	//function prints "Person is a Mango Seller!"
	newQueue := nameGraph[name]
	searchedList := []string{}
	//if the newQueue is non-empty => search
	for len(newQueue) != 0 {
		//person is the leftmost in newQueue
		//newQueue is shortened by one
		person := newQueue[0]
		newQueue = newQueue[1:]
		fmt.Println("checking :", person)
		// see if person is a seller and not searched
		if !person_was_searched(person, searchedList) && person_is_seller(person) {
			fmt.Println(person, " is a Mango Seller!")
			return
		} else {
			//add person to searched
			//add more nodes  to newQueue
			searchedList = append(searchedList, person)
			newQueue = append(newQueue, nameGraph[person]...)

		}
	}

}

func main() {
	stringGraph := make(map[string][]string)
	stringGraph["you"] = []string{"alice", "bob", "claire"}
	stringGraph["bob"] = []string{"anuj", "peggy"}
	stringGraph["alice"] = []string{"peggy"}
	stringGraph["claire"] = []string{"thom", "jonny"}
	stringGraph["anuj"] = []string{""}
	stringGraph["peggy"] = []string{""}
	stringGraph["thom"] = []string{""}
	stringGraph["jonny"] = []string{""}

	//fmt.Println()
	searchSeller("you", stringGraph)
	//fmt.Println(person_is_seller("thom"))
	//shortArr(stringGraph["you"])

}
