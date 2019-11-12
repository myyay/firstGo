package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getLine(l int) string {
	prefix := "P87638888888888"
	return prefix + strconv.Itoa(l) + "|2|20190903163625|SUCCESS|0.01^0.01"

}

func main() {

	fileName := "D:\\export\\data\\channel\\tradeStatus\\20190904\\fedz\\CAP_TEMP-911147-20190904_112115.csv"

	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, "bank_order_no|trade_amount|send_time|trade_status|redeem_amounts")
	fmt.Fprintln(writer, "P876381909030820295918|2|20190903153429|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909030820303666|2|20190903153449|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909030820261964|3|20190903152114|FAIL|0.01^0.01^0.01")
	fmt.Fprintln(writer, "P876381909030820311124|2|20190903155700|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909030820325372|3|20190903160725|FAIL|")
	fmt.Fprintln(writer, "P876381909030820277703|2|20190903153309|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909030820280451|2|20190903153334|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909030820348104|2|20190903161040|FAIL|")
	fmt.Fprintln(writer, "P876381909030820351863|2|20190903163625|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909030820395013|2|20190903185957|SUCCESS|0.01^0.01")
	fmt.Fprintln(writer, "P876381909030820331901|15000|20190903161020|SUCCESS|")
	fmt.Fprintln(writer, "P876381909030820375229|2|20190903183056|FAIL|0.01^0.01")
	fmt.Fprintln(writer, "P876381909030820363270|2|20190903173336|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909030820389182|2|20190903185017|SUCCESS|0.02")
	fmt.Fprintln(writer, "P876381909040820423589|2|20190904122531|FAIL|")
	fmt.Fprintln(writer, "P876381909040820406676|2|20190904100354|FAIL|")
	fmt.Fprintln(writer, "P876381909040820431243|2|20190904133652|SUCCESS|0.01^0.01")
	fmt.Fprintln(writer, "P876381909040820447404|2|20190904134122|SUCCESS|0.01^0.01")
	fmt.Fprintln(writer, "P876381909040820410300|2|20190904122246|SUCCESS|0.01^0.01")
	fmt.Fprintln(writer, "P876381909040820410322|10|20190904122200|SUCCESS|0.01^0.01^0.01^0.01^0.01^0.01^0.01^0.01^0.01^0.01")

	l := 1000000

	for i := 0; i < 250000; i++ {
		fmt.Fprintln(writer, getLine(l))
		l++
	}

}
