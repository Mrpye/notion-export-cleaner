package notion

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var uuids map[string]string = make(map[string]string)

//******************************************
//extract the uuid from the filename
//the uuid is the last part of the filename
//the uuid is 36 characters long
//******************************************
func getUUID(filename string) string {
	parts := strings.Split(filename, " ")
	//Get the last part of the filename
	if len(parts) > 0 {
		data := parts[len(parts)-1]
		if len(" "+data) == len(" 3bf9a712efcd43a696ef5eb0b209c943") {
			return " " + data
		}
	}
	return ""
}

//************************************************
//get file name without the extension from a path
//get the uuid from the file name
//remove the uuid from the file name
//check if the uuid is already in the map
//if it is not in the map then add it to the map
//return the cleaned file name and the uuid
//************************************************
func getCleanFileName(path string, uuids map[string]string) (string, string) {
	file_no_ext := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	uuid := getUUID(file_no_ext)
	cleaned_file_name := strings.ReplaceAll(path, uuid, "")
	//loop through map to check if the uuid is already in the map
	for _, v := range uuids {
		cleaned_file_name = strings.ReplaceAll(cleaned_file_name, v, "")
	}

	return cleaned_file_name, uuid
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

//*************************************
//unzip a zip file
//remove the uuid from the file names
//return a list of the file names
//*************************************
func UnzipCleanFileNames(src string, dest string) ([]string, error) {

	//************************
	//check if the file exists
	//************************
	if !FileExists(src) {
		return nil, fmt.Errorf("file does not exist")
	}

	//*****************************************
	//Create the directory if it does not exist
	//*****************************************
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return nil, err
	}

	var filenames []string

	//*****************
	//Open the zip file
	//*****************
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	//**************************************
	//Loop through the files in the zip file
	//**************************************
	for _, f := range r.File {
		// Store filename/path for returning and using later on
		temp_filename, uuid := getCleanFileName(f.Name, uuids)
		if uuid != "" {
			uuids[uuid] = uuid
		}

		fpath := filepath.Join(dest, temp_filename)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	//fix the urls in files
	readFiles(dest, uuids)
	return filenames, nil
}

//******************************************
//recurse through a directory and read file
//replace the uuid in the file
//save the file
//******************************************
func readFiles(root string, uuids map[string]string) {
	filepath.Walk(root, func(root string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		//get the extension of the file
		ext := filepath.Ext(info.Name())
		if strings.ToLower(ext) == ".md" {
			replaceValue(root, uuids)
		}

		return nil
	})
}

//*******************************************
//open a file and replace a value and save it
//*******************************************
func replaceValue(path string, uuids map[string]string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	tmp_data := string(data)
	for _, v := range uuids {
		replace_space := strings.ReplaceAll(v, " ", "%20")
		tmp_data = strings.Replace(tmp_data, replace_space, "", -1)
	}

	err = ioutil.WriteFile(path, []byte(tmp_data), 0)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
}
