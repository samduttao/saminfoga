package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Struct to hold the API keys
type Config struct {
	NumVerifyApiKey      string `json:"numverify_api_key"`
	TwilioAccountSid     string `json:"twilio_account_sid"`
	TwilioAuthToken      string `json:"twilio_auth_token"`
	TwilioLookupApiUrl   string `json:"twilio_lookup_api_url"`
	IPGeolocationApiKey  string `json:"ipgeolocation_api_key"`
}

func main() {
	// Parse command-line arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: saminfoga <phone number>")
		os.Exit(1)
	}
	phoneNumber := args[0]

	// Load API keys from config.json file
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		os.Exit(1)
	}

	// Validate phone number using NumVerify API
	fmt.Println("Validating phone number...")
	numVerifyResp, err := http.Get(fmt.Sprintf("http://apilayer.net/api/validate?access_key=%s&number=%s&country_code=&format=1", config.NumVerifyApiKey, phoneNumber))
	if err != nil {
		fmt.Println("Error calling NumVerify API:", err)
		os.Exit(1)
	}
	defer numVerifyResp.Body.Close()
	numVerifyBody, err := ioutil.ReadAll(numVerifyResp.Body)
	if err != nil {
		fmt.Println("Error reading NumVerify API response:", err)
		os.Exit(1)
	}
	var numVerifyResult map[string]interface{}
	err = json.Unmarshal(numVerifyBody, &numVerifyResult)
	if err != nil {
		fmt.Println("Error parsing NumVerify API response:", err)
		os.Exit(1)
	}
	if numVerifyResult["valid"].(bool) == false {
		fmt.Println("Invalid phone number.")
		os.Exit(1)
	}

	// Get carrier information using Twilio Lookup API
	fmt.Println("Getting carrier information...")
	twilioClient := &http.Client{}
	twilioReq, err := http.NewRequest("GET", config.TwilioLookupApiUrl, nil)
	if err != nil {
		fmt.Println("Error creating Twilio Lookup API request:", err)
		os.Exit(1)
	}
	twilioReq.SetBasicAuth(config.TwilioAccountSid, config.TwilioAuthToken)
	q := twilioReq.URL.Query()
	q.Add("PhoneNumber", phoneNumber)
	twilioReq.URL.RawQuery = q.Encode()
	twilioResp, err := twilioClient.Do(twilioReq)
	if err != nil {
		fmt.Println("Error calling Twilio Lookup API:", err)
		os.Exit(1)
	}
	defer twilioResp.Body.Close()
	twilioBody, err := ioutil.ReadAll(twilioResp.Body)
	if err != nil {
		fmt.Println("Error reading Twilio Lookup API response:", err)
		os.Exit(1)
	}
	var twilioResult map[string]interface{}
	err = json.Unmarshal(twilioBody, &twilioResult)
	if err != nil {
		fmt.Println("Error parsing Twilio Lookup API response:", err)
		os.Exit(1)
	}
	if twilioResult["carrier"] != nil {
result["carrier"] = twilioResult["carrier"].(map[string]interface{})["name"]
} else {
result["carrier"] = "N/A"
}
	func main() {
	// Parse command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: saminfoga <phone number>")
		fmt.Println("Example: saminfoga +14155552671")
		return
	}
	phoneNumber := os.Args[1]

	// Load API keys from config file
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Perform phone number validation
	fmt.Println("Performing phone number validation...")
	valid, err := validatePhoneNumber(phoneNumber, config.NumverifyAPIKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !valid {
		fmt.Println("Invalid phone number")
		return
	}

	// Perform spam call check
	fmt.Println("Performing spam call check...")
	spam, err := checkSpamCall(phoneNumber, config.TwilioAPIURL, config.TwilioAccountSID, config.TwilioAuthToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	if spam {
		fmt.Println("Spam call detected")
		return
	}

	// Perform reverse phone number lookup
	fmt.Println("Performing reverse phone number lookup...")
	name, address, err := reverseLookup(phoneNumber, config.TelnyxAPIURL, config.TelnyxAPIKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Address: %s\n", address)
	}

	// Perform IP address and geolocation lookup
	fmt.Println("Performing IP address and geolocation lookup...")
	ip, location, err := lookupIPAndLocation(phoneNumber, config.IPGeolocationAPIURL, config.IPGeolocationAPIKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("IP address: %s\n", ip)
		fmt.Printf("Location: %s\n", location)
	}

	// Perform social media lookup
	fmt.Println("Performing social media lookup...")
	html, err := socialMediaLookup(phoneNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		username := getTwitterUsername(html)
		if username != "" {
			fmt.Printf("Twitter username: %s\n", username)
		} else {
			fmt.Println("Twitter username not found")
		}
	}
}

