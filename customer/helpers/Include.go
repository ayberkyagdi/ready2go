package helpers

import (
	"path/filepath"
	"strings"
	"fmt"
)

func Include(path string) []string {
	
	if strings.Contains(path,"store") || strings.Contains(path,"customer"){
		files, _ := filepath.Glob("customer/templates/"+path+"/*.html")
		return files
	}else{
		fmt.Println("admin buraya da gel di !")
		files, _ := filepath.Glob("admin/views/templates/*.html")
		path_files, _ := filepath.Glob("admin/views/dashboard/"+path+"/*.html")
		for _, file := range path_files {
			files = append(files, file)
		}
		return files
	}
	
}
