package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/validator"
)

//counterfeiter:generate -o ./fakes/file_artifacter.go --fake-name FileArtifacter . FileArtifacter
type FileArtifacter interface {
	Name() string
	SHA256() string
}

//counterfeiter:generate -o ./fakes/stemcell_artifacter.go --fake-name StemcellArtifacter . StemcellArtifacter
type StemcellArtifacter interface {
	Slug() string
	Version() string
}

//counterfeiter:generate -o ./fakes/product_downloader_service.go --fake-name ProductDownloader . ProductDownloader
type ProductDownloader interface {
	Name() string
	GetAllProductVersions(slug string) ([]string, error)
	GetLatestProductFile(slug, version, glob string) (FileArtifacter, error)
	DownloadProductToFile(fa FileArtifacter, file *os.File) error
	GetLatestStemcellForProduct(fa FileArtifacter, downloadedProductFileName string) (StemcellArtifacter, error)
}

type DownloadProductOptions struct {
	Source     string   `long:"source"                     short:"s" description:"enables download from external sources when set to [s3|gcs|azure|pivnet]" default:"pivnet"`
	ConfigFile string   `long:"config"                     short:"c" description:"path to yml file for configuration (keys must match the following command line flags)"`
	OutputDir  string   `long:"output-directory"           short:"o" description:"directory path to which the file will be outputted. File Name will be preserved from Pivotal Network" required:"true"`
	VarsEnv    []string `long:"vars-env" env:"OM_VARS_ENV"           description:"load variables from environment variables matching the provided prefix (e.g.: 'MY' to load MY_var=value)" experimental:"true"`
	VarsFile   []string `long:"vars-file"                  short:"l" description:"load variables from a YAML file"`
	Vars       []string `long:"var"                                  description:"Load variable from the command line. Format: VAR=VAL"`

	PivnetProductSlug   string `long:"pivnet-product-slug"   short:"p"                          description:"path to product" required:"true"`
	PivnetDisableSSL    bool   `long:"pivnet-disable-ssl"                                       description:"whether to disable ssl validation when contacting the Pivotal Network"`
	PivnetToken         string `long:"pivnet-api-token"      short:"t"                          description:"API token to use when interacting with Pivnet. Can be retrieved from your profile page in Pivnet."`
	FileGlob            string `long:"file-glob"             short:"f" alias:"pivnet-file-glob" description:"glob to match files within Pivotal Network product to be downloaded." required:"true"`
	ProductVersion      string `long:"product-version"       short:"v"                          description:"version of the product-slug to download files from. Incompatible with --product-version-regex flag."`
	ProductVersionRegex string `long:"product-version-regex" short:"r"                          description:"regex pattern matching versions of the product-slug to download files from. Highest-versioned match will be used. Incompatible with --product-version flag."`

	Bucket       string `long:"blobstore-bucket"        alias:"s3-bucket,gcs-bucket,azure-container"                   description:"bucket name where the product resides in the s3|gcs|azure compatible blobstore"`
	ProductPath  string `long:"blobstore-product-path"  alias:"s3-product-path,gcs-product-path,azure-product-path"    description:"specify the lookup path where the s3|gcs|azure product artifacts are stored"`
	StemcellPath string `long:"blobstore-stemcell-path" alias:"s3-stemcell-path,gcs-stemcell-path,azure-stemcell-path" description:"specify the lookup path where the s3|gcs|azure stemcell artifacts are stored"`

	GCSServiceAccountJSON string `long:"gcs-service-account-json" alias:"gcp-service-account-json" description:"the service account key JSON"`
	GCSProjectID          string `long:"gcs-project-id"           alias:"gcp-project-id"           description:"the project id for the bucket's gcp account"`

	S3AccessKeyID     string `long:"s3-access-key-id"                 description:"access key for the s3 compatible blobstore"`
	S3AuthType        string `long:"s3-auth-type"                     description:"can be set to \"iam\" in order to allow use of instance credentials" default:"accesskey"`
	S3SecretAccessKey string `long:"s3-secret-access-key"             description:"secret key for the s3 compatible blobstore"`
	S3RegionName      string `long:"s3-region-name"                   description:"bucket region in the s3 compatible blobstore. If not using AWS, this value is 'region'"`
	S3Endpoint        string `long:"s3-endpoint"                      description:"the endpoint to access the s3 compatible blobstore. If not using AWS, this is required"`
	S3DisableSSL      bool   `long:"s3-disable-ssl"                   description:"whether to disable ssl validation when contacting the s3 compatible blobstore"`
	S3EnableV2Signing bool   `long:"s3-enable-v2-signing"             description:"whether to use v2 signing with your s3 compatible blobstore. (if you don't know what this is, leave blank, or set to 'false')"`

	AzureStorageAccount string `long:"azure-storage-account" description:"the name of the storage account where the container exists"`
	AzureKey            string `long:"azure-storage-key"     description:"the access key for the storage account"`

	DeprecatedStemcell bool   `long:"download-stemcell" description:"**DEPRECATED**: no-op for backwards compatibility"`
	StemcellIaas       string `long:"stemcell-iaas"     description:"download the latest available stemcell for the product for the specified iaas. for example 'vsphere' or 'vcloud' or 'openstack' or 'google' or 'azure' or 'aws'. Can contain globbing patterns to match specific files in a stemcell release on Pivnet"`
	StemcellHeavy      bool   `long:"stemcell-heavy" description:"force the downloading of a heavy stemcell, will fail if non exists"`
}

