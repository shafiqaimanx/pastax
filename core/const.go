package core

const (
	
	// Output Directories
	RECENT_DIR_PATH = "pastax/recent"
	USERS_DIR_PATH = "pastax/users"

	// Scrape Regex
	PUBLIC_RGX = `<span class="status -public"></span><a href="[^"]+">`
	USER_PASTE_RGX = `<span class="status -public" title="[^"]+"></span>\s+<a href="[^"]+">`
	HREF_RGX = `/[a-zA-Z0-9]+"`
	
	// PasteBin Urls
	ARCHIVE_URL = "https://pastebin.com/archive"
	USERNAME_URL = "https://pastebin.com/u/"
	DOWNLOAD_URL = "https://pastebin.com/raw/"

)