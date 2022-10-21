package main

// define struct VerifyFileNames
type VerifyFileNames struct {
	style       string
	baseDir     string
	extensions  string
	ignore      string
	ignoreFiles string
}

func (v *VerifyFileNames) GetName() string {
	return "VerifyFileNames"
}

func (v *VerifyFileNames) InitSettings(with map[interface{}]interface{}) {
	v.style = with["style"].(string)
	v.baseDir = with["baseDir"].(string)
	v.extensions = with["extensions"].(string)
	v.ignore = with["ignore"].(string)
	v.ignoreFiles = with["ignoreFiles"].(string)
}

func (v *VerifyFileNames) Run(checkInputs CheckInputs) error {
	for _, file := range checkInputs.GetFilePaths() {

		bytes, err := checkInputs.ReadFile(file)
		if err != nil {
			return err
		}

		str := string(bytes)
		str = str
		// TODO: implement
	}
	return nil
}
