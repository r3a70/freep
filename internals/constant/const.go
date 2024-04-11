package constant

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	GRAY   = "\033[37m"
	WHITE  = "\033[97m"
)

// Max file size is 10 MB
const ALLOW_FILE_SZIE = 1999 * 1024 * 1024

const MULTY_PART_MAX_SIZE = 2 << 20
