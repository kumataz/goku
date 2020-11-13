/*   
// golang internel function
import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

func main() {
    data := []byte("kumata")

    sum256 := sha256.Sum256(data)
    fmt.Println(hex.EncodeToString(sum256[:]))

}
*/


/*
// golang source code

TODO output sha256中轮计算的第13轮data和w0处理的第5轮data,
add function:  Getww5()   Gets13()
add tool function： StrFormat32()
*/
package crypto

import (
	"fmt"
	"encoding/hex"
	"strings"
)

// The size of a SHA256 checksum in bytes.
const Size = 32

const (
	chunk     = 64
	init0     = 0x6A09E667
	init1     = 0xBB67AE85
	init2     = 0x3C6EF372
	init3     = 0xA54FF53A
	init4     = 0x510E527F
	init5     = 0x9B05688C
	init6     = 0x1F83D9AB
	init7     = 0x5BE0CD19 
)


// digest represents the partial evaluation of a checksum.
type digest struct {
	h     [8]uint32
	x     [chunk]byte
	nx    int
	len   uint64
}


var _K = []uint32{
	0x428a2f98,
	0x71374491,
	0xb5c0fbcf,
	0xe9b5dba5,
	0x3956c25b,
	0x59f111f1,
	0x923f82a4,
	0xab1c5ed5,
	0xd807aa98,
	0x12835b01,
	0x243185be,
	0x550c7dc3,
	0x72be5d74,
	0x80deb1fe,
	0x9bdc06a7,
	0xc19bf174,
	0xe49b69c1,
	0xefbe4786,
	0x0fc19dc6,
	0x240ca1cc,
	0x2de92c6f,
	0x4a7484aa,
	0x5cb0a9dc,
	0x76f988da,
	0x983e5152,
	0xa831c66d,
	0xb00327c8,
	0xbf597fc7,
	0xc6e00bf3,
	0xd5a79147,
	0x06ca6351,
	0x14292967,
	0x27b70a85,
	0x2e1b2138,
	0x4d2c6dfc,
	0x53380d13,
	0x650a7354,
	0x766a0abb,
	0x81c2c92e,
	0x92722c85,
	0xa2bfe8a1,
	0xa81a664b,
	0xc24b8b70,
	0xc76c51a3,
	0xd192e819,
	0xd6990624,
	0xf40e3585,
	0x106aa070,
	0x19a4c116,
	0x1e376c08,
	0x2748774c,
	0x34b0bcb5,
	0x391c0cb3,
	0x4ed8aa4a,
	0x5b9cca4f,
	0x682e6ff3,
	0x748f82ee,
	0x78a5636f,
	0x84c87814,
	0x8cc70208,
	0x90befffa,
	0xa4506ceb,
	0xbef9a3f7,
	0xc67178f2,
}

//var sh []string
var preS13 = ""
var preww5 = ""

func blockGeneric(dig *digest, p []byte) {
	var w [64]uint32
	h0, h1, h2, h3, h4, h5, h6, h7 := dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7]
	for len(p) >= chunk {
		// Can interlace the computation of w with the
		// rounds below if needed for speed.
		for i := 0; i < 16; i++ {
			j := i * 4
			w[i] = uint32(p[j])<<24 | uint32(p[j+1])<<16 | uint32(p[j+2])<<8 | uint32(p[j+3])	
			

			if i == 15{
				for count:=0;count<11;count++{
					ww1 := fmt.Sprintf(StrFormat32(fmt.Sprintf("%x",w[5+count]),"0",8))
					preww5 += ww1
				}	
			}

		}
		
		for i := 16; i < 64; i++ {
			v1 := w[i-2]
			t1 := (v1>>17 | v1<<(32-17)) ^ (v1>>19 | v1<<(32-19)) ^ (v1 >> 10)
			v2 := w[i-15]
			t2 := (v2>>7 | v2<<(32-7)) ^ (v2>>18 | v2<<(32-18)) ^ (v2 >> 3)
			w[i] = t1 + w[i-7] + t2 + w[i-16]
			
			if i < 21{
				ww2 := fmt.Sprintf(StrFormat32(fmt.Sprintf("%x",w[i]),"0",8))
				preww5 += ww2
			}

		}

		a, b, c, d, e, f, g, h := h0, h1, h2, h3, h4, h5, h6, h7
		
		for i := 0; i < 64; i++ {
			t1 := h + ((e>>6 | e<<(32-6)) ^ (e>>11 | e<<(32-11)) ^ (e>>25 | e<<(32-25))) + ((e & f) ^ (^e & g)) + _K[i] + w[i]

			t2 := ((a>>2 | a<<(32-2)) ^ (a>>13 | a<<(32-13)) ^ (a>>22 | a<<(32-22))) + ((a & b) ^ (a & c) ^ (b & c))

			h = g
			g = f
			f = e
			e = d + t1
			d = c
			c = b
			b = a
			a = t1 + t2
			
			if i == 12{
				
				aa := StrFormat32(fmt.Sprintf("%x",a),"0",8)
				bb := StrFormat32(fmt.Sprintf("%x",b),"0",8)
				cc := StrFormat32(fmt.Sprintf("%x",c),"0",8)
				dd := StrFormat32(fmt.Sprintf("%x",d),"0",8)
				ee := StrFormat32(fmt.Sprintf("%x",e),"0",8)
				ff := StrFormat32(fmt.Sprintf("%x",f),"0",8)
				gg := StrFormat32(fmt.Sprintf("%x",g),"0",8)
				hh := StrFormat32(fmt.Sprintf("%x",h),"0",8)
				preS13 = fmt.Sprintf("%s%s%s%s%s%s%s%s",aa,bb,cc,dd,ee,ff,gg,hh)
				
			}

		}

		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
		h5 += f
		h6 += g
		h7 += h

		p = p[chunk:]
	}

	dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7] = h0, h1, h2, h3, h4, h5, h6, h7
}


