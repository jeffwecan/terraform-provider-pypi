package pypi

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	api "github.com/jeffwecan/go-pypi/pypi"
)

func dataSourcePackageFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePackageFileRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
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
			"output_size": {
				Type:     schema.TypeInt,
				Computed: true,
				ForceNew: true,
			},
			"output_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "SHA1 checksum of output file",
			},
			"output_base64sha256": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "Base64 Encoded SHA256 checksum of output file",
			},
			"output_md5": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Description: "MD5 of output file",
			},
		},
	}
}

func dataSourcePackageFileRead(d *schema.ResourceData, meta interface{}) (err error) {
	client := meta.(*api.PackageIndex)
	name := d.Get("name").(string)
	outputDir := d.Get("output_dir").(string)
	// log.Printf("[DEBUG] outputDir: %s", outputDir)
	if outputDir != "" {
		if _, err := os.Stat(outputDir); err != nil {
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return err
			}
		}
	}

	var filename string
	if version, ok := d.GetOk("version"); ok {
		filename, err = client.DownloadRelease(outputDir, name, version.(string))
	} else {
		filename, err = client.DownloadLatest(outputDir, name)
	}
	if err != nil {
		return err
	}

	outputPath := path.Join(outputDir, filename)
	d.Set("output_path", outputPath)

	// Generate archived file stats
	fi, err := os.Stat(outputPath)
	if err != nil {
		return err
	}
	sha1, base64sha256, md5, err := genFileShas(outputPath)
	if err != nil {

		return fmt.Errorf("could not generate file checksum sha256: %s", err)
	}
	d.Set("output_sha", sha1)
	d.Set("output_base64sha256", base64sha256)
	d.Set("output_md5", md5)

	d.Set("output_size", fi.Size())
	d.SetId(d.Get("output_sha").(string))

	return nil
}

func genFileShas(filename string) (string, string, string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", "", "", fmt.Errorf("could not compute file '%s' checksum: %s", filename, err)
	}
	h := sha1.New()
	h.Write([]byte(data))
	sha1 := hex.EncodeToString(h.Sum(nil))

	h256 := sha256.New()
	h256.Write([]byte(data))
	shaSum := h256.Sum(nil)
	sha256base64 := base64.StdEncoding.EncodeToString(shaSum[:])

	md5 := md5.New()
	md5.Write([]byte(data))
	md5Sum := hex.EncodeToString(md5.Sum(nil))

	return sha1, sha256base64, md5Sum, nil
}
