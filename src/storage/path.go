package storage

import (
	"path"
)

var StorageBasePath = "/data"

func GetPackagePath(packageName string) string {
	return path.Join(StorageBasePath, packageName)
}

func GetPackageVersionPath(packageName string, version string) string {
	return path.Join(GetPackagePath(packageName), version)
}

func GetPackageVersionPubspec(packageName string, version string) string {
	return path.Join(GetPackageVersionPath(packageName, version), "pubspec.yaml")
}

func GetPackageVersionArchive(packageName string, version string, filename string) string {
	return path.Join(GetPackageVersionPath(packageName, version), filename)
}
