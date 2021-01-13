package keeper

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/x/cert/types"
)

// SetCertificate stores a certificate using its ID field.
func (k Keeper) SetCertificate(ctx sdk.Context, certificate types.Certificate) {
	store := ctx.KVStore(k.storeKey)
	bz := k.MustMarshalCertificate(certificate)
	store.Set(types.CertificateStoreKey(certificate.ID().Bytes()), bz)
}

// MustMarshalCertificate attempts to encode a Certificate object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalCertificate(certificate types.Certificate) []byte {
	// marshals a Certificate interface. If the given type implements
	// the Marshaler interface, it is treated as a Proto-defined message and
	// serialized that way. Otherwise, it falls back on the internal Amino codec.
	bz, err := k.cdc.MarshalInterface(certificate)
	if err != nil {
		panic(fmt.Errorf("failed to encode evidence: %w", err))
	}
	return bz
}

// DeleteCertificate deletes a certificate using its ID field.
func (k Keeper) DeleteCertificate(ctx sdk.Context, certificate types.Certificate) error {
	if !k.HasCertificateByID(ctx, certificate.ID()) {
		return types.ErrCertificateNotExists
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.CertificateStoreKey(certificate.ID().Bytes()))
	return nil
}

// HasCertificateByID checks if a certificate exists given an ID.
func (k Keeper) HasCertificateByID(ctx sdk.Context, id types.CertificateID) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.CertificateStoreKey(id.Bytes()))
}

// GetCertificateByID retrieves a certificate given an ID.
func (k Keeper) GetCertificateByID(ctx sdk.Context, id types.CertificateID) (types.Certificate, error) {
	store := ctx.KVStore(k.storeKey)
	certificateData := store.Get(types.CertificateStoreKey(id.Bytes()))
	if certificateData == nil {
		return nil, types.ErrCertificateNotExists
	}

	var cert types.Certificate
	err := k.cdc.UnmarshalInterface(certificateData, &cert)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

// GetNewCertificateID gets an unused certificate ID for a new certificate.
func (k Keeper) GetNewCertificateID(ctx sdk.Context, certType types.CertificateType,
	certContent types.RequestContent) (types.CertificateID, error) {
	var i uint8
	var certID types.CertificateID
	var err error
	// Find an unoccupied key
	for {
		certID = types.GetCertificateID(certType, certContent, i)
		_, err = k.GetCertificateByID(ctx, certID)
		if err == types.ErrCertificateNotExists {
			break
		}
		if i == math.MaxUint8 {
			return "", errors.New("index overflow")
		}
		i++
	}
	return certID, nil
}

// GetCertificateType gets type of a certificate by certificate ID.
func (k Keeper) GetCertificateType(ctx sdk.Context, id types.CertificateID) (types.CertificateType, error) {
	certificate, err := k.GetCertificateByID(ctx, id)
	if err != nil {
		return types.CertificateTypeNil, err
	}
	return certificate.Type(), nil
}

// IsCertified checks if a certificate of given type and content exists.
func (k Keeper) IsCertified(ctx sdk.Context, requestContentType string, content string, certType string) bool {
	requestContent, err := types.NewRequestContent(requestContentType, content)
	if err != nil {
		return false
	}
	certificateType := types.CertificateTypeFromString(certType)
	certificates := k.GetCertificatesByTypeAndContent(ctx, certificateType, requestContent)
	return len(certificates) > 0
}

// IsContentCertified checks if a certificate of given content exists.
func (k Keeper) IsContentCertified(ctx sdk.Context, requestContent string) bool {
	for _, requestContentType := range types.RequestContentTypes {
		requestContent := types.RequestContent{RequestContentType: requestContentType, RequestContent: requestContent}
		if len(k.GetCertificatesByContent(ctx, requestContent)) > 0 {
			return true
		}
	}
	return false
}

// IssueCertificate issues a certificate.
func (k Keeper) IssueCertificate(ctx sdk.Context, c types.Certificate) (types.CertificateID, error) {
	if !k.IsCertifier(ctx, c.Certifier()) {
		return "", types.ErrUnqualifiedCertifier
	}

	certificateID, err := k.GetNewCertificateID(ctx, c.Type(), c.RequestContent())
	if err != nil {
		return "", err
	}
	c.SetCertificateID(certificateID)

	txhash := hex.EncodeToString(tmhash.Sum(ctx.TxBytes()))
	c.SetTxHash(txhash)

	k.SetCertificate(ctx, c)

	return c.ID(), nil
}

// IterateAllCertificate iterates over the all the stored certificates and performs a callback function.
func (k Keeper) IterateAllCertificate(ctx sdk.Context, callback func(certificate types.Certificate) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.CertificatesStoreKey())

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var cert types.Certificate
		err := k.cdc.UnmarshalInterface(iterator.Value(), &cert)
		if err != nil {
			panic(err)
		}

		if callback(cert) {
			break
		}
	}
}

// IterateCertificatesByContent iterates over certificates with identical given certifier,
// certificate type, and certificate content.
func (k Keeper) IterateCertificatesByContent(ctx sdk.Context, certType types.CertificateType,
	content types.RequestContent, callback func(certificate types.Certificate) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := types.CertificateStoreContentKey(certType, content.RequestContentType, content.RequestContent)
	iterator := sdk.KVStorePrefixIterator(store, prefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var cert types.Certificate
		err := k.cdc.UnmarshalInterface(iterator.Value(), &cert)
		if err != nil {
			panic(err)
		}

		if callback(cert) {
			break
		}
	}
}

