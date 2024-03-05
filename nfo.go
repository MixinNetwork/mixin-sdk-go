package kernel

import "github.com/MixinNetwork/mixin/crypto"

func IsNonFungibleAsset(id string) bool { // only support native, not support bridged yet
	am := map[string]bool{
		"2b9be927-106f-368c-8971-443624b9df84":                             true, // 0x3c8c161a18ae2c8b14fda1216fff7da88c419b5d
		"1700941284a95f31b25ec8c546008f208f88eee4419ccdcdbe6e3195e60128ca": true,
	}
	mixinId := crypto.Sha256Hash([]byte(id)).String()
	return am[id] || am[mixinId]
}