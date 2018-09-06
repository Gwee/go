package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//unit tests can be done by mocking the filesystem by using afero for example: https://github.com/spf13/afero
func Test_searchWord_FindMockFile_MockFileFound(t *testing.T) {
	//afero can be used here to mock filesystem
}
func Test_searchWord_PassEmptyPath_EmptyPathErrorReturned(t *testing.T) {
	var buf []byte
	_, err := searchWord("mock", "", &buf)
	if err == nil {
		t.Error(err)
		t.Fatal("Empty path passed and should not be allowed")
	}
}
func Test_searchWord_PassEmptyTerm_EmptyTermErrorReturned(t *testing.T) {
	var buf []byte
	_, err := searchWord("", ".", &buf)
	if err == nil {
		t.Error(err)
		t.Fatal("Empty path passed and should not be allowed")
	}
}


//integration tests

//this tests the basic functionality for a matching file
func Test_searchWord_FindMatchingFile_MatchingFileReturned(t *testing.T) {
	term := "findme"
	fileName := "findme"
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	//create test Dir
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	//create test File
	_, err = os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		t.Error("File could not be created")
	}
	var buf []byte
	result, err := searchWord(term, dirPath, &buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Results found in test: " + result)
	if !strings.Contains(result, term) {
		t.Fatal("Test failed! Expected " + fileName + " but found nothing")
	}
}

//test the ability of our method to find a substring within the middle of the name
func Test_searchWord_FindSubStringInName_MatchingFileReturned(t *testing.T) {
	term := "findme"
	fileName := "IAmAfindmeSubString.txt"
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	_, err = os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		t.Error("File could not be created")
	}
	var buf []byte
	result, err := searchWord(term, dirPath, &buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Results found in test: " + result)
	if !strings.Contains(result, term) {
		t.Fatal("Test failed! Expected " + fileName + " but found nothing")
	}

}

//test the ability of our method to find a matching directory name
func Test_searchWord_FindMatchingDir_MatchingDirFound(t *testing.T) {
	term := "findme"
	subDir := "findme"
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	err = os.Mkdir(filepath.Join(dirPath, subDir), os.ModePerm)
	if err != nil {
		t.Error("Dir could not be created")
	}
	var buf []byte
	result, err := searchWord(term, dirPath, &buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Results found in test: " + result)
	if !strings.Contains(result, term) {
		t.Fatal("Test failed! Expected " + subDir + " but found nothing")
	}
}

//test the method to see that nothing is returned when nothing is found
func Test_searchWord_NoMatchingFiles_EmptyResultReturned(t *testing.T) {
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	var buf []byte
	result, err := searchWord("nothing", dirPath, &buf)
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Fatal("No result expected but found a match")
	}
}