// IterateCertificatesByType iterates over certificates with identical given certificate type.
func (k Keeper) IterateCertificatesByType(ctx sdk.Context, certType types.CertificateType, callback func(certificate types.Certificate) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	prefix := append(types.CertificatesStoreKey(), certType.Bytes()...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var cert types.Certificate
		err := k.cdc.UnmarshalInterface(iterator.Value(), &cert)
		if err != nil {
			panic(err)
		}

		if callback(cert) {
			break
		}
	}
}

// GetAllCertificates gets all certificates.
func (k Keeper) GetAllCertificates(ctx sdk.Context) (certificates []types.Certificate) {
	k.IterateAllCertificate(ctx, func(certificate types.Certificate) bool {
		certificates = append(certificates, certificate)
		return false
	})
	return certificates
}

// GetCertificatesByCertifier gets certificates certified by a given certifier.
func (k Keeper) GetCertificatesByCertifier(ctx sdk.Context, certifier sdk.AccAddress) []types.Certificate {
	certificates := []types.Certificate{}
	k.IterateAllCertificate(ctx, func(certificate types.Certificate) bool {
		if certificate.Certifier().Equals(certifier) {
			certificates = append(certificates, certificate)
		}
		return false
	})
	return certificates
}

// GetCertificatesByContent retrieves all certificates with given content.
func (k Keeper) GetCertificatesByContent(ctx sdk.Context, requestContent types.RequestContent) []types.Certificate {
	certificates := []types.Certificate{}
	for _, certType := range types.CertificateTypes {
		k.IterateCertificatesByContent(
			ctx,
			certType,
			requestContent,
			func(certificate types.Certificate) bool {
				if certificate.RequestContent() == requestContent {
					certificates = append(certificates, certificate)
				}
				return false
			},
		)
	} // for each certificate type

	return certificates
}

// GetCertificatesByTypeAndContent retrieves all certificates with given certificate type and content.
func (k Keeper) GetCertificatesByTypeAndContent(ctx sdk.Context, certType types.CertificateType,
	requestContent types.RequestContent) []types.Certificate {
	certificates := []types.Certificate{}
	k.IterateCertificatesByContent(
		ctx,
		certType,
		requestContent,
		func(certificate types.Certificate) bool {
			if certificate.RequestContent() == requestContent &&
				certificate.Type() == certType {
				certificates = append(certificates, certificate)
			}
			return false
		},
	)
	return certificates
}

// GetCertificatesFiltered gets certificates filtered.
func (k Keeper) GetCertificatesFiltered(ctx sdk.Context, params types.QueryCertificatesParams) (uint64, []types.Certificate, error) {
	filteredCertificates := []types.Certificate{}
	callback := func(certificate types.Certificate) bool {
		if len(params.Certifier) != 0 && !certificate.Certifier().Equals(params.Certifier) {
			return false
		}
		if params.ContentType != "" &&
			(strings.ToUpper(params.ContentType) != strings.ToUpper(certificate.RequestContent().RequestContentType.String()) ||
				certificate.RequestContent().RequestContent != params.Content) {
			return false
		}
		filteredCertificates = append([]types.Certificate{certificate}, filteredCertificates...)
		return false
	}

	// Choose an efficient iteration mechanism.
	if len(params.Certifier) != 0 {
		if params.ContentType != "" && params.Content != "" {
			for _, certType := range types.CertificateTypes {
				requestContent, err := types.NewRequestContent(params.ContentType, params.Content)
				if err != nil {
					return 0, nil, err
				}
				k.IterateCertificatesByContent(ctx, certType, requestContent, callback)
			}
		} else {
			k.IterateAllCertificate(ctx, callback)
		}
	} else if params.ContentType != "" && params.Content != "" {
		requestContent, err := types.NewRequestContent(params.ContentType, params.Content)
		if err != nil {
			return 0, nil, err
		}
		filteredCertificates = k.GetCertificatesByContent(ctx, requestContent)
	} else {
		k.IterateAllCertificate(ctx, callback)
	}

	// Post-processing
	total := uint64(len(filteredCertificates))

	start, end := client.Paginate(len(filteredCertificates), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		filteredCertificates = []types.Certificate{}
	} else {
		filteredCertificates = filteredCertificates[start:end]
	}

	return total, filteredCertificates, nil
}

// RevokeCertificate revokes a certificate.
func (k Keeper) RevokeCertificate(ctx sdk.Context, certificate types.Certificate, revoker sdk.AccAddress) error {
	if !k.IsCertifier(ctx, revoker) {
		return types.ErrUnqualifiedRevoker
	}
	return k.DeleteCertificate(ctx, certificate)
}

// GetCertifiedIdentities returns a list of addresses certified as identities.
func (k Keeper) GetCertifiedIdentities(ctx sdk.Context) []sdk.AccAddress {
	identities := []sdk.AccAddress{}
	k.IterateCertificatesByType(ctx, types.CertificateTypeIdentity, func(certificate types.Certificate) (stop bool) {
		addr, _ := sdk.AccAddressFromBech32(certificate.RequestContent().RequestContent)
		identities = append(identities, addr)
		return false
	})
	return identities
}
