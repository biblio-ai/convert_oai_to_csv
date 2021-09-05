package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Records was generated 2021-07-29 10:26:02 by justin on justin-HP-EliteBook-8570w.
type Records struct {
	XMLName xml.Name `xml:"Records"`
	Text    string   `xml:",chardata"`
	Xsi     string   `xml:"xsi,attr"`
	Record  []struct {
		Text   string `xml:",chardata"`
		Xmlns  string `xml:"xmlns,attr"`
		Header struct {
			Text       string `xml:",chardata"`
			Status     string `xml:"status,attr"`
			Identifier struct {
				Text string `xml:",chardata"` // tle:93535558-3bc8-46ca-8d...
			} `xml:"identifier"`
			Datestamp struct {
				Text string `xml:",chardata"` // 2017-09-19T05:43:51Z, 201...
			} `xml:"datestamp"`
			SetSpec struct {
				Text string `xml:",chardata"` // 7a95f13b-f4f1-4a12-a3fd-e...
			} `xml:"setSpec"`
		} `xml:"header"`
		Metadata struct {
			Text string `xml:",chardata"`
			Dc   struct {
				Text           string `xml:",chardata"`
				SchemaLocation string `xml:"schemaLocation,attr"`
				OaiDc          string `xml:"oai_dc,attr"`
				Dc             string `xml:"dc,attr"`
				Xsi            string `xml:"xsi,attr"`
				Identifier     struct {
					Text string `xml:",chardata"` // 93535558-3bc8-46ca-8d9d-7...
				} `xml:"identifier"`
				Title struct {
					Text string `xml:",chardata"` // "Ellerslie", Harcourt Str...
				} `xml:"title"`
				Date struct {
					Text string `xml:",chardata"`
				} `xml:"date"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
				TitleDetails struct {
					Text string `xml:",chardata"`
				} `xml:"title.details"`
				Publisher struct {
					Text string `xml:",chardata"` // :, :, :, :, :, :, :, :, :...
				} `xml:"publisher"`
				Edition struct {
					Text string `xml:",chardata"`
				} `xml:"edition"`
				Rights struct {
					Text string `xml:",chardata"` // Swinburne has been unable...
				} `xml:"rights"`
				Source struct {
					Text string `xml:",chardata"` // Subject Materials, Subjec...
				} `xml:"source"`
				Description struct {
					Text string `xml:",chardata"`
				} `xml:"description"`
				Location struct {
					Text string `xml:",chardata"` // https://commons.swinburne...
				} `xml:"location"`
				LocationAttachment struct {
					Text string `xml:",chardata"` // teoh-ellerslie_16_harcour...
				} `xml:"location.attachment"`
			} `xml:"dc"`
		} `xml:"metadata"`
		About struct {
			Text string `xml:",chardata"`
		} `xml:"about"`
	} `xml:"record"`
}

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("./input-xml/swin-all.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var records Records
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &records)

        var url string
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for _, ch := range records.Record {
		fmt.Println("Header ID: " + ch.Header.Identifier.Text)
		fmt.Println("Header datestampe: " + ch.Header.Datestamp.Text)
		fmt.Println("ID: " + ch.Metadata.Dc.Identifier.Text)
		fmt.Println("Location: " + ch.Metadata.Dc.Location.Text)
		fmt.Println("Attachment: " + ch.Metadata.Dc.LocationAttachment.Text)
		url = ch.Metadata.Dc.Location.Text + ch.Metadata.Dc.LocationAttachment.Text
		fmt.Println("URL: " + url)
	}

}
