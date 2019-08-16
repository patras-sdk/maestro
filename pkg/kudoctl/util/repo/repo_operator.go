package repo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/kudobuilder/kudo/pkg/apis/kudo/v1alpha1"
	"github.com/kudobuilder/kudo/pkg/kudoctl/bundle"
	"github.com/kudobuilder/kudo/pkg/kudoctl/http"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// Repository is an abstraction for a service that can retrieve package bundles
type Repository interface {
	GetBundle(name string, version string) (bundle.Bundle, error)
}

// OperatorRepository represents an operator repository
type OperatorRepository struct {
	Config *RepositoryConfiguration
	Client http.Client
}

// NewOperatorRepository constructs OperatorRepository
func NewOperatorRepository(conf *RepositoryConfiguration) (*OperatorRepository, error) {
	_, err := url.Parse(conf.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid repository URL: %s", conf.URL)
	}

	client := http.NewClient()

	return &OperatorRepository{
		Config: conf,
		Client: *client,
	}, nil
}

// downloadIndexFile fetches the index file from a repository.
func (r *OperatorRepository) downloadIndexFile() (*IndexFile, error) {
	var indexURL string
	parsedURL, err := url.Parse(r.Config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "parsing config url")
	}
	parsedURL.Path = fmt.Sprintf("%s/index.yaml", strings.TrimSuffix(parsedURL.Path, "/"))

	indexURL = parsedURL.String()

	resp, err := r.Client.Get(indexURL)
	if err != nil {
		return nil, errors.Wrap(err, "getting index url")
	}

	indexBytes, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, errors.Wrap(err, "reading index response")
	}

	indexFile, err := parseIndexFile(indexBytes)
	return indexFile, err
}

// getPackageReaderByFullPackageName downloads the tgz file from the remote repository and unmarshals it to the package CRDs
func (r *OperatorRepository) getPackageReaderByFullPackageName(fullPackageName string) (io.Reader, error) {
	var fileURL string
	parsedURL, err := url.Parse(r.Config.URL)
	if err != nil {
		return nil, errors.Wrap(err, "parsing config url")
	}
	parsedURL.Path = fmt.Sprintf("%s/%s.tgz", parsedURL.Path, fullPackageName)

	fileURL = parsedURL.String()
	return r.getPackageReaderByURL(fileURL)
}

func (r *OperatorRepository) getPackageReaderByURL(packageURL string) (io.Reader, error) {
	resp, err := r.Client.Get(packageURL)
	if err != nil {
		return nil, errors.Wrap(err, "getting package url")
	}

	return resp, nil
}

// GetPackageReader provides an io.Reader for a provided package name and optional version
func (r *OperatorRepository) GetPackageReader(name string, version string) (io.Reader, error) {
	// Construct the package name and download the index file from the remote repo
	indexFile, err := r.downloadIndexFile()
	if err != nil {
		return nil, errors.WithMessage(err, "could not download repository index file")
	}

	bundleVersion, err := indexFile.GetByNameAndVersion(name, version)
	if err != nil {
		return nil, errors.Wrapf(err, "getting %s in index file", name)
	}

	packageName := bundleVersion.Name + "-" + bundleVersion.Version

	return r.getPackageReaderByFullPackageName(packageName)
}

// GetBundle provides an Bundle for a provided package name and optional version
func (r *OperatorRepository) GetBundle(name string, version string) (bundle.Bundle, error) {
	reader, err := r.GetPackageReader(name, version)
	if err != nil {
		return nil, err
	}
	return bundle.NewBundleFromReader(reader), nil
}

// GetOperatorVersionDependencies helper method returns a slice of strings that contains the names of all
// dependency Operators
func GetOperatorVersionDependencies(ov *v1alpha1.OperatorVersion) ([]string, error) {
	var dependencyOperators []string
	if ov.Spec.Dependencies != nil {
		for _, v := range ov.Spec.Dependencies {
			dependencyOperators = append(dependencyOperators, v.Name)
		}
	}
	return dependencyOperators, nil
}

// IndexDirectory creates an index file for the operators in the path
func (r *OperatorRepository) IndexDirectory(fs afero.Fs, path string, target string, url string, now *time.Time) (*IndexFile, error) {
	archives, err := afero.Glob(fs, filepath.Join(path, "*.tgz"))
	if err != nil {
		return nil, err
	}
	if len(archives) == 0 {
		return nil, errors.New("no packages discovered")
	}
	index := newIndexFile(now)
	ops := bundle.MapPaths(fs, archives)
	pvs := Map(ops, url)
	for _, pv := range pvs {
		index.addBundleVersion(pv)
	}
	index.sortPackages()
	return index, nil
}
