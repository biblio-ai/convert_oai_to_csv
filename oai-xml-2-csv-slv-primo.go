package main

import (
	"encoding/xml"
        "encoding/csv"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

// Response was generated 2021-07-28 22:34:49 by justin on justin-HP-EliteBook-8570w.
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
				Text string `xml:",chardata"` // oai:rosetta.slv.vic.gov.a...
			} `xml:"identifier"`
			Datestamp struct {
				Text string `xml:",chardata"` // 2021-01-21T10:57:19Z, 202...
			} `xml:"datestamp"`
			SetSpec []struct {
				Text string `xml:",chardata"` // Primo, Primo, Primo, Prim...
			} `xml:"setSpec"`
		} `xml:"header"`
		Metadata struct {
			Text string `xml:",chardata"`
			Dc   struct {
				Text           string `xml:",chardata"`
				SchemaLocation string `xml:"schemaLocation,attr"`
				OaiDc          string `xml:"oai_dc,attr"`
				Slvterms       string `xml:"slvterms,attr"`
				Dnx            string `xml:"dnx,attr"`
				Dc             string `xml:"dc,attr"`
				Mets           string `xml:"mets,attr"`
				Dcterms        string `xml:"dcterms,attr"`
				Xsi            string `xml:"xsi,attr"`
				Record         struct {
					Text   string `xml:",chardata"`
					Xmlns  string `xml:"xmlns,attr"`
					Issued struct {
						Text    string `xml:",chardata"` // 2020-11-20, 2020-12-02, 2...
						Content string `xml:"content,attr"`
					} `xml:"issued"`
					Identifier []struct {
						Text    string `xml:",chardata"` // 1684770, ONE, FL20641823,...
						Type    string `xml:"type,attr"`
						RepCode string `xml:"repCode,attr"`
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
						Mods    string `xml:"mods,attr"`
					} `xml:"identifier"`
					Title []struct {
						Text    string `xml:",chardata"` // Black Spur. Vic, Puffing ...
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
						Content string `xml:"content,attr"`
						Xmlns   string `xml:"xmlns,attr"`
						Xlink   string `xml:"xlink,attr"`
					} `xml:"title"`
					Description []struct {
						Text    string `xml:",chardata"` // Man walking along dirt ro...
						Content string `xml:"content,attr"`
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
						Xmlns   string `xml:"xmlns,attr"`
						Xlink   string `xml:"xlink,attr"`
					} `xml:"description"`
					FileName []struct {
						Text  string `xml:",chardata"` // a05110, gq000422, pi02591...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xmlns string `xml:"xmlns,attr"`
						Xlink string `xml:"xlink,attr"`
					} `xml:"fileName"`
					Created []struct {
						Text  string `xml:",chardata"` // 1903/1909, 1967-01/1967-0...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xmlns string `xml:"xmlns,attr"`
						Xlink string `xml:"xlink,attr"`
					} `xml:"created"`
					Format []struct {
						Text  string `xml:",chardata"` // TIFF, TIFF, TIFF, TIFF, T...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xmlns string `xml:"xmlns,attr"`
						Xlink string `xml:"xlink,attr"`
					} `xml:"format"`
					Type []struct {
						Text  string `xml:",chardata"` // StillImage, StillImage, S...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xmlns string `xml:"xmlns,attr"`
						Xlink string `xml:"xlink,attr"`
					} `xml:"type"`
					Genre []struct {
						Text     string `xml:",chardata"` // Prints, Photographs, Phot...
						Mods     string `xml:"mods,attr"`
						Ns       string `xml:"ns,attr"`
						Xmlns    string `xml:"xmlns,attr"`
						Xlink    string `xml:"xlink,attr"`
						Slvterms string `xml:"slvterms,attr"`
					} `xml:"genre"`
					Rights []struct {
						Text  string `xml:",chardata"` // This work is out of copyr...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xmlns string `xml:"xmlns,attr"`
						Xlink string `xml:"xlink,attr"`
					} `xml:"rights"`
					AccessRights []struct {
						Text  string `xml:",chardata"` // No copyright restrictions...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xmlns string `xml:"xmlns,attr"`
						Xlink string `xml:"xlink,attr"`
					} `xml:"accessRights"`
					DigitalPages []struct {
						Text  string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xlink string `xml:"xlink,attr"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"digitalPages"`
					Accession []struct {
						Text  string `xml:",chardata"` // H37831, H83.156/422, H201...
						Mods  string `xml:"mods,attr"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"accession"`
					DigitalCollection []struct {
						Text  string `xml:",chardata"` // Pictoria, Smith trams and...
						Mods  string `xml:"mods,attr"`
						Ns    string `xml:"ns,attr"`
						Xlink string `xml:"xlink,attr"`
					} `xml:"digitalCollection"`
					IsPartOf []struct {
						Text    string `xml:",chardata"` // 1415415, Smith Trams and ...
						Type    string `xml:"type,attr"`
						Content string `xml:"content,attr"`
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
						Xmlns   string `xml:"xmlns,attr"`
						Xlink   string `xml:"xlink,attr"`
					} `xml:"isPartOf"`
					Creator []struct {
						Text    string `xml:",chardata"` // Smith A E active 1966-196...
						Content string `xml:"content,attr"`
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
						Xmlns   string `xml:"xmlns,attr"`
						Xlink   string `xml:"xlink,attr"`
					} `xml:"creator"`
					Extent []struct {
						Text    string `xml:",chardata"` // 3 negatives : flexible ba...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"extent"`
					TableOfContents []struct {
						Text    string `xml:",chardata"` // Close-up view of the NA c...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"tableOfContents"`
					Abstract []struct {
						Text    string `xml:",chardata"` // Three views taken from th...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
						Mods    string `xml:"mods,attr"`
					} `xml:"abstract"`
					Subject []struct {
						Text    string `xml:",chardata"` // Puffing Billy Preservatio...
						Content string `xml:"content,attr"`
						Scheme  string `xml:"scheme,attr"`
						Type    string `xml:"type,attr"`
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"subject"`
					IndexTerms []struct {
						Text    string `xml:",chardata"` // Close-up view of the NA c...
						Content string `xml:"content,attr"`
						Scheme  string `xml:"scheme,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"indexTerms"`
					Contributor []struct {
						Text    string `xml:",chardata"` // A. & K. Henderson archite...
						Content string `xml:"content,attr"`
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
						Xmlns   string `xml:"xmlns,attr"`
					} `xml:"contributor"`
					Alternative []struct {
						Text    string `xml:",chardata"` // Japanese Zero at Buna air...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"alternative"`
					CreatedOriginalResource []struct {
						Text    string `xml:",chardata"` // [ca. 1914], 1907-1914., 1...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"created.originalResource"`
					Publisher []struct {
						Text    string `xml:",chardata"` // Littleton, N.H. : Kilburn...
						Content string `xml:"content,attr"`
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"publisher"`
					Provenance struct {
						Text    string `xml:",chardata"` // The files which now compr...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"provenance"`
					Spatial []struct {
						Text    string `xml:",chardata"` // Brunswick, Victoria, Lane...
						Mods    string `xml:"mods,attr"`
						Ns      string `xml:"ns,attr"`
						Content string `xml:"content,attr"`
						Xmlns   string `xml:"xmlns,attr"`
					} `xml:"spatial"`
					ParentID struct {
						Text string `xml:",chardata"` // COLES014, COLES014, COLES...
						Mods string `xml:"mods,attr"`
						Ns   string `xml:"ns,attr"`
					} `xml:"parentID"`
					HasPart struct {
						Text    string `xml:",chardata"` // Panoramic view of the Cit...
						Content string `xml:"content,attr"`
					} `xml:"hasPart"`
					DigitalCollection2 []struct {
						Text string `xml:",chardata"` // Pictoria, Postcards, Pict...
						Mods string `xml:"mods,attr"`
					} `xml:"digitalCollection2"`
					IssuedOriginalResource []struct {
						Text    string `xml:",chardata"` // [between 1900 and 1910?],...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"issued.originalResource"`
					Temporal struct {
						Text    string `xml:",chardata"` // d  1840 Aug. d 1847 June....
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"temporal"`
					Replaces []struct {
						Text    string `xml:",chardata"` // Transactions and proceedi...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"replaces"`
					IsReplacedBy []struct {
						Text    string `xml:",chardata"` // Transactions of the Royal...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"isReplacedBy"`
					References []struct {
						Text    string `xml:",chardata"` // http://handle.slv.vic.gov...
						Type    string `xml:"type,attr"`
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"references"`
					DisplayOrder []struct {
						Text    string `xml:",chardata"` // vp1273, vp0655, vp0547, v...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"displayOrder"`
					IsReferencedBy []struct {
						Text  string `xml:",chardata"` // Ferguson 14519, http://ha...
						Mods  string `xml:"mods,attr"`
						Xlink string `xml:"xlink,attr"`
						Type  string `xml:"type,attr"`
					} `xml:"isReferencedBy"`
					AccrualPeriodicity struct {
						Text    string `xml:",chardata"` // Annual, To be published o...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"accrualPeriodicity"`
					IsFormatOf struct {
						Text    string `xml:",chardata"` // English oil colour painte...
						Content string `xml:"content,attr"`
						Ns      string `xml:"ns,attr"`
					} `xml:"isFormatOf"`
					LoadCode []struct {
						Text string `xml:",chardata"` // PDF, PDF, PDF, PDF, PDF, ...
						Mods string `xml:"mods,attr"`
						Ns   string `xml:"ns,attr"`
					} `xml:"loadCode"`
					Date struct {
						Text string `xml:",chardata"` // 1872, 1869/1874, 1860
						Mods string `xml:"mods,attr"`
					} `xml:"date"`
					SourceID struct {
						Text  string `xml:",chardata"` // 972836, 972836, 972836, 9...
						Xlink string `xml:"xlink,attr"`
					} `xml:"sourceID"`
				} `xml:"record"`
			} `xml:"dc"`
		} `xml:"metadata"`
		About struct {
			Text string `xml:",chardata"`
		} `xml:"about"`
	} `xml:"record"`
}

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("./input-xml/slv-primo.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

        csvFile, err := os.Create("./output-csv/slv-primo.csv")
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
        var blank string
        var record_accession_id string
        var record_handle_id string
        var record_cms_id string
        var record_file_id string
        url = "https://rosetta.slv.vic.gov.au/delivery/DeliveryManagerServlet?dps_func=stream&dps_pid="
        // we unmarshal our byteArray which contains our
        // xmlFiles content into 'users' which we defined above
        xml.Unmarshal(byteValue, &records)

        // we iterate through every user within our users array and
        // print out the user Type, their name, and their facebook url
        // as just an example
        var data = []string{"Header_Identifier", "Date_Latest", "Metadata_Identifier","Metadata_Identifier_Handle_ID", "Metadata_Identifier_CMS_ID", "Metadata_Identifier_Accession_ID","Metadata_Identifier_File_ID", "URL"}
        if err := writer.Write(data) ; err != nil{
          fmt.Println(err)
        }
        for _, ch := range records.Record {
          /* date latest
          fmt.Println("Header datestamp: " + ch.Header.Datestamp.Text)
          fmt.Print(", "
          fmt.Println("Date edited " + ch.Metadata.Dc.Record.Issued.Text)
          for _, chii := range ch.Metadata.Dc.Record.Identifier {
            */
            record_id := strings.Split(ch.Header.Identifier.Text,":")
            fmt.Print(", ")
            fmt.Print(ch.Header.Datestamp.Text)
            fmt.Print(", ")
            if len(ch.Metadata.Dc.Record.Accession) > 0 {
              record_accession_id = ch.Metadata.Dc.Record.Accession[0].Text
            } else {
              record_accession_id = blank
            }
            fmt.Print(", ")
            for _, chii := range ch.Metadata.Dc.Record.Identifier {
              if chii.Type == "Handle" {
                //	fmt.Println("URL: " + url + chii.Text)
                record_handle_id = chii.Text
              }
              if chii.Type == "voyagerBibId" {
                //	fmt.Println("URL: " + url + chii.Text)
                record_cms_id = chii.Text
              }
              if chii.Type == "fileID" {
                record_file_id = chii.Text
              }
            }
            fmt.Print(record_handle_id)
            fmt.Print(", ")
            fmt.Print(record_cms_id)
            fmt.Print(", ")
            fmt.Println(url + record_file_id)
            var data = []string{ch.Header.Identifier.Text, ch.Header.Datestamp.Text, record_id[2], record_handle_id, record_cms_id, record_accession_id,record_file_id,url+record_file_id}
            err := writer.Write(data)
            if err != nil {
              fmt.Println(err)
            }
          }
          //fmt.Print("\n")

        }
