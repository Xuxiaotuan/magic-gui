package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getAllMsgDbnames(folderPath string) ([]string, error) {
	var filenames []string
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return filenames, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".db") && strings.HasPrefix(file.Name(), "msg") {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames, nil
}

func getContactDbnames(folderPath string) ([]string, error) {
	var filenames []string
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return filenames, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".db") {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames, nil
}

func moveFileAndSqlcipher(folderPath string, key string) {
	results := make(map[string]string)

	sql := fmt.Sprintf("PRAGMA key = \"x'%s'\";\nPRAGMA cipher_compatibility = 3;\nPRAGMA kdf_iter = 64000;\nPRAGMA cipher_page_size = 1024;", key)
	filenames, err := getAllMsgDbnames(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range filenames {
		fmt.Println(filename)
		results[filename] = sql
	}

	for dbPath, sql := range results {
		plainPath := fmt.Sprintf("plain_%s", filepath.Base(dbPath))
		os.Remove(plainPath)

		fmt.Printf("Decrypting %s into %s ...\n", dbPath, plainPath)

		cmd := exec.Command("sqlcipher", dbPath)

		cmd.Stdin = strings.NewReader(fmt.Sprintf(`%s
	ATTACH DATABASE '%s' AS plaintext KEY '';
	SELECT sqlcipher_export('plaintext');
	DETACH DATABASE plaintext;`, sql, plainPath))

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}

		fmt.Println(string(out))
	}
}
