/*

   Copyright 2016 Wenhui Shen <www.webx.top>

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

*/

package codec

import "crypto/des"

var (
	_ Codec = NewDESCBC()
	_ Codec = NewDESECB()
)

func GenDESKey(key []byte) []byte {
	size := des.BlockSize
	kkey := make([]byte, 0, size)
	ede2Key := []byte(key)
	length := len(ede2Key)
	if length > size {
		kkey = append(kkey, ede2Key[:size]...)
	} else {
		div := size / length
		mod := size % length
		for i := 0; i < div; i++ {
			kkey = append(kkey, ede2Key...)
		}
		kkey = append(kkey, ede2Key[:mod]...)
	}
	return kkey
}
