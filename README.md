# saminfoga 🕵️‍♂️

saminfoga is a command-line tool for gathering information about phone numbers. It uses various APIs to perform reverse phone number lookups, phone number validation, spam call detection, IP address lookup, geolocation lookup, and social media lookup.

## Usage

```
$ saminfoga <phone number>
```

Replace `<phone number>` with the phone number you want to lookup, in the E.164 format (+ followed by country code and phone number). For example:

```
$ saminfoga +14155552671
```

## Installation

1. Make sure you have Go installed. If not, [install Go](https://golang.org/doc/install).
2. Clone the repository: `git clone https://github.com/yourusername/saminfoga.git`
3. Navigate to the project directory: `cd saminfoga`
4. Build the executable: `go build`
5. Run the tool: `./saminfoga <phone number>`

## API Keys

Some of the APIs used by saminfoga require API keys. You can obtain free API keys from the following websites:

- Telnyx (for reverse phone number lookup): https://developers.telnyx.com/docs/v2/anonymous-phone-number-lookup/lookup-phone-numbers
- NumVerify (for phone number validation): https://numverify.com/dashboard
- IPGeolocation (for IP address lookup and geolocation lookup): https://ipgeolocation.io/signup.html

After obtaining the API keys, add them to the `config.json` file in the following format:

```json
{
  "telnyx_api_key": "your_telnyx_api_key",
  "numverify_api_key": "your_numverify_api_key",
  "ipgeolocation_api_key": "your_ipgeolocation_api_key"
}
```

## Credits

This tool was inspired by the Infoga tool (https://github.com/m4ll0k/Infoga) and uses various APIs, libraries, and regular expressions to gather information about phone numbers.

