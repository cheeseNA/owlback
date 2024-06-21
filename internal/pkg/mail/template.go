package mail

import "fmt"

func NewFailedToCrawlMail(siteURL string) string {
	return fmt.Sprint(`<html>
<head>
</head>
<body>
<h1>Failed to crawl</h1>
<p>Failed to crawl the site: <a href="`, siteURL, `">`, siteURL, `</a></p>
<p>Please check the site and try again.</p>
</body>
</html>`)
}

func NewTriggeredMail(siteURL, condition string) string {
	return fmt.Sprint(`<html>
<head>
</head>
<body>
<h1>Triggered</h1>
<p>Your tracked site has been updated: <a href="`, siteURL, `">`, siteURL, `</a></p>
<p>The condition you set: <code>`, condition, `</code> has been met.</p>
</body>
</html>`)
}
