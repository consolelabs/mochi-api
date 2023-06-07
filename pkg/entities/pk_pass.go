package entities

import (
	"fmt"

	"github.com/alvinbaena/passkit"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) PassHandle(req request.GeneratePkPassRequest) ([]byte, error) {
	// Create a new pass and set the pass type
	c := passkit.NewGenericPass()

	// Add header fields value
	c.AddHeaderField(passkit.Field{
		Key:   req.Category,
		Value: req.Type,
	})

	// Add secondary fields value
	c.AddSecondaryFields(passkit.Field{
		Key:   req.Type,
		Value: req.Content,
		Label: req.Category,
	})

	//create pass file with value from apple dev account
	pass := passkit.Pass{
		FormatVersion:      1,
		Description:        req.QrValue,
		TeamIdentifier:     e.cfg.PkpassMochiPassTeamIdentifier,
		OrganizationName:   e.cfg.PkpassMochiOrganizationName,
		PassTypeIdentifier: e.cfg.PkpassMochiPassTypeIdentifier,
		BackgroundColor:    "rgb(254, 233, 217)",
		ForegroundColor:    "rgb(17,24,39)",
		SerialNumber:       util.RandomString(8),
		Generic:            c,
		Barcodes: []passkit.Barcode{
			{
				Format:          passkit.BarcodeFormatQR,
				Message:         req.QrValue,
				MessageEncoding: "iso-8859-1",
				AltText:         req.Content,
			},
		},
	}

	// Set buffer template, so we don't need to save it in file
	template := passkit.NewInMemoryPassTemplate()
	template.AddAllFiles("images/pkpass")

	// Create signInfo from cert files
	signInfo, err := passkit.LoadSigningInformationFromFiles(e.cfg.CertDir+"/Certificates.p12", e.cfg.PKpassMochiKeyStorePass, e.cfg.CertDir+"/AppleWWDRCAG3.cer")
	if err != nil {
		return nil, fmt.Errorf("fail to LoadSigningInformationFromFiles pkpass: %v", err)
	}

	// Create singer, signed file with signInfo and zipped this to []byte data
	signer := passkit.NewMemoryBasedSigner()
	z, err := signer.CreateSignedAndZippedPassArchive(&pass, template, signInfo)
	if err != nil {
		return nil, fmt.Errorf("fail to CreateSignedAndZippedPassArchive pkpass: %v", err)
	}

	return z, nil

}