type DownloadProduct struct {
	environFunc    func() []string
	progressWriter io.Writer
	stderr         *log.Logger
	stdout         *log.Logger
	downloadClient ProductDownloader
	Options        DownloadProductOptions
}

func NewDownloadProduct(
	environFunc func() []string,
	stdout *log.Logger,
	stderr *log.Logger,
	progressWriter io.Writer,
) *DownloadProduct {
	return &DownloadProduct{
		environFunc:    environFunc,
		stderr:         stderr,
		stdout:         stdout,
		progressWriter: progressWriter,
	}
}

func (c DownloadProduct) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "This command attempts to download a single product file from Pivotal Network. The API token used must be associated with a user account that has already accepted the EULA for the specified product",
		ShortDescription: "downloads a specified product file from Pivotal Network",
		Flags:            c.Options,
	}
}

func (c *DownloadProduct) Execute(args []string) error {
	err := loadConfigFile(args, &c.Options, c.environFunc)
	if err != nil {
		return fmt.Errorf("could not parse download-product flags: %s", err)
	}

	err = c.validate()
	if err != nil {
		return err
	}

	err = c.createClient()
	if err != nil {
		return err
	}

	productVersion, err := c.determineProductVersion()
	if err != nil {
		return err
	}

	productFileName, productFileArtifact, err := c.downloadProductFile(
		c.Options.PivnetProductSlug,
		productVersion,
		c.Options.FileGlob,
		fmt.Sprintf("[%s,%s]", c.Options.PivnetProductSlug, productVersion),
	)
	if err != nil {
		return fmt.Errorf("could not download product: %s", err)
	}

	if c.Options.StemcellIaas == "" {
		return c.writeDownloadProductOutput(productFileName, productVersion, "", "")
	}

	c.stderr.Printf("Downloading stemcell")

	nameParts := strings.Split(productFileName, ".")
	if nameParts[len(nameParts)-1] != "pivotal" {
		c.stderr.Printf("the downloaded file is not a .pivotal file. Not determining and fetching required stemcell.")
		return c.writeDownloadProductOutput(productFileName, productVersion, "", "")
	}

	stemcell, err := c.downloadClient.GetLatestStemcellForProduct(productFileArtifact, productFileName)
	if err != nil {
		return fmt.Errorf("could not get information about stemcell: %s", err)
	}

	stemcellGlobs := []string{
		fmt.Sprintf("light*bosh*%s*", c.Options.StemcellIaas),
		fmt.Sprintf("bosh*%s*", c.Options.StemcellIaas),
	}

	if c.Options.StemcellHeavy {
		stemcellGlobs = []string{
			fmt.Sprintf("bosh*%s*", c.Options.StemcellIaas),
		}
	}

	stemcellFileName, err := "", nil
	for _, stemcellGlob := range stemcellGlobs {
		stemcellFileName, _, err = c.downloadProductFile(
			stemcell.Slug(),
			stemcell.Version(),
			stemcellGlob,
			fmt.Sprintf("[%s,%s]", stemcell.Slug(), stemcell.Version()),
		)
		if err == nil {
			break
		}
	}
	if err != nil {
		isHeavy := ""
		if c.Options.StemcellHeavy {
			isHeavy = "heavy "
		}
		return fmt.Errorf("could not download stemcell: %s\nNo %sstemcell identified for IaaS \"%s\" on Pivotal Network. Correct the `stemcell-iaas` option to match the IaaS portion of the stemcell filename, or remove the option", err, isHeavy, c.Options.StemcellIaas)
	}

	err = c.writeDownloadProductOutput(productFileName, productVersion, stemcellFileName, stemcell.Version())
	if err != nil {
		return err
	}

	return c.writeAssignStemcellInput(productFileName, stemcell.Version())
}

