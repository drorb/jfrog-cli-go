package commands

import (
    "github.com/JFrogDev/bintray-cli-go/bintray/utils"
)

func DownloadFile(versionDetails *utils.VersionDetails, path string, bintrayDetails *utils.BintrayDetails) {
    if bintrayDetails.User == "" {
        bintrayDetails.User = versionDetails.Subject
    }
    utils.DownloadBintrayFile(bintrayDetails, versionDetails, path)
}