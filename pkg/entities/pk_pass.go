package entities

import (
	"github.com/alvinbaena/passkit"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) PassHandle(req request.GeneratePkPassRequest) ([]byte, error) {
	// Create a new pass and set the pass type
	c := passkit.NewGenericPass()
	c.AddHeaderField(passkit.Field{
		Key:   req.Category,
		Value: req.Type,
	})
	c.AddSecondaryFields(passkit.Field{
		Key:   req.Type,
		Value: req.Content,
		Label: req.Category,
	})
	pass := passkit.Pass{
		FormatVersion: 1,
		Description:   req.QrValue,

		TeamIdentifier:     "W777S7V8TN",
		OrganizationName:   "Mochi",
		PassTypeIdentifier: "pass.so.console.mochi",
		BackgroundColor:    "rgb(254, 233, 217)",
		ForegroundColor:    "rgb(17,24,39)",
		SerialNumber:       util.RandomString(8),
		Generic:            c,
		Barcodes: []passkit.Barcode{
			{
				Format:          passkit.BarcodeFormatQR,
				Message:         req.QrValue,
				MessageEncoding: "iso-8859-1",
				AltText:         req.QrValue,
			},
		},
	}
	template := passkit.NewInMemoryPassTemplate()
	template.AddAllFiles("images/pkpass")
	signInfo, err := passkit.LoadSigningInformationFromFiles("cert/Certificates.p12", "consolelabs", "cert/AppleWWDRCAG3.cer")
	if err != nil {
		panic(err)
	}

	signer := passkit.NewMemoryBasedSigner()
	z, err := signer.CreateSignedAndZippedPassArchive(&pass, template, signInfo)
	if err != nil {
		panic(err)
	}

	return z, nil

}
