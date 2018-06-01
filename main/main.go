package main

import (
	"net/http"
	"fmt"
	"time"
)

//-----------------------------------
//  HTTP Site Status Checker
//
// Checks if urls can have successful
// HTTP connectivity
//
// Asks user to enter url
// Stores url in slice
// Ask if want to add more
// Asks for additional url if Yes
// Checks HTTP connection of all
// urls in slice if No
//
// Uses different Go routines for each check
// Waits 3 seconds between the checks
//------------------------------------



func main(){

	var links []string          			  // slice of strings for the site

	addUrlToList(&links)  					  // add first url to list

	ctn := "T"								  // continue program is true

	for ctn == "T"{							  // while true

		askAddMore()						  // ask if user wants to enter more

		ans := getCommand()					  // store response

		if *ans == "Y" || *ans == "y" {		  // Scenario 1: Y - add more url

			addUrlToList(&links)    		  // add additional  url to list

		} else if *ans == "N" || *ans == "n" {// Scenario 2: N - check url in slice

			c := make(chan string) 			  // create a new channel c for strings

			for _, link := range links {      // loop through slice
				go checkLink(link, c)
			}
			for l := range c { 				   // iterate though slice

				go func(link string) {	       // function literal to take the string
					time.Sleep(3 * time.Second)// delay child routine for 3 seconds
					checkLink(link, c)         // call checkLink on each link in slice using new Go routine
				}(l)
			}
			ctn = "F"						  // dont continue asking
		}
	}
}

/*
	askUrl

	ask user for url
 */
func askUrl(){

	fmt.Println("Enter the URL of the website for the status check (ie: google.com)")
	fmt.Println("Press Enter Key to submit URL")
	fmt.Print("http://")

	}


/**
 * getCommand
 *
 * Stores command entered by user
 *
 *return command : command url from user
 */
func getCommand()*string {

	var command string 		    // define variable to store url

	fmt.Scanln(&command) 		// stores character from user

	return &command 			// returns users url
}
/*
	checkLink

	Check if a link is responding to traffic

	Input: link - the link to check
	Input: c    - channel to communicate between routines
 */
func checkLink(link string, c chan string){

	_,err := http.Get(link) 						// checks if link is up

	if err != nil {									// if there is an error returned

		fmt.Println("Failed Connection: ", err)  // print error
		c <- link				    				// send message to channel
		return
	}

	fmt.Println("Successful Connection: ", link) // Else: print sucess to console
		c <- link									// send message to channel
		}
/**
 * addMore
 *
 * asks user if they want to add more
 *
 */
func askAddMore() {

	fmt.Println("Do you want to check another url? ( Y / N ) ")
	fmt.Println("Press Enter Key to submit")

}

/**
 * addUrlToList
 *
 * asks user for url and
 * adds url entered to slice
 *
 *Input: List of urls
 */
func addUrlToList( pL *[]string){

	askUrl() 		            // ask user for url

	pLink := getCommand()       // store pointer to user url

	link := "http://"+ *pLink   // convert user url to full url

	*pL = append(*pL, link)     // add to slice


}