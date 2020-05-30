package pypi

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	api "github.com/jeffwecan/go-pypi/pypi"
)

func dataSourceRequirementsFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRequirementsFileRead,
		Schema: map[string]*schema.Schema{
			"requirements_file": {
				Type:     schema.TypeString,
				Required: true,
			},
			"output_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
			"requirements_sha": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			// "output_size": {
			// 	Type:     schema.TypeInt,
			// 	Computed: true,
			// 	ForceNew: true,
			// },
			"output_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				// ForceNew:    true,
				Description: "SHA1 checksum of requirements file",
			},
		},
	}
}

func dataSourceRequirementsFileRead(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*api.PackageIndex)
	outputDir := d.Get("output_dir").(string)
	requirementsFile := d.Get("requirements_file").(string)
	log.Printf("[DEBUG] outputDir: %s", outputDir)
	if outputDir != "" {
		if _, err := os.Stat(outputDir); err != nil {
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return err
			}
		}
	}

	log.Printf("[DEBUG] starting downloads from: %s", requirementsFile)
	reqs, err := client.DownloadFromRequirementsFile(outputDir, requirementsFile)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] generating sha1 hash of requirements file: %s", requirementsFile)
	data, err := ioutil.ReadFile(requirementsFile)
	if err != nil {
		return fmt.Errorf("could not compute file '%s' checksum: %s", requirementsFile, err)
	}
	reqsH := sha1.New()
	reqsH.Write([]byte(data))
	requirementsSha1 := hex.EncodeToString(reqsH.Sum(nil))
	if err != nil {
		return fmt.Errorf("could not generate file checksum sha256: %s", err)
	}
	log.Printf("[DEBUG] generated sha1 hash of requirements file: %s", requirementsSha1)
	d.Set("requirements_sha", requirementsSha1)

	log.Printf("[DEBUG] generating sha1 hash of returned list of downloaded requirements")
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%v", reqs)))
	sha1 := hex.EncodeToString(h.Sum(nil))
	log.Printf("[DEBUG] generated sha1 hash: %s", sha1)
	d.Set("output_sha", sha1)
	d.SetId(d.Get("output_sha").(string))

	return nil
}
