/*
 * ZAnnotate Copyright 2017 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {

	var conf GlobalConf
	flags := flag.NewFlagSet("flags", flag.ExitOnError)
	flags.StringVar(&conf.InputFilePath, "input-file", "-", "ip addresses to read")
	flags.StringVar(&conf.InputFileType, "input-file-type", "ips", "ips, csv, json")
	flags.StringVar(&conf.OutputFilePath, "output-file", "-", "where should JSON output be saved")
	flags.StringVar(&conf.MetadataFilePath, "metadata-file", "", "where should JSON metadata be saved")
	flags.StringVar(&conf.LogFilePath, "log-file", "", "where should JSON logs be saved")
	flags.IntVar(&conf.Verbosity, "verbosity", 3, "log verbosity: 1 (lowest)--5 (highest)")
	flags.IntVar(&conf.Threads, "threads", 5, "how many processing threads to use")
	// MaxMind GeoIP2
	flags.BoolVar(&conf.GeoIP2City, "maxmind-geoip2", false, "geolocate")
	flags.StringVar(&conf.GeoIP2DatabasePath, "maxmind-geoip2-database", "", "path to MaxMind GeoIP2 database")
	flags.StringVar(&conf.GeoIP2Mode, "maxmind-mode", "mmap", "how to open database: mmap or memory")
	flags.StringVar(&conf.GeoIP2Language, "maxmind-language", "en", "how to open database: mmap or memory")
	flags.StringVar(&conf.GeoIP2Fields, "maxmind-fields", "*", "city, continent, country, location, postal, registered_country, subdivisions, traits")
	// Routing Table AS Data
	flags.BoolVar(&conf.Routing, "routing", false, "add routing data")
	flags.StringVar(&conf.RoutingTablePath, "routing-table-path", "", "geolocate")
	flags.StringVar(&conf.ASData, "as-data", "", "geolocate")

	flags.Parse(os.Args[2:])
	if conf.LogFilePath != "" {
		f, err := os.OpenFile(gc.LogFilePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalf("Unable to open log file (%s): %s", gc.LogFilePath, err.Error())
		}
		log.SetOutput(f)
	}
	// Translate the assigned verbosity level to a logrus log level.
	switch conf.Verbosity {
	case 1: // Fatal
		log.SetLevel(log.FatalLevel)
	case 2: // Error
		log.SetLevel(log.ErrorLevel)
	case 3: // Warnings  (default)
		log.SetLevel(log.WarnLevel)
	case 4: // Information
		log.SetLevel(log.InfoLevel)
	case 5: // Debugging
		log.SetLevel(log.DebugLevel)
	default:
		log.Fatal("Unknown verbosity level specified. Must be between 1 (lowest)--5 (highest)")
	}

}
