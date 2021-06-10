package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Listing struct {
	ItemName                string `tsv:"item-name"`
	ItemDescription         string `tsv:"item-description"`
	ListingId               string `tsv:"listing-id"`
	SellerSku               string `tsv:"seller-sku"`
	Price                   string `tsv:"price"`
	Quantity                string `tsv:"quantity"`
	OpenDate                string `tsv:"open-date"`
	ImageUrl                string `tsv:"image-url"`
	ItemIsMarketplace       string `tsv:"item-is-marketplace"`
	ProductIdType           string `tsv:"product-id-type"`
	ZshopShippingFee        string `tsv:"zshop-shipping-fee"`
	ItemNote                string `tsv:"item-note"`
	ItemCondition           string `tsv:"item-condition"`
	ZshopCategory1          string `tsv:"zshop-category1"`
	ZshopBrowsePath         string `tsv:"zshop-browse-path"`
	ZshopStorefrontFeature  string `tsv:"zshop-storefront-feature"`
	Asin1                   string `tsv:"asin1"`
	Asin2                   string `tsv:"asin2"`
	Asin3                   string `tsv:"asin3"`
	WillShipInternationally string `tsv:"will-ship-internationally"`
	ExpeditedShipping       string `tsv:"expedited-shipping"`
	ZshopBoldface           string `tsv:"zshop-boldface"`
	ProductId               string `tsv:"product-id"`
	BidForFeaturedPlacement string `tsv:"bid-for-featured-placement"`
	AddDelete               string `tsv:"add-delete"`
	PendingQuantity         string `tsv:"pending-quantity"`
	FulfillmentChannel      string `tsv:"fulfillment-channel"`
	MerchantShippingGroup   string `tsv:"merchant-shipping-group"`
	Status                  string `tsv:"status"`
}

// func main() {
// 	var tsvFile = "fuck.tsv"
// 	file, err := os.Open(tsvFile)

// 	if err != nil {
// 		log.Printf("Cannot open text file: %s, err:[%v]", tsvFile, err)
// 	}
// 	defer file.Close()

// 	data := Listing{}
// 	parser, _ := tsv.NewParser(file, &data)

// 	// parser.Next()
// 	for {
// 		eof, err := parser.Next()
// 		if eof {
// 			return
// 		}
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(data)
// 	}

// 	// scanner := bufio.NewScanner(file)
// 	// scanner.Scan()
// 	// line := scanner.Text()
// 	// GenerateStruct(line, "Listing")
// }

func GenerateStruct(header string, typeName string) string {

	fileds := strings.Split(header, "\t")

	struct_str := "type " + typeName + " struct{\n"

	for _, value := range fileds {
		var nameArr = strings.Split(value, "-")
		// var nameWithBlank = strings.Trim(fmt.Sprint(nameArr), "[]")
		var nameTile = ""
		for _, v := range nameArr {
			nameTile += strFirstToUpper(v)
		}
		// var nameTile = strings.ToTitle(nameWithBlank)
		var name = strings.Trim(nameTile, " ")

		struct_str += "\t" + name + " string `tsv:\"" + value + "\"`\n"
	}
	struct_str += "}\n"
	fmt.Print(struct_str)
	handler, _ := os.OpenFile(typeName+".go", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer handler.Close()

	write := bufio.NewWriter(handler)
	write.WriteString(struct_str)
	write.WriteString("\n")
	write.Flush()
	handler.Close()
	return struct_str
}
func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
