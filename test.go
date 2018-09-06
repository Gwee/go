//package main
//
//import (
//	"errors"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"log"
//	"os"
//	"path/filepath"
//
//	"strings"
//	"testing"
//
//)
//type mockFiles []os.FileInfo
//
////var fs fileSystem = osFS{}
//
//type fileSystem interface {
//	Open(name string) (File, error)
//	Stat(name string) (os.FileInfo, error)
//}
////type file interface {
////	io.Closer
////	io.Reader
////	io.ReaderAt
////	io.Seeker
////	Stat() (os.FileInfo, error)
////}
//type File interface {
//	io.Closer
//	io.Reader
//	io.ReaderAt
//	io.Seeker
//	io.Writer
//	io.WriterAt
//
//	Name() string
//	Readdir(count int) ([]os.FileInfo, error)
//	Readdirnames(n int) ([]string, error)
//	Stat() (os.FileInfo, error)
//	Sync() error
//	Truncate(size int64) error
//	WriteString(s string) (ret int, err error)
//}
//type osFS struct {}
//
//func (osFS) Stat(name string) (os.FileInfo, error) {
//	panic("implement me")
//}
//
//type mockFS struct {
//	osFS
//	reportErr bool
//	reportSize int64
//}
//
//type mockFileInfo struct {
//	os.FileInfo
//	size int64
//	name string
//	mode mockFileMode
//}
//
//type mockFileMode struct {
//	os.FileMode
//	mode uint32
//	isThisADir bool
//}
//type Afero struct {
//	Fs
//}
//
//type Fs interface {
//	// Create creates a file in the filesystem, returning the file and an
//	// error, if any happens.
//	Create(name string) (File, error)
//
//	// Mkdir creates a directory in the filesystem, return an error if any
//	// happens.
//	Mkdir(name string, perm os.FileMode) error
//
//	// MkdirAll creates a directory path and all parents that does not exist
//	// yet.
//	MkdirAll(path string, perm os.FileMode) error
//
//	// Open opens a file, returning it or an error, if any happens.
//	Open(name string) (File, error)
//
//	// OpenFile opens a file using the given flags and the given mode.
//	OpenFile(name string, flag int, perm os.FileMode) (File, error)
//
//	// Remove removes a file identified by name, returning an error, if any
//	// happens.
//	Remove(name string) error
//
//	// RemoveAll removes a directory path and any children it contains. It
//	// does not fail if the path does not exist (return nil).
//	RemoveAll(path string) error
//
//	// Rename renames a file.
//	Rename(oldname, newname string) error
//
//	// Stat returns a FileInfo describing the named file, or an error, if any
//	// happens.
//	Stat(name string) (os.FileInfo, error)
//
//	// The name of this FileSystem
//	Name() string
//
//	//Chmod changes the mode of the named file to mode.
//	Chmod(name string, mode os.FileMode) error
//}
//
//func (a Afero) ReadDir(dirname string) ([]os.FileInfo, error) {
//	return ReadDir(a.Fs, dirname)
//}
//
//func ReadDir(fs Fs, dirname string) ([]os.FileInfo, error) {
//	f, err := fs.Open(dirname)
//	if err != nil {
//		return nil, err
//	}
//	list, err := f.Readdir(-1)
//	f.Close()
//	if err != nil {
//		return nil, err
//	}
//	return list, nil
//}
//
//
////func (m mockFileInfo) Size() int64 { return m.size }
////
////func (m mockFS) Stat(name string) (os.FileInfo, error) {
////	if m.reportErr {
////		return nil, os.ErrNotExist
////	}
////	return mockFileInfo{size: m.reportSize}, nil
////}
//////func (osFS) Open(name string) (file, error) {return os.Open(name)}
////func (m mockFileMode) IsDir() bool {return m.isThisADir}
////func (m mockFileInfo) Name() string {return m.name}
////func (m mockFileInfo) Mode() mockFileMode {return m.mode}
//
//func main() {
//	searchWord_FindMatchingDirMock_MatchingDirReturnedMock()
//	path := "/Users/guy/Desktop"
//	var buf []byte
//	itemsFound, err := searchWord("test",path, &buf)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(itemsFound)
//}
////func getSize(name string)(int64, error){
////	stat, err := fs.Stat(name)
////	if err != nil {
////		return 0, err
////	}
////	return stat.Size(), nil
////}
//func searchWord(term string, path string, buffer * []byte)(string, error){
//	if  len(path) == 0  {
//		return "",errors.New("Please enter a valid path")
//	}
//	if len(term) == 0 {
//		return "",errors.New("Please enter a valid search term")
//	}
//	files, err := ioutil.ReadDir(path)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, file := range files{
//		path := filepath.Join(path, file.Name())
//		result,err := filepath.Abs(path)
//		if err != nil {
//			log.Fatal(err)
//		}
//		if strings.Contains(file.Name(), term) {
//			//fmt.Println(result)
//			*buffer = append(*buffer,result+"\n"...)
//		}
//		if file.Mode().IsDir() {
//			searchWord(term, path, buffer)
//		}
//	}
//	return string(*buffer),err
//}
//
////unit tests
//func searchWord_PassEmptyPath_EmptyPathErrorReturned(t *testing.T){
//
//}
//func searchWord_PassEmptyTerm_EmptyTermErrorReturned(t *testing.T){
//
//}
//
//func searchWord_FindMatchingDirMock_MatchingDirReturnedMock(){
//	//oldFs := fileSystem()
//	//mfs := &mockFS{}
//	//fs = mfs
//	//defer func() {
//	//	fs = oldFs
//	//}()
//	var buf []byte
//	output, err := searchWord("test", "/usr/test",&buf)
//	//output, err := getSize("guyguy.go")
//
//	if err != nil{
//		log.Fatal(err)
//	}
//	fmt.Println(output)
//}
//
////integration tests
//func searchWord_FindMatchingFile_MatchingFileReturned(t *testing.T){
//	//chdir, check if file exists, if so delete, create file, test by finding it, remove file
//	file,err:= os.Create("findme.txt")
//	if err !=nil {
//		log.Fatal(err)
//	}
//	fmt.Println(file)
//
//}
//func searchWord_FindMatchingDir_MatchingDirReturned(t *testing.T){
//
//}
//func searchWord_NoMatchingFiles_EmptyResultReturned(t *testing.T){
//
//}
//func searchWord_FindNestedFile_NestedFileFoundReturned(t *testing.T){
//
//}
//func searchWord_MatchingFileWithinMatchingDir_NoDuplicateResultReturned(t *testing.T){
//
//}
//func searchWord_NonExistingPath_MissingPathErrorReturned(t *testing.T){
//
//}
////Optional
//func searchWord_CheckInvalidPermissions_InvalidPermissionErrorReturned(t *testing.T){
//
//}
////Optional
//func searchWord_InvalidPath_InvalidPathErrorReturned(t *testing.T){
//
//}