//test the method to find a matching file within a nested directory structure. (Optional): Add a more complex structure to be more wide and not only depth
func Test_searchWord_FindNestedFile_NestedFileFound(t *testing.T) {
	term := "findnested"
	fileName := "findnested.txt"
	nestedDirPath := "Findme/Findme1/Findme2/Findme3/Findme4/Findme5/Findme6/Findme7/Findme8"
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	//create test dir
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	//created nested dir structure
	nestedDirPath, err = filepath.Abs(filepath.Join(dirPath, nestedDirPath))
	if err != nil {
		t.Error("Couldn't get absolute path for nested directory")
	}
	err = os.MkdirAll(nestedDirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create nested directory")
	}

	//create file in nested dir
	_, err = os.Create(filepath.Join(nestedDirPath, fileName))
	if err != nil {
		t.Error("File could not be created")
	}
	//test function
	var buf []byte
	result, err := searchWord(term, dirPath, &buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Results found in test: " + result)
	if !strings.Contains(result, term) {
		t.Fatal("Test failed! Expected " + fileName + " but found nothing")
	}
}

//test the method to see that when a matching file exists within a matching dir, both results are returned
func Test_searchWord_MatchingFileWithinMatchingDir_BothResultsReturned(t *testing.T) {
	term := "findme"
	subDir := "findme"
	fileName := "findme"
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	//create test Dir
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	//crete sub Dir
	subDirPath, err := filepath.Abs(filepath.Join(dirPath,subDir))
	err = os.Mkdir(subDirPath, os.ModePerm)
	if err != nil {
		t.Error("Sub dir could not be created")
	}
	//create file in nested dir
	_, err = os.Create(filepath.Join(subDirPath, fileName))
	if err != nil {
		t.Error("File could not be created")
	}
	var buf []byte
	result, err := searchWord(term, dirPath, &buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Results found in test: " + result)
	if strings.Count(result,"\n") != 2 {
		t.Fatal("Test failed! Didn't find 2 results")
	}
}

//test the method to find a match that have special characters
func Test_searchWord_FindSpecialCharFile_SpecialCharFileFound(t *testing.T) {
	term := "_+findme!"
	fileName := "!@#$%^&*()_+findme!@#$%^&*()_+"
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	_, err = os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		t.Error("File could not be created")
	}
	var buf []byte
	result, err := searchWord(term, dirPath, &buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Results found in test: " + result)
	if !strings.Contains(result, term) {
		t.Fatal("Test failed! Expected " + fileName + " but found nothing")
	}
}

//test the method to find case sensitive files. (Optional): Can add OS detection and test according to OS (Linux/Windows)
func Test_searchWord_FindCaseSensitiveFile_FileFound(t *testing.T) {
	//can add OS detection here
	term := "findMe"
	fileName := "findMe"
	testingDir := "TestDir"
	dirPath, err := filepath.Abs(testingDir)
	if err != nil {
		t.Error("Couldn't get absolute path")
	}
	//remove test Dir if exists
	if _, err := os.Stat(dirPath); err == nil {
		os.RemoveAll(dirPath)
	}
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		t.Error("Couldn't create test directory")
	}
	_, err = os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		t.Error("File could not be created")
	}
	var buf []byte
	result, err := searchWord(term, dirPath, &buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Results found in test: " + result)
	if !strings.Contains(result, term) {
		t.Fatal("Test failed! Expected " + fileName + " but found nothing")
	}
}


//func Test_searchWord_NonExistingPath_PathNotFoundErrorReturned(t *testing.T) {
//	fakeDir:= "this_dir_does_not_exist"
//	dirPath, err := filepath.Abs(fakeDir)
//	if nil != err {
//		t.Error("Couldn't get absolute path")
//	}
//	//check if fake dir exists, if so remove it
//	if _, err := os.Stat(dirPath); err == nil {
//		os.RemoveAll(dirPath)
//	}
//	var buf []byte
//	_, err = searchWord("mock", dirPath, &buf)
//	if err != nil {
//		t.Log(err)
//	}
//}

//
//func Test_searchWord_CheckInvalidPermissions_InvalidPermissionErrorReturned(t *testing.T) {
//	term := "findme"
//	fileName := "findme"
//	testingDir := "TestDir"
//	dirPath, err := filepath.Abs(testingDir)
//	if err != nil {
//		t.Error("Couldn't get absolute path")
//	}
//	//remove test Dir if exists
//	if _, err := os.Stat(dirPath); err == nil {
//		os.RemoveAll(dirPath)
//	}
//	err = os.Mkdir(dirPath, os.ModePerm)
//	if err != nil {
//		t.Error("Couldn't create test directory")
//	}
//	_, err = os.Create(filepath.Join(dirPath, fileName))
//	if err != nil {
//		t.Error("File could not be created")
//	}
//	//remove read permissions from file
//	err = os.Chmod(dirPath, 0333)
//	if err != nil {
//		t.Error(err)
//	}
//	var buf []byte
//
//	fmt.Println("BEFORE")
//	result, err1 := searchWord(term, dirPath, &buf)
//	if err1 != nil{
//		fmt.Println("IN ERROR")
//	}
//	fmt.Println("AFTER")
//
//
//
//	t.Log(err1)
//	fmt.Println("Results found in test: " + result)
//	if !strings.Contains(result, "permission denied") {
//		t.Fatal("Test failed! Expected permission denied error but did not get")
//	}
//}




//use this for mixed dir struct
//for i:=1; i<10; i++ {
//if i%2==0 {
//subDirPath, err = filepath.Abs(filepath.Join(filepath.Join(filepath.Join(subDir,subDir+strconv.Itoa(i+1),subDir+strconv.Itoa(i+2))),subDir+strconv.Itoa(i+3)))
//if err != nil {
//t.Error("Couldn't get absolute path")
//}
//} else {
//subDirPath, err = filepath.Abs(subDir+strconv.Itoa(i))
//}
//t.Log(subDirPath)
//if i ==8{
//nestedDirPath=subDirPath
//}
//err = os.MkdirAll(filepath.Join(dirPath,subDirPath),os.ModePerm)
//}
//t.Log("nested path = " +nestedDirPath)
