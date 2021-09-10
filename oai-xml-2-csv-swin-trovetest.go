package main

import (
	"encoding/xml"
        "encoding/csv"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)
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
                        Identifier string `xml:"identifier"`
                        Datestamp  string `xml:"datestamp"`
                } `xml:"header"`
                Metadata struct {
                        Text string `xml:",chardata"`
                        Dc   struct {
                                Text               string   `xml:",chardata"`
                                SchemaLocation     string   `xml:"schemaLocation,attr"`
                                OaiDc              string   `xml:"oai_dc,attr"`
                                Dc                 string   `xml:"dc,attr"`
                                Xsi                string   `xml:"xsi,attr"`
                                Identifier         string   `xml:"identifier"`
                                Title              string   `xml:"title"`
                                Creator            []string `xml:"creator"`
                                Description        string   `xml:"description"`
                                Date               string   `xml:"date"`
                                CoverageTemporal   string   `xml:"coverage.temporal"`
                                CoverageSpatial    string   `xml:"coverage.spatial"`
                                Type               string   `xml:"type"`
                                FormatMedium       string   `xml:"format.medium"`
                                Rights             string   `xml:"rights"`
                                RightsURI          string   `xml:"rights.uri"`
                                Subject            []string `xml:"subject"`
                                Relation           string   `xml:"relation"`
                                Source             string   `xml:"source"`
                                Publisher          string   `xml:"publisher"`
                                Location           string   `xml:"location"`
                                LocationAttachment []string `xml:"location.attachment"`
                                Thumbnail          string   `xml:"thumbnail"`
                        } `xml:"dc"`
                } `xml:"metadata"`
                About string `xml:"about"`
        } `xml:"record"`
}

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("./input-xml/swin-trovetest.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

        csvFile, err := os.Create("./output-csv/swin-trovetest.csv")
	if err != nil {
		fmt.Println(err)
	}
        defer csvFile.Close()
        writer := csv.NewWriter(csvFile)
        defer writer.Flush()


        // read our opened xmlFile as a byte array.
        byteValue, _ := ioutil.ReadAll(xmlFile)

        // we initialize our Users array
        var records Records
        var url string
        var record_file_id string
        url = "https://commons.swinburne.edu.au/file/"
        // we unmarshal our byteArray which contains our
        // xmlFiles content into 'users' which we defined above
        xml.Unmarshal(byteValue, &records)

        // we iterate through every user within our users array and
        // print out the user Type, their name, and their facebook url
        // as just an example
        var data = []string{"Header_Identifier", "Date_Latest", "Metadata_Identifier","Metadata_Identifier_File_ID", "URL"}
        if err := writer.Write(data) ; err != nil{
          fmt.Println(err)
        }
        for _, ch := range records.Record {
          /* date latest
          fmt.Println("Date edited " + ch.Metadata.Dc.Record.Issued.Text)
          for _, chii := range ch.Metadata.Dc.Record.Identifier {
            */
            fmt.Println("Header datestamp: " + ch.Header.Datestamp)
            for _, chii := range ch.Metadata.Dc.LocationAttachment {
              if strings.Contains(chii,".jpg") || strings.Contains(chii,".png") {
                record_file_id = chii

                fmt.Print(ch.Header.Identifier)
                fmt.Print(", ")
                fmt.Print(ch.Header.Datestamp)
                fmt.Print(", ")
                fmt.Print(record_file_id)
                fmt.Print(", ")
                fmt.Println(url + record_file_id)

                var data = []string{ch.Header.Identifier, ch.Header.Datestamp,ch.Metadata.Dc.Identifier,record_file_id,url+record_file_id}
                err := writer.Write(data)
                if err != nil {
                  fmt.Println(err)
                }
              } else {
                fmt.Println("Not image: " + chii)
              }
            }
          }
          //fmt.Print("\n")

        }
