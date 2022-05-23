package recon

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func Routers() map[string]common.Address {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(list), &m)
	if err != nil {
		panic(err)
	}
	result := make(map[string]common.Address)
	for name, addy := range m {
		a := common.HexToAddress(addy)
		name = strings.ToLower(name)
		for existingName, existingAddy := range result {
			if name == existingName {
				panic("name registered twice")
			}
			if a == existingAddy {
				panic("address registered twice")
			}
		}
		result[name] = a
	}
	return result
}

const list = `{
  "agile": "0x9AadB4A3BFaCff8aa60a2c63735Cb8B94De7C57d",
  "annex": "0x77befde82eba4bdc4d3e4a853bf3ea4ffb49dd58",
  "cougardex": "0x3c1997d8738dcaB7Ed099105fCd61A9fe5f351Dd",
  "crodex": "0xeC0A7a0C2439E8Cb67b992b12ecd020Ea943c7Be",
  "cronaswap": "0xcd7d16fB918511BF7269eC4f48d61D79Fb26f918",
  "crowfi": "0xd30d3aC04E2325E19A2227cfE6Bc860376Ba20b1",
  "cyborgswap": "0x5bFc95C3BbF50579bD57957cD074fa96a4d5fF9F",
  "ducky": "0x28a10fE91d4a8D0637999a903eEf9Ad5b1D9947C",
  "elkfinance": "0xdB02A597b283eACb9436Cd2a2d15039a11A3299d",
  "empiredex": "0xdADaae6cDFE4FA3c35d54811087b3bC3Cd60F348",
  "kryptodex": "0xDE25A2d1AD885aF4df667132f2214b920b225e52",
  "meerkat": "0x145677FC4d9b8F19B5D56d1820c48e0443049a30",
  "photon": "0x69004509291F4a4021fA169FafdCFc2d92aD02Aa",
  "smolswap": "0x8118DD9fED86523Bf724e2EC5f56055Da0668AF4",
  "stabil": "0xB2C8D1e0626245A06BD5901063aB52f86F772BC4",
  "swapp": "0x600d0b65C2A25b64C9b517A43B7a44592448d285",
  "vvs": "0x145863Eb42Cf62847A6Ca784e6416C1682b1b2Ae"
}`
