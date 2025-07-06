package main

import (
	"log"
	"os"

	"github.com/HadasAmar/analytics-load-tool/Reader"
	Simulator "github.com/HadasAmar/analytics-load-tool/Simulator"
	"github.com/HadasAmar/analytics-load-tool/configuration"
	formatter "github.com/HadasAmar/analytics-load-tool/formatter"
)

func main() {
	
	
	errGlobalConsul := configuration.InitGlobalConsul()
	if errGlobalConsul != nil {
		panic(errGlobalConsul)
	}
	if len(os.Args) < 2 {
		log.Fatal("Pass a path to the log file as a parameter")
	}
	logFile := os.Args[1]

	// load the reader for the log file
	reader, err := Reader.GetReader(logFile)
	if err != nil {
		log.Fatalf("❌ error finding the reader: %v", err)
	}

	// reads the log file and parses it into records
	records, err := reader.Read(logFile)
	if err != nil {
		log.Fatalf("❌ error reading the reader: %v", err)
	}

	errSimulateReplay := Simulator.SimulateReplay(records)
	if errSimulateReplay != nil {
		log.Fatalf("error simulating: %v", errSimulateReplay)
	}

	f, err := os.Create("output.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	count := 0
	for _, record := range records {
		if record == nil || record.Parsed == nil {
			continue
		}

		// creates SQL → formats it → writes to file
		raw := formatter.BuildSQLQuery(record.Parsed)
		pretty := formatter.PrettySQL(raw)

		count++

		_, err := f.WriteString(pretty + "\n\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}
