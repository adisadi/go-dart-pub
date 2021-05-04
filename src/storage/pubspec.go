package storage

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strings"

	"github.com/icza/dyno"
	"gopkg.in/yaml.v2"
)

func PubspecBytesFromPackage(file *multipart.FileHeader) (*[]byte, error) {

	if !strings.HasSuffix(file.Filename, "tar.gz") {
		return nil, errors.New("invalid fileending")
	}

	ff, _ := file.Open()

	defer ff.Close()

	uncompressedStream, err := gzip.NewReader(ff)
	if err != nil {
		return nil, err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		if header.Name == "pubspec.yaml" {

			data, err := io.ReadAll(tarReader)
			if err != nil {
				return nil, err
			}

			return &data, nil
		}
	}

	return nil, errors.New("no pubspec.yaml found")
}

func StoragePubspecAsJson(packageName string, version string) *string {

	pubspecJson := ""
	content, err := ioutil.ReadFile(GetPackageVersionPubspec(packageName, version))
	if err != nil {
		log.Println(err.Error())
		return &pubspecJson
	}

	var body interface{}
	if err := yaml.Unmarshal(content, &body); err != nil {
		log.Println(err.Error())
		return &pubspecJson
	}

	body = dyno.ConvertMapI2MapS(body)

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(body)
	pubspecJson = bf.String()
	return &pubspecJson
}
