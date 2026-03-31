//Extracting Hashtags Automatically 

package utils

import "regexp"

func ExtractHashtags(content string) []string {
	re := regexp.MustCompile(`#\w+`)
	return re.FindAllString(content, -1)
}