func putUint32(x []byte, s uint32) {
	_ = x[3]
	x[0] = byte(s >> 24)
	x[1] = byte(s >> 16)
	x[2] = byte(s >> 8)
	x[3] = byte(s)
}


func putUint64(x []byte, s uint64) {
	_ = x[7]
	x[0] = byte(s >> 56)
	x[1] = byte(s >> 48)
	x[2] = byte(s >> 40)
	x[3] = byte(s >> 32)
	x[4] = byte(s >> 24)
	x[5] = byte(s >> 16)
	x[6] = byte(s >> 8)
	x[7] = byte(s)
}


func (d *digest) Reset() {

	d.h[0] = init0
	d.h[1] = init1
	d.h[2] = init2
	d.h[3] = init3
	d.h[4] = init4
	d.h[5] = init5
	d.h[6] = init6
	d.h[7] = init7
	
	d.nx = 0
	d.len = 0
}


func (d *digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	d.len += uint64(nn)
	if d.nx > 0 {
		n := copy(d.x[d.nx:], p)
		d.nx += n
		if d.nx == chunk {
			blockGeneric(d, d.x[:])
			d.nx = 0
		}
		p = p[n:]
	}
	if len(p) >= chunk {
		n := len(p) &^ (chunk - 1)
		blockGeneric(d, p[:n])
		p = p[n:]
	}
	if len(p) > 0 {
		d.nx = copy(d.x[:], p)
		//fmt.Println(d.x[:])
	}
	return
}


func (d *digest) checkSum() [Size]byte {
	len := d.len
	// Padding. Add a 1 bit and 0 bits until 56 bytes mod 64.
	var tmp [64]byte
	tmp[0] = 0x80
	if len%64 < 56 {
		d.Write(tmp[0 : 56-len%64])
	} else {
		d.Write(tmp[0 : 64+56-len%64])
	}

	// Length in bits.
	len <<= 3
	putUint64(tmp[:], len)
	d.Write(tmp[0:8])

	if d.nx != 0 {
		panic("d.nx != 0")
	}

	var digest [Size]byte

	putUint32(digest[0:], d.h[0])
	putUint32(digest[4:], d.h[1])
	putUint32(digest[8:], d.h[2])
	putUint32(digest[12:], d.h[3])
	putUint32(digest[16:], d.h[4])
	putUint32(digest[20:], d.h[5])
	putUint32(digest[24:], d.h[6])
	putUint32(digest[28:], d.h[7])

	return digest
}


func Sum256(data []byte) [Size]byte {
	var d digest
	d.Reset()
	d.Write(data)
	return d.checkSum()
}

/*--------add---------*/


func StrFormat32(raw ,sub string,slen int) string {

     switch {
     case len(raw) == slen:
         return raw
     case len(raw) > slen:
         return fmt.Sprintf("%.32s",raw[:slen])
     case len(raw) < slen:
         return fmt.Sprintf("%s%s",strings.Repeat(sub,slen-len(raw)),raw)
     default:
         /*should never meet here*/
         return raw
     }
 }



func Getww5(infos string)string{
	preww5 = ""
	data, _ := hex.DecodeString(infos)
	Sum256(data)
	return preww5
}


func Gets13(infos string)string{
	preS13 = ""
	data, _ := hex.DecodeString(infos)
	Sum256(data)
	return preS13
}


func Presha(infos string)string{
	preS13 = ""
	preww5 = ""
	data, _ := hex.DecodeString(infos)
	Sum256(data)
	return preww5 + preS13
}

/*

func main(){
	//test := "616263"
	test := "05fa169729f18f2da892825de22991503bd313e0fbf16e8e88cfc6d1ae821465b26e9412000000009c1b15930000000000000000000000"
//	data, _ := hex.DecodeString(test)
//	Sum256(data)
//	fmt.Printf("HashResult: %x\n\n", h) 
	ww5 := Getww5(test)
	s13 := Gets13(test)
}

*/





