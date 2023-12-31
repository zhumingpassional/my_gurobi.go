package setup

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
)

/*
gurobi_version_info.go
*/

// Type Definitions
// ================

type GurobiVersionInfo struct {
	MajorVersion    int
	MinorVersion    int
	TertiaryVersion int
}

/*
CreateGurobiHomeDirectory
Description:

	Creates the Gurobi Home directory where your gurobi program and helper files should be installed when you run
	the Gurobi Installer.
*/
func CreateGurobiHomeDirectory(versionInfo GurobiVersionInfo) (string, error) {
	// Create Base Home Directory
	gurobiHome := fmt.Sprintf(
		"/Library/gurobi%v%v%v",
		versionInfo.MajorVersion,
		versionInfo.MinorVersion,
		versionInfo.TertiaryVersion,
	)

	// Create
	switch runtime.GOOS {
	case "darwin":
		// Decide on if we are on intel processors or apple silicon.
		switch runtime.GOARCH {
		case "amd64":
			gurobiHome := fmt.Sprintf("%v/mac64", gurobiHome)
			return gurobiHome, nil
		case "arm64":
			gurobiHome := fmt.Sprintf("%v/macos_universal2", gurobiHome)
			return gurobiHome, nil
		default:
			return "", fmt.Errorf("There was an error finding the architecture on your mac! What is \"%v\"? Send an issue to the gurobi.go repository.", runtime.GOARCH)
		}
	default:
		return "", fmt.Errorf("The operating system that you are using is not recognized: \"%v\".", runtime.GOOS)
	}

}

/*
StringsToGurobiVersionInfoList
Description:

	Receives a set of strings which should be in the format of valid gurobi installation directories
	and returns a list of GurobiVersionInfo objects.

Assumptions:

	Assumes that a valid gurobi name is given.
*/
func StringsToGurobiVersionInfoList(gurobiDirectoryNames []string) ([]GurobiVersionInfo, error) {

	// Convert Directories into Gurobi Version Info
	gurobiVersionList := []GurobiVersionInfo{}
	for _, directory := range gurobiDirectoryNames {
		tempGVI, err := StringToGurobiVersionInfo(directory)
		if err != nil {
			return gurobiVersionList, err
		}
		gurobiVersionList = append(gurobiVersionList, tempGVI)
	}
	// fmt.Println(gurobiVersionList)

	return gurobiVersionList, nil

}

/*
StringToGurobiVersionInfo
Assumptions:

	Assumes that a valid gurobi name is given.
*/
func StringToGurobiVersionInfo(gurobiDirectoryName string) (GurobiVersionInfo, error) {
	// Collect just the version numbers
	versionNumbersString := gurobiDirectoryName[len("gurobi"):]

	//Locate major and minor version indices in gurobi directory name
	majorVersionAsString := string(versionNumbersString[:len(versionNumbersString)-2])
	minorVersionAsString := string(versionNumbersString[len(versionNumbersString)-2])
	tertiaryVersionAsString := string(versionNumbersString[len(versionNumbersString)-1])

	// Convert using strconv to integers
	majorVersion, err := strconv.Atoi(majorVersionAsString)
	if err != nil {
		return GurobiVersionInfo{}, err
	}

	minorVersion, err := strconv.Atoi(minorVersionAsString)
	if err != nil {
		return GurobiVersionInfo{}, err
	}

	tertiaryVersion, err := strconv.Atoi(tertiaryVersionAsString)
	if err != nil {
		return GurobiVersionInfo{}, err
	}

	return GurobiVersionInfo{
		MajorVersion:    majorVersion,
		MinorVersion:    minorVersion,
		TertiaryVersion: tertiaryVersion,
	}, nil

}

/*
// Iterate through all versions in gurobiVersionList and find the one with the largest major or minor version.
*/
func FindHighestVersion(gurobiVersionList []GurobiVersionInfo) (GurobiVersionInfo, error) {

	// Input Checking
	if len(gurobiVersionList) == 0 {
		return GurobiVersionInfo{}, errors.New("No gurobi versions were provided to FindHighestVersion().")
	}

	// Perform search
	highestVersion := gurobiVersionList[0]
	if len(gurobiVersionList) == 1 {
		return highestVersion, nil
	}

	for _, gvi := range gurobiVersionList {
		// Compare Major version numbers
		if gvi.MajorVersion > highestVersion.MajorVersion {
			highestVersion = gvi
			continue
		}

		if gvi.MajorVersion == highestVersion.MajorVersion {
			// Compare minor version numbers
			if gvi.MinorVersion > highestVersion.MinorVersion {
				highestVersion = gvi
				continue
			}

			if gvi.MinorVersion == highestVersion.MinorVersion {
				// Compare tertiary version numbers
				if gvi.TertiaryVersion > highestVersion.TertiaryVersion {
					highestVersion = gvi
					continue
				}
			}
		}
	}

	return highestVersion, nil

}