func (c *DownloadProduct) determineProductVersion() (string, error) {
	if c.Options.ProductVersionRegex != "" {
		re, err := regexp.Compile(c.Options.ProductVersionRegex)
		if err != nil {
			return "", fmt.Errorf("could not compile regex '%s': %s", c.Options.ProductVersionRegex, err)
		}

		productVersions, err := c.downloadClient.GetAllProductVersions(c.Options.PivnetProductSlug)
		if err != nil {
			return "", err
		}

		var versions version.Collection
		for _, productVersion := range productVersions {
			if !re.MatchString(productVersion) {
				continue
			}

			v, err := version.NewVersion(productVersion)
			if err != nil {
				c.stderr.Printf(fmt.Sprintf("warning: could not parse semver version from: %s", productVersion))
				continue
			}
			versions = append(versions, v)
		}

		sort.Sort(versions)

		if len(versions) == 0 {
			existingVersions := strings.Join(productVersions, ", ")
			if existingVersions == "" {
				existingVersions = "none"
			}
			return "", fmt.Errorf("no valid versions found for product '%s' and product version regex '%s'\nexisting versions: %s", c.Options.PivnetProductSlug, c.Options.ProductVersionRegex, existingVersions)
		}

		return versions[len(versions)-1].Original(), nil
	}
	return c.Options.ProductVersion, nil
}

func (c *DownloadProduct) createClient() error {
	plugin, ok := plugins[c.Options.Source]
	if !ok {
		return fmt.Errorf("could not find valid source for '%s'", c.Options.Source)
	}

	value, err := plugin(c.Options, c.progressWriter, c.stdout, c.stderr)
	if err != nil {
		return err
	}

	c.downloadClient = value
	return nil
}

func (c *DownloadProduct) validate() error {
	if c.Options.ProductVersionRegex != "" && c.Options.ProductVersion != "" {
		return fmt.Errorf("cannot use both --product-version and --product-version-regex; please choose one or the other")
	}

	if c.Options.ProductVersionRegex == "" && c.Options.ProductVersion == "" {
		return fmt.Errorf("no version information provided; please provide either --product-version or --product-version-regex")
	}

	if c.Options.PivnetToken == "" && c.Options.Source == "pivnet" {
		return fmt.Errorf(`could not execute "download-product": could not parse download-product flags: missing required flag "--pivnet-api-token"`)
	}

	if c.Options.StemcellHeavy == true && c.Options.StemcellIaas == "" {
		return fmt.Errorf("--stemcell-heavy requires --stemcell-iaas to be defined")
	}

	return nil
}

func (c DownloadProduct) writeDownloadProductOutput(productFileName string, productVersion string, stemcellFileName string, stemcellVersion string) error {
	downloadProductFilename := "download-file.json"
	c.stderr.Printf("Writing a list of downloaded artifact to %s", downloadProductFilename)
	downloadProductPayload := struct {
		ProductPath     string `json:"product_path,omitempty"`
		ProductSlug     string `json:"product_slug,omitempty"`
		ProductVersion  string `json:"product_version,omitempty"`
		StemcellPath    string `json:"stemcell_path,omitempty"`
		StemcellVersion string `json:"stemcell_version,omitempty"`
	}{
		ProductPath:     productFileName,
		StemcellPath:    stemcellFileName,
		ProductSlug:     c.Options.PivnetProductSlug,
		ProductVersion:  productVersion,
		StemcellVersion: stemcellVersion,
	}

	outputFile, err := os.Create(filepath.Join(c.Options.OutputDir, downloadProductFilename))
	if err != nil {
		return fmt.Errorf("could not create %s: %s", downloadProductFilename, err)
	}
	defer outputFile.Close()

	err = json.NewEncoder(outputFile).Encode(downloadProductPayload)
	if err != nil {
		return fmt.Errorf("could not encode JSON for %s: %s", downloadProductFilename, err)
	}

	return nil
}

