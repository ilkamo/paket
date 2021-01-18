//important: You can edit this file. However, you need to know what you are doing.
// *panic* may occur.

package main

import (
	paket "github.com/SeanTolstoyevski/paket/pengine"
)

//The map vault for datas. The init function writing the required data.
var Data = make(paket.Datas)

// The name of the folder from which the files were was taken. Information is writing by init.
var foldername string

func init() {
	foldername = "datas"
	Data["Gustaf Oloveson - Gnossienne No.1.ogg"] = paket.Values{"0", "3254444", "3254428", "3254444", "b9c0ad65aaa316ae74cdf55b657aabe737fd27ee925fdd5970114b2a68151aea", "8b2d0c9a455d4fb4b352b7913ce88c60c44cd71e962246b054b42afc25fb4ff6"}
	Data["start.wav"] = paket.Values{"3254444", "3300418", "45958", "45974", "2d868326c7e21fa744d95911320cfc9b52b76a6f17819ba3a301b3d68c62bb68", "ee8c829ced6ed4d01966bf4e1a6ce532985c53f483254529480f74d05441786a"}
}