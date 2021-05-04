package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"sort"
	"time"

	"github.com/hashicorp/go-version"
	"gopkg.in/yaml.v2"
)

type PackageInfo struct {
	Name     string        `json:"name"`
	Latest   VersionInfo   `json:"latest"`
	Versions []VersionInfo `json:"versions"`
}

type VersionInfo struct {
	Version    string           `json:"version"`
	Pubspec    *json.RawMessage `json:"pubspec"`
	ArchiveUrl string           `json:"archive_url"`
	Published  time.Time        `json:"published"`
}

func FindPackageVersions(packageName string, url *url.URL) (*PackageInfo, error) {
	if _, err := os.Stat(GetPackagePath(packageName)); err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(GetPackagePath(packageName))
	if err != nil {
		return nil, err
	}

	var versions []VersionInfo

	for _, f := range files {
		pubspecRaw := json.RawMessage(*StoragePubspecAsJson(packageName, f.Name()))
		versions = append(versions, VersionInfo{
			Version:    f.Name(),
			Pubspec:    &pubspecRaw,
			ArchiveUrl: fmt.Sprintf("%v://%v/packages/%v/versions/%v.tar.gz", url.Scheme, url.Host, packageName, f.Name()),
			Published:  f.ModTime(),
		})
	}

	sort.Slice(versions, func(i, j int) bool {
		v1, _ := version.NewVersion(versions[i].Version)
		v2, _ := version.NewVersion(versions[j].Version)
		return v2.LessThan(v1)
	})

	pi := PackageInfo{
		Name:     packageName,
		Latest:   versions[0],
		Versions: versions,
	}

	return &pi, nil
}

func UploadPackage(file *multipart.FileHeader) error {

	var err error

	pubspecBytes, err := PubspecBytesFromPackage(file)
	if err != nil {
		return err
	}

	yml := struct {
		Name    string
		Version string
	}{}

	err = yaml.Unmarshal(*pubspecBytes, &yml)
	if err != nil {
		return err
	}

	err = os.MkdirAll(GetPackageVersionPath(yml.Name, yml.Version), os.ModePerm)
	if err != nil {
		return err
	}

	outFile, err := os.Create(GetPackageVersionPubspec(yml.Name, yml.Version))
	if err != nil {
		return err
	}
	defer outFile.Close()

	if _, err := outFile.Write(*pubspecBytes); err != nil {
		return err
	}

	archive, err := file.Open()
	if err != nil {
		return err
	}

	outArchive, err := os.Create(GetPackageVersionArchive(yml.Name, yml.Version, "package.tar.gz"))
	if err != nil {
		return err
	}

	_, err = io.Copy(outArchive, archive)
	if err != nil {
		return err
	}

	return nil
}

func Download(packageName string, version string) (string, error) {
	if _, err := os.Stat(GetPackageVersionArchive(packageName, version, "package.tar.gz")); err == nil {
		return GetPackageVersionArchive(packageName, version, "package.tar.gz"), nil
	} else if os.IsNotExist(err) {
		return "", fmt.Errorf("package '%v' version '%v' doesnt exists: %v", packageName, version, err)
	} else {
		return "", err
	}
}
