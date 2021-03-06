package location

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	CONTENT_FS = iota + 1
	CONTENT_URL
	CONTENT_INLINE
)

type Location struct {
	Path        string
	Kind        int
	Workingpath string
}

type Locations struct {
	Items []*Location
}

func NewLocation(ref string, cwl string) (*Location, error) {
	loc := &Location{}
	if v, _ := url.Parse(ref); v.Scheme == "http" || v.Scheme == "https" {
		loc.Kind = CONTENT_URL
		loc.Path = ref
	} else if _, err := os.Stat(ref); err == nil {
		var err error
		loc.Kind = CONTENT_FS
		loc.Path, err = filepath.Abs(ref)
		if err != nil {
			return loc, err
		}
	} else if _, err := os.Stat(filepath.Join(cwl, ref)); err == nil {
		loc.Kind = CONTENT_FS
		loc.Path = filepath.Join(cwl, ref)
	} else if v, _ := url.Parse(cwl); v.Scheme == "http" || v.Scheme == "https" {
		loc.Kind = CONTENT_URL
		loc.Path = cwl + ref
	} else {
		loc.Kind = CONTENT_INLINE
		loc.Path = ref

		// return fmt.Errorf("unknow location %s -- current working location is %s", ref, cwl)
	}

	loc.Workingpath = cwl
	return loc, nil
}

func (l *Location) Expand() (*Locations, int, error) {
	ls := &Locations{}
	if l.Kind == CONTENT_FS {
		subpaths, err := expandFilePath(l.Path)
		if err != nil {
			return ls, 0, err
		}

		for _, subpath := range subpaths {
			subloc := &Location{
				Path:        subpath,
				Workingpath: l.Workingpath,
				Kind:        l.Kind,
			}
			ls.Items = append(ls.Items, subloc)
		}
	}

	return ls, len(ls.Items), nil
}

func (l *Locations) Add(ref string, cwl string) error {
	loc, err := NewLocation(ref, cwl)
	if err != nil {
		return err
	}

	locs, count, err := loc.Expand()
	if err != nil {
		return err
	}

	if count > 0 {
		l.Items = append(l.Items, locs.Items...)
	} else {
		l.Items = append(l.Items, loc)
	}

	return nil
}

func (l *Location) Content() ([]byte, string, error) {
	return l.ContentWithOptions(map[string]string{})
}
func (l *Location) ContentWithOptions(options map[string]string) ([]byte, string, error) {
	var content []byte
	var cwl string
	var err error

	switch l.Kind {
	case CONTENT_INLINE:
		content = []byte(l.Path)
		cwl = l.Workingpath

	case CONTENT_URL:
		response, err := http.Get(l.Path)
		if err != nil {
			return content, cwl, err
		} else {
			content, err = ioutil.ReadAll(response.Body)
			response.Body.Close()
			if err != nil {
				return content, cwl, err
			}
		}

		uriSegments := strings.Split(l.Path, "/")
		cwl = strings.Join(uriSegments[:len(uriSegments)-1], "/") + "/"

	case CONTENT_FS:

		// si location est relatif
		if false == filepath.IsAbs(l.Path) {
			l.Path = filepath.Join(l.Workingpath, l.Path)
		}

		content, err = ioutil.ReadFile(l.Path)
		if err != nil {
			return content, cwl, fmt.Errorf(`Error while reading "%s" [%s]`, l.Path, err)
		}
		cwl = filepath.Dir(l.Path)
	}

	// find ${FOO:default value} and replace with
	// var["FOO"] if found
	// environnement variaable FOO if env variable exists
	// default value, empty when not provided
	contentString := string(content)
	r, _ := regexp.Compile(`\${([a-zA-Z_\-0-9]+):?([^"'}]*)}`)
	envVars := r.FindAllStringSubmatch(contentString, -1)
	for _, envVar := range envVars {
		varText := envVar[0]
		varName := envVar[1]
		varDefaultValue := envVar[2]

		if value, ok := options[varName]; ok {
			contentString = strings.Replace(contentString, varText, value, -1)
			continue
		}

		// Lookup for env
		if value, found := os.LookupEnv(varName); found {
			contentString = strings.Replace(contentString, varText, value, -1)
			continue
		}
		// Set default value
		contentString = strings.Replace(contentString, varText, varDefaultValue, -1)
		continue
	}
	content = []byte(contentString)

	return content, cwl, err

}

func expandFilePath(path string) ([]string, error) {
	locs := []string{}
	if fi, err := os.Stat(path); err == nil {

		if false == fi.IsDir() {
			locs = append(locs, path)
			return locs, nil
		}
		files, err := filepath.Glob(filepath.Join(path, "*.*"))
		if err != nil {
			return locs, fmt.Errorf("error %s", err.Error())

		}
		//use each file
		for _, file := range files {
			switch strings.ToLower(filepath.Ext(file)) {
			case ".conf":
				locs = append(locs, file)
				continue
			default:

			}
		}
	} else {
		return locs, fmt.Errorf("%s not found", path)
	}
	return locs, nil
}
