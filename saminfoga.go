package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type PhoneInfo struct {
	CountryCode    string `json:"country_code"`
	NationalNumber string `json:"national_number"`
	Carrier        string `json:"carrier"`
	LineType       string `json:"line_type"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <phone number>")
		return
	}

	phoneNumber := os.Args[1]

	if !isValidPhoneNumber(phoneNumber) {
		fmt.Println("Invalid phone number!")
		return
	}

	fmt.Printf("Phone number: %s\n", phoneNumber)

	// Reverse phone number lookup
	if phoneInfo := reverseLookup(phoneNumber); phoneInfo != nil {
		fmt.Printf("Country code: %s\n", phoneInfo.CountryCode)
		fmt.Printf("National number: %s\n", phoneInfo.NationalNumber)
		fmt.Printf("Carrier: %s\n", phoneInfo.Carrier)
		fmt.Printf("Line type: %s\n", phoneInfo.LineType)
	}

	// Phone number validation
	if result := validatePhoneNumber(phoneNumber); result["valid"].(bool) {
		fmt.Println("Phone number is valid!")
	} else {
		fmt.Println("Phone number is invalid!")
	}

	// Spam call detection
	if isSpam(phoneNumber) {
		fmt.Println("This is a spam call!")
	} else {
		fmt.Println("This is not a spam call.")
	}

	// IP address lookup
	if ip, err := getIPAddress(phoneNumber); err == nil {
		fmt.Printf("IP address: %s\n", ip)

		// Geolocation lookup
		if geoInfo, err := getGeolocation(ip); err == nil {
			fmt.Printf("Geolocation: %s\n", geoInfo)
		}
	}

	// Social media lookup
	if socialMediaInfo := getSocialMediaInfo(phoneNumber); len(socialMediaInfo) > 0 {
		fmt.Println("Social media accounts:")
		for platform, account := range socialMediaInfo {
			fmt.Printf("- %s: %s\n", platform, account)
		}
	} else {
		fmt.Println("No social media accounts found.")
	}
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Check if the phone number matches the E.164 format
	valid, err := regexp.MatchString(`^\+?[1-9]\d{1,14}$`, phoneNumber)
	return err == nil && valid
}

func reverseLookup(phoneNumber string) *PhoneInfo {
	fmt.Println("Performing reverse phone number lookup...")
	resp, err := http.Get(fmt.Sprintf("https://api.telnyx.com/anonymous/v2/number_lookup/%s", phoneNumber))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var phoneInfo PhoneInfo
	err = json.Unmarshal(body, &phoneInfo)
	if err != nil {
		return nil
	}
	return &phoneInfo
}

func validatePhoneNumber(phoneNumber string) map[string]interface{} {
	fmt.Println("Performing phone number validation...")
	resp, err := http.Get(fmt.Sprintf("https://numvalidate.com/api/validate?number=%s", phoneNumber))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
func getTwitterUsername(html string) string {
	re := regexp.MustCompile(`https://twitter.com/([^/]+)/`)
	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func getInstagramUsername(html string) string {
	re := regexp.MustCompile(`https://www.instagram.com/([^/]+)/`)
	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func getFacebookUsername(html string) string {
	re := regexp.MustCompile(`https://www.facebook.com/([^/]+)/`)
	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func getSocialMediaInfo(phoneNumber string) map[string]string {
	fmt.Println("Performing social media lookup...")
	socialMediaInfo := make(map[string]string)

	// Twitter lookup
	html := getHTML(fmt.Sprintf("https://www.google.com/search?q=%s+twitter", phoneNumber))
	if username := getTwitterUsername(html); username != "" {
		socialMediaInfo["Twitter"] = username
	}

	// Instagram lookup
	html = getHTML(fmt.Sprintf("https://www.google.com/search?q=%s+instagram", phoneNumber))
	if username := getInstagramUsername(html); username != "" {
		socialMediaInfo["Instagram"] = username
	}

	// Facebook lookup
	html = getHTML(fmt.Sprintf("https://www.google.com/search?q=%s+facebook", phoneNumber))
	if username := getFacebookUsername(html); username != "" {
		socialMediaInfo["Facebook"] = username
	}

	return socialMediaInfo
}
