package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://frankfurt.kapeli.com/feeds/AngularJS.xml")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	type Version struct {
		Name string `xml:"name"`
	}

	type Entry struct {
		Version       string    `xml:"version"`
		IOSVersion    string    `xml:"ios_version"`
		Url           []string  `xml:"url"`
		OtherVersions []Version `xml:"other-versions>version"`
	}

	v := Entry{
		Version:    "",
		IOSVersion: "",
		Url:        nil,
	}

	// 	<entry>
	//     <version>1.8.2</version>
	//     <ios_version>1</ios_version>
	//     <url>http://sanfrancisco.kapeli.com/feeds/AngularJS.tgz</url>
	//     <url>http://london.kapeli.com/feeds/AngularJS.tgz</url>
	//     <url>http://newyork.kapeli.com/feeds/AngularJS.tgz</url>
	//     <url>http://tokyo.kapeli.com/feeds/AngularJS.tgz</url>
	//     <url>http://frankfurt.kapeli.com/feeds/AngularJS.tgz</url>
	//     <other-versions>
	//         <version><name>1.8.2</name></version>
	//         <version><name>1.7.9</name></version>
	//         <version><name>1.6.10</name></version>
	//         <version><name>1.5.8</name></version>
	//         <version><name>1.4.9</name></version>
	//         <version><name>1.3.15</name></version>
	//         <version><name>1.2.26</name></version>
	//         <version><name>1.0.8</name></version>
	//     </other-versions>
	// </entry>

	if xml.Unmarshal(data, &v); err != nil {
		fmt.Println("XML Unmarshl Error", err)
		os.Exit(1)
	}

	fmt.Printf("Version:\t%s\n", v.Version)
	fmt.Printf("IOS Version:\t%s\n", v.IOSVersion)

	fmt.Println("Urls:")
	for i := range v.Url {
		fmt.Printf("\t> %s\n", v.Url[i])
	}

	fmt.Println("Other Versions:")
	for i := 0; i < len(v.OtherVersions); i++ {
		fmt.Printf("\t> %s\n", v.OtherVersions[i].Name)
	}
}
