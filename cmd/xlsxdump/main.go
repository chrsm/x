package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/tealeg/xlsx/v2"
)

var (
	src string
	dst string
)

func main() {
	flag.StringVar(&src, "src", "", "xlsx to parse")
	flag.StringVar(&dst, "dst", "", "dir to write csv files to")
	flag.Parse()

	if src == "" || dst == "" {
		log.Fatal("provide -src and -dst")
	}

	xlfi, err := xlsx.OpenFile(src)
	if err != nil {
		log.Fatalf("failed to open xlsx: %s", err)
	}

	for i := range xlfi.Sheets {
		sheet := xlfi.Sheets[i]
		log.Printf("sheet[%d](%s): %d rows", i, sheet.Name, len(sheet.Rows))

		fp, err := os.OpenFile(dst+"/"+sheet.Name+".csv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			log.Fatalf("failed to open file for %s: %s", sheet.Name, err)
		}

		wsheet(sheet, fp)
	}
}

func wsheet(s *xlsx.Sheet, w io.Writer) {
	enc := csv.NewWriter(w)

	for i := range s.Rows {
		r := s.Rows[i]

		rec := make([]string, len(r.Cells))
		for j := range r.Cells {
			c := r.Cells[j]

			if c.String() != "" {
				rec[j] = c.String()
			}
		}

		if allempty(rec) {
			continue
		}

		enc.Write(rec)
	}

	enc.Flush()
}

func allempty(s []string) bool {
	for i := range s {
		if s[i] != "" {
			return false
		}
	}

	return true
}
