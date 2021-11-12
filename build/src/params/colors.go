package params

// Colors to be used in logger output

var colorReset string = "\033[0m"
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"
var colorYellow string = "\033[33m"
var colorBlue string = "\033[34m"

var ErrorLog string = colorRed + "ERROR: " + colorReset
var WarnLog string = colorYellow + "WARN: " + colorReset
var InfoLog string = colorBlue + "INFO: " + colorReset
var SuccessLog string = colorGreen + "SUCCESS: " + colorReset