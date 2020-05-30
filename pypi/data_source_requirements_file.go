package pypi

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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
			"additional_files": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"output_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
			"output_path": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
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
				ForceNew:    true,
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

	log.Printf("[DEBUG] generate sha1 hash of returned list of downloaded requirements")
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%v", reqs)))
	sha1 := hex.EncodeToString(h.Sum(nil))
	log.Printf("[DEBUG] generated sha1 hash: %s", sha1)
	d.Set("output_sha", sha1)
	d.SetId(d.Get("output_sha").(string))

	return nil
}