func (c DownloadProduct) writeAssignStemcellInput(productFileName string, stemcellVersion string) error {
	assignStemcellFileName := "assign-stemcell.yml"
	c.stderr.Printf("Writing a assign stemcell artifact to %s", assignStemcellFileName)
	metadata, err := getProductMetadata(productFileName)
	if err != nil {
		return fmt.Errorf("cannot parse product metadata: %s", err)
	}

	assignStemcellPayload := struct {
		Product  string `json:"product"`
		Stemcell string `json:"stemcell"`
	}{
		Product:  metadata.ProductName,
		Stemcell: stemcellVersion,
	}

	outputFile, err := os.Create(filepath.Join(c.Options.OutputDir, assignStemcellFileName))
	if err != nil {
		return fmt.Errorf("could not create %s: %s", assignStemcellFileName, err)
	}
	defer outputFile.Close()

	err = json.NewEncoder(outputFile).Encode(assignStemcellPayload)
	if err != nil {
		return fmt.Errorf("could not encode JSON for %s: %s", assignStemcellFileName, err)
	}

	return nil
}

func (c *DownloadProduct) downloadProductFile(slug, version, glob, prefixPath string) (string, FileArtifacter, error) {
	fileArtifact, err := c.downloadClient.GetLatestProductFile(slug, version, glob)
	if err != nil {
		return "", nil, err
	}

	var productFilePath string
	if c.Options.Source != "pivnet" || c.Options.Bucket == "" {
		productFilePath = filepath.Join(c.Options.OutputDir, filepath.Base(fileArtifact.Name()))
	} else {
		productFilePath = filepath.Join(c.Options.OutputDir, prefixPath+filepath.Base(fileArtifact.Name()))
	}

	c.stderr.Printf("attempting to download the file %s from source %s", fileArtifact.Name(), c.downloadClient.Name())

	// check for already downloaded file
	exist, err := checkFileExists(productFilePath)
	if err != nil {
		return productFilePath, nil, err
	}

	if exist {
		if ok, _ := c.shasumMatches(productFilePath, fileArtifact.SHA256()); ok {
			c.stderr.Printf("%s already exists, skip downloading", productFilePath)
			return productFilePath, fileArtifact, nil
		} else {
			c.stderr.Printf("%s already exists, sha sum does not match, re-downloading", productFilePath)
		}
	}

	partialProductFilePath := productFilePath + ".partial"
	// create a new file to download
	productFile, err := os.Create(partialProductFilePath)
	if err != nil {
		return "", nil, fmt.Errorf("could not create file %s: %s", productFilePath, err)
	}
	defer productFile.Close()

	err = c.downloadClient.DownloadProductToFile(fileArtifact, productFile)
	if err != nil {
		return productFilePath, fileArtifact, err
	}

	// check for correct sha on newly downloaded file
	if ok, calculatedSum := c.shasumMatches(partialProductFilePath, fileArtifact.SHA256()); !ok {
		e := fmt.Sprintf("the sha (%s) from %s does not match the calculated sha (%s) for the file %s",
			fileArtifact.SHA256(),
			c.downloadClient.Name(),
			calculatedSum,
			productFilePath,
		)
		c.stderr.Print(e)
		_ = os.Remove(partialProductFilePath)
		return productFilePath, fileArtifact, fmt.Errorf(e)
	}

	_ = os.Rename(partialProductFilePath, productFilePath)
	return productFilePath, fileArtifact, nil
}

func (c *DownloadProduct) shasumMatches(path, exepectedSum string) (bool, string) {
	if exepectedSum == "" {
		return true, ""
	}

	c.stderr.Printf("calculating sha sum for %s", path)
	validate := validator.NewSHA256Calculator()
	calculatedSum, err := validate.Checksum(path)
	if err != nil {
		return false, ""
	}

	return calculatedSum == exepectedSum, calculatedSum
}

func checkFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, fmt.Errorf("failed to get file information: %s", err)
		}
	}

	return true, nil
}

type ProductClientRegistration func(
	c DownloadProductOptions,
	progressWriter io.Writer,
	stdout *log.Logger,
	stderr *log.Logger,
) (ProductDownloader, error)

var plugins = make(map[string]ProductClientRegistration)

func RegisterProductClient(name string, f ProductClientRegistration) {
	plugins[name] = f
}
