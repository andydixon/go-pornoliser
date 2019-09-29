package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	swear := 50
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(proseConverter(nameConverter(text, swear), swear))
}

func nameConverter(text string, swearibility int) string {

	names := []string{"Big Cock", "Fuck me for a Buck", "Dirk Diggler", "Big Dick", "Give it to me", "Cock Sucker", "Up the Arse", 
	"Cockboy", "Omar Pussy", "Fat Ass", "Superdick", "Mistress Shiva", "Bite Me", "Mistress Anal", "Long Finger", "Plugin", "Anal", 
	"Ass-stitcher", "Jar Jar", "Bust-a-Cunt", "Buzzwordbaby", "Sniff-my-Ass", "Dripper Dick", "Butplug", "The-Champ", "Fill me up",
	 "Suck my tits dry", "Ball Buster", "Dildo", "Motherfucker", "Dickwad", "Fuckface", "Jerkoff", "Bitch", "Bastard", "Asshole", 
	 "Hard-on", "Pimp Mastah", "Son of a whore", "Muffdiver", "Rugmuncher", "Admiral Browning", "Afterburner", "Airing the Orchid", 
	 "Aphrodite's Evostick", "Ballbuffer", "Muffmuncher", "Cuntcleaner", "Mount Thrushmore", "Scrotscrubber", "Muffminer", "Anusapple", 
	 "Mouth-full-o'-cock", "Nobgoblin", "Fannyfarmer", "Spunksupper", "Clitcollector", "Bumbanger", "Assrush", "Bonebagger", "Saggysack", 
	 "Cumbucket", "King Choad"}
	match := 0
	returnText := ""
	textArray := strings.Fields(text)

	for _, word := range textArray {
		if word == strings.Title(word) {
			match = match + 1
		} else {
			match = 0
		}

		if match == 2 {
			rand.Seed(time.Now().Unix())
			if rand.Intn(len(names)) <= swearibility {
				returnText = returnText + " \"" + names[rand.Intn(len(names))] + "\""
			}
			match = 0
		}
		returnText = returnText + " " + word
	}

	return strings.TrimSpace(returnText)
}

func proseConverter(text string, swearibility int) string {

	dict := []string{
		"fuck", "smooch", "smack", "peck", "unclefuck", "spank", "motherfuck", "deep throat", "ballbust", "spew", "dripp", 
		"thrust", "cocksuck", "fistfuck", "fist", "suck", "squirt", "wank", "bang", "brown", "cuntlick", "cuntlapp", "felch", 
		"fingerfuck", "gangbang", "plow", "raunch", "screw", "sex fight", "titty fuck", "enter", "raid", "jerk", "finger", 
		"shaft", "blow", "lick", "asslick", "fuck", "assfuck", "wad pull", "muff sniff", "aardvark", "ball", "barf", "charver", 
		"cream", "fart", "fomp", "gamahuche"}
	returnText := ""
	textArray := strings.Fields(text)
	suffix := ""

	for _, word := range textArray {
		if strings.HasSuffix(word, "ing") || strings.HasSuffix(word, "ed") || strings.HasSuffix(word, "s") {
			rand.Seed(time.Now().Unix())
			if rand.Intn(len(dict)) <= swearibility {

				if strings.HasSuffix(word, "ing") {
					suffix = "ing"
				}
				if strings.HasSuffix(word, "ed") {
					suffix = "ed"
				}
				if strings.HasSuffix(word, "s") {
					suffix = "s"
				}

				word = dict[rand.Intn(len(dict))] + suffix
				suffix = ""
			}
		}
		returnText = returnText + " " + word
	}
	return strings.TrimSpace(returnText)
}
