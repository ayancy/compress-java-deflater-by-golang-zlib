///////////////////////////////////////////////////////////////////////////////////////////////////
/*                Tool using when go server reqest java server                                   */
/*Runtime environment:  GO version go1.8.3 linux/amd64,JDK version:1.8.                          */
/*Function: text compressed by go-zlib which could be decompressed by java.util.zip.Deflater.    */
/*create at 2018.05.07                                                                           */
/*Author: Yancy Lee                                                                              */
/*warning: Please identify before using,for not test with all character set.                     */
///////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
      "os" 
      "io"
      "bytes"
      "compress/zlib"  
      //"fmt" 
      "hash/adler32"
      "encoding/binary"
 
)
 
func main() {
  arr:="A,b,.,\u0001,_,-,1,\n"
  goZlib2javaDeflater(arr)

}
 
//input arr string:your own text to compress ,tested sucessful of "A,b,.,\u0001,_,-,1,\n".
//output string: the compress text;
//       err: if  not nil,failed to compress.
func goZlib2javaDeflater(arr string )(string,error) {
		
  var err error
  var b bytes.Buffer
  
  input:= []byte(arr) 
  ret := adler32.Checksum(input)
  compressor, err := zlib.NewWriterLevel(&b, 9)
  if err != nil {  
     return  "",err
  }  
  compressor.Write(input)  
  compressor.Flush()
  defer compressor.Close()
  
  
  b2:=b.Bytes()
  lenc:=len(b.Bytes())
  
  buf := new(bytes.Buffer)
  
  binary.Write(buf, binary.BigEndian, ret)
  b3:=buf.Bytes() 
  b2[2]=b2[2]+1   // not check the protocol,just by the experimental rule 
  b2[lenc-4]=b3[0]//filled last 4byte by the  BigEndian adler32 of text
  b2[lenc-3]=b3[1]
  b2[lenc-2]=b3[2]
  b2[lenc-1]=b3[3]
  
  
  bs:=b.String()
  // write result to decompress by java.util.zip.Deflater for testing
  f0, _ := os.Create("zlibtest0507") 
  _, _ = io.WriteString(f0, bs) 
  
  return bs,err
       
}